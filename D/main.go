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
	buf := bufio.NewReader(os.Stdin)

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

	incomeCountry := 0
	parents := 0
	income := 0
	result := make([]int, q, q)
	for i := 0; i < q; i++ {
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
	}

	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}
