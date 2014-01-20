/**
 *
 */
package main

import (
	"bufio"
	"fmt"
	"time"
	"regexp"
	"strings"
	"net/theatlantis/octopus"
)

func main() {
	start := time.Now()
	fmt.Printf("Hello world!\n")
	fmt.Printf("Second line!\n")
	start.Day();
	s := strings.NewReader("Mrs. Leonora Spocky")
	re1, _ := regexp.Compile(`[Le]{2}`)
	result:= re1.FindReaderIndex(octopus.NewDebugReader(bufio.NewReader(s)))
	fmt.Println(result)
}
