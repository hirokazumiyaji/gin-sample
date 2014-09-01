package main

import (
    "github.com/gin-gonic/gin"
    "./user"
)

func main() {
    r := gin.Default()

    r.POST("/signup", user.Signup)
    r.GET("/signin", user.Signin)
    r.GET("/search/:username", user.Search)

    r.Run(":8080")
}
