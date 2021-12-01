package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.Get("http://localhost:8080/temp/")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))

	resp_get, err := http.Get("http://localhost:8080/?var1=9&var2=9")
	if err != nil {
		panic(err)
	}

	defer resp_get.Body.Close()

	data_get, err := ioutil.ReadAll(resp_get.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data_get))

	resp_post, err := http.PostForm("http://localhost:8080", url.Values{"var1": {"9"}, "var2": {"9"}})
	if err != nil {
		panic(err)
	}

	defer resp_post.Body.Close()

	data_post, err := ioutil.ReadAll(resp_post.Body)
	if err == nil {
		str := string(data_post)
		println(str)
	}
}
