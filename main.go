package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"myapp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

type ExampleSelect struct {
	Title string
}

func main() {
	r := gin.New()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	r.Static("/wwwroot/css", "./wwwroot/css")
	r.Static("/wwwroot/js", "./wwwroot/js")

	r.GET("/", func(c *gin.Context) {
		// CONNECT TO DATABASE WITH SPECIAL VALUES EXAMPLE
		content, err := ioutil.ReadFile("database.config")
		if err != nil {
			fmt.Println(err)
		}
		dsn := string(content)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println(err)
		}

		/*
			// ADD TO DATABASE EXAMPLE
			user := map[string]interface{}{}
			user["id"] = 12
			user["name"] = "bill"
			user["email"] = "bill@bill.com"
			db.Table("user").Create(&user)
		*/

		// READ FROM DATABASE EXAMPLE
		result := map[string]interface{}{}
		db.Table("user").Take(&result)
		if err != nil {
			fmt.Println(err)
		}

		var user models.User

		user = user.User(
			result["id"].(int32),
			result["name"].(string),
			result["email"].(string),
		)

		// EXAMPLE SELECT OPTIONS
		data := []ExampleSelect{
			{Title: "Bullet 1"},
			{Title: "Bullet 2"},
		}

		// READ FROM DATABASE EXAMPLE MULTIPLE ITEMS
		var results []map[string]interface{}
		db.Table("user").Find(&results)

		fmt.Println(results[0])

		c.HTML(http.StatusOK, "/views/index.tmpl", gin.H{
			"Id":      user.GetId(),
			"Name":    user.GetName(),
			"Email":   user.GetEmail(),
			"Example": data,
		})
	})

	r.Run("localhost:9000")
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
