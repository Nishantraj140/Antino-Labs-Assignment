package main

import (
	"Antino-Labs-Assignment/internal/blog"
	"Antino-Labs-Assignment/internal/config"
	"Antino-Labs-Assignment/pkg/sql"
	"flag"
	"log"
)

var configFile = flag.String("config", "conf/config.json", "config file")

func main() {
	flag.Parse()
	cnf := config.ReadConfig(*configFile)
	err := sql.DBConn(cnf.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer sql.DB.Close()
	sql.DB.AutoMigrate(&blog.Article{})
}
