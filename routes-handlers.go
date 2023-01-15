package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func HandleFacebookLogin(c echo.Context) error {
	accessToken := "YOUR_ACCESS_TOKEN" //fill this string with your access token
	fbUserDetails, fbUserDetailsError := GetUserInfoFromFacebook(accessToken)

	if fbUserDetailsError != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
	}
	res, err := json.Marshal(fbUserDetails)
	if err != nil {
		log.Println(err.Error())
	}
	return c.String(http.StatusOK, string(res))
}
