package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	urlLoki       string
	containerName string
	serverName    string
)

type LokiStruct struct {
}

func init() {
	flag.StringVar(&urlLoki, "urlLoki", "10.20.28.11", "path to Loki")
	flag.StringVar(&containerName, "containerName", "reverse-archive", "container name")
	flag.StringVar(&serverName, "serverName", "test", "server name")
}

func main() {
	//....
	flag.Parse()
	url := "http://10.20.28.11:3100/loki/api/v1/push"
	method := "POST"

	payload := strings.NewReader(`{"1":"2"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
