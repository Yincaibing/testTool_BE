package handler

import (
	"SearchTable/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// PolicyHandler 定义一个 db 字段，以便于所有的方法使用
type PolicyHandler struct {
	db *gorm.DB
}

// NewPolicyHandler 在创建 PolicyHandler 的时候，传入 db 参数
func NewPolicyHandler(db *gorm.DB) *PolicyHandler {
	return &PolicyHandler{db: db}
}

func (h *PolicyHandler) HandleGetPolicy() gin.HandlerFunc {
	return func(c *gin.Context) {
		any_id := c.Param("any_id")
		if any_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing any_id"})
			return
		}
		policy, err := domain.GetPolicyByAnyID(h.db, any_id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"policy": policy})
	}
}
