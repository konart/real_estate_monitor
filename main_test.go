package main

import (
	"testing"
	"log"
	"net/url"
)

func TestCreateFlat(t *testing.T) {
	address := "Сходненская 6, к.1, кв.24"
	source := "https://www.cian.ru/sale/flat/152890949/"
	rooms := 2
	height := 2.75
	bathrooms := 1
	separateBathrooms := 1

	testFlatMap := make(map[string]string)

	testFlatMap["address"] = address
	testFlatMap["source"] = source
	testFlatMap["rooms"] = "2"
	testFlatMap["height"] = "2.75"
	testFlatMap["bathrooms"] = "1"
	testFlatMap["reparate_bathrooms"] = "1"

	testFlat := createFlat(testFlatMap)
	if testFlat.address != address {
		log.Fatal("Flat did not recive the right address")
		t.FailNow()
	}

	if u, _ := url.Parse(source); testFlat.sources[0].url != *u {
		log.Fatal("Flat did not recive the right source")

		t.FailNow()
	}

	if testFlat.rooms != rooms {
		log.Fatal("Flat did not recive the right number of rooms")
		t.FailNow()
	}

	if testFlat.height != height {
		log.Fatal("Flat did not recive the right height")
		t.FailNow()
	}

	if testFlat.bathroomCount != bathrooms {
		log.Fatal("Flat did not recive the right number of bathrooms")
		t.FailNow()
	}

	if testFlat.separateBathroom != separateBathrooms {
		log.Fatal("Flat did not recive the right number of separate bathrooms")
		t.FailNow()
	}
}

func TestCreateSource(t *testing.T) {
	source := "https://www.cian.ru/sale/flat/152890949/"
	name := "www.cian.ru"
	url, _ := url.Parse(source)
	s := createSource(source)

	if s.name != name {
		log.Fatal("Source did not recive the right name")
		t.FailNow()
	}

	if s.url != *url {
		log.Fatal("Flat did not recive the right Url")
		t.FailNow()
	}
}
