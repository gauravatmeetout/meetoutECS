package main

import (
	"fmt"
	"meetout-ecr/libs"
	"meetout-ecr/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	dao := libs.ConnectToDB("meetout-ecs", "Pass@123", "meetout-ecs-01.ckreqp0.mongodb.net", "meetout")
	dao.ListDatabaseNames(false)
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		userModel := model.UserModel{Tablename: "users"}
		users, err := userModel.GetAllUsers(dao)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, users)
		}

	})

	r.POST("/users/add", func(c *gin.Context) {
		var user model.User
		if err := c.BindJSON(&user); err == nil {
			userModel := model.UserModel{
				Tablename: "users",
				User:      user,
			}
			users, err := userModel.Insert(dao)
			if err != nil {
				fmt.Println("Error Occured", err)
				c.JSON(http.StatusInternalServerError, err)
			} else {
				c.JSON(http.StatusOK, users)
			}
		} else {
			fmt.Println(err, user)
			c.JSON(http.StatusInternalServerError, err)
		}
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
