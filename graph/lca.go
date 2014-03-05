package graph

const (
	ancLevel = 18
	maxN     = 1000000 + 10
)

var (
	anc [maxN][ancLevel]int
	dep [maxN]int
)

/*
	Build:

		dep[p] = dep[anc[p][0]] + 1
		for lv := 1; lv < ancLevel; lv++ {
			anc[p][lv] = anc[anc[p][lv-1]][lv-1]
		}
*/

func GetLca(x int, y int) int {
	a, b := x, y
	if dep[x] < dep[y] {
		a, b = b, a
	}

	for lv := ancLevel - 1; lv >= 0; lv-- {
		if dep[a]-(1<<uint(lv)) >= dep[b] {
			a = anc[a][lv]
		}
	}
	if a == b {
		return a
	}

	for lv := ancLevel - 1; lv >= 0; lv-- {
		if anc[a][lv] != anc[b][lv] {
			a = anc[a][lv]
			b = anc[b][lv]
		}
	}
	return anc[a][0]
}
