package controllers

import (
	"log"

	"github.com/labstack/echo"
)

// DashboardHandler ...
func DashboardHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("DashboardHandler")

		data := map[string]interface{}{}
		data["userData"] = nil
		data["media"] = nil
		data["media_count"] = 0
		data["hasMedia"] = false
		return RenderTemplate(c, "dashboard/index", data)
	}
}
