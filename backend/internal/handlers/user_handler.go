package handlers

import (
	"net/http"
	"strconv"

	"github.com/Geovannisouza23/go-user-api/internal/repositories"
	"github.com/Geovannisouza23/go-user-api/internal/services"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Me(c *gin.Context) {
	userID, _ := c.Get("user_id")

	service := services.UserService{
		UserRepo: &repositories.UserRepository{},
	}

	user, err := service.GetByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}

func ListUsers(c *gin.Context) {
	service := services.UserService{
		UserRepo: &repositories.UserRepository{},
	}

	users, err := service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []UserResponse
	for _, user := range users {
		response = append(response, UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	c.JSON(http.StatusOK, response)
}

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	service := services.UserService{
		UserRepo: &repositories.UserRepository{},
	}

	user, err := service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}

type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := services.UserService{
		UserRepo: &repositories.UserRepository{},
	}

	user, err := service.Update(uint(id), req.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	service := services.UserService{
		UserRepo: &repositories.UserRepository{},
	}

	if err := service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
