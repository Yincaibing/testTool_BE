package handler

import (
	"SearchTable/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// ClaimHandler 定义一个 db 字段，以便于所有的方法使用
type ClaimHandler struct {
	db *gorm.DB
}

// NewClaimHandler 在创建 ClaimHandler 的时候，传入 db 参数
func NewClaimHandler(db *gorm.DB) *ClaimHandler {
	return &ClaimHandler{db: db}
}

func (h *ClaimHandler) HandleGetClaim() gin.HandlerFunc {
	return func(c *gin.Context) {
		//这个查询参数可能是claim表中任一参数：id,display_id，insure_process_id，external_id，policy_id
		anyId := c.Param("any_id")
		if anyId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing any_id"})
			return
		}
		claim, err := domain.GetClaimByAnyId(h.db, anyId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"claim": claim})
	}
}
