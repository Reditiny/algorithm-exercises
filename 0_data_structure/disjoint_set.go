package ds

type DisjointSet struct {
	parent map[int]int
	rank   map[int]int
}

func NewDisjointSet() *DisjointSet {
	return &DisjointSet{
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}
}

func (d *DisjointSet) MakeSet(x int) {
	d.parent[x] = x
	d.rank[x] = 0
}

func (d *DisjointSet) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DisjointSet) Union(x, y int) {
	rootX, rootY := d.Find(x), d.Find(y)
	if rootX == rootY {
		return
	}
	if d.rank[rootX] > d.rank[rootY] {
		d.parent[rootY] = rootX
	} else {
		d.parent[rootX] = rootY
		if d.rank[rootX] == d.rank[rootY] {
			d.rank[rootY]++
		}
	}
}

func (d *DisjointSet) IsSameSet(x, y int) bool {
	return d.Find(x) == d.Find(y)
}
