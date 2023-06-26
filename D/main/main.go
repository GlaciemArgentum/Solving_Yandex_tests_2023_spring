package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func MyScan(str string) []string {
	str = strings.ReplaceAll(str, "\n", "")
	return strings.Split(str, " ")
}

func RealMain() {
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
	incomeCountryStr := MyScan(in)
	in, _ = buf.ReadString('\n')
	eduCountry := MyScan(in)
	in, _ = buf.ReadString('\n')
	parentsCountry := MyScan(in)

	in, _ = buf.ReadString('\n')
	q, _ := strconv.Atoi(MyScan(in)[0])
	in, _ = buf.ReadString('\n')
	incomeStudent := MyScan(in)
	in, _ = buf.ReadString('\n')
	eduStudent := MyScan(in)
	in, _ = buf.ReadString('\n')
	parentsStudent := MyScan(in)

	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	incomeCountry := 0
	parents := 0
	income := 0
	result := make([]int, q, q)
	for i := 0; i < q; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			income, _ = strconv.Atoi(incomeStudent[i])
			parents, _ = strconv.Atoi(parentsStudent[i])
			if parents != 0 && parentsCountry[parents-1] == "1" {
				result[i] = parents
			}
			for j := 0; j < n; j++ {
				if result[i] != 0 && result[i]-1 <= j {
					break
				}
				if eduCountry[j] == "1" && eduStudent[i] == "0" {
					continue
				}
				incomeCountry, _ = strconv.Atoi(incomeCountryStr[j])
				if income >= incomeCountry {
					result[i] = j + 1
					break
				}
			}
		}(i, wg, mu)
	}

	wg.Wait()

	file, _ := os.Create("output.txt")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(result), "[]")))
}

func main() {
	RealMain()
}
