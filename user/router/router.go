package router

import (
	"net/http"
	"strconv"

	"user/model"
	"user/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var repo *repository.UserRepository

func SetupRouter(db *sqlx.DB) *gin.Engine {
	r := gin.Default()
	repo = repository.NewUserRepository(db)

	r.GET("/health_check", healthCheck)
	r.GET("/users", getUsers)
	r.POST("/users/", createUser)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)
	return r
}

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "hello from bookbarter!")
}

// get all users
func getUsers(c *gin.Context) {
	users, err := repo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Message": "Error getting users."})
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// create a user
func createUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repo.Create(user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, user)
}

// get user by ID
func getUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := repo.GetUser(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Message": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func updateUser(c *gin.Context) {
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := repo.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = repo.UpdateUser(user, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := repo.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	repo.DeleteUser(id)
	c.JSON(http.StatusNoContent, gin.H{})
}
