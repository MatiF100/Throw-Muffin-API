package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/MatiF100/Throw-Muffin-API/ent"
	"github.com/gin-gonic/gin"
	"gopkg.in/fsnotify.v1"

	"github.com/gofor-little/env"

	_ "github.com/lib/pq"
)

type App struct {
	client     *ent.Client
	local_mode bool
}

func main() {
	app := App{}
	setup_env(&app)
	setup_db(&app)
	defer app.client.Close()

	run_server()
}

func setup_env(app *App) {
	if err := env.Load(".env"); err != nil {
		log.Println("Failed to locate .env file. Using system variables and/or defaults")
		app.local_mode = true
	}
}

func setup_db(app *App) {
	dbString := env.Get("dbString", fmt.Sprintf("host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"))
	client, err := ent.Open("postgres", dbString)
	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Printf("failed creating schema resources: %v", err)
	}

	app.client = client

}

func run_server() {
	old_main()
}

func old_main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Hello from Go and Gin running on Azure App Service",
			"link":  "/json",
		})
	})

	router.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"foo": "bar",
		})
	})

	router.Static("/public", "./public")

	// creates a new file watcher for App_offline.htm
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	// watch for App_offline.htm and exit the program if present
	// This allows continuous deployment on App Service as the .exe will not be
	// terminated otherwise
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if strings.HasSuffix(event.Name, "app_offline.htm") {
					fmt.Println("Exiting due to app_offline.htm being present")
					os.Exit(0)
				}
			}
		}
	}()

	// get the current working directory and watch it
	currentDir, err := os.Getwd()
	if err := watcher.Add(currentDir); err != nil {
		fmt.Println("ERROR", err)
	}

	// Azure App Service sets the port as an Environment Variable
	// This can be random, so needs to be loaded at startup
	port := os.Getenv("HTTP_PLATFORM_PORT")

	// default back to 8080 for local dev
	if port == "" {
		port = "8080"
	}

	router.Run("0.0.0.0:" + port)
}
