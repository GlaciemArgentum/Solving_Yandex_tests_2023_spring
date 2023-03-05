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

type Deal struct {
	total float64
	buy   int
	sell  int
}

type TwoDeals struct {
	total float64
	buy1  int
	sell1 int
	buy2  int
	sell2 int
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
	lenData := len(input)
	if lenData < 3 {
		return input
	}
	output := make([]Data, 0, lenData)
	output = append(output, input[0])

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
	output = append(output, input[lenData-1])
	return output
}

// GetMin возвращает minPosition, min
func GetMin(input []Data) (int, Data) {
	if len(input) == 0 {
		return -1, Data{}
	}

	min := input[0]
	minPosition := 0
	for i, data := range input {
		if data.curr < min.curr {
			min = data
			minPosition = i
		}
	}
	return minPosition, min
}

// GetMax возвращает maxPosition, max
func GetMax(input []Data) (int, Data) {
	if len(input) == 0 {
		return -1, Data{}
	}

	max := input[0]
	maxPosition := 0
	for i, data := range input {
		if data.curr > max.curr {
			max = data
			maxPosition = i
		}
	}
	return maxPosition, max
}

// CheckData возвращает false, если дат > 2, иначе true. Также выводит []Data{}, которые приведут к увеличению прибыли.
func CheckData(input []Data) (bool, []int) {
	lenData := len(input)
	if lenData >= 3 {
		return false, []int{}
	}
	if lenData == 2 {
		if input[0].curr < input[1].curr {
			return true, []int{input[0].day, input[1].day}
		}
	}
	return true, []int{}
}

func Input() (int, []Data) {
	//buf := bufio.NewReader(os.Stdin)

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
	dataStr := MyScan(in)

	data := make([]Data, n, n)
	for i := 0; i < n; i++ {
		data[i].curr, _ = strconv.ParseFloat(dataStr[i], 64)
		data[i].day = i + 1
	}
	return n, data
}

func Output(days []int) {
	lenData := len(days)
	if lenData == 4 {
		fmt.Printf("2\n%d %d\n%d %d\n", days[0], days[1], days[2], days[3])
		return
	} else if lenData == 2 {
		fmt.Printf("1\n%d %d\n", days[0], days[1])
		return
	} else if lenData == 0 {
		fmt.Printf("0\n")
		return
	}
}

// BestDeal возвращает
func BestDeal(input []Data) Deal {
	input = TrimReduce(input)

	lenData := len(input)
	if lenData == 2 {
		return Deal{input[1].curr / input[0].curr, input[0].day, input[1].day}
	}
	if lenData == 0 || lenData == 1 {
		return Deal{0, -1, -1}
	}

	minPosition, min := GetMin(input)
	maxPosition, max := GetMax(input)
	if min.day < max.day {
		return Deal{max.curr / min.curr, min.day, max.day}
	}

	_, localMin := GetMin(input[:maxPosition])
	_, localMax := GetMax(input[minPosition+1:])

	leftDeal := Deal{max.curr / localMin.curr, localMin.day, max.day}
	rightDeal := Deal{localMax.curr / min.curr, min.day, localMax.day}
	middleDeal := BestDeal(input[maxPosition+1 : minPosition])

	switch {
	case leftDeal.total >= rightDeal.total && leftDeal.total >= middleDeal.total:
		return leftDeal
	case rightDeal.total >= middleDeal.total:
		return rightDeal
	}
	return middleDeal
}

func RealMain() {
	n, data := Input()
	if flag, days := CheckData(data); flag {
		Output(days)
		return
	}
	data = DeleteRepeat(TrimReduce(data))
	if flag, days := CheckData(data); flag {
		Output(days)
		return
	}

	oneDeal := BestDeal(data)

	twoDealsSlice := make([]TwoDeals, 0, n)
	for i := 2; i < len(data)-1; i++ {
		leftDeal := BestDeal(data[:i])
		rightDeal := BestDeal(data[i:])
		twoDealsSlice = append(twoDealsSlice, TwoDeals{leftDeal.total * rightDeal.total, leftDeal.buy, leftDeal.sell, rightDeal.buy, rightDeal.sell})
	}

	twoDeal := twoDealsSlice[0]
	for _, deal := range twoDealsSlice {
		if deal.total > twoDeal.total {
			twoDeal = deal
		}
	}

	switch {
	case oneDeal.total <= 1 && twoDeal.total <= 1:
		Output([]int{})
	case oneDeal.total >= twoDeal.total:
		Output([]int{oneDeal.buy, oneDeal.sell})
	default:
		Output([]int{twoDeal.buy1, twoDeal.sell1, twoDeal.buy2, twoDeal.sell2})
	}
}

func main() {
	RealMain()
}
