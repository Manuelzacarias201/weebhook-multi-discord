package infraestructure

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {

	routes_w := router.Group("github")

	{
		routes_w.POST("/webhook", HandleGitHubEvents)
	}

}
