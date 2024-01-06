package handlers

import (
	"encoding/json"
	"github.com/Hoaper/golang_university/app/models"
	"github.com/Hoaper/golang_university/app/services"
	"github.com/Hoaper/golang_university/app/utils"
	"net/http"
)

type AuthHandler struct {
	UserService services.UserService
}

func NewAuthHandler(userService services.UserService) *AuthHandler {
	return &AuthHandler{UserService: userService}
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = h.UserService.CreateUser(&user)
	if err != nil {
		println(err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.UserService.AuthenticateUser(loginRequest.Login, loginRequest.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := utils.GenerateToken(user.ID.Hex(), user.Login, user.Role)

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Login successful", "token": token})
}

func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Logout successful"})
}
