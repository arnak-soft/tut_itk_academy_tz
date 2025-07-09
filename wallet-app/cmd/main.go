package main

import (
    "context"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "wallet-app/internal/handlers"
    "wallet-app/internal/repository"
    "wallet-app/internal/services"
)

func main() {
    // Считываем переменные окружения
    dsn := os.Getenv("DATABASE_URL")

    // Создаем репозиторий
    repo, err := repository.NewMySQLRepository(dsn)
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    // Создаем сервис и хендлер
    svc := services.NewWalletService(repo)
    h := handlers.NewWalletHandler(svc)

    r := gin.Default()
    r.POST("/api/v1/wallet", h.HandleWalletOperation)
    r.GET("/api/v1/wallets/:walletId", h.GetWalletBalance)

    r.Run(":8080")
}