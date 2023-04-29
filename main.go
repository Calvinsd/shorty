package main

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var hashToUrls map[string]string = make(map[string]string, 100)

var lookUpTable map[string]string = make(map[string]string, 100)

type LongUrl struct {
	Url string `json:"url"`
}

func main() {

	const port = ":8080"

	fmt.Println("Starting server on :8080")

	http.ListenAndServe(port, http.HandlerFunc(handleRoutes))

}

func handleRoutes(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL.Path)
	log.Println("Lookup :", lookUpTable)
	log.Println("hash Table :", hashToUrls)

	// handle get routes
	if r.Method == http.MethodGet {
		log.Println(strings.Split(r.URL.Path, "/")[1])
		if pathParams := strings.Split(r.URL.Path, "/"); len(pathParams) == 3 {

			log.Println("path params", pathParams)
			if longUrl, ok := hashToUrls[pathParams[2]]; ok {
				http.Redirect(w, r, longUrl, http.StatusMovedPermanently)
				return
			}

			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		return
	}

	// handle post routes
	if r.Method == http.MethodPost && r.URL.Path == "/url" {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Fatal(err)
		}

		var longUrl LongUrl

		err = json.Unmarshal(body, &longUrl)

		if err != nil {
			log.Fatal(err)
		}

		_, ok := lookUpTable[longUrl.Url]

		if ok {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(lookUpTable[longUrl.Url]))
			return
		}

		table := crc32.MakeTable(crc32.IEEE)

		checksum := crc32.Checksum([]byte("input1"), table)

		hashToUrls[strconv.Itoa(int(checksum))] = longUrl.Url

		lookUpTable[longUrl.Url] = strconv.Itoa(int(checksum))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(strconv.Itoa(int(checksum))))

		return
	}

	w.WriteHeader(http.StatusNotFound)

	return
}

// if dest, ok := pathsToUrls[path]; ok {
// 	http.Redirect(w, r, dest, http.StatusFound)
// 	return
