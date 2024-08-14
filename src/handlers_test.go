package src_test

import (
	"bytes"
	"encoding/json"
	. "github.com/dating-api/src"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app = fiber.New()
var dbMock = databaseMock{Mock: &mock.Mock{}}
var handlerMock = NewHandler(&dbMock)

func userData() User {
	hashPassword, _ := HashPassword("123456")
	expectUser := User{
		ID:       0,
		Email:    "test@test.com",
		Password: hashPassword,
		Name:     "test",
		Gender:   "male",
		Address:  "jakarta",
	}
	return expectUser
}

func TestHandler_Ping(t *testing.T) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	t.Run("test ping should return pong", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/ping", nil)
		resp, err := app.Test(request)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	})
}

func TestHandler_Register(t *testing.T) {

	app.Get("/register", handlerMock.Register)

	expectUser := userData()

	ID := int64(expectUser.ID)
	dbMock.Mock.On("StoreUser", mock.Anything).Return(&ID, nil)

	registerRequest := RegisterRequest{
		Email:                "test@gmail.com",
		Name:                 "test",
		Gender:               "male",
		Address:              "jakarta",
		Password:             "123456",
		PasswordConfirmation: "123456",
	}

	t.Run("test register should return 201", func(t *testing.T) {

		dbMock.Mock.On("GetUserByEmail", mock.Anything).Return(&expectUser, nil)

		requestJson, _ := json.Marshal(registerRequest)
		request := httptest.NewRequest("GET", "/register", bytes.NewReader(requestJson))
		request.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(request, 10000)

		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	})

	t.Run("test register existing email should error ", func(t *testing.T) {

		expectUser.Email = "test@gmail.com"
		dbMock.Mock.On("GetUserByEmail", expectUser.Email).Return(&expectUser, nil)

		registerRequest := RegisterRequest{
			Email:                "test@gmail.com",
			Name:                 "test",
			Gender:               "male",
			Address:              "jakarta",
			Password:             "123456",
			PasswordConfirmation: "123456",
		}

		requestJson, _ := json.Marshal(registerRequest)
		request := httptest.NewRequest("GET", "/register", bytes.NewReader(requestJson))
		request.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(request, 10000)

		bodyData := make([]byte, resp.ContentLength)
		_, _ = resp.Body.Read(bodyData)

		var respBody map[string]interface{}
		_ = json.Unmarshal(bodyData, &respBody)

		code := int(respBody["code"].(float64))
		message := respBody["message"].(string)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, fiber.StatusBadRequest, code)
		assert.Equal(t, "email is already registered", message)
	})

	t.Run("test register with invalid email format should return 400 invalid param", func(t *testing.T) {

		registerRequest.Email = "invalid_email"

		requestJson, _ := json.Marshal(registerRequest)
		request := httptest.NewRequest("GET", "/register", bytes.NewReader(requestJson))
		request.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(request, 10000)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		bodyData := make([]byte, resp.ContentLength)
		_, _ = resp.Body.Read(bodyData)

		var respBody map[string]interface{}
		_ = json.Unmarshal(bodyData, &respBody)

		assert.Equal(t, fiber.StatusBadRequest, int(respBody["code"].(float64)))
		assert.Equal(t, "invalid param", respBody["message"])
	})

	t.Run("test register with invalid password length should return 400 invalid param", func(t *testing.T) {

		registerRequest.Password = "123"
		registerRequest.PasswordConfirmation = "123"

		requestJson, _ := json.Marshal(registerRequest)
		request := httptest.NewRequest("GET", "/register", bytes.NewReader(requestJson))
		request.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(request, 10000)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		bodyData := make([]byte, resp.ContentLength)
		_, _ = resp.Body.Read(bodyData)

		var respBody map[string]interface{}
		_ = json.Unmarshal(bodyData, &respBody)

		assert.Equal(t, fiber.StatusBadRequest, int(respBody["code"].(float64)))
		assert.Equal(t, "invalid param", respBody["message"])
	})

	t.Run("test register with unmatched password should return 400 invalid param", func(t *testing.T) {

		registerRequest.Password = "123456"
		registerRequest.PasswordConfirmation = "123"

		requestJson, _ := json.Marshal(registerRequest)
		request := httptest.NewRequest("GET", "/register", bytes.NewReader(requestJson))
		request.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(request, 10000)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		bodyData := make([]byte, resp.ContentLength)
		_, _ = resp.Body.Read(bodyData)

		var respBody map[string]interface{}
		_ = json.Unmarshal(bodyData, &respBody)

		assert.Equal(t, fiber.StatusBadRequest, int(respBody["code"].(float64)))
		assert.Equal(t, "invalid param", respBody["message"])
	})

}

func TestHandler_Login(t *testing.T) {
	app.Get("/login", handlerMock.Login)

	hashPassword, _ := HashPassword("123456")
	expectUser := User{
		ID:       0,
		Email:    "test@test.com",
		Password: hashPassword,
		Name:     "test",
		Gender:   "male",
		Address:  "jakarta",
	}

	dbMock.Mock.On("GetUserByEmail", expectUser.Email).Return(&expectUser, nil)

	loginRequest := LoginRequest{
		Email:    "test@test.com",
		Password: "123456",
	}

	t.Run("test login should return 200", func(t *testing.T) {

		requestJson, _ := json.Marshal(loginRequest)
		request := httptest.NewRequest("GET", "/login", bytes.NewReader(requestJson))
		request.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(request)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		bodyData := make([]byte, resp.ContentLength)
		_, _ = resp.Body.Read(bodyData)

		var respBody map[string]interface{}
		_ = json.Unmarshal(bodyData, &respBody)

		code := int(respBody["code"].(float64))
		message := respBody["message"].(string)
		data := respBody["data"].(string)

		assert.Equal(t, fiber.StatusOK, code)
		assert.Equal(t, "login successful", message)
		assert.Contains(t, data, "Bearer")

	})

	t.Run("test login with invalid email should return 400", func(t *testing.T) {

		loginRequest.Email = "invalid_email"

		requestJson, _ := json.Marshal(loginRequest)
		request := httptest.NewRequest("GET", "/login", bytes.NewReader(requestJson))
		request.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(request)

		bodyData := make([]byte, resp.ContentLength)
		_, _ = resp.Body.Read(bodyData)

		var respBody map[string]interface{}
		_ = json.Unmarshal(bodyData, &respBody)

		code := int(respBody["code"].(float64))
		message := respBody["message"].(string)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, fiber.StatusBadRequest, code)
		assert.Equal(t, "invalid param", message)

	})

	t.Run("test login with invalid password length should return 400", func(t *testing.T) {

		loginRequest.Password = "123"

		requestJson, _ := json.Marshal(loginRequest)
		request := httptest.NewRequest("GET", "/login", bytes.NewReader(requestJson))
		request.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(request)

		bodyData := make([]byte, resp.ContentLength)
		_, _ = resp.Body.Read(bodyData)

		var respBody map[string]interface{}
		_ = json.Unmarshal(bodyData, &respBody)

		code := int(respBody["code"].(float64))
		message := respBody["message"].(string)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, fiber.StatusBadRequest, code)
		assert.Equal(t, "invalid param", message)

	})

}

func TestHandler_AccessHome(t *testing.T) {
	app.Get("/home", AuthMiddleware, handlerMock.Home)

	t.Run("test without token should error", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/home", nil)
		request.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(request)

		bodyData := make([]byte, resp.ContentLength)
		_, _ = resp.Body.Read(bodyData)

		var respBody map[string]interface{}
		_ = json.Unmarshal(bodyData, &respBody)

		code := int(respBody["code"].(float64))
		message := respBody["message"].(string)

		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
		assert.Equal(t, fiber.StatusUnauthorized, code)
		assert.Equal(t, "forbidden access", message)

	})

	t.Run("test with token should success", func(t *testing.T) {

		hashPassword, _ := HashPassword("123456")
		user := User{
			ID:       0,
			Email:    "test@test.com",
			Password: hashPassword,
			Name:     "test",
			Gender:   "male",
			Address:  "jakarta",
		}

		token, _ := GenerateToken(&user)

		request := httptest.NewRequest(http.MethodGet, "/home", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", token)
		resp, _ := app.Test(request)

		bodyData := make([]byte, resp.ContentLength)
		_, _ = resp.Body.Read(bodyData)

		var respBody map[string]interface{}
		_ = json.Unmarshal(bodyData, &respBody)

		code := int(respBody["code"].(float64))
		message := respBody["message"].(string)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		assert.Equal(t, fiber.StatusOK, code)
		assert.Equal(t, "success", message)

	})
}
