package user

import "github.com/gin-gonic/gin"

type siginupRequest struct {
    Name string `json:"name"`
}

type signupResponse struct {
    IsSuccess bool `json:"is_success"`
    User map[string]string `json:"user"`
}

func Signup(c *gin.Context) {
    request := siginupRequest{}
    response := signupResponse{}

    c.Bind(&request)

    user := User{}
    user.Signup(request.Name)

    response.IsSuccess = true
    response.User = user.ToMap()
    c.JSON(200, response)
}
