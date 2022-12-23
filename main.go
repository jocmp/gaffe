package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jocmp/gaffe/pkg/gaffe"
)

type Config struct {
	CSVPath string
	WebPort string
}

//go:embed templates/*
var f embed.FS

func main() {
	config := fetchConfig()
	indexHandler := buildIndexHandler(config)

	router := gin.Default()
	initTemplates(router)

	router.GET("/", indexHandler)

	router.Run(fmt.Sprintf(":%s", config.WebPort))
}

func initTemplates(router *gin.Engine) {
	if gin.Mode() == gin.ReleaseMode {
		t := template.Must(template.New("").ParseFS(f, "templates/*.html"))
		router.SetHTMLTemplate(t)
	} else {
		router.LoadHTMLGlob("templates/*")
	}
}

func buildIndexHandler(config Config) func(c *gin.Context) {
	return func(c *gin.Context) {
		benchmarks, err := gaffe.FetchBenchmarks(config.CSVPath)
		view := gaffe.PresentBenchmarks(benchmarks)

		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"view": view,
		})
	}
}

func fetchConfig() Config {
	config := Config{
		CSVPath: os.Getenv("CSV_PATH"),
		WebPort: os.Getenv("WEB_PORT"),
	}

	return config
}
