package main

import (
	"bufio"
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

	if max.day > min.day {
		maxWealth[1] = max.curr / min.curr
	}

	//
	//max1 := data[n-2]
	//max2 := data[n-1]
	//min1 := data[0]
	//min2 := data[1]

}

func main() {
	RealMain()
}
