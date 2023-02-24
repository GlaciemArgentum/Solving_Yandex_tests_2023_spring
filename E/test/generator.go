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

	nConst = 100_000
	kConst = 1_000_000_000
)

func Generate(name string) {
	rand.Seed(time.Now().Unix())

	//n := nConst
	n := rand.Intn(nConst) + 1
	//k := kConst
	k := rand.Intn(kConst) + 1

	book := ""
	p := make([]int, 0, n)
	d := make([]int, 0, n)

	for i := 0; i < n; i++ {
		book += string(rune(rand.Intn(26)) + 'a')
		p = append(p, i+1)
		d = append(d, rand.Intn(26))
	}

	rand.Shuffle(len(p), func(i, j int) { p[i], p[j] = p[j], p[i] })

	file, _ := os.Create("test" + name + ".txt")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	_, _ = file.WriteString(fmt.Sprintf("%d %d\n", n, k))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", book))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(p), "[]")))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(d), "[]")))
}

func main() {
	for i := 1; i <= countOfFiles; i++ {
		Generate(strconv.Itoa(i))
	}
}
