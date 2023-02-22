package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	countOfFiles = 1

	countryConst  = 200_000       //колличество стран
	studentsConst = 200_000       //колличество студентов
	incomeConst   = 1_000_000_000 //максимальный доход
)

func Generate(name string) {
	//rand.Seed(time.Now().Unix())

	n := countryConst //rand.Intn(countryConst) + 1

	incomes := make([]int, 0, n)
	edus := make([]int, 0, n)
	parents := make([]int, 0, n)

	for i := 0; i < n; i++ {
		incomes = append(incomes, rand.Intn(incomeConst+1))
		edus = append(edus, rand.Intn(2))
		parents = append(parents, rand.Intn(2))
	}

	k := studentsConst //rand.Intn(studentsConst) + 1

	incomesC := make([]int, 0, n)
	edusC := make([]int, 0, n)
	parentsC := make([]int, 0, n)

	for i := 0; i < k; i++ {
		incomesC = append(incomesC, rand.Intn(incomeConst+1))
		edusC = append(edusC, rand.Intn(2))
		parentsC = append(parentsC, rand.Intn(n+1))
	}

	file, _ := os.Create("test" + name + ".txt")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	_, _ = file.WriteString(fmt.Sprintf("%d\n", n))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(incomes), "[]")))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(edus), "[]")))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(parents), "[]")))
	_, _ = file.WriteString(fmt.Sprintf("%d\n", k))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(incomesC), "[]")))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(edusC), "[]")))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(parentsC), "[]")))
}

func main() {
	for i := 1; i <= countOfFiles; i++ {
		Generate(strconv.Itoa(i))

	}
}
