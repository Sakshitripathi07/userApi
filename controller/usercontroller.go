package controller

import (
	"net/http"
	"strconv"
	"userApi/database"
	"userApi/model"

	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	var users []model.User
	err := database.DB.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})

	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})

}

func FindUserbyID(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	user.ID = id
	err = database.DB.Where("id = ?", user.ID).Find(&user).Error

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "could not load the data",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}
