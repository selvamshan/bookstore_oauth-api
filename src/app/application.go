package app

import (
	"github.com/gin-gonic/gin"
	"github.com/selvamshan/bookstore_oauth-api/src/repository/db"
	"github.com/selvamshan/bookstore_oauth-api/src/repository/rest"
	"github.com/selvamshan/bookstore_oauth-api/src/services/access_token"
	"github.com/selvamshan/bookstore_oauth-api/src/http"
	// "github.com/selvamshan/bookstore_oauth-api/src/clients/cassandra"
)

var (
	router = gin.Default()
)

func StartApplication() {
	

	atService := access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository())
	atHandler := http.NewAccessTokenHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.PUT("/oauth/access_token", atHandler.UpdateExpirationTime)
	router.Run(":8080")
}