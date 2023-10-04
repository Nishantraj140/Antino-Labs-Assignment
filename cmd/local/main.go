package main

import (
	"Antino-Labs-Assignment/internal/blog"
	"Antino-Labs-Assignment/internal/config"
	"Antino-Labs-Assignment/pkg/logger"
	"Antino-Labs-Assignment/pkg/sql"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

var configFile = flag.String("config", "conf/config.json", "config file")

func main() {

	flag.Parse()
	cnf := config.ReadConfig(*configFile)
	logger.InitLogger("server_log.txt")
	err := sql.DBConn(cnf.DBConfig)

	if err != nil {
		log.Fatal(err)
	}

	defer sql.DB.Close()

	sql.DB.SetLogger(logger.DBLogger)
	sql.DB.LogMode(true)

	gin.DefaultWriter = logger.File

	router := gin.Default()

	router.POST("/article/create", blog.CreateArticle)
	router.GET("/get/article/:id", blog.GetArticleByID)
	router.GET("/get/articles", blog.GetAllArticles)
	router.PATCH("/update/article/:id", blog.UpdateArticle)
	router.DELETE("/delete/article/:id", blog.DeleteArticle)

	logger.InfoLogger.Printf("msg:%v", "server starting....")

	router.Run(":8080")
}
