package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Contents of the array are below</h1>"))

	var arr []string
	i := 0
	for i < 8 {
		inst := getter()

		if strings.HasSuffix(inst, "jpg") {
			arr = append(arr, inst)
			i++

		}

	}
	//print the obtained content for now

	for dog := range arr {
		w.Write([]byte(`<img src="`))
		w.Write([]byte(arr[dog]))
		w.Write([]byte(`" width="300"`))
		w.Write([]byte("<h3> .. </h3>"))

	}

}

type dog struct {
	URL string `json:"url"`
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)

}

func getter() string {
	url := "https://random.dog/woof.json"

	dogClient := http.Client{
		Timeout: time.Second * 3, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Dog-Getter", "doglist-webapp")

	res, getErr := dogClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	dog1 := dog{}
	jsonErr := json.Unmarshal(body, &dog1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	adog := dog1.URL
	return adog

}
