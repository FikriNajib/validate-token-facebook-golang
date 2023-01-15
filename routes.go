package main

import (
	"github.com/labstack/echo/v4"
	"log"
)

// AddApproutes will add the routes for the application
func AddApproutes() *echo.Echo {

	log.Println("Loadeding Routes...")
	c := echo.New()
	c.GET("/", HandleFacebookLogin)

	log.Println("Routes are Loaded.")
	return c
}
