package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MyScan() []string {
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	str = strings.ReplaceAll(str, "\n", "")

	return strings.Split(str, " ")
}

func main() {
	n, _ := strconv.Atoi(MyScan()[0])
	ids := MyScan()
	rows := MyScan()
	k, _ := strconv.Atoi(MyScan()[0])
	text := MyScan()

	mapRows := make(map[string]string, n)
	for i := 0; i < n; i++ {
		mapRows[ids[i]] = rows[i]
	}

	preInd := mapRows[text[0]]
	counter := 0
	for i := 1; i < k; i++ {
		if mapRows[text[i]] != preInd {
			counter++
			preInd = mapRows[text[i]]
		}
	}
	fmt.Println(counter)
}
