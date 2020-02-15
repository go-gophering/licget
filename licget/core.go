package licget

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	// "os"
	// "strings"
)

var baseURL = "https://api.github.com/licenses"

func request(url string, object interface{}) error {
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/vnd.github.drax-preview+json")
	res, getErr := client.Do(req)
	if getErr != nil {
		return getErr
	}
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}
	// fmt.Println(string(body))

	jsonErr := json.Unmarshal(body, &object)
	if jsonErr != nil {
		return jsonErr
	}

	return nil
}

func list_licenses() {
	licenses := Licenses{}
	err := request(baseURL, &licenses)
	if err != nil {
		log.Fatal(err)
	}

	for _, license := range licenses {
		fmt.Println(license.Key)
	}
}

func get_license(licenseKey string) {
	license := License{}
	err := request(baseURL+"/"+licenseKey, license)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(license)
}

func Run(license string, list_all bool) {
	switch list_all {
	case true:
		list_licenses()
	case false:
		get_license(license)
	}

}
