package router

import (
	"database/sql"
	"fmt"
	company "gin/controllers"
	"gin/controllers/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var db = make(map[string]string)
type Product struct {
	Time *DeleteTime
}
type DeleteTime sql.NullInt64

func (this *DeleteTime) Scan(value interface{}) error {
	fmt.Println("DeleteTime Scan ")
	fmt.Println(value)
	this.Int64, this.Valid = value.(int64)
	return nil
}
func (this DeleteTime) MarshalJSON() (int64, error) {
	fmt.Println("MarshalJSON start")
	if !this.Valid {
		fmt.Println("MarshalJSON ni.Valid ")
		return 0, nil
	}
	fmt.Println("MarshalJSON")
	return this.Int64, nil
}

func (this *DeleteTime) UnmarshalJSON(data []byte) error {
	fmt.Println("UnmarshalJSON")
	if data == nil || string(data) == `null` {
		this.Valid = false
		return nil
	}
	val, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		this.Valid = false
		return err
	}
	this.Int64 = val
	this.Valid = true
	return nil
}

func SetupRouter(engine *gin.Engine) *gin.Engine {




	engine.GET("/user", user.Index)
	engine.GET("/company", company.Index)
	// Ping test
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong-ssss")
	})

	// Get user value
	engine.GET("/user/:name", func(c *gin.Context) {
		//user := c.Params.ByName("name")
		//value, ok := db[user]
		//if ok {
		//	c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		//} else {
		//	c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		//}
		c.JSON(http.StatusOK, Product{Time: &DeleteTime{Int64: 11111, Valid: true}})

	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := engine.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return engine
}
