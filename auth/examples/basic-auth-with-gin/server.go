package main

import "github.com/gin-gonic/gin"

const realm = "web-dev-auth"

var users = map[string]string{
	"root": "123456",
	// other users
}

func main() {
	r := gin.Default()

	r.GET("/protected", gin.BasicAuthForRealm(users, realm), func(ctx *gin.Context) {
		ctx.String(200, "Aha! You got me!")
	})

	r.Run(":3000")
}
