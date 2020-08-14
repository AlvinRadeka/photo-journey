package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AlvinRadeka/photo-journey/internal/renderers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	renderers.Register(e)

	e.GET("/sorry", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home", nil)
	})

	e.File("/favicon.ico", "files/images/sadsmiley.ico")
	e.File("/background", "files/images/pine-tree.jpg")
	e.File("/styles", "files/css/home.css")

	go startApp(e)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("shutting down the server...")
	if err := e.Shutdown(ctx); err != nil {
		fmt.Println("server shut down unexpectedly.")
		e.Logger.Fatal(err)
	}
}

func startApp(e *echo.Echo) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":" + port))
}
