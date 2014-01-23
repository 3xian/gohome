package container

type SegTree struct {
	val []int
	n   int
}

func NewSegTree(n int) *SegTree {
	size := calcTreeSize(n)
	return &SegTree{
		val: make([]int, size),
		n:   n,
	}
}

func (s *SegTree) Add(l int, r int, val int) {
	s.add(1, 0, s.n-1, l, r, val)
}

func (s *SegTree) Sum(i int) int {
	return s.sum(1, 0, s.n-1, i)
}

func (s *SegTree) add(i int, l int, r int, ql int, qr int, qv int) {
	if l >= ql && r <= qr {
		s.val[i] += qv
		return
	}
	m := (l + r) / 2
	if qr <= m {
		s.add(i*2, l, m, ql, qr, qv)
	} else if ql > m {
		s.add(i*2+1, m+1, r, ql, qr, qv)
	} else {
		s.add(i*2, l, m, ql, m, qv)
		s.add(i*2+1, m+1, r, m+1, qr, qv)
	}
}

func (s *SegTree) sum(i int, l int, r int, qi int) int {
	if l == qi && r == qi {
		return s.val[i]
	}
	m := (l + r) / 2
	if qi <= m {
		return s.sum(i*2, l, m, qi) + s.val[i]
	} else {
		return s.sum(i*2+1, m+1, r, qi) + s.val[i]
	}
}

func calcTreeSize(length int) int {
	size := 1
	for size < length {
		size *= 2
	}
	return size * 2
}
