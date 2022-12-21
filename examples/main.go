package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/thealish/gocto"
)

func main() {
	octoClient := gocto.New(gocto.Options{
		OctoShopID:    123,
		OctoSecret:    "octo secret",
		ReturnURL:     "return url",
		NotifyURL:     "notify url",
		BindNotifyURL: "bind url",
		HTTPClient:    &http.Client{},
	})

	res, err := octoClient.BindCard(context.Background(), gocto.BindCardRequest{
		ShopTransactionId: "",
		Pan:               "",
		Exp:               "",
		Phone:             "",
		CardHolderName:    "",
		Cvc2:              "",
		Method:            "",
		ReturnUrl:         "",
		NotifyUrl:         "",
		BindNotifyUrl:     "",
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", res)
}
