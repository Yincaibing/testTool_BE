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
	qaDB, stagingDB := db.GetDB()
	// 初始化QAHandler
	QquotationHandler := handler.NewQuotationHandler(qaDB)
	QapplicationHandler := handler.NewApplicationHandler(qaDB)
	Qtransaction_handler := handler.NewTransactionHandler(qaDB)
	QpolicyHandler := handler.NewPolicyHandler(qaDB)
	QclaimHandler := handler.NewClaimHandler(qaDB)
	QreimburseHandler := handler.NewReimburseHandler(qaDB)
	QengineFlowProcessHandler := handler.NewEngineFlowProcessHandler(qaDB)
	QallTablesHandler := handler.NewAllTablesHandler(qaDB)

	// 初始化StagingHandler
	SquotationHandler := handler.NewQuotationHandler(stagingDB)
	SapplicationHandler := handler.NewApplicationHandler(stagingDB)
	Stransaction_handler := handler.NewTransactionHandler(stagingDB)
	SpolicyHandler := handler.NewPolicyHandler(stagingDB)
	SclaimHandler := handler.NewClaimHandler(stagingDB)
	SreimburseHandler := handler.NewReimburseHandler(stagingDB)
	SengineFlowProcessHandler := handler.NewEngineFlowProcessHandler(stagingDB)
	SallTablesHandler := handler.NewAllTablesHandler(stagingDB)
	// 挂载路由
	//qa环境的
	r.GET("/qa/quotation/:any_id", QquotationHandler.HandleGetQuotation())
	r.GET("/qa/application/:any_id", QapplicationHandler.HandleGetApplication())
	r.GET("/qa/transaction/:any_id", Qtransaction_handler.HandleGetTransaction())
	r.GET("/qa/policy/:any_id", QpolicyHandler.HandleGetPolicy())
	r.GET("/qa/claim/:any_id", QclaimHandler.HandleGetClaim())
	r.GET("/qa/reimburse/:any_id", QreimburseHandler.HandleGetReimburse())
	r.GET("/qa/engine_flow_process/:any_id", QengineFlowProcessHandler.HandleGetEngineFlowProcess())
	r.GET("/qa/search/:any_id", QallTablesHandler.HandleGetAllTables())

	// staging环境的
	r.GET("/staging/quotation/:any_id", SquotationHandler.HandleGetQuotation())
	r.GET("/staging/application/:any_id", SapplicationHandler.HandleGetApplication())
	r.GET("/staging/transaction/:any_id", Stransaction_handler.HandleGetTransaction())
	r.GET("/staging/policy/:any_id", SpolicyHandler.HandleGetPolicy())
	r.GET("/staging/claim/:any_id", SclaimHandler.HandleGetClaim())
	r.GET("/staging/reimburse/:any_id", SreimburseHandler.HandleGetReimburse())
	r.GET("/staging/engine_flow_process/:any_id", SengineFlowProcessHandler.HandleGetEngineFlowProcess())
	r.GET("/staging/search/:any_id", SallTablesHandler.HandleGetAllTables())

	r.Run() // listen and serve on 0.0.0.0:8080

}
