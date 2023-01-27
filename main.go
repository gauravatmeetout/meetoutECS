package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
   "meetout-ecr/model"
)

func main() {
  var user model.UserObject
  user=model.UserObject{
    Name:"Gaurav",
    Marks:100,
    Role:model.Role{
      Name:"Admin",
      Id:"334",
    },
  }

  r := gin.Default()
  r.GET("/",func(c *gin.Context){
    c.JSON(http.StatusOK,user)
  })

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}