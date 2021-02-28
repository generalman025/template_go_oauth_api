package rest

import (
	"encoding/json"
	"fmt"

	"github.com/generalman025/template_go_oauth_api/domain/users"
	"github.com/generalman025/template_go_util_lib_api/rest_errors"
	"github.com/go-resty/resty/v2"
)

const (
	baseURL = "http://localhost:8081"
)

var (
	// usersRestClient = rest.RequestBuilder{
	// 	BaseURL: "http://localhost:8081",
	// 	Timeout: 100 * time.Millisecond,
	// }

	usersRestClient = resty.New().R()
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

type usersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response, _ := usersRestClient.SetBody(request).Post(fmt.Sprintf("%s/users/login", baseURL))
	if response == nil {
		return nil, rest_errors.NewInternalServerError("invalid resclient response when trying to login user", rest_errors.NewError("error when trying to login user"))
	}
	if response.StatusCode() > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Body())
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", rest_errors.NewError("error when trying to login user"))
		}
		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal users response", rest_errors.NewError("error when trying to login user"))
	}

	return &user, nil
}
