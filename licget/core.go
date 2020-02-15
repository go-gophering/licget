package licget

import (
	"encoding/json"
	"fmt"
	// "io"
	"log"
	"net/http"
	// "os"
	// "strings"
)

var baseURL = "https://api.opensource.org/licenses/"

func request(url string) []License {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var licenses []License
	json.NewDecoder(res.Body).Decode(&licenses)
	if err != nil {
		log.Fatal(err)
	}
	return licenses
}

func list_licenses() {
	licenses := request(baseURL)
	for _, lic := range licenses {
		fmt.Println(lic.Id)
	}
}

func get_license(license string) {
	licenses := request(baseURL + license)
	fmt.Println(licenses)
}

func Run(license string, list_all bool) {
	switch list_all {
	case true:
		list_licenses()
	case false:
		get_license(license)
	}

}
