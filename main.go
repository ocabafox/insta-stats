package main

import (
	"log"
	"net/http"

	"github.com/a-fis/insta-stats/app/controllers"
	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// init ..
func init() {
	log.SetFlags(log.Lshortfile)

}
func main() {
	// Make an engine
	engine := echo.New()

	engine.Renderer = NewTemplate("app/views/")

	// use actual files
	engine.Static("/assets", "app/data/assets")
	//engine.Use(middleware.Recover())

	engine.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf_token",
		CookiePath:  "/",
	}))
	//gzip
	engine.Use(middleware.Gzip())
	//session
	store := session.NewCookieStore([]byte("secret-key-kaskdfkjhkedknjjkelaasdfjkajkeasd"))
	store.MaxAge(86400)
	engine.Use(session.Sessions("GSESSION", store))

	engine.GET("/favicon.ico", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/assets/favicon.ico")
	})

	engine.GET("/", controllers.DashboardHandler())
	engine.POST("/", controllers.DashboardPostHandler())

	// cool custom logging
	engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${method} | ${status} | ${uri} -> ${latency_human}` + "\n",
	}))

	engine.Start(":3000")
	return

}
