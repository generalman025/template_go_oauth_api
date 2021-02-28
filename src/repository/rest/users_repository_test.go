package rest

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	// rest.FlushMockups()
	// rest.AddMockups(&rest.Mock{
	// 	URL:          "https://api.template.com/users/login",
	// 	HTTPMethod:   http.MethodPost,
	// 	ReqBody:      `{"email":"email@gmail.com", "password":"the-password"}`,
	// 	RespHTTPCode: -1,
	// 	RespBody:     `{}`,
	// })
	// repository := usersRepository{}

	// user, err := repository.LoginUser("email@gmail.com", "the-password")
	// assert.Nil(t, user)
	// assert.NotNil(t, err)
	// assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	// assert.EqualValues(t, "invalid resclient response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	// rest.FlushMockups()
	// rest.AddMockups(&rest.Mock{
	// 	URL:          "https://api.template.com/users/login",
	// 	HTTPMethod:   http.MethodPost,
	// 	ReqBody:      `{"email":"email@gmail.com", "password":"the-password"}`,
	// 	RespHTTPCode: http.StatusNotFound,
	// 	RespBody:     `{"message":"invalid login credentials", "status":404, "error":"not_found"}`,
	// })
	// repository := usersRepository{}

	// user, err := repository.LoginUser("email@gmail.com", "the-password")
	// assert.Nil(t, user)
	// assert.NotNil(t, err)
	// assert.EqualValues(t, http.StatusNotFound, err.Status)
	// assert.EqualValues(t, "invalid resclient response when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredential(t *testing.T) {
	// rest.FlushMockups()
	// rest.AddMockups(&rest.Mock{
	// 	URL:          "https://api.template.com/users/login",
	// 	HTTPMethod:   http.MethodPost,
	// 	ReqBody:      `{"email":"email@gmail.com", "password":"the-password"}`,
	// 	RespHTTPCode: http.StatusNotFound,
	// 	RespBody:     `{"message":"invalid login credentials", "status":404, "error":"not_found"}`,
	// })
	// repository := usersRepository{}

	// user, err := repository.LoginUser("email@gmail.com", "the-password")
	// assert.Nil(t, user)
	// assert.NotNil(t, err)
	// assert.EqualValues(t, http.StatusNotFound, err.Status)
	// assert.EqualValues(t, "invalid login credentials ", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	// rest.FlushMockups()
	// rest.AddMockups(&rest.Mock{
	// 	URL:          "https://api.template.com/users/login",
	// 	HTTPMethod:   http.MethodPost,
	// 	ReqBody:      `{"email":"email@gmail.com", "password":"the-password"}`,
	// 	RespHTTPCode: http.StatusOK,
	// 	RespBody:     `{"id": 6, "first_name": "Kiattisak", "last_name": "Chaisomboon", "email": "bass.kiattisak@gmail.com"}`,
	// })

	// repository := usersRepository{}

	// user, err := repository.LoginUser("email@gmail.com", "the-password")
	// assert.Nil(t, user)
	// assert.NotNil(t, err)
	// assert.EqualValues(t, http.StatusNotFound, err.Status)
	// assert.EqualValues(t, "error when trying to unmarshal users response ", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	// rest.FlushMockups()
	// rest.AddMockups(&rest.Mock{
	// 	URL:          "https://api.template.com/users/login",
	// 	HTTPMethod:   http.MethodPost,
	// 	ReqBody:      `{"email":"email@gmail.com", "password":"the-password"}`,
	// 	RespHTTPCode: http.StatusOK,
	// 	RespBody:     `{"id": 1, "first_name": "Kiattisak", "last_name": "Chaisomboon", "email": "bass.kiattisak@gmail.com"}`,
	// })

	// repository := usersRepository{}

	// user, err := repository.LoginUser("email@gmail.com", "the-password")
	// assert.Nil(t, user)
	// assert.NotNil(t, err)
	// assert.EqualValues(t, 1, user.ID)
	// assert.EqualValues(t, "Kiattisak", user.FirstName)
	// assert.EqualValues(t, "Chaisomboon", user.LastName)
	// assert.EqualValues(t, "bass.kiattisak@gmail.com", user.Email)
}
