package gen

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	nConst   = 200000
	idConst  = 1_000_000_000
	rowConst = 1_000_000_000
)

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Generate(name string) {
	rand.Seed(time.Now().Unix())

	n := rand.Intn(nConst) + 1

	ids := make([]string, 0, n)
	rows := make([]string, 0, n)
	id := ""

	for i := 0; i < n; i++ {
		id = strconv.Itoa(rand.Intn(idConst + 1))
		if Contains(ids, id) {
			i--
			continue
		}
		ids = append(ids, id)
		rows = append(rows, strconv.Itoa(rand.Intn(rowConst)+1))
	}

	k := rand.Intn(nConst) + 1

	text := make([]string, 0, k)
	for i := 0; i < k; i++ {
		text = append(text, ids[rand.Intn(n)])
	}

	file, _ := os.Create(name + ".txt")
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	_, _ = file.WriteString(fmt.Sprintf("%d\n", n))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Join(ids, " ")))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Join(rows, " ")))
	_, _ = file.WriteString(fmt.Sprintf("%d\n", k))
	_, _ = file.WriteString(fmt.Sprintf("%s\n", strings.Join(text, " ")))
}

func main() {
	for i := 0; i < 10; i++ {
		Generate(strconv.Itoa(i))
	}
}
