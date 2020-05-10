package rest

import (
	"fmt"
	//"flag"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)



func TestMain(m *testing.M) {
	fmt.Println("about to start test cases ...")	
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	// defer rest.StopMockupServer()
	// rest.StartMockupServer()
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8082/users/login",
		ReqBody:      `{"email":"email@abc.com","password":"passwd"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("eamil@abc.com", "passwd")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message())
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	// defer rest.StopMockupServer()
	// rest.StartMockupServer()
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8080/users/login",
		ReqBody:      `{"email":"email1@abc.com","password":"passwd"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "no user found with given credintials", "status": "404", "error": "not_found"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("eamil1@abc.com", "passwd")
	fmt.Println("IN LINE 66",err)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message())
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	// defer rest.StopMockupServer()
	// rest.StartMockupServer()
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8082/users/login",
		ReqBody:      `{"email":"email@abc.com","password":"passwd"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "no user found with given credintials", "status": 404, "error": "not_found"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("eamil@abc.com", "passwd")
	//fmt.Println("IN LINE 88",err)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "no user found with given credintials", err.Message())

}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	// defer rest.StopMockupServer()
	// rest.StartMockupServer()
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8082/users/login",
		ReqBody:      `{"email":"email@abc.com","password":"passwd"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "7", "first_name": "jeeth", "last_name": "selva", "email": "jeeth@abc.com"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("eamil@abc.com", "passwd")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "rest_errors when trying to unmarshal users response", err.Message())
}

func TestLoginUserNoError(t *testing.T) {
	// defer rest.StopMockupServer()
	// rest.StartMockupServer()
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost:8082/users/login",
		ReqBody:      `{"email":"jeeth@abc.com","password":"passwd"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 7, "first_name": "jeeth", "last_name": "selva", "email": "jeeth@abc.com"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("eamil@abc.com", "passwd")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 7, user.Id)
	assert.EqualValues(t, "jeeth", user.FirstName)
	assert.EqualValues(t, "selva", user.LastName)
	assert.EqualValues(t, "jeeth@abc.com", user.Email)
}
