package main

import (
	"fmt"
	"strconv"
)

func solve(R, C, RR, CC int, insts string) (int, int) {
	instMap := map[byte][]int{'E': {0, 1}, 'W': {0, -1}, 'N': {-1, 0}, 'S': {1, 0}}
	rowItvs, colItvs := make([]map[string][]int, R+1), make([]map[string][]int, C+1)
	for i := 1; i <= R; i++ {
		rowItvs[i] = make(map[string][]int)
	}
	for i := 1; i <= C; i++ {
		colItvs[i] = make(map[string][]int)
	}
	rowItvs[RR], colItvs[CC] = map[string][]int{str(RR, CC): {CC, CC}}, map[string][]int{str(RR, CC): {RR, RR}}
	currX, currY := RR, CC
	for i := 0; i < len(insts); i++ {
		move := instMap[insts[i]]
		nextX, nextY := currX+move[0], currY+move[1]
		rowItv, colItv := findRowItv(rowItvs, nextX, nextY), findColItv(colItvs, nextX, nextY)
		if rowItv == nil { // colItv must be zero in this condition. Move to <nextX, nextY>
			currX, currY = nextX, nextY
		} else {
			if insts[i] == 'E' {
				currX, currY = nextX, rowItv[1]+1
			} else if insts[i] == 'W' {
				currX, currY = nextX, rowItv[0]-1
			} else if insts[i] == 'N' {
				currX, currY = colItv[0]-1, nextY
			} else {
				currX, currY = colItv[1]+1, nextY
			}
		}
		updateItvs(rowItvs, colItvs, currX, currY)
	}
	return currX, currY
}

func updateItvs(rowItvs, colItvs []map[string][]int, currX, currY int) {
	var itv []int
	leftStr, rightStr := str(currX, currY-1), str(currX, currY+1)
	leftItv, rightItv := rowItvs[currX][leftStr], rowItvs[currX][rightStr]
	if leftItv != nil && rightItv != nil {
		itv = []int{leftItv[0], rightItv[1]}
	} else if leftItv != nil {
		itv = []int{leftItv[0], currY}
	} else if rightItv != nil {
		itv = []int{currY, rightItv[1]}
	} else {
		itv = []int{currY, currY}
	}
	// Must do it, even it's not on the edge of the interval. But it decide
	// whether the slot is visited.
	rowItvs[currX][str(currX, currY)] = itv

	rowItvs[currX][str(currX, itv[0])] = itv
	rowItvs[currX][str(currX, itv[1])] = itv

	upStr, downStr := str(currX-1, currY), str(currX+1, currY)
	upItv, downItv := colItvs[currY][upStr], colItvs[currY][downStr]
	if upItv != nil && downItv != nil {
		itv = []int{upItv[0], downItv[1]}
	} else if upItv != nil {
		itv = []int{upItv[0], currX}
	} else if downItv != nil {
		itv = []int{currX, downItv[1]}
	} else {
		itv = []int{currX, currX}
	}
	colItvs[currY][str(currX, currY)] = itv
	colItvs[currY][str(itv[0], currY)] = itv
	colItvs[currY][str(itv[1], currY)] = itv

}

func findRowItv(rowItvs []map[string][]int, x, y int) []int {
	return rowItvs[x][str(x, y)]
}

func findColItv(colItvs []map[string][]int, x, y int) []int {
	return colItvs[y][str(x, y)]
}

func str(x, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for ti := 1; ti <= t; ti++ {
		var N, R, C, RR, CC int
		fmt.Scanf("%d %d %d %d %d", &N, &R, &C, &RR, &CC)
		var insts string
		fmt.Scanf("%s", &insts)
		r, c := solve(R, C, RR, CC, insts)
		fmt.Printf("Case #%d: %d %d\n", ti, r, c)
	}
}
