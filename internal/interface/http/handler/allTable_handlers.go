package handler

import (
	"SearchTable/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// AllHandler 定义一个 db 字段，以便于所有的方法使用
type AllTablesHandler struct {
	db *gorm.DB
}

// NewAllTablesHandler 在创建 AllTablesHandler 的时候，传入 db 参数
func NewAllTablesHandler(db *gorm.DB) *AllTablesHandler {
	return &AllTablesHandler{db: db}
}

func (h *AllTablesHandler) HandleGetAllTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		anyId := c.Param("any_id")
		if anyId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing any_id"})
			return
		}

		policy, err := domain.GetPolicyByAnyID(h.db, anyId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		claim, err := domain.GetClaimByAnyId(h.db, anyId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		reimburse, err := domain.GetReimburseByAnyID(h.db, anyId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		application, err := domain.GetApplicationByAnyId(h.db, anyId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		engine_flow_process, err := domain.GetEngineFlowProcessByAnyId(h.db, anyId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		transaction, err := domain.GetTransactionByAnyID(h.db, anyId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var isCollapsePolicy, isCollapseClaim, isCollapseReimburse, isCollapseTransaction, isCollapseApplication, isCollapseEngineFlowProcess bool
		if policy == nil { // 这里你需要根据实际数据模型调整来判断数据是否为空
			isCollapsePolicy = true
		}

		if claim == nil { // 这里你需要根据实际数据模型调整来判断数据是否为空
			isCollapseClaim = true
		}

		if reimburse == nil { // 这里你需要根据实际数据模型调整来判断数据是否为空
			isCollapseReimburse = true
		}

		if transaction == nil { // 这里你需要根据实际数据模型调整来判断数据是否为空
			isCollapseTransaction = true
		}

		if application == nil { // 这里你需要根据实际数据模型调整来判断数据是否为空
			isCollapseApplication = true
		}

		if engine_flow_process == nil { // 这里你需要根据实际数据模型调整来判断数据是否为空
			isCollapseEngineFlowProcess = true
		}

		allTables := []domain.AllTables{
			{
				Name:       "Application",
				Data:       application,
				IsCollapse: isCollapseApplication,
			},
			{
				Name:       "Policy",
				Data:       policy,
				IsCollapse: isCollapsePolicy,
			},
			{
				Name:       "Claim",
				Data:       claim,
				IsCollapse: isCollapseClaim,
			},
			{
				Name:       "Reimburse",
				Data:       reimburse,
				IsCollapse: isCollapseReimburse,
			},
			{
				Name:       "Transaction",
				Data:       transaction,
				IsCollapse: isCollapseTransaction,
			},
			{
				Name:       "engine_flow_process",
				Data:       engine_flow_process,
				IsCollapse: isCollapseEngineFlowProcess,
			},
		}

		c.JSON(http.StatusOK, gin.H{"AllTables": allTables})
	}
}
