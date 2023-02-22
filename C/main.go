package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	day  int
	curr int
}

func MyScan(str string) []string {
	str = strings.ReplaceAll(str, "\n", "")
	return strings.Split(str, " ")
}

func main() {
	buf := bufio.NewReader(os.Stdin)

	in, _ := buf.ReadString('\n')
	n, _ := strconv.Atoi(in)
	in, _ = buf.ReadString('\n')
	dataStr := MyScan(in)

	data := make([]Data, n, n)
	for i := 0; i < n; i++ {
		data[i].curr, _ = strconv.Atoi(dataStr[i])
		data[i].day = i + 1
	}

	//sort.SliceStable(data, func(i, j int) bool { return data[i].curr < data[j].curr })
	//
	//max1 := data[n-2]
	//max2 := data[n-1]
	//min1 := data[0]
	//min2 := data[1]

}
