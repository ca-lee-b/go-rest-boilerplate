package handlers

import (
	"net/http"
	"time"

	"github.com/ca-lee-b/go-rest-boilerplate/internal/repository"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	UserRepository    *repository.UserRepository
	SessionRepository *repository.SessionRepository
}

func newAuthHandler(userRepo *repository.UserRepository, sessionRepo *repository.SessionRepository) *AuthHandler {
	return &AuthHandler{
		UserRepository:    userRepo,
		SessionRepository: sessionRepo,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	type loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var data loginData
	err := c.Bind(&data)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	if data.Email == "" {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if data.Password == "" {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	//Check if user exists
	user := h.UserRepository.GetUserByEmail(data.Email)
	if user == nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	//Compare password
	comparePassword := compare(user.Password, data.Password)
	if !comparePassword {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	//Create session
	session := h.SessionRepository.Create()

	cookie := new(http.Cookie)
	cookie.Name = "sid"
	cookie.Value = session.Id
	cookie.Expires = session.Expiry
	cookie.MaxAge = int(session.Expiry.Unix() - time.Now().Unix())
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, session)
}

func (h *AuthHandler) Register(c echo.Context) error {
	type registerData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var data registerData
	err := c.Bind(&data)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if data.Email == "" {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if data.Password == "" {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if data.Username == "" {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	hashedPassword, err := hashPassword(data.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if err := h.UserRepository.Create(data.Username, data.Email, hashedPassword); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.String(http.StatusCreated, "Success")
}

// Helper functions
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	return string(bytes), err
}

func compare(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
