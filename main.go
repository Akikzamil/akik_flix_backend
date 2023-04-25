package main

import (
	"akikflix/database"
	"akikflix/middleware"
	"akikflix/route"
	"fmt"
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func main() {
	finish := make(chan bool)
	database.InitDatabase()
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // this is the default limit of 4MB
	})
	app.Static("/", "files")
	middleware.InitMiddleware(app)
	route.InitRoutes(app)
	go func() {
		app.Listen(":8000")
	}()

	go func() {
		runFileServer()
	}()

	<-finish
}

func runFileServer() {
	// configure the songs directory name and port
	const songsDir = "files"
	const port = 8080

	// add a handler for the song files
	http.Handle("/", addHeaders(http.FileServer(http.Dir(songsDir))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", songsDir, port);

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
