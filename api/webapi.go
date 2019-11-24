package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<h1>General statistics:</h1>")
	fmt.Fprintln(w, "</body></html>")
}

func RamShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeram := vars["type"]

	dat, err := os.Open("/proc/meminfo")
	check(err)

	scanner := bufio.NewScanner(dat)

	memstring := ""

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), typeram) {
			memstring = string(scanner.Text())
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	page := "<html><head></head><body><h1>RAM statistics:</h1><br /><p>" + memstring + "</p></body></html>"

	fmt.Fprintln(w, page)
}

func RamIndex(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("/proc/meminfo")
	check(err)

	fmt.Fprintln(w, "<html><head></head><body>")
	fmt.Fprintln(w, "<h1>RAM statistics:</h1><br />")
	fmt.Fprintln(w, "<p>")
	fmt.Fprintln(w, string(dat))
	//fmt.Print(string(dat))
	fmt.Fprintln(w, "</p>")
	fmt.Fprintln(w, "</body></html>")
}

func DiskShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mount := vars["mount"]
	fmt.Fprintln(w, "Disk show:", mount)
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/ram/{type}", RamShow)
	router.HandleFunc("/ram", RamIndex)
	router.HandleFunc("/disk/{mount}", DiskShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
