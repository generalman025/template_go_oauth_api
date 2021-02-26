package app

import (
	"github.com/generalman025/template_go_oauth_api/clients/cassandra"
	"github.com/generalman025/template_go_oauth_api/http"
	"github.com/generalman025/template_go_oauth_api/repository/db"
	"github.com/generalman025/template_go_oauth_api/repository/rest"
	"github.com/generalman025/template_go_oauth_api/services/access_token"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	// test
)

func StartApplication() {
	defer cassandra.GetSession().Close()
	atHandler := http.NewAccessTokenHandler(access_token.NewService(rest.NewRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
