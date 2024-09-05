package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//go:embed template/*
var templateFs embed.FS

func main() {

	r := gin.Default()

	t, err := template.ParseFS(templateFs, "template/*")
	if err != nil {
		panic(err)
	}

	r.SetHTMLTemplate(t)

	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	err = r.Run(":8082")
	if err != nil {
		panic(err)
	}
}
