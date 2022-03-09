package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func ReadFile() {

	fmt.Printf("\n\nReading a file in Go lang\n")
	fileName := "dns.txt"

	// The ioutil package contains inbuilt
	// methods like ReadFile that reads the
	// filename and returns the contents.
	data, err := ioutil.ReadFile("dns.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nData: \n%s", data)

}

func readdomains() {

}

func main() {

	content, err := ioutil.ReadFile("dns.txt")
	if err != nil {
		fmt.Println("Error unable to read file!")
	}
	domains := strings.Split(string(content), "\n")

	// strings := []string{"www.google.com", "www.facebook.com", "www.hotmail.com"}
	for _, s := range domains {
		iprecords, _ := net.LookupIP(s)
		fmt.Println("A Records for:\n", s)
		for _, ip := range iprecords {
			fmt.Println(ip)
		}
	}

	// Take www. and replace with nothing to allow mx record lookup
	domainsmx := strings.Replace(string(content), "www.", "", -1)
	domainssplit := strings.Split(string(domainsmx), "\n")
	fmt.Println("Domains:\n", domainsmx)

	for _, a := range domainssplit {
		mxrecords, _ := net.LookupMX(a)
		for _, mx := range mxrecords {
			fmt.Println("MX Records for:\n", mx.Host, "Priority:", mx.Pref)
		}
	}
}
