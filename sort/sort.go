package sort

type Pair struct {
	key   int
	value int
}

type PairSlice []Pair

func (p PairSlice) Len() int {
	return len(p)
}

func (p PairSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PairSlice) Less(i, j int) bool {
	return p[i].key < p[j].key
}
