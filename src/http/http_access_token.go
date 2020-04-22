package http

import (
	"net/http"
	// "strings"
	"github.com/gin-gonic/gin"
	"github.com/selvamshan/bookstore_oauth-api/src/domain/access_token"
	"github.com/selvamshan/bookstore_oauth-api/src/utils/errors"
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
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}


func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := handler.service.Create(&at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, at)
}


func (handler *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {

	c.JSON(http.StatusNotImplemented, "implement me")
}