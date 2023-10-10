package main

import (
	"context"
	"github.com/ac-kurniawan/gosong"
	"github.com/labstack/echo/v4"
	"net/http"
)

// HelloWorldService Define Hello World service
type HelloWorldService struct {
}

func (h HelloWorldService) HelloWorld(ctx context.Context) string {
	return "hello world!"
}

// Controller Defining Controller
type Controller struct {
	Echo    *echo.Echo
	Service *HelloWorldService `import:"HelloWorldService"`
}

func (c Controller) Start() {
	c.Echo.GET("/hello-world", func(ctx echo.Context) error {
		result := c.Service.HelloWorld(ctx.Request().Context())
		return ctx.String(http.StatusOK, result)
	})

	c.Echo.Logger.Fatal(c.Echo.Start(":1231"))
}

func main() {
	apps := gosong.Application{
		Name: "EchoApps",
	}

	e := echo.New()
	controller := Controller{
		Echo: e,
	}
	apps.AddSingleton("HelloWorldService", &HelloWorldService{})
	apps.AddController("Controller", &controller)

	apps.AddEntry(controller.Start)

	gosong.RunApplications(apps)
}
