package app

import (
	"github.com/PS-07/bookstore_oauth-api/src/http"
	"github.com/PS-07/bookstore_oauth-api/src/repository/db"
	"github.com/PS-07/bookstore_oauth-api/src/repository/rest"
	"github.com/PS-07/bookstore_oauth-api/src/services/accesstoken"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// StartApp func
func StartApp() {
	dbRepository := db.NewRepository()
	restRepository := rest.NewRestUsersRepository()
	atService := accesstoken.NewService(restRepository, dbRepository)
	atHandler := http.NewAccessTokenHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
