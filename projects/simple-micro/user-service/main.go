package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/labstack/echo"
	"gopkg.in/resty.v1"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", hello)
	e.GET("/healthcheck", healthcheck)
	e.GET("/user/info", UserInfo)

	// Register order-service with Consul service registry -> This helps services can know together and access
	registerServiceWithConsul()

	// Register user-service with Kong gateway so that the gateway can redirect the client's request to the correct service
	registerKong()

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

// Register user-service with kong admin. This func is registering by code
// If do not want to regster by code, can register via konga dashboard (http://localhost:1337/#!/services)
// This service runs on port 3000
func registerKong() {
	fmt.Println("=======START KONG=======")
	client := resty.New()
	res, _ := client.R().
		SetFormData(map[string]string{
			"name": "user-service",
			"path": "/user-service",
			"url":  "http://10.0.0.16:3000",
		}).Post("http://localhost:8001/services/")

	fmt.Println(res)
	fmt.Println("=======START KONG=======")
}

// Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"code":    http.StatusOK,
		"message": "Welcome to User Service",
	})
}

func healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "Good!")
}

func UserInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"userId":   "123456",
		"fullName": "Ryan Nguyen",
		"avatar":   "https://genknews.genkcdn.vn/2018/8/23/anh-0-1535019031645146400508.jpg",
		"email":    "code4func@gmail.com",
	})
}

func registerServiceWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "user-service"   //replace with service id
	registration.Name = "user-service" //replace with service name
	//address := hostname()              // return "Vuongs-MacBook-Pro.local"
	address := "10.0.0.16"
	registration.Address = address
	if err != nil {
		log.Fatalln(err)
	}
	registration.Port = 3000
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck",
		address, 3000)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"
	consul.Agent().ServiceRegister(registration)
}

func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}

func LookupServiceWithConsul(serviceID string) (string, error) {
	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)
	if err != nil {
		return "", err
	}
	services, err := client.Agent().Services()
	if err != nil {
		return "", err
	}
	srvc := services[serviceID]
	address := srvc.Address
	port := srvc.Port
	return fmt.Sprintf("http://%s:%v", address, port), nil
}
