package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//dat, err := ioutil.ReadFile("/proc/meminfo")
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

	//fmt.Print(splitdat)

}
