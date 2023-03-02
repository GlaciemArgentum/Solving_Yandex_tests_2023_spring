package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	day  int
	curr float64
}

func MyScan(str string) []string {
	str = strings.ReplaceAll(str, "\n", "")
	return strings.Split(str, " ")
}

func RealMain() {
	buf := bufio.NewReader(os.Stdin)

	//f, err := os.Open("tests/test" + "2" + ".txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer func(f *os.File) {
	//	_ = f.Close()
	//}(f)
	//buf := bufio.NewReader(f)

	in, _ := buf.ReadString('\n')
	n, _ := strconv.Atoi(MyScan(in)[0])
	in, _ = buf.ReadString('\n')
	dataStr := MyScan(in)

	data := make([]Data, n, n)
	for i := 0; i < n; i++ {
		data[i].curr, _ = strconv.ParseFloat(dataStr[i], 64)
		data[i].day = i + 1
	}

	dataCopy := make([]Data, n, n)
	_ = copy(dataCopy, data)
	sort.SliceStable(dataCopy, func(i, j int) bool { return dataCopy[i].curr < dataCopy[j].curr })
	max := dataCopy[n-1]
	min := dataCopy[0]

	maxWealth := make([]float64, 3, 3)
	maxWealth[0] = 1
	days := make([][]int, 3, 3)

	if max.day > min.day {
		maxWealth[1] = max.curr / min.curr
		days[1] = []int{min.day, max.day}
	} else {
		var minLocal Data
		if max.day > 1 {
			dataCopyMax := make([]Data, max.day-1, max.day-1)
			_ = copy(dataCopyMax, data[:max.day-1])
			sort.SliceStable(dataCopyMax, func(i, j int) bool { return dataCopyMax[i].curr < dataCopyMax[j].curr })
			minLocal = dataCopyMax[0]
		} else {
			minLocal.day = -1
			minLocal.curr = -1
		}

		var maxLocal Data
		if min.day < n {
			dataCopyMin := make([]Data, n-min.day, n-min.day)
			_ = copy(dataCopyMin, data[min.day:])
			sort.SliceStable(dataCopyMin, func(i, j int) bool { return dataCopyMin[i].curr < dataCopyMin[j].curr })
			maxLocal = dataCopyMin[n-min.day-1]
		} else {
			maxLocal.day = -1
			maxLocal.curr = -1
		}

		curr1 := max.curr / minLocal.curr
		curr2 := maxLocal.curr / min.curr
		if minLocal.day > 0 && curr1 >= curr2 {
			maxWealth[1] = curr1
			days[1] = []int{minLocal.day, max.day}
		} else if maxLocal.day > 0 && curr2 >= curr1 {
			maxWealth[1] = curr2
			days[1] = []int{min.day, maxLocal.day}
		}
	}

	if maxWealth[0] >= maxWealth[1] && maxWealth[0] >= maxWealth[2] {
		fmt.Printf("0\n")
	} else if maxWealth[1] >= maxWealth[0] && maxWealth[1] >= maxWealth[2] {
		fmt.Printf("1\n%d %d\n", days[1][0], days[1][1])
	} else if maxWealth[2] >= maxWealth[0] && maxWealth[2] >= maxWealth[1] {
		fmt.Printf("2\n%d %d\n%d %d\n", days[2][0], days[2][1], days[2][2], days[2][3])
	}

}

func main() {
	RealMain()
}
