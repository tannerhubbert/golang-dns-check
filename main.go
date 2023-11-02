package main

import (
	"fmt"

	"github.com/lixiangzhong/dnsutil"
)

func main() {
	var dig dnsutil.Dig
	dig.SetDNS("8.8.8.8")
	var name string
	fmt.Println("What domain would you like to look up?")
	fmt.Scanf("%s", &name)
	a, err := dig.A(name)
	fmt.Println(a, err)
}
