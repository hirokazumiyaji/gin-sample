package user

import "github.com/gin-gonic/gin"

type SigninResponse struct {
    IsSuccess bool `json:"is_success"`
}

func Signin(c *gin.Context) {
    response := SigninResponse{false}
    response.IsSuccess = true
    c.JSON(200, response)
}
