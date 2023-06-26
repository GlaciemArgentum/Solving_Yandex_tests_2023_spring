package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MyScan(str string) []string {
	str = strings.ReplaceAll(str, "\n", "")
	return strings.Split(str, " ")
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	buf := bufio.NewReader(f)

	in, _ := buf.ReadString('\n')
	n, _ := strconv.Atoi(MyScan(in)[0])

	in, _ = buf.ReadString('\n')
	ids := MyScan(in)

	in, _ = buf.ReadString('\n')
	rows := MyScan(in)

	in, _ = buf.ReadString('\n')
	k, _ := strconv.Atoi(MyScan(in)[0])

	in, _ = buf.ReadString('\n')
	text := MyScan(in)

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
