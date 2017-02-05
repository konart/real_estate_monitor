package parsers_test

import (
	cian "github.com/konart/cian/parsers"
	"testing"
	"net/url"
	"os"
)

const testFile = "../examples/cian_example.html"

func TestParseFlatPage(t *testing.T) {
	url, _ := url.Parse("https://www.cian.ru/sale/flat/152890949/")
	cian.GetFlatPage(url)
	testFlatMap := make(map[string]string)

	//testFlatMap["address"] = "Сходненская 6, к.1, кв.24"
	//testFlatMap["source"] = "https://www.cian.ru/sale/flat/152890949/"
	testFlatMap["rooms"] = "2"
	testFlatMap["height"] = "2,74"
	testFlatMap["bathrooms"] = "1"
	testFlatMap["reparate_bathrooms"] = "1"

	page, _ := os.Open(testFile)
	//reader := bufio.NewReader(f)
	//reader.

	flatMap := cian.ParseFlatPage(page)
	for k, v := range flatMap {
		if flatMap[k] != testFlatMap[k] {
			t.Logf("Failed to parse the example, key: %s, value: %s [expected: %s]", k, v, testFlatMap[k])
			t.Fail()
		}
	}
}