package main
// bufio.Scanner、ioutil.ReadFile和ioutil.WriteFile都使用*os.File的Read和Write方法
import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	flag := 0
	for input.Scan() {
		if flag > 2 {
			break
		}
		flag++
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}