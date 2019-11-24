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

	//dat, err := ioutil.ReadFile("/proc/meminfo")
	//check(err)

	dat, err := os.Open("/proc/meminfo")
	check(err)
	//	splitdat := strings.Split(string(dat), " ")

	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "MemFree:") {
			fmt.Println(scanner.Text())
		}
		//fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	memstring := ""

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), typeram) {
			//fmt.Println(scanner.Text())
			memstring = string(scanner.Text())
		}
		//fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, "<html><head></head><body>")
	fmt.Fprintln(w, "<h1>RAM statistics:</h1><br />")
	fmt.Fprintln(w, "<p>")
	fmt.Fprintln(w, memstring)
	fmt.Fprintln(w, "</p>")
	fmt.Fprintln(w, "</body></html>")
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
