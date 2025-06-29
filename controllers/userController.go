package controllers

import (
	"hola/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{ID: 1, Name: "William", Email: "Willuuuuuuuu@mail.com", BloodType: "B-"},
	{ID: 2, Name: "Sori", Email: "Soriw6@mail.com"},
}

var nextID = 3

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = nextID
	nextID++
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
}
