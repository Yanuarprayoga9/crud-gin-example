package handler

import (
	"day1/model/web"
	"day1/service"
	"day1/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var request web.UserCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(c, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(request)
	if err != nil {
		utils.ErrorResponse(c, "Failed to create user", http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(c, "User created successfully", "user", user, http.StatusCreated)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		utils.ErrorResponse(c, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(c, "Users retrieved successfully", "users", users, http.StatusOK)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, "User not found", http.StatusNotFound)
		return
	}

	utils.SuccessResponse(c, "User found", "user", user, http.StatusOK)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var request web.UserUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(c, "Invalid input", http.StatusBadRequest)
		return
	}

	request.Id = id

	updatedUser, err := h.service.UpdateUser(request)
	if err != nil {
		utils.ErrorResponse(c, "Failed to update user", http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(c, "User updated successfully", "user", updatedUser, http.StatusOK)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteUser(uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			utils.ErrorResponse(c, "User not found", http.StatusNotFound)
		} else {
			utils.ErrorResponse(c, "Failed to delete user", http.StatusInternalServerError)
		}
		return
	}

	utils.SuccessResponse(c, "User deleted successfully", "user", nil, http.StatusOK)
}
