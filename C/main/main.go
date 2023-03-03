package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	day  int
	curr float64
}

type TwoDeals struct {
	total float64
	buy   Data
	sell  Data
}

func MyScan(str string) []string {
	str = strings.ReplaceAll(str, "\n", "")
	return strings.Split(str, " ")
}

// TrimReduce убирает падение с обоих концов среза
func TrimReduce(input []Data) []Data {
	lenInput := len(input)

	if lenInput < 2 {
		return []Data{}
	}
	if lenInput == 2 {
		if input[0].curr < input[1].curr {
			return input
		} else {
			return []Data{}
		}
	}

	preCurr := 10_001.0
	minBorder := 0
	maxBorder := lenInput
	flag := 0
	value := 0.0
	for i := 0; i < lenInput; i++ {
		value = input[i].curr
		if value > preCurr {
			minBorder = i - 1
			flag++
			break
		}
		preCurr = value
	}
	if flag == 0 {
		return []Data{}
	}
	preCurr = 0
	for i := lenInput - 1; i >= 0; i-- {
		value = input[i].curr
		if value < preCurr {
			maxBorder = i + 2
			break
		}
		preCurr = value
	}
	return input[minBorder:maxBorder]
}

// DeleteRepeat удаляет дни, в которые повторяется рост или падение
func DeleteRepeat(input []Data) []Data {
	if len(input) < 3 {
		return input
	}
	output := make([]Data, 0, len(input))

	prePre := input[0].curr
	pre := input[1]
	value := 0.0
	for _, i := range input[2:] {
		value = pre.curr
		if (prePre-value)*(i.curr-value) > 0 {
			output = append(output, pre)
			prePre = value
		}
		pre = i
	}
	return output
}

// GetMinMax возвращает min, max
func GetMinMax(input []Data) (Data, Data) {
	if len(input) == 0 {
		return Data{}, Data{}
	}

	min := input[0]
	max := input[0]
	var value float64
	for _, i := range input {
		value = i.curr
		if value > max.curr {
			max = i
		} else if value < min.curr {
			min = i
		}
	}
	return min, max
}

func Input() (int, []Data) {
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
	return n, data
}

func Output(k int, days []int) {
	if k == 2 {
		fmt.Printf("2\n%d %d\n%d %d\n", days[0], days[1], days[2], days[3])
		return
	} else if k == 1 {
		fmt.Printf("1\n%d %d\n", days[0], days[1])
		return
	} else if k == 0 {
		fmt.Printf("0\n")
		return
	}
}

func RealMain() {
	n, data := Input()

	lenData := len(data)
	if lenData < 3 {
		switch {
		case lenData == 1:
			Output(0, []int{})
			return
		case lenData == 2:
			if data[0].curr < data[1].curr {
				Output(1, []int{1, 2})
			} else {
				Output(0, []int{})
			}
			return
		}
	}

	data = DeleteRepeat(TrimReduce(data))
	lenData = len(data)

	min, max := GetMinMax(data)

	twoDealsSlice := make([]TwoDeals, 0, n)
}

func main() {
	RealMain()
}
