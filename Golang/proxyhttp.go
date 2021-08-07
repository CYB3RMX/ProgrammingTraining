package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func MakeReq(given_proxy string) {
	// Setting up proxy
	proxyUrl, err := url.Parse(given_proxy)
	if err != nil {
		panic(err)
	}
	httpCli := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	// HTTP Get request
	httpresp, err := httpCli.Get("http://ifconfig.me")
	if err != nil {
		panic(err)
	}

	// Parsing data and return value
	data, err := ioutil.ReadAll(httpresp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Tested Proxy: %s || IP Address: %s\n", proxyUrl, string(data))
}

func main() {
	proxy_list := []string{
		"http://89.135.188.201:9090",
		"http://149.248.50.206:80",
		"http://58.234.116.100:80",
		"http://136.243.211.104:80",
	}
	fmt.Println("[+] Querying your public IP...")
	for _, prox := range proxy_list {
		go MakeReq(prox)
	}
	time.Sleep(10 * time.Second) // Wait for 10 seconds to finish all goroutines
}
