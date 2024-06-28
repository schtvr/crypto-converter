package main

import (
	"fmt"
	"net/http"
)

func main() {
	c := http.Client{}
	u := "https://api.coinbase.com/v2/exchange-rates?currency=USD"

	resp, err := c.Get(u)
	if err != nil {
		fmt.Println(err)
	}
	body := resp.Body

	fmt.Printf(body.Close().Error())

}
