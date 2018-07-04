package controllers

// POST: /login

import (
	"log"

	"github.com/a-fis/insta-stats/app/helpers"
	"github.com/labstack/echo"
)

// DashboardPostHandler ...
func DashboardPostHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("DashboardPostHandler")

		userIG := c.Request().PostFormValue("search")

		if userIG == "" {
			return Redirect(c, "/")
		}

		IGJsonData, _, err := helpers.GetIGJsonData("https://www.instagram.com/"+userIG, "GoogleBot")

		log.Print(err)
		log.Print("@@@@@@@@@@@@@@")
		log.Print(IGJsonData.EntryData.ProfilePage[0].Graphql.User.FullName)

		totalLike := 0
		totalComment := 0
		for _, media := range IGJsonData.EntryData.ProfilePage[0].Graphql.User.EdgeOwnerToTimelineMedia.Edges {
			totalLike += media.Node.EdgeLikedBy.Count
			totalComment += media.Node.EdgeMediaToComment.Count
		}

		data := map[string]interface{}{}
		data["userData"] = IGJsonData.EntryData.ProfilePage[0].Graphql.User
		data["media"] = IGJsonData.EntryData.ProfilePage[0].Graphql.User.EdgeOwnerToTimelineMedia.Edges
		data["media_count"] = IGJsonData.EntryData.ProfilePage[0].Graphql.User.EdgeOwnerToTimelineMedia.Count
		if IGJsonData.EntryData.ProfilePage[0].Graphql.User.EdgeOwnerToTimelineMedia.Count > 0 {
			data["hasMedia"] = true
		} else {
			data["hasMedia"] = false
		}
		data["totalLike"] = totalLike
		data["totalComment"] = totalComment

		return RenderTemplate(c, "dashboard/index", data)
	}
}
