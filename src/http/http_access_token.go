package http

import (
	"net/http"
	// "strings"
	"github.com/gin-gonic/gin"
	atDomain "github.com/selvamshan/bookstore_oauth-api/src/domain/access_token"
	"github.com/selvamshan/bookstore_oauth-api/src/services/access_token"
	"github.com/selvamshan/bookstore_utils-go/rest_errors"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context) 
	UpdateExpirationTime(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := c.Param("access_token_id")
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}


func (handler *accessTokenHandler) Create(c *gin.Context) {
	var atr atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&atr); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	at, err := handler.service.Create(&atr);
	if err != nil {

		c.JSON(err.Status(), err.Error())
		return
	}

	c.JSON(http.StatusCreated, at)
}


func (handler *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {

	c.JSON(http.StatusNotImplemented, "implement me")
}