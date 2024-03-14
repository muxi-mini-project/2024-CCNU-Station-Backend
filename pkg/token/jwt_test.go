package token

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestParseToken(t *testing.T) {
	//token, _ := GenerateToken("2023214414")
	//token, _ := GenerateToken("2022214214")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdHVpZCI6IjIwMjMyMTQ0MTQiLCJleHAiOjE3MDYxNjE2MzksImlzcyI6Im15LXByb2plY3QifQ.WQziSdcvqKNmovN8uPOsE1YLiCLp3S6_knxYWPhAxXg"
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdHVpZCI6IiIsImV4cCI6MTcwNjE2MTU3MSwiaXNzIjoibXktcHJvamVjdCJ9.9B0JYUqJT5wMOVMaeFwEe16m-6TWNa6RmK696TmcAS4"
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdHVpZCI6IiIsImV4cCI6MTcwNjE2MTMwNywiaXNzIjoibXktcHJvamVjdCJ9.oVo_U83_ysBW-PuaRO9QNLe3o21f6p17p09wa-7LZpw"
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdHVpZCI6IiIsImV4cCI6MTcwNjEwNTQ3MCwiaXNzIjoibXktcHJvamVjdCJ9.UokIQmTVzZUTCW7KlEUJBoTiRWk0pchSozcEuTdIQNk"
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdHVpZCI6IjIwMjIyMTQyMTQiLCJleHAiOjE3MDYxMTEwMzIsImlzcyI6Im15LXByb2plY3QifQ.gXDABYd9W8TXKGhPzHG3PrGgfYi1C3KwSX4onii4ql4"
	fmt.Println("token:", token)
	claims, _ := ParseToken(token)
	fmt.Println(claims.StuID)
}
func TestJWTAuthMiddleware(t *testing.T) {
	r := gin.Default()
	r.GET("/test", JWTAuthMiddleware(), func(c *gin.Context) {
		stuid, _ := c.Get("stuid")
		fmt.Println(stuid.(string))
	})
}
