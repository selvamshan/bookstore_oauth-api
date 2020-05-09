package access_token

import (
	"strings"
	"github.com/selvamshan/bookstore_utils-go/rest_errors"
	"github.com/selvamshan/bookstore_oauth-api/src/repository/db"
	"github.com/selvamshan/bookstore_oauth-api/src/repository/rest"
	"github.com/selvamshan/bookstore_oauth-api/src/domain/access_token"
)

// type Repository interface {
// 	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
// 	Create(*access_token.AccessTokenRequest) *rest_errors.RestErr
// 	UpdateExpirationTime(*AccessToken) *rest_errors.RestErr
// }

type Service interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestErr)
	Create(*access_token.AccessTokenRequest) (*access_token.AccessToken, rest_errors.RestErr)
	UpdateExpirationTime(*access_token.AccessToken) rest_errors.RestErr
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo       db.DbRepository
}

func NewService(usersRepo rest.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: usersRepo, 
		dbRepo:        dbRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, rest_errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0{
		return nil, rest_errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil{
		return nil, err
	}
	return accessToken, nil
	
}

func (s *service) Create(atr *access_token.AccessTokenRequest)(*access_token.AccessToken, rest_errors.RestErr) {
	
	if err := atr.Validate(); err != nil {
		return nil, err 
	}
	
	//TODO: Support both client_credendtials and password
	// Auterntiacate the user against the User API
	user, err := s.restUsersRepo.LoginUser(atr.Username, atr.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	// Save the acces token in Cassenda
	if err := s.dbRepo.Create(&at); err != nil {
		return nil, err
	}

	return &at, nil
}



func (s *service) UpdateExpirationTime(at *access_token.AccessToken) rest_errors.RestErr {
	if err := at.Validate(); err != nil {
		return err 
	}
	return s.dbRepo.UpdateExpirationTime(at)
}

