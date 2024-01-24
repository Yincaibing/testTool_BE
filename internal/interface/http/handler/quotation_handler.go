package handler

import (
	"SearchTable/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// QuotationHandler 定义一个 db 字段，以便于所有的方法使用
type QuotationHandler struct {
	db *gorm.DB
}

// NewQuotationHandler 在创建 QuotationHandler 的时候，传入 db 参数
func NewQuotationHandler(db *gorm.DB) *QuotationHandler {
	return &QuotationHandler{db: db}
}

func (h *QuotationHandler) HandleGetQuotation() gin.HandlerFunc {
	return func(c *gin.Context) {
		any_id := c.Param("any_id")
		if any_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing any_id"})
			return
		}
		quotation, err := domain.GetQuotationByAnyID(h.db, any_id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Quotation": quotation})
	}
}
