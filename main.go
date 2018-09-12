package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/myanrichal/VeWeb/database"
	"flag"
)

func main() {
	fmt.Println("Ve Web");
	port := flag.Int("port", 5000, "port to run on.")

	db := database.Default()
	db.MigrateSQL()
	
	r := gin.Default()
	r.Run(fmt.Sprintf("localhost:%v", *port))
}
