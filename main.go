package main

//import "net/http"
import "flag"
import "fmt"
import (
	"strings"
	"net/http"
	"os"
	"io/ioutil"
)

var ip string
var method string
var headers string
var data string

func main()  {
	flag.StringVar(&ip,"host", "127.0.0.1", "set the hostname of ip (default: 127.0.0.1)")
	flag.StringVar(&method, "method", "GET", "set the HTTP method (default: GET)")
	flag.StringVar(&headers, "headers", "", "set the header variables, comma-separated")
	flag.StringVar(&data, "data", "", "set the header variables, comma-separated")
	flag.Parse()
	client := http.Client{}
	req, err := http.NewRequest(method, ip, nil)
	if err != nil {
		fmt.Printf("Can't create request %v for %v a some reason.\n", method, ip)
		os.Exit(0)
	}
	if headers != "" {
		result := strings.Split(headers, ",")
		for _, pair := range result {
			items := strings.Split(pair, "=")
			key, value := items[0], items[1]
			req.Header.Add(key, value)
		}
	}
	fmt.Println("IP has value ", ip)
	fmt.Println("Method has value ", method)
	fmt.Println("Headers has value ", headers)
	fmt.Println("Data has value ", data)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Got error %v\n", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Printf(bodyString)
}

