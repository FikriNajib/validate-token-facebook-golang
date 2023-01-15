package main

import (
	"github.com/huandu/facebook"
	"log"
)

func GetUserInfoFromFacebook(token string) (FacebookUserDetails, error) {
	res, err := facebook.Get("/me", facebook.Params{
		"fields":       "id,name,email",
		"access_token": token,
	})
	if err != nil {
		log.Println(err.Error())
	}
	var result FacebookUserDetails
	res.Decode(&result)

	return result, nil
}
