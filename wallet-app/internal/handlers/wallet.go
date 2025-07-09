package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "wallet-app/internal/models"
    "wallet-app/internal/services"
)

type WalletHandler struct {
    service *services.WalletService
}

func NewWalletHandler(service *services.WalletService) *WalletHandler {
    return &WalletHandler{service: service}
}

func (h *WalletHandler) HandleWalletOperation(c *gin.Context) {
    var req models.WalletOperationRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    walletId, err := uuid.Parse(req.WalletId)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Wallet ID"})
        return
    }

    ctx := c.Request.Context()
    err = h.service.HandleOperation(ctx, walletId, req.OperationType, req.Amount)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *WalletHandler) GetWalletBalance(c *gin.Context) {
    walletId, err := uuid.Parse(c.Param("walletId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Wallet ID"})
        return
    }

    ctx := c.Request.Context()
    balance, err := h.service.GetBalance(ctx, walletId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"balance": balance})
}