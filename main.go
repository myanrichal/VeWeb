package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"flag"
)

func main() {
	fmt.Println("Ve Web");
	
	port := flag.Int("port", 8080, "port to run on.")
	r := gin.Default()
	r.Run(fmt.Sprintf("localhost:%v", *port))
}
