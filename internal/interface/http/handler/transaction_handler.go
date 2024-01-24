package handler

import (
	"SearchTable/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// ApplicationHandler 定义一个 db 字段，以便于所有的方法使用
type TransactionHandler struct {
	db *gorm.DB
}

// NewApplicationHandler 在创建 ApplicationHandler 的时候，传入 db 参数
func NewTransactionHandler(db *gorm.DB) *TransactionHandler {
	return &TransactionHandler{db: db}
}

func (h *TransactionHandler) HandleGetTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		anyId := c.Param("any_id")
		if anyId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing any_id"})
			return
		}
		transaction, err := domain.GetTransactionByAnyID(h.db, anyId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Transaction": transaction})
	}
}
