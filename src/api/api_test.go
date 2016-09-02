package api_test

import (
	"fmt"
    "log"
    "api"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server   *httptest.Server
	reader   io.Reader
	animalsUrl string
)

func init() {
    router := api.NewRouter()

    log.Fatal( http.ListenAndServe( ":8080", router ) )

	animalsUrl = fmt.Sprintf("%s/animals", router.Path)
}

func TestCreateAnimal( t *testing.T ) {
	animalJson := `{ "name": "bear", "leg_count": 4, "life_span", 30, "is_endangered", true }`

	reader = strings.NewReader(animalJson)

	request, err := http.NewRequest("POST", animalsUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}


func TestListAnimals(t *testing.T) {
	reader = strings.NewReader("")

	request, err := http.NewRequest("GET", animalsUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
