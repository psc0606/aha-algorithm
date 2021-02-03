package tree

// Disjoint set is a very useful tree data structure but organized by array.
// https://en.wikipedia.org/wiki/Disjoint-set_data_structure
// Time: see wikipedia
type DisjointSet struct {
	// the max number of elements.
	n int
	// father[x] = y means element x's parent is y.
	father []int
	// rank[x] means the depth or level of the element node.
	rank []int
}

// Initialize n node without any edges.
func (d *DisjointSet) Init(n int) {
	d.n = n
	d.father = make([]int, n)
	d.rank = make([]int, n)
}

func (d *DisjointSet) MakeSet(x int) {
	d.father[x] = x
	d.rank[x] = 1
}

// To find the root element of the tree.
// return the root parent of its set.
func (d *DisjointSet) Find(x int) int {
	if x == d.father[x] {
		return x
	}
	// no optimized
	// return d.Find(d.father[x])

	// optimized-1: use path compress to balance the tree.
	// make the x as direct child of root father.
	d.father[x] = d.Find(d.father[x])
	return d.father[x]
}

// Make sure the inputted `x` and `y` are legal.
// `x` - Whether x is in the same set with y.
// `y` - Whether y is in the same set with x.
func (d *DisjointSet) Same(x int, y int) bool {
	return d.Find(x) == d.Find(y)
}

// Make sure the inputted `x` and `y` are legal.
// `x` - Union the two set
// `y` - Union the two set
func (d *DisjointSet) Union(x int, y int) {
	// root parent of x
	px := d.Find(x)
	// root parent of y
	py := d.Find(y)
	if px != py {
		// how to determine x is parent of y, or y is parent of x ?
		// which is better?
		// d.father[px] = py
		if d.rank[px] < d.rank[py] {
			d.father[px] = py
		} else {
			d.father[py] = px

			// equals should increment depth, or no need.
			if d.rank[px] == d.rank[py] {
				d.rank[py]++
			}
		}
	}
}
