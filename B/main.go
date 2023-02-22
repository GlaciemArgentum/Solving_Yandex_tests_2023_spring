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

func main() {
	buf := bufio.NewReader(os.Stdin)

	in, _ := buf.ReadString('\n')
	params := MyScan(in)
	n, _ := strconv.Atoi(params[0])
	xInt, _ := strconv.Atoi(params[1])
	x := float64(xInt)
	tInt, _ := strconv.Atoi(params[2])
	t := float64(tInt)
	in, _ = buf.ReadString('\n')
	dataStr := MyScan(in)
	data := make([]Data, n, n)

	num := 0
	for i := 0; i < n; i++ {
		num, _ = strconv.Atoi(dataStr[i])
		data[i].weight = math.Abs(x - float64(num))
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

	fmt.Println(len(onTime))
	if len(onTime) > 0 {
		fmt.Println(strings.Trim(fmt.Sprint(onTime), "[]"))
	}
}
