package main

import (
	"fmt"

	"github.com/lixiangzhong/dnsutil"
)

func main() {
	var dig dnsutil.Dig
	dig.SetDNS("8.8.8.8")
	var domain string
	fmt.Println("What domain would you like to look up?")
	fmt.Scanf("%s", &domain)
	a, err := dig.A(domain)
	fmt.Println(a, err)
}
