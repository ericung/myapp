package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

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
		c.HTML(http.StatusOK, "/views/index.tmpl", gin.H{
			"Foo": time.Now(),
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
