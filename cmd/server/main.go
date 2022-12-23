package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/jocmp/gaffe"
	"github.com/jocmp/gaffe/cmd"
)

type Config struct {
	CSVPath string
	WebPort string
}

func main() {
	config := fetchConfig()
	indexHandler := buildIndexHandler(config)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", indexHandler)

	router.Run(fmt.Sprintf(":%s", config.WebPort))
}

func buildIndexHandler(config Config) func(c *gin.Context) {
	return func(c *gin.Context) {
		benchmarks, err := gaffe.FetchBenchmarks(config.CSVPath)
		view := cmd.PresentBenchmarks(benchmarks)

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
