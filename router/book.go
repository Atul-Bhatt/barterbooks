package router

import (
	"net/http"

	"barterbooks/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(db *sqlx.DB) *gin.Engine {
	r := gin.Default()
	repo := repository.NewBookRepository(db)

	r.GET("/health_check", func(c *gin.Context) {
		c.String(http.StatusOK, "hello from bookbarter!")
	})

	r.GET("/books/", func(c *gin.Context) {
		books, err := repo.GetAll()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"Message": "Error getting books."})
		} else {
			c.JSON(http.StatusOK, books)
		}
		// user := c.Params.ByISBN("isbn")
		// value, ok := db[user]
		// if ok {
		// 	c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		// } else {
		// 	c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		// }
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	// authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	// authorized.POST("admin", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)

	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}

	// 	if c.Bind(&json) == nil {
	// 		db[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })

	return r
}
