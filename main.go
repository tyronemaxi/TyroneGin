package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	IniteRoute(r)
	r.Run()
}

