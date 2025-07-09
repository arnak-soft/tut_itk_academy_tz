package services

import (
    "context"
    "github.com/google/uuid"
    "wallet-app/internal/repository"
)

type WalletService struct {
    repo repository.WalletRepository
}

func NewWalletService(repo repository.WalletRepository) *WalletService {
    return &WalletService{repo: repo}
}

func (s *WalletService) HandleOperation(ctx context.Context, walletId uuid.UUID, operationType string, amount float64) error {
    // Проверка существования кошелька, если нет - создаем
    if err := s.repo.CreateWallet(ctx, walletId); err != nil {
        return err
    }

    // Логика обработки операций DEPOSIT и WITHDRAW
    if operationType == "DEPOSIT" {
        return s.repo.UpdateBalance(ctx, walletId, amount)
    } else if operationType == "WITHDRAW" {
        return s.repo.UpdateBalance(ctx, walletId, -amount)
    }

    return nil
}

func (s *WalletService) GetBalance(ctx context.Context, walletId uuid.UUID) (float64, error) {
    return s.repo.GetBalance(ctx, walletId)
}