package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
)

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		Jar: jar,
	}
	for i := 0; i < 2; i++ {
		res, err := client.Get("http://localhost:18888/cookie")
		if err != nil {
			log.Fatal(err)
		}

		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(dump))
	}
}