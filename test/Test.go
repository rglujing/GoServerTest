package main;

import "strconv"
import "fmt"

func main() {
	var i int64
	i = 123456780011111111

	fmt.Printf("%s", strconv.FormatInt(i, 10))

}


