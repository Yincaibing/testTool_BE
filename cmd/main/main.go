package main

import (
	"SearchTable/config"
	"SearchTable/internal/infra/db"
	"SearchTable/internal/interface/http/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql" //MySQL dialect
	"log"
)

func main() {
	cfgFile := "././deployment/config.yml"         //相对于当前执行程序main.go文件的目录的路径
	cfg, err := config.LoadConfigWithFile(cfgFile) // 加载配置文件
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	err = cfg.Init() // 初始化配置
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
		return
	}

	r := gin.Default()

	// 跨域请求配置
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	db.Init() // 你的数据库连接函数
	DB := db.GetDB()
	// 初始化Handler
	policyHandler := handler.NewPolicyHandler(DB)
	claimHandler := handler.NewClaimHandler(DB)
	allTablesHandler := handler.NewAllTablesHandler(DB)
	// 挂载路由
	r.GET("/policy/:any_id", policyHandler.HandleGetPolicy())
	r.GET("/claim/:any_id", claimHandler.HandleGetClaim())
	r.GET("/reimburse/:any_id", claimHandler.HandleGetClaim())
	r.GET("/search/:any_id", allTablesHandler.HandleGetAllTables())

	r.Run() // listen and serve on 0.0.0.0:8080

}
