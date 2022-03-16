package main

import (
	"flag"
	"fmt"
	"net"
)

var ipAdres string
var NameServer string

func init() {
	flag.StringVar(&ipAdres, "IP", "8.8.8.8", "The IPadress")
	flag.StringVar(&NameServer, "NS", "google.com", "The nameserver")
	flag.Parse()
}

func main() {

	ipRecord, err := net.LookupIP(NameServer)
	fmt.Println("--------IPADRESS----------")
	if err != nil {
		fmt.Println("None")
	}
	for _, ip := range ipRecord {
		fmt.Println("IPadress from", NameServer, ":", ip)
	}

	adres, err := net.LookupAddr(ipAdres)
	fmt.Println("--------DNS----------")
	if err != nil {
		fmt.Println("None")
	}
	for _, adrs := range adres {
		fmt.Println("DNS from", ipAdres, ":", adrs)
	}

	NSrecords, err := net.LookupNS(NameServer)
	fmt.Println("--------NAMESERVER----------")
	if err != nil {
		fmt.Println("None")
	}
	for _, ns := range NSrecords {
		fmt.Println("ns record ", ns)
	}

}
