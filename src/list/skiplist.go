package list

import (
	"fmt"
	"math/rand"
	"time"
)

// optional config
const (
	branching = 4
	maxLevel  = 32
)

// Skiplist https://leetcode-cn.com/problems/design-skiplist/
type Skiplist struct {
	// current level of skiplist, no maximum limited.
	level int
	// head val not store any valid data.
	head *SkiplistNode
}

type SkiplistNode struct {
	val  int
	next []*SkiplistNode
}

// ConstructorSkiplist `Constructor` duplicate with other package, so change to ConstructorSkiplist
func ConstructorSkiplist() Skiplist {
	rand.Seed(time.Now().UnixNano())
	return Skiplist{head: new(SkiplistNode)}
}

// Search search the elements.
func (this *Skiplist) Search(target int) bool {
	// empty skiplist
	if this.level < 1 {
		return false
	}
	node := this.head
	for i := this.level - 1; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].val < target {
			node = node.next[i]
		}
	}
	node = node.next[0]
	if node != nil && node.val == target {
		return true
	}
	return false
}

func (this *Skiplist) Add(num int) {
	node := this.head
	prev := make([]*SkiplistNode, this.level)
	for i := this.level - 1; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].val < num {
			node = node.next[i]
		}
		prev[i] = node
	}
	newNode := &SkiplistNode{val: num}
	randomLevel := this.RandomLevel()
	newNode.next = make([]*SkiplistNode, randomLevel)
	for i := 0; i < randomLevel; i++ {
		// more higher level
		if i+1 > this.level {
			this.head.next = append(this.head.next, newNode)
		} else {
			tmp := prev[i].next[i]
			prev[i].next[i] = newNode
			newNode.next[i] = tmp
		}
	}
	if this.level < randomLevel {
		this.level = randomLevel
	}
}

func (this *Skiplist) Erase(num int) bool {
	// empty skiplist
	if this.level < 1 {
		return false
	}
	node := this.head
	prev := make([]*SkiplistNode, this.level)
	for i := this.level - 1; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].val < num {
			node = node.next[i]
		}
		prev[i] = node
	}
	node = node.next[0]
	if node != nil && node.val == num {
		for i := 0; i < len(node.next); i++ {
			prev[i].next[i] = node.next[i]
		}

		for this.level > 1 && this.head.next[this.level-1] == nil {
			this.level--
		}
		this.head.next = this.head.next[0:this.level]
		return true
	}
	return false
}

func (this *Skiplist) print() {
	for i := this.level - 1; i >= 0; i-- {
		node := this.head.next[i]
		fmt.Printf("Level: %d - ", i)
		for node != nil {
			fmt.Printf("%d ", node.val)
			node = node.next[i]
		}
		fmt.Println()
	}
	fmt.Println()
}

// RandomLevel return the random level.
func (this *Skiplist) RandomLevel() int {
	level := 1
	// for rand.Intn(branching) == 0 {
	for level < maxLevel && rand.Intn(branching) == 0 {
		level++
	}
	return level
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
