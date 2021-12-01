package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ip, err := getIP(r)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("No valid ip"))
	}
	w.WriteHeader(200)
	fmt.Println("::Client address   :" + ip)
	fmt.Println("::Request command  :" + r.Method)
	fmt.Println("::Request path     :" + r.URL.RequestURI())
	fmt.Println("::Request version  :" + r.Proto)

	switch r.Method {
	case "GET":
		query := r.URL.Query()
		raw_url := r.URL.RequestURI()
		if strings.Contains(raw_url, "?") {
			var1, present := query["var1"]
			if !present || len(var1) == 0 {
				fmt.Println("wrong expression")
			}
			var2, present := query["var2"]
			if !present || len(var2) == 0 {
				fmt.Println("Wrong expression")
			}
			var1_str := strings.Join(var1, "")
			var2_str := strings.Join(var2, "")
			var1_int, _ := strconv.Atoi(var1_str)
			var2_int, _ := strconv.Atoi(var2_str)
			result := var1_int * var2_int
			get_response := fmt.Sprintf("## GET request for calculation => %d x %d = %d", var1_int, var2_int, result)
			fmt.Fprint(w, get_response)
			fmt.Println(get_response)
		} else {
			get_response := fmt.Sprintf("## GET request for directory => %s", r.URL.RequestURI())
			fmt.Fprint(w, get_response)
			fmt.Println(get_response)
		}
		w.WriteHeader(200)
		fmt.Println("")
	case "POST":
		// query := r.URL.Query()
		// raw_url := r.URL.RequestURI()

		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		fmt.Println("## POST request data => ", string(b))
		vars := strings.Split(string(b), "&")
		var_ints := []int{}
		for _, var_tuple := range vars {
			real_vars := strings.Split(var_tuple, "=")
			var_int, _ := strconv.Atoi(real_vars[1])
			var_ints = append(var_ints, var_int)
		}
		result := var_ints[0] * var_ints[1]
		get_response := fmt.Sprintf("## POST request for calculation => %d x %d = %d", var_ints[0], var_ints[1], result)
		fmt.Fprint(w, get_response)
		fmt.Println(get_response)
		w.WriteHeader(200)
		fmt.Println("")
	}

}

func getIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("No valid ip found")
}

// read all header string
func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)

}

// read a Accep-Encoding in header string
func headersEncoding(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Accept-Encoding"]
	fmt.Fprintln(w, h)
}

// read a Accep-Encoding in header string by get method
func headersGet(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, h)
}

// read a body string
func bodyString(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	fmt.Printf("## HTTP server started at http://%s\n", server.Addr)
	http.HandleFunc("/", handler)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/headers/encoding", headersEncoding)
	http.HandleFunc("/headers/get/encoding", headersGet)
	http.HandleFunc("/body", bodyString)

	server.ListenAndServe()
}
