package controller


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Thanakon015/sa-project-66/entity"
)

// POST /users
//สร้างข้อมูล
func CreateUser(c *gin.Context) {
	//ชื่อตาราง
	var user entity.User
	
	if err := c.ShouldBindJSON(&user); err != nil {
	
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	
	return
	
	}
	
	if err := entity.DB().Create(&user).Error; err != nil {
	
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	
	return
	
	}
	
	c.JSON(http.StatusOK, gin.H{"data": user})
	
	} 
