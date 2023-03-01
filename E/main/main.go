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
	inScan := MyScan(in)
	n, _ := strconv.Atoi(inScan[0])
	k, _ := strconv.Atoi(inScan[1])
	in, _ = buf.ReadString('\n')
	//
	//fmt.Print(in)
	//
	sRune := []rune(strings.ReplaceAll(in, "\n", ""))
	in, _ = buf.ReadString('\n')
	pStr := MyScan(in)
	in, _ = buf.ReadString('\n')
	dStr := MyScan(in)

	datas := make(map[int]Data, n) // Входные данные.
	//letterLibrary := make(map[int][]int, 26) // Позиции, на которых расположена та или иная буква.
	for i := 0; i < n; i++ {
		var data Data
		data.letter = int(sRune[i] - 'a')
		pBuf, _ := strconv.Atoi(pStr[i])
		data.next = pBuf - 1
		data.d, _ = strconv.Atoi(dStr[i])
		data.position = i
		datas[i] = data
		//letterLibrary[data.letter] = append(letterLibrary[data.letter], i)
	}

	var power int                        // Мощность заклинания.
	letterStart := datas[datas[0].next]  // Первая буква заклинания
	libraryStart := make(map[int]int, n) // Буквы (их позиции в book), с которых начиналось заклинание.
	libraryStart[letterStart.position]++
	//var preSpell []rune
	//preMeets := make(map[int]int, n)
	//preLibrary := make(map[int]int, n)
	for i := 0; i < n; i++ {
		letter := letterStart           // Выбранная буква.
		spell := make([]rune, 0, 26)    // Текст заклинания.
		meets := make(map[int]int, n)   // Буквы, которые были прочитаны из book (и сколько раз).
		library := make(map[int]int, n) // Буквы, которые были добавлены в spell (И сколько раз).
		preLen := 0                     // Ранее зафиксированная длина library.
		counter := 0                    // Счётчик, сколько раз не изменялась длина library.
		//if len(preSpell) != 0 {
		//	spell = preSpell[1:]
		//	meets = preMeets
		//	meets[letterStart.letter]--
		//	library = preLibrary
		//	for z := 0; z < len(letterLibrary[letterStart.letter]); z++ {
		//		spell[letterLibrary[letterStart.letter][z]] = rune((int(spell[letterLibrary[letterStart.letter][z]]-'a')-(z+1)*datas[letterLibrary[letterStart.letter][z]].d+26)%26) + 'a'
		//	}
		//}
		for j := len(spell); j < k; j++ {
			meets[letter.letter]++
			if meets[letter.letter] > 1 {
				letter.letter = (letter.letter + (meets[letter.letter]-1)*letter.d) % 26
			}
			library[letter.letter]++
			spell = append(spell, rune(letter.letter)+'a')
			letter = datas[letter.next]
			if len(library) == preLen {
				counter++
				if counter == 26*n {
					break
				}
			} else {
				counter = 0
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

		//
		//fmt.Println(strings.Repeat(" ", i), word)
		//

		letterStart = datas[letterStart.next]
		if _, inMap := libraryStart[letterStart.position]; inMap {
			//
			//fmt.Println()
			//
			for _, newLetterStart := range datas {
				if _, inMap2 := libraryStart[newLetterStart.position]; !inMap2 {
					letterStart = newLetterStart
					break
				}
			}
			//preSpell = []rune{}
		}
		libraryStart[letterStart.position]++

		//preSpell = spell
		//preMeets = meets
		//preLibrary = library
	}
	fmt.Println(power)
}

func main() {
	RealMain()
}
