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
		log.Warn("Error select user")
		return types.UserInfo{}, fmt.Errorf("error: %s: %v", op, err)
	}

	query = fmt.Sprintf("SELECT name, quantiry FROM %s JOIN %s ON %s.item_id = %s.id WHERE user_id = $1", purchasesTable, itemsTable, purchasesTable, itemsTable)
	rows, err := r.db.Query(query, userID)
	if err != nil {
		log.Warn("Error select inventory")
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
		log.Warn("Error select transactiions")
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

	return fmt.Errorf("not implemented")
}

func (r *UserRepository) BuyItem(userID int, itemName string) error {
	const op = "postgres.BuyItem"
	log := slog.With(
		slog.String("op", op),
	)
	log.Info("Call BuyItem")

	return fmt.Errorf("not implemented")
}
