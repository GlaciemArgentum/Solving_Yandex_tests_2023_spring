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

	dataBig := make([]Data, n, n)
	for i := 0; i < n; i++ {
		dataBig[i].curr, _ = strconv.ParseFloat(dataStr[i], 64)
		dataBig[i].day = i + 1
	}

	preCurr := 10_001.0
	minBorder := 0
	maxBorder := n
	flag := 0
	for i := 0; i < n; i++ {
		if dataBig[i].curr > preCurr {
			minBorder = i - 1
			flag++
			break
		}
		preCurr = dataBig[i].curr
	}
	if flag == 0 {
		fmt.Printf("0\n")
	}
	preCurr = 0
	for i := n - 1; i >= 0; i-- {
		if dataBig[i].curr < preCurr {
			maxBorder = i + 2
			break
		}
		preCurr = dataBig[i].curr
	}
	data := dataBig[minBorder:maxBorder]
	lenData := len(data)

	dataCopy := make([]Data, lenData, lenData)
	_ = copy(dataCopy, data)
	sort.SliceStable(dataCopy, func(i, j int) bool { return dataCopy[i].curr < dataCopy[j].curr })
	max := dataCopy[lenData-1]
	min := dataCopy[0]

	maxWealth := make([]float64, 2, 2)
	days := make([][]int, 2, 2)

	if max.day > min.day {
		maxWealth[0] = max.curr / min.curr
		days[0] = []int{min.day, max.day}
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
		if min.day < lenData {
			dataCopyMin := make([]Data, lenData-min.day, lenData-min.day)
			_ = copy(dataCopyMin, data[min.day:])
			sort.SliceStable(dataCopyMin, func(i, j int) bool { return dataCopyMin[i].curr < dataCopyMin[j].curr })
			maxLocal = dataCopyMin[lenData-min.day-1]
		} else {
			maxLocal.day = -1
			maxLocal.curr = -1
		}

		curr1 := max.curr / minLocal.curr
		curr2 := maxLocal.curr / min.curr
		if minLocal.day > 0 && curr1 >= curr2 {
			maxWealth[0] = curr1
			days[0] = []int{minLocal.day, max.day}
		} else if maxLocal.day > 0 && curr2 >= curr1 {
			maxWealth[0] = curr2
			days[0] = []int{min.day, maxLocal.day}
		}
	}

	if maxWealth[0] >= maxWealth[1] {
		fmt.Printf("1\n%d %d\n", days[0][0], days[0][1])
	} else {
		fmt.Printf("2\n%d %d\n%d %d\n", days[1][0], days[1][1], days[1][2], days[1][3])
	}

}

func main() {
	RealMain()
}
