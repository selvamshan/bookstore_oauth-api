package db

import (	
	"github.com/gocql/gocql"
	"github.com/selvamshan/bookstore_oauth-api/src/domain/access_token"
	"github.com/selvamshan/bookstore_oauth-api/src/utils/errors"
	"github.com/selvamshan/bookstore_oauth-api/src/clients/cassandra"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccesToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?);"
	queryUpdateExpirationTime = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(*access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(*access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {

}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (db *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session := cassandra.GetSession()
	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId, 
		&result.ClientId,
		&result.Expires,
	); err != nil{
		if err == gocql.ErrNotFound {
			return nil, errors.NewInternalServerError("no access_token found for given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}	
	return &result, nil
}

func (db *dbRepository) Create(at *access_token.AccessToken) *errors.RestErr {
	session := cassandra.GetSession()	
	
	if err := session.Query(queryCreateAccesToken,
		 at.AccessToken, 
		 at.UserId, 
		 at.ClientId, 
		 at.Expires,
		 ). Exec(); err != nil {
			return errors.NewInternalServerError(err.Error())	
		 }
	return nil
}


func (db *dbRepository) UpdateExpirationTime(at *access_token.AccessToken) *errors.RestErr {
	session := cassandra.GetSession()
	if err := session.Query(queryUpdateExpirationTime,
		at.Expires,
		at.AccessToken,		
		 ). Exec(); err != nil {
			return  errors.NewInternalServerError(err.Error())	
		 }
	return nil
}


