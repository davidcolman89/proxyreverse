package test

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main()  {
	resp, err := http.Get("http://localhost:9999/")

	if err != nil {
		fmt.Println("Error:  ",err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("OK")
	fmt.Println(string(body))

}
