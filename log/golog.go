package main

import (
	"log"
	"net/http"
	"os"
)

//go原生的log使用

func SetupLogger() {
	logFileLocation, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}


func SimpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s:%s", url, err.Error())
	} else {
		log.Printf("Status Code for:%s:%d", url, resp.StatusCode)
	}
}

func main() {
	//原生log的使用
	SetupLogger()
	SimpleHttpGet("https://www.google.com")
	SimpleHttpGet("https://www.baidu.com")

}