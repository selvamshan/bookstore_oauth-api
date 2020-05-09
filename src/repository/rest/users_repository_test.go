package rest

import (
	//"fmt"
	//"flag"
	"net/http"
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/mercadolibre/golang-restclient/rest"
)

// var _ = func () bool {
// 	testing.Init()
// 	return true
// }()

// func init() {
// 	flag.Parse()
	
// }

func TestMain(m *testing.M) {
	//fmt.Println("about to start test cases ...")
	defer rest.StopMockupServer()
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	// defer rest.StopMockupServer()
	// rest.StartMockupServer()
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "http://localhost:8080/users/login",		
		ReqBody: `{"email":"email@abc.com","password":"passwd"}`,
		RespHTTPCode:-1,
		RespBody: `{}`,
	})
	repository := usersRepository{}

	user, err:= repository.LoginUser("eamil@abc.com", "passwd")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user",err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T){
	// defer rest.StopMockupServer()
	// rest.StartMockupServer()
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "http://localhost:8080/users/login",		
		ReqBody: `{"email":"email@abc.com","password":"passwd"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"message": "no user found with given credintials", "status": "404", "error": "not_found"}`,
	})
	repository := usersRepository{}

	user, err:= repository.LoginUser("eamil@abc.com", "passwd")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user",err.Message)
}


func TestLoginUserInvalidLoginCredentials(t *testing.T){
	// defer rest.StopMockupServer()
	// rest.StartMockupServer()
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "http://localhost:8080/users/login",		
		ReqBody: `{"email":"email@abc.com","password":"passwd"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"message": "no user found with given credintials", "status": 404, "error": "not_found"}`,
	})
	repository := usersRepository{}

	user, err:= repository.LoginUser("eamil@abc.com", "passwd")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "no user found with given credintials",err.Message)

}

func TestLoginUserInvalidUserJsonResponse(t *testing.T){
	// defer rest.StopMockupServer()
	// rest.StartMockupServer()
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "http://localhost:8080/users/login",		
		ReqBody: `{"email":"email@abc.com","password":"passwd"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id": "7", "first_name": "jeeth", "last_name": "selva", "email": "jeeth@abc.com"}`,
	})
	repository := usersRepository{}

	user, err:= repository.LoginUser("eamil@abc.com", "passwd")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "errors when trying to unmarshal users response",err.Message)
}

func TestLoginUserNoError(t *testing.T){
	// defer rest.StopMockupServer()
	// rest.StartMockupServer()
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "http://localhost:8080/users/login",		
		ReqBody: `{"email":"email@abc.com","password":"passwd"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id": 7, "first_name": "jeeth", "last_name": "selva", "email": "jeeth@abc.com"}`,
	})
	repository := usersRepository{}

	user, err:= repository.LoginUser("eamil@abc.com", "passwd")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "jeeth", user.FirstName)
	assert.EqualValues(t, "selva", user.LastName)
	assert.EqualValues(t, "jeeth@abc.com", user.Email)
}