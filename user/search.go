package user

import "github.com/gin-gonic/gin"

func Search(c *gin.Context) {
    userName := c.Params.ByName("username")
    user := FindUser(userName)
    c.JSON(200, user)
}
