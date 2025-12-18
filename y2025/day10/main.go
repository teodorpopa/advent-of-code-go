package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

const NoSolution = math.MaxInt

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func parseLine(
	textline string,
	display *int,
	buttons *[]int,
	other *[]int,
) {
	r := strings.NewReader(textline)
	read := func() (rune, bool) {
		ch, _, err := r.ReadRune()
		return ch, err == nil
	}
	readNonSpace := func() (rune, bool) {
		for {
			ch, ok := read()
			if !ok {
				return 0, false
			}
			if ch != ' ' && ch != '\n' && ch != '\t' {
				return ch, true
			}
		}
	}

	*display = 0
	displaySize := 0
	read()
	for {
		ch, ok := read()
		if !ok {
			break
		}
		if ch == '.' {
			*display <<= 1
		} else if ch == '#' {
			*display = (*display << 1) | 1
		} else {
			break
		}
		displaySize++
	}

	*buttons = []int{}
	for {
		ch, ok := readNonSpace()
		if !ok || ch != '(' {
			break
		}
		button := 0
		for {
			temp := 0
			fmt.Fscan(r, &temp)
			ch, _ = readNonSpace()
			button |= (1 << (displaySize - temp - 1))
			if ch == ')' {
				break
			}
		}
		*buttons = append(*buttons, button)
	}

	*other = []int{}
	for {
		temp := 0
		_, err := fmt.Fscan(r, &temp)
		if err != nil {
			break
		}
		ch, _ := readNonSpace()
		*other = append(*other, temp)
		if ch == '}' {
			break
		}
	}
}

func permutation(
	display int,
	buttons []int,
	buttonIdx int,
	depth int,
	minDepth *int,
	stack []int,
) {
	if depth > *minDepth {
		return
	}
	if buttonIdx != -1 {
		display ^= buttons[buttonIdx]
		if display == 0 {
			if depth < *minDepth {
				*minDepth = depth
			}
			return
		}
	}
	for i := buttonIdx + 1; i < len(buttons); i++ {
		stack = append(stack, i)
		permutation(display, buttons, i, depth+1, minDepth, stack)
		stack = stack[:len(stack)-1]
	}
}

func processLine(textline string, sum *int) {
	if strings.TrimSpace(textline) == "" {
		return
	}
	var display int
	var buttons []int
	var other []int
	parseLine(textline, &display, &buttons, &other)
	if display == 0 {
		*sum += 0
		return
	}
	minDepth := NoSolution
	permutation(display, buttons, -1, 0, &minDepth, []int{})
	if minDepth == NoSolution {
		return
	}
	*sum += minDepth
}

func parseButtonRaw(b []byte, l int) []int {
	var (
		res     = make([]int, l)
		i, next int
	)
	b = b[1 : len(b)-1]

	next = bytes.IndexByte(b, ',')
	for next != -1 {
		res[utils.ToInt(string(b[i:i+next]))] = 1
		i += next + 1
		next = bytes.IndexByte(b[i:], ',')
	}
	res[utils.ToInt(string(b[i:]))] = 1

	return res
}

func parseJoltage(b []byte) []int {
	var (
		nums  []int
		i     int
		split [][]byte
	)
	i = bytes.IndexByte(b, '{')
	split = bytes.Split(b[i+1:len(b)-1], []byte{','})
	nums = make([]int, len(split))
	for i = range nums {
		nums[i] = utils.ToInt(string(split[i]))
	}
	return nums
}

func transpose[T utils.Number](m [][]T) [][]T {
	r := make([][]T, len(m[0]))
	for i := range r {
		r[i] = make([]T, len(m))
	}
	for i := range m {
		for j := range r {
			r[j][i] = m[i][j]
		}
	}
	return r
}

func swapRows[T utils.Number](m [][]T, col int) [][]T {
	firstCol := firstNonzero(m[col])
	if firstCol == col {
		return m
	}

	maxCol := firstCol
	maxRow := col
	for i := col + 1; i < len(m); i++ {
		t := firstNonzero(m[i])
		if t == col {
			maxRow = i
			break
		} else if t < maxCol {
			maxCol = t
			maxRow = i
		}
	}
	if maxRow == len(m) {
		return m
	}
	m[col], m[maxRow] = m[maxRow], m[col]
	return m
}

func eliminate[T utils.Number](m [][]T, beg, idx int) {
	for i := beg + 1; i < len(m); i++ {

		if m[i][idx] == 0 {
			continue
		}
		ratio := m[i][idx] / m[beg][idx]
		for j := range m[i] {
			m[i][j] -= ratio * m[beg][j]
		}
	}
}

func firstNonzero[T utils.Number](src []T) int {
	var i int
	for i < len(src) && (src[i] == 0 || utils.AbsT(float64(src[i])) < .0001) {
		i++
	}
	return i
}

func minimize[T utils.Number](m [][]T) [][]T {
	n := make([][]T, 0, len(m))
	for row := range m {
		zc := 0
		for col := range m[row][:len(m[row])-1] {
			zc += b2i(m[row][col] != 0)
		}
		if zc == 0 || zc == 1 && m[row][len(m[row])-1] == 0 {
			continue
		}
		n = append(n, make([]T, len(m[row])))
		copy(n[len(n)-1], m[row])
	}
	return n
}

func rowReduce[T utils.Number](m [][]T) [][]T {
	for i := range len(m) {
		if i >= len(m[i]) {
			break
		}
		m = swapRows(m, i)
		k := firstNonzero(m[i])

		if k >= len(m[i])-1 {
			break
		}

		t := m[i][k]
		for j := 0; j < len(m[i]); j++ {
			m[i][j] /= t
		}
		eliminate(m, i, k)
	}

	for i := len(m) - 1; i > -1; i-- {
		k := firstNonzero(m[i])
		if k >= len(m[i])-1 {
			continue
		}
		t := m[i][k]

		for j := i - 1; j > -1; j-- {
			ratio := m[j][k] / t
			for h := k; h < len(m[i]); h++ {
				m[j][h] -= ratio * m[i][h]
			}
		}
	}

	return m
}

func emptyOrFree[T utils.Number](m [][]T) ([]int, []int) {
	var empty, free []int
	for col := range m[0][:len(m[0])-1] {
		count := 0
		for row := range m {
			count += b2i(m[row][col] != 0)
		}
		if count == 0 {
			empty = append(empty, col)
		} else if count > 1 {
			free = append(free, col)
		}
	}
	return empty, free
}

func sumFinal[T utils.Number](m [][]T) T {
	var n T
	for row := range m {
		n += m[row][len(m[row])-1]
	}
	return n
}

func evaluateRow[T utils.Number](row []T, freeVals []fv[T]) T {
	var res T
	i := firstNonzero(row)
	if i == len(row) {
		return 0
	}
	res = row[len(row)-1]
	for j := i + 1; j < len(row)-1; j++ {
		res -= row[j] * freeVals[j].val
	}
	res /= row[i]

	return res
}

type fv[T utils.Number] struct {
	val   T
	upper T
}

func paramDFS[T utils.Number](m [][]T, freeVals []fv[T], freeIndices []bool, i int) (T, bool) {
	var res, t T
	res = 512000
	var ok, okt bool
	if i == len(freeVals) {
		res = 0
		for row := range m {
			t = evaluateRow(m[row], freeVals)
			if t < 0 {
				return 512000, false
			}
			f0 := float64(t)
			f1 := float64(int(t))
			if f0 != f1 {
				if utils.AbsT(f0-f1) > 0.1 {
					return 512000, false
				}
			}
			res += t
		}
		for _, v := range freeVals {
			res += v.val
		}
		return res, true
	}
	if !freeIndices[i] {
		return paramDFS(m, freeVals, freeIndices, i+1)
	}

	for p := T(0); p <= freeVals[i].upper; p++ {
		t, okt = paramDFS(m, freeVals, freeIndices, i+1)
		if okt {
			res = min(res, t)
			ok = true
		}
		freeVals[i].val++
	}
	freeVals[i].val = 0
	return res, ok
}

func upperBound[T utils.Number](row []T, cur T) T {
	all := true
	for i := range row {
		if row[i] < 0 {
			all = false
			break
		}
	}
	if all {
		if cur == 0 {
			return row[len(row)-1]
		}
		return min(cur, row[len(row)-1])
	}
	if cur == 0 {
		return 4000
	}
	return cur
}

func solveFree[T utils.Number](m [][]T, cols []int) T {
	for row := range m {
		for {
			t := getSmallestDecimal(m[row])
			if t == 1 || t == 0 {
				break
			}
			t = 1 / t
			for col := range m[row] {
				m[row][col] *= t
			}
		}
	}

	freeVals := make([]fv[T], len(m[0]))

	freeCols := make([]bool, len(m[0]))
	for col := range cols {
		freeCols[cols[col]] = true
		for row := range m {
			if m[row][cols[col]] != 0 {
				freeVals[cols[col]].upper = upperBound(m[row], freeVals[cols[col]].upper)
			}
		}
	}

	res, ok := paramDFS(m, freeVals, freeCols, 0)
	if !ok {
		panic(ok)
	}

	return res
}

func getSmallestDecimal[T utils.Number](row []T) T {
	res := T(1)

	for i := range row {
		f0 := float64(row[i])
		f1 := float64(int(row[i]))
		if f0 == f1 {
			continue
		}
		if utils.AbsT(f0-f1) < 0.00001 {
			row[i] = T(f1)
			continue
		}
		res = min(res, utils.AbsT(T(f0-f1)))
	}

	return res
}

func part1(input string) int {
	var sum int = 0
	for _, line := range utils.ReadLines(input) {
		processLine(line, &sum)
	}
	return sum
}

func part2(input string) int {
	var (
		res, i   int
		b        []byte
		pieces   [][]byte
		r        = bytes.NewReader([]byte(input))
		scanner  = bufio.NewScanner(r)
		buttons  = make([][]int, 0, 64)
		joltages []int
	)

	for scanner.Scan() {
		b = scanner.Bytes()
		if len(b) == 0 {
			continue
		}
		pieces = bytes.Split(b, []byte{' '})

		buttons = make([][]int, 0, 12)
		for i = 1; i < len(pieces)-1; i++ {
			buttons = append(buttons, parseButtonRaw(pieces[i], len(pieces[0])-2))
		}
		joltages = parseJoltage(pieces[len(pieces)-1])

		buttons = append(buttons, joltages)
		b2 := make([][]float64, len(buttons))
		for row := range buttons {
			b2[row] = make([]float64, len(buttons[row]))
			for col := range buttons[row] {
				b2[row][col] = float64(buttons[row][col])
			}
		}
		b2 = transpose(b2)
		rowReduce(b2)
		b2 = minimize(b2)
		_, free := emptyOrFree(b2)
		if len(free) > 0 {
			ans := solveFree(b2, free)
			res += int(ans)
		} else {
			res += int(sumFinal(b2))
		}

	}

	return res
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part to run? 1 or 2")
	flag.Parse()
	res := 0
	if part == 1 {
		res = part1(input)
	} else {
		res = part2(input)
	}
	fmt.Println("Result: ", res)
}
