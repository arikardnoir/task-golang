package tests

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"task-golang/controllers"
	"task-golang/repositories"
	"task-golang/services"
	"testing"
)

func TestListUsers(t *testing.T) {
	// Setting Echo
	e := echo.New()
	SetupTestDB()

	userRepo := &repositories.UserRepository{}
	userService := &services.UserService{UserRepo: userRepo}
	userController := &controllers.UserController{UserService: userService}

	// Generate JWT token for a auth
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzM2NjY4NDcsInVzZXJJRCI6IjJmNGIxZDNjLTA5N2EtNDIwMS1iODhlLWFlNWY4OWMzMjQwMyJ9.MgAUIMxedXpg_IcsiFY1nY0030lkvsmDhk2DGLXARSk"

	// create simulated request
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderAuthorization, token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// execute handler
	if assert.NoError(t, userController.ListUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var users []map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &users)
		assert.NotEmpty(t, users) // Verifica se a lista de usuários foi retornada
	}
}
