package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"task-golang/controllers"
	"task-golang/repositories"
	"task-golang/routes"
	"task-golang/services"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	// Setting Echo
	e := echo.New()
	SetupTestDB()
	userRepo := &repositories.UserRepository{}
	userService := &services.UserService{UserRepo: userRepo}
	userController := &controllers.UserController{UserService: userService}

	// Set Router
	routes.SetupRoutes(e)

	// Request Body
	body := map[string]interface{}{
		"name":             "Test User",
		"email":            "test1@example.com",
		"password":         "password123",
		"confirm_password": "password123",
		"cep":              "88804060",
	}
	jsonBody, _ := json.Marshal(body)

	print("Body: ", string(jsonBody))
	// Created simulated request
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Executar handler
	err := userController.Register(c)

	// Debugging
	if err != nil {
		fmt.Println("Handler Error:", err)
	}
	fmt.Println("Status Code:", rec.Code)
	fmt.Println("Response Headers:", rec.Header())
	fmt.Println("Response Body:", rec.Body.String())

	// Verificar resultados esperados
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, "user registered successfully", response["message"])
	}
}

func TestLoginUser(t *testing.T) {
	// Setting Echo
	e := echo.New()
	SetupTestDB()

	userRepo := &repositories.UserRepository{}
	authService := &services.AuthService{UserRepo: userRepo}
	authController := &controllers.AuthController{AuthService: authService}

	// Request Body
	body := map[string]interface{}{
		"email":    "test@example.com",
		"password": "password123",
	}
	jsonBody, _ := json.Marshal(body)

	// Created simulated request
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Execute handler
	if assert.NoError(t, authController.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]string
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NotEmpty(t, response["token"]) // Verifica se o token JWT foi gerado
	}
}
