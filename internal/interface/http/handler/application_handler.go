package handler

import (
	"SearchTable/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// ApplicationHandler 定义一个 db 字段，以便于所有的方法使用
type ApplicationHandler struct {
	db *gorm.DB
}

// NewApplicationHandler 在创建 ApplicationHandler 的时候，传入 db 参数
func NewApplicationHandler(db *gorm.DB) *ApplicationHandler {
	return &ApplicationHandler{db: db}
}

func (h *ApplicationHandler) HandleGetApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		anyId := c.Param("any_id")
		if anyId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing any_id"})
			return
		}
		application, err := domain.GetApplicationByAnyId(h.db, anyId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Application": application})
	}
}
