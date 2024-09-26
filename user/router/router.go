package router

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"user/model"
	"user/repository"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var repo *repository.UserRepository

const HashingCost = 14
const userRole = "user"

func SetupRouter(db *sqlx.DB) *gin.Engine {
	r := gin.Default()
	repo = repository.NewUserRepository(db)

	r.GET("/health_check", healthCheck)
	r.GET("/users", getUsers)
	r.POST("/users/", createUser)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.POST("login", login)
	r.POST("signup", signUp)
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

func signUp(c *gin.Context) {
	var suPayload model.SignUpRequest
	var user model.User

	if err := c.ShouldBindJSON(&suPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userExists, userExistsErr := repo.UsernameExists(suPayload.Username)
	if userExistsErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": userExistsErr.Error()})
		return
	}

	if userExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	hashBytes, hashError := bcrypt.GenerateFromPassword([]byte(suPayload.Password), HashingCost)
	if hashError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": hashError.Error()})
		return
	}

	user.Username = suPayload.Username
	user.FirstName = suPayload.FirstName
	user.LastName = suPayload.LastName
	user.Password = string(hashBytes)
	user.Role = userRole

	if err := repo.Create(user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, user)
}

func login(c *gin.Context) {
	var creds model.LoginRequest

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashBytes, hashError := bcrypt.GenerateFromPassword([]byte(creds.Password), HashingCost)
	if hashError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": hashError.Error()})
		return
	}

	if authError := repo.CheckPassword(creds.Username, string(hashBytes)); authError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": authError.Error()})
		return
	}

	user, userErr := repo.GetUserByUsername(creds.Username)
	if userErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": userErr.Error()})
		return
	}

	token, tokenErr := getJWTToken(user)
	if tokenErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": tokenErr.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"token": token})
}

func getJWTToken(user model.User) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["firstname"] = user.FirstName
	claims["lastname"] = user.LastName
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// generate encoded token and send it as response
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return t, err
	}
	return t, nil
}
