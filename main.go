package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Username  string     `json:"username" gorm:"column:username;"`
	Password  string     `json:"password" gorm:"column:pass;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (User) TableName() string { return "users" }
func main() {
	dsn := "root:secret@tcp(host.docker.internal:3306)/viki?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected to MySQL:", db)
	router := gin.Default()
	user := router.Group("/user")
	{
		user.GET("/check", check())
		// user.POST("/signup", signup(db))              // create item
		user.POST("/login", login(db)) // create item
		// user.GET("/items", getListOfItems(db))        // list items
		// user.GET("/items/:id", readItemById(db))      // get an item by ID
		// user.PUT("/items/:id", editItemById(db))      // edit an item by ID
		// user.DELETE("/items/:id", deleteItemById(db)) // delete an item by ID
	}
	router.Run()
}

func check() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello checker"})
	}
}

// func an() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"data": "Hello an"})
// 	}
// }

// func signup(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var dataUser User

// 		if err := c.ShouldBind(&dataUser); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		// preprocess title - trim all spaces
// 		dataUser.Username = strings.TrimSpace(dataUser.Username)
// 		dataUser.Password = strings.TrimSpace(dataUser.Password)

// 		if dataUser.Username == "" || dataUser.Password == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Username or password is empty"})
// 			return
// 		}

// 		if err := db.Create(&dataUser).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"data": dataUser.Id})
// 	}
// }

func login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataUser User
		var compareUser User

		if err := c.ShouldBind(&dataUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// preprocess title - trim all spaces
		dataUser.Username = strings.TrimSpace(dataUser.Username)
		dataUser.Password = strings.TrimSpace(dataUser.Password)

		if dataUser.Username == "" || dataUser.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username or password is empty"})
			return
		}

		if err := db.Where("username = ?", dataUser.Username).First(&compareUser).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if dataUser.Password != compareUser.Password {
			c.JSON(http.StatusOK, gin.H{"validate": false})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": compareUser.Id})
		}

	}
}

// func readItemById(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var dataUser User

// 		id, err := strconv.Atoi(c.Param("id"))

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Where("id = ?", id).First(&dataUser).Error; err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"data": dataUser})
// 	}
// }

// func getListOfItems(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		type DataPaging struct {
// 			Page  int   `json:"page" form:"page"`
// 			Limit int   `json:"limit" form:"limit"`
// 			Total int64 `json:"total" form:"-"`
// 		}

// 		var paging DataPaging

// 		if err := c.ShouldBind(&paging); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if paging.Page <= 0 {
// 			paging.Page = 1
// 		}

// 		if paging.Limit <= 0 {
// 			paging.Limit = 10
// 		}

// 		offset := (paging.Page - 1) * paging.Limit

// 		var result []User

// 		if err := db.Table(User{}.TableName()).
// 			Count(&paging.Total).
// 			Offset(offset).
// 			Order("id desc").
// 			Find(&result).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"data": result})
// 	}
// }

// func editItemById(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		var dataUser User

// 		if err := c.ShouldBind(&dataUser); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Where("id = ?", id).Updates(&dataUser).Error; err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"data": true})
// 	}
// }

// func deleteItemById(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Table(User{}.TableName()).
// 			Where("id = ?", id).
// 			Delete(nil).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"data": true})
// 	}
// }
