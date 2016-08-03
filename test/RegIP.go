package main;

import (
	"fmt"
	"regexp"
	)

func main() {

	td := "127.0.0.1"
	allowips := []string{"127.0.[0-5].*", "127.[[0.0.0", "127.0.1.*", "127.*.*.*", "127.0.[0-1].[0-1]"}

	for _,ip := range allowips {
		
		b, _ := regexp.MatchString(ip, td)
		
		if b {
			fmt.Printf("%s matched %s\n", td, ip)
		} else {
			fmt.Printf("%s not matched %s\n", td, ip)
		}
	}
}
