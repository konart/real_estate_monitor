package main

import (
	"fmt"
	//"io/ioutil"
	//"net/http"
	//"net"
	"log"
	"net/url"
	"strconv"
)

type flat struct {
	id               int
	sources          []*source
	address          string
	rooms            int
	height           float64
	bathroomCount    int
	separateBathroom int
}

func createFlat(d map[string]string) *flat {
	f := new(flat)
	f.address = d["address"]
	rooms, err := strconv.Atoi(d["rooms"])
	if err != nil {
		fmt.Println(err)
	}
	f.rooms = rooms

	height, err := strconv.ParseFloat(d["height"], 64)
	if err != nil {
		fmt.Println("FUCK2")
	}
	f.height = height

	bathNum, err := strconv.Atoi(d["bathrooms"])
	if err != nil {
		fmt.Println("FUCK3")
	}
	f.bathroomCount = bathNum

	sepBathNum, err := strconv.Atoi(d["reparate_bathrooms"])
	if err != nil {
		fmt.Println("FUCK4")
	}
	f.separateBathroom = sepBathNum
	f.addSource(d["source"])
	return f
}

func (f *flat) addSource(s string) {
	src := createSource(s)
	if len(f.sources) > 0 {
		f.sources = append(f.sources, src)
	} else {
		f.sources = []*source{src}
	}
}

//update fetches the latest info about the flat
func (f *flat) update() {
	//get html by id
	//parse
	//...and update!
}

type source struct {
	name string
	url  url.URL
}

//createSource creates new source from the provided url
func createSource(u string) *source {

	url, err := url.Parse(u)
	if err != nil {
		log.Fatal(err)
	}

	//resp, _ := http.Get(u)
	//
	//bytes, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println("HTML:\n\n", string(bytes))

	//resp.Body.Close()
	s := source{url.Host, *url}
	return &s
}

func main() {
}
