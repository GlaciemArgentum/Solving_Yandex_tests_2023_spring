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

type Data struct {
	letter   int
	next     int
	d        int
	position int
}

func RealMain() {
	//buf := bufio.NewReader(os.Stdin)

	f, err := os.Open("tests/test" + "1" + ".txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	buf := bufio.NewReader(f)

	in, _ := buf.ReadString('\n')
	inScan := MyScan(in)
	n, _ := strconv.Atoi(inScan[0])
	k, _ := strconv.Atoi(inScan[1])
	in, _ = buf.ReadString('\n')
	sRune := []rune(strings.ReplaceAll(in, "\n", ""))
	in, _ = buf.ReadString('\n')
	pStr := MyScan(in)
	in, _ = buf.ReadString('\n')
	dStr := MyScan(in)

	datas := make(map[int]Data, n)
	for i := 0; i < n; i++ {
		var data Data
		data.letter = int(sRune[i] - 'a')
		pBuf, _ := strconv.Atoi(pStr[i])
		data.next = pBuf - 1
		data.d, _ = strconv.Atoi(dStr[i])
		data.position = i
		datas[i] = data
	}

	var power int
	for i := 0; i < n; i++ {
		spell := make([]rune, 0, n)
		letter := datas[i]
		meets := make(map[int]int, n)
		library := make(map[int]int, n)
		preLen := 0
		counter := 0
		for j := 0; j < k; j++ {
			meets[letter.letter]++
			if meets[letter.letter] > 1 {
				letter.letter = (letter.letter + (meets[letter.letter]-1)*letter.d) % 26
			}
			library[letter.letter]++
			spell = append(spell, rune(letter.letter)+'a')
			letter = datas[letter.next]
			if len(library) == preLen {
				counter++
				if counter == 26 {
					break
				}
			} else {
				counter = 1
			}
			preLen = len(library)
		}
		word := string(spell)
		lenWord := len(word)
		lastPower := 0

		for h := 1; h <= lenWord; h++ {
			for o := 'a'; o < 'a'+26; o++ {
				if strings.ContainsRune(word[:h], o) {
					power++
					if h == lenWord {
						lastPower++
					}
				}
			}
		}
		power += (k - lenWord) * lastPower
	}
	fmt.Println(power)
}

func main() {
	RealMain()
}
