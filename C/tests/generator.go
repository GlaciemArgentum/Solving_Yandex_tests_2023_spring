package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	countOfFiles = 1

	nConst = 20 //300_000
	kConst = 10 //10_000
)

func Generate(name string) {
	rand.Seed(time.Now().Unix())

	n := nConst
	//n := rand.Intn(nConst) + 1
	k := kConst
	//k := rand.Intn(kConst) + 1

	d := make([]int, 0, n)

	for i := 0; i < n; i++ {
		d = append(d, rand.Intn(k)+1)
	}

	file, _ := os.Create("test" + name + ".txt")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	_, _ = file.WriteString(fmt.Sprintf("%d\n", n))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(d), "[]")))
}

func main() {
	for i := 1; i <= countOfFiles; i++ {
		Generate(strconv.Itoa(i))
	}
}
