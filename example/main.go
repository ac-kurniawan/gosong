package main

import (
	"fmt"

	"github.com/ac-kurniawan/gosong"
	"github.com/labstack/echo/v4"
)

// Define service interface and struct
type IService interface {
	GetWallets()
	GetWallet() string
}

type TestService struct {
	Repo IRepository `import:"testRepo"`
}

func (s TestService) GetWallets() {
	fmt.Println("test Service")
	s.Repo.GetWallets()
}
func (s TestService) GetWallet() string {
	return s.Repo.GetWallet()
}

// Define repository interface and struct
type IRepository interface {
	GetWallets()
	GetWallet() string
}

type TestRepoistory struct {
}

func (s TestRepoistory) GetWallet() string {
	return "WALLET Repo"
}
func (r TestRepoistory) GetWallets() {
	fmt.Println("test Repository")

}

// Define controller interface and struct
type Test2Controller struct {
	Service IService   `import:"testService"`
	Echo    *echo.Echo `import:"echoContext"`
}

func (ct Test2Controller) Run() {
	ct.Echo.GET("/test", func(c echo.Context) error {
		result := ct.Service.GetWallet()
		return c.String(200, result)
	})
}

func main() {
	asd := gosong.Application{
		Name: "asd",
	}
	asd.AddComponents("testRepo", &TestRepoistory{})
	asd.AddProviders("testService", &TestService{})
	ec := echo.New()
	asd.AddComponents("echoContext", ec)

	controller2 := Test2Controller{}
	asd.AddControllers("test2Controller", &controller2)
	asd.AddEntries(controller2.Run)

	ec.HideBanner = true

	apps := []gosong.Application{asd}

	gosong.RunApplications(apps...)

	if err := ec.Start(":3222"); err != nil {
		fmt.Printf("[HTTP Server] - %s", err.Error())
	}
}
