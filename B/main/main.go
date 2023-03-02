package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	id     int
	weight float64
}

func MyScan(str string) []string {
	str = strings.ReplaceAll(str, "\n", "")
	return strings.Split(str, " ")
}

func RealMain() {
	buf := bufio.NewReader(os.Stdin)

	//f, err := os.Open("tests/test" + "1" + ".txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer func(f *os.File) {
	//	_ = f.Close()
	//}(f)
	//buf := bufio.NewReader(f)

	in, _ := buf.ReadString('\n')
	params := MyScan(in)
	n, _ := strconv.Atoi(params[0])
	x, _ := strconv.ParseFloat(params[1], 64)
	t, _ := strconv.ParseFloat(params[2], 64)
	in, _ = buf.ReadString('\n')
	dataStr := MyScan(in)
	data := make([]Data, n, n)

	num := 0.0
	for i := 0; i < n; i++ {
		num, _ = strconv.ParseFloat(dataStr[i], 64)
		data[i].weight = math.Abs(x - num)
		data[i].id = i + 1
	}

	sort.SliceStable(data, func(i, j int) bool { return data[i].weight < data[j].weight })

	onTime := make([]int, 0, n)
	for _, d := range data {
		t -= d.weight
		if t < 0 {
			break
		}
		onTime = append(onTime, d.id)
	}
	sort.Ints(onTime)

	file, _ := os.Create("output.txt")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	_, _ = file.WriteString(fmt.Sprintf("%d\n", len(onTime)))
	//fmt.Println(len(onTime))

	if len(onTime) > 0 {
		_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(onTime), "[]")))
		//fmt.Println(strings.Trim(fmt.Sprint(onTime), "[]"))
	}
}

func main() {
	RealMain()
}
