package main

import (
	"github.com/Thanakon015/sa-project-66/entity"
	"github.com/Thanakon015/sa-project-66/controller"
	"github.com/gin-gonic/gin"
)



func main() {

	entity.SetupDatabase()
	
	r := gin.Default()
	
	r.Use(CORSMiddleware())
	
	// User Routes
	//ฟังชั่นที่สร้าง
	r.POST("/users", controller.CreateUser)

	
	// Run the server
	
	r.Run()
	
	}
	
	
	func CORSMiddleware() gin.HandlerFunc {
	
	return func(c *gin.Context) {
	
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	
	
	if c.Request.Method == "OPTIONS" {
	
	c.AbortWithStatus(204)
	
	return
	
	}
	
	
	c.Next()
	
	}
	
	}