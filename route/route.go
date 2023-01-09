package route

import (
	"github.com/labstack/echo/v4"
)

func Routing(g *echo.Group) {
	g.GET("/command", GetCommands)
	g.POST("/command", CreateCommand)
	g.PUT("/command/:uuid", UpdateCommandByUUID)
	g.DELETE("/command/:uuid", DeleteCommandByUUID)
	g.GET("/request", GetRequests)
	g.POST("/request", CreateRequest)
	g.PUT("/request/:uuid", UpdateRequestByUUID)
	g.DELETE("/request/:uuid", DeleteRequestByUUID)
}
