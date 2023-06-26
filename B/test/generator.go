package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	countOfFiles = 10

	nConst = 200_000
	xConst = 1_000_000_000
	tConst = 300_000_000_000_000
)

func Generate(name string) {
	//rand.Seed(time.Now().Unix())

	n := nConst
	//n := rand.Intn(nConst) + 1
	x := xConst
	//x := rand.Intn(xConst) + 1
	t := tConst
	//t := rand.Intn(tConst) + 1

	scul := make([]int, 0, n)

	for i := 0; i < n; i++ {
		scul = append(scul, rand.Intn(x)+1)
	}

	file, _ := os.Create("test" + name + ".txt")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	_, _ = file.WriteString(fmt.Sprintf("%d %d %d\n", n, x, t))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(scul), "[]")))
}

func main() {
	for i := 1; i <= countOfFiles; i++ {
		Generate(strconv.Itoa(i))
	}
}
