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
	nConst   = 200000
	idConst  = 1_000_000_000
	rowConst = 1_000_000_000

	countConst = 1
)

func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Generate(name string) {
	rand.Seed(time.Now().Unix())

	n := nConst
	//n := rand.Intn(nConst) + 1

	ids := make([]int, 0, n)
	rows := make([]int, 0, n)
	id := 0

	for i := 0; i < n; i++ {
		id = rand.Intn(idConst + 1)
		if Contains(ids, id) {
			i--
			continue
		}
		ids = append(ids, id)
		rows = append(rows, rand.Intn(rowConst)+1)
	}

	k := rand.Intn(nConst) + 1

	text := make([]int, 0, k)
	for i := 0; i < k; i++ {
		text = append(text, ids[rand.Intn(n)])
	}

	file, _ := os.Create(name + ".txt")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	_, _ = file.WriteString(fmt.Sprintf("%d\n", n))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(ids), "[]")))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(rows), "[]")))
	_, _ = file.WriteString(fmt.Sprintf("%d\n", k))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(text), "[]")))
}

func main() {
	for i := 0; i < countConst; i++ {
		Generate(strconv.Itoa(i))
	}
}
