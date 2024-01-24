package main

import (
	"SearchTable/internal/infra/db"
	"SearchTable/internal/interface/http/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql" //MySQL dialect
	_ "gitlab.iglooinsure.com/axinan/backend/common/gip_platform_pkg/pkg/config"
	_ "log"
)

func main() {

	r := gin.Default()

	// 跨域请求配置
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	db.Init() // 你的数据库连接函数
	DB := db.GetDB()
	// 初始化Handler
	applicationHandler := handler.NewApplicationHandler(DB)
	transaction_handler := handler.NewTransactionHandler(DB)
	policyHandler := handler.NewPolicyHandler(DB)
	claimHandler := handler.NewClaimHandler(DB)
	reimburseHandler := handler.NewReimburseHandler(DB)
	engineFlowProcessHandler := handler.NewEngineFlowProcessHandler(DB)
	allTablesHandler := handler.NewAllTablesHandler(DB)

	// 挂载路由
	r.GET("/application/:any_id", applicationHandler.HandleGetApplication())
	r.GET("/transaction/:any_id", transaction_handler.HandleGetTransaction())
	r.GET("/policy/:any_id", policyHandler.HandleGetPolicy())
	r.GET("/claim/:any_id", claimHandler.HandleGetClaim())
	r.GET("/reimburse/:any_id", reimburseHandler.HandleGetReimburse())
	r.GET("/engine_flow_process/:any_id", engineFlowProcessHandler.HandleGetEngineFlowProcess())

	r.GET("/search/:any_id", allTablesHandler.HandleGetAllTables())

	r.Run() // listen and serve on 0.0.0.0:8080

}
