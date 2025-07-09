package repository

import (
	"database/sql"
	"context"
	"github.com/google/uuid"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type WalletRepository interface {
	UpdateBalance(ctx context.Context, walletId uuid.UUID, amount float64) error
	GetBalance(ctx context.Context, walletId uuid.UUID) (float64, error)
}

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(dsn string) (*MySQLRepository, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &MySQLRepository{db: db}, nil
}

func (r *MySQLRepository) UpdateBalance(ctx context.Context, walletId uuid.UUID, amount float64) error {
	_, err := r.db.ExecContext(ctx, "UPDATE wallets SET balance = balance + ? WHERE id = ?", amount, walletId)
	if err != nil {
		log.Printf("Error updating balance: %v", err)
		return err
	}
	return nil
}

func (r *MySQLRepository) GetBalance(ctx context.Context, walletId uuid.UUID) (float64, error) {
	var balance float64
	err := r.db.QueryRowContext(ctx, "SELECT balance FROM wallets WHERE id = ?", walletId).Scan(&balance)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error getting balance: %v", err)
		return 0, err
	}
	return balance, nil
}