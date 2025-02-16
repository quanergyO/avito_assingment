package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/quanergyO/avito_assingment/types"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserInfo(userID int) (types.UserInfo, error) {
	const op = "postgres.GetUserInfo"
	log := slog.With(
		slog.String("op", op),
	)
	log.Info("Call GetUserInfo")

	var userInfo types.UserResponse
	query := fmt.Sprintf("SELECT username, coins FROM %s WHERE id=$1", userTable)
	err := r.db.QueryRow(query, userID).Scan(&userInfo.Username, &userInfo.Coins)
	if err != nil {
		log.Info("Error select user")
		return types.UserInfo{}, fmt.Errorf("error: %s: %v", op, err)
	}

	query = fmt.Sprintf("SELECT name, quantity FROM %s JOIN %s ON %s.item_id = %s.id WHERE user_id = $1", purchasesTable, itemsTable, purchasesTable, itemsTable)
	rows, err := r.db.Query(query, userID)
	if err != nil {
		log.Info("Error select inventory")
		return types.UserInfo{}, fmt.Errorf("error: %s: %v", op, err)
	}

	inventory := make([]types.PurchasesResponse, 0)
	for rows.Next() {
		var item types.PurchasesResponse
		if err := rows.Scan(&item.ItemName, &item.Quantity); err != nil {
			return types.UserInfo{}, err
		}
		inventory = append(inventory, item)
	}
	rows.Close()

	transactions := make([]types.TransactionsResponse, 0)
	query = fmt.Sprintf("SELECT from_user_id, to_user_id, timestamp, amount FROM %s WHERE from_user_id = $1 OR to_user_id = $2", transactionsTable)
	rows, err = r.db.Query(query, userID, userID)
	if err != nil {
		log.Info("Error select transactiions")
		return types.UserInfo{}, fmt.Errorf("error: %s: %v", op, err)
	}
	for rows.Next() {
		var transaction types.TransactionsResponse
		if err := rows.Scan(&transaction.FromUserId, &transaction.ToUserId, &transaction.Timestamp, &transaction.Amount); err != nil {
			return types.UserInfo{}, err
		}
		transactions = append(transactions, transaction)
	}

	return types.UserInfo{
		UserInfo:        userInfo,
		PurchasesInfo:   inventory,
		TransactionInfo: transactions,
	}, nil
}

func (r *UserRepository) SendCoins(senderID, receiverID int, amount int) error {
	const op = "postgres.SendCoins"
	log := slog.With(
		slog.String("op", op),
	)
	log.Info("Call SendCoins")

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("error start transaction")
	}

	var senderDAO types.UserDAO
	query := fmt.Sprintf("SELECT username, coins FROM %s WHERE id = $1", userTable)
	if err := tx.QueryRow(query, senderID).Scan(&senderDAO.Username, &senderDAO.Coins); err != nil {
		log.Info("Can't get senderDAO from db")
		tx.Rollback()
		return fmt.Errorf("can't get senderDAO from db")
	}
	if senderDAO.Coins < amount {
		tx.Rollback()
		log.Info("NO MONEY NO HONEY")
		return fmt.Errorf("NO MONEY NO HONEY") // TODO change exception
	}

	query = fmt.Sprintf("UPDATE %s SET coins = $1 WHERE id = $2", userTable)
	if _, err := tx.Exec(query, senderDAO.Coins-amount, senderID); err != nil {
		tx.Rollback()
		log.Info("can't update sender coins")
		return fmt.Errorf("can't update sender coins")
	}

	query = fmt.Sprintf("UPDATE %s SET coins = coins + $1 WHERE id = $2", userTable)
	if _, err := tx.Exec(query, amount, receiverID); err != nil {
		tx.Rollback()
		log.Info("can't update receiver coins")
		return fmt.Errorf("can't update receiver coins")
	}

	query = fmt.Sprintf("INSERT INTO %s (from_user_id, to_user_id, amount) VALUES ($1, $2, $3)", transactionsTable)
	if _, err := tx.Exec(query, senderID, receiverID, amount); err != nil {
		tx.Rollback()
		log.Info("can't insert in transaction db", slog.Any("error", err.Error()))
		return fmt.Errorf("can't insert in transaction db")
	}

	return tx.Commit()
}

func (r *UserRepository) BuyItem(userID int, itemName string) error {
	const op = "postgres.BuyItem"
	log := slog.With(
		slog.String("op", op),
	)
	log.Info("Call BuyItem")

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("error start transaction")
	}

	query := fmt.Sprintf("SELECT coins FROM %s WHERE id=$1", userTable)
	var coins int
	if err := tx.QueryRow(query, userID).Scan(&coins); err != nil {
		tx.Rollback()
		return fmt.Errorf("error get balance")
	}

	var item types.ItemDAO
	query = fmt.Sprintf("SELECT id, price FROM %s WHERE name=$1", itemsTable)
	if err := tx.QueryRow(query, itemName).Scan(&item.Id, &item.Price); err != nil {
		tx.Rollback()
		return fmt.Errorf("error get price")
	}

	if coins < item.Price {
		tx.Rollback()
		return fmt.Errorf("NO MONEY NO HONEY") // TODO change exception
	}

	query = fmt.Sprintf("INSERT INTO %s (user_id, item_id) VALUES ($1, $2)", purchasesTable)
	if _, err := tx.Exec(query, item.Id, item.Price); err != nil {
		tx.Rollback()
		return fmt.Errorf("error insert in purchases")
	}

	query = fmt.Sprintf("UPDATE %s SET coins = coins - $1 WHERE id = $2", userTable)
	if _, err := tx.Exec(query, item.Price, userID); err != nil {
		tx.Rollback()
		return fmt.Errorf("error update balance")
	}

	return tx.Commit()
}
