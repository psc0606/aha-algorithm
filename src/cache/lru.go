package cache

// https://leetcode-cn.com/problems/lru-cache/
// Implement by double linked list.
type LRUCache struct {
	capacity int
	size     int
	ht       map[int]*LRUCacheNode
	head     *LRUCacheNode
	tail     *LRUCacheNode
}

type LRUCacheNode struct {
	key  int
	vale int
	pre  *LRUCacheNode
	next *LRUCacheNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		ht:       make(map[int]*LRUCacheNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.ht[key]
	if ok {
		if v != this.head && v != this.tail {
			// move `v` to head
			v.pre.next, v.next.pre = v.next, v.pre

			// make `v` as head
			v.next = this.head
			this.head.pre = v
			this.head = v
		} else if v != this.head {
			// `v` is the tail.
			this.tail = this.tail.pre
			this.tail.next = nil

			// make `v` as head
			v.next = this.head
			v.pre = nil
			this.head.pre = v
			this.head = this.head.pre
		} else {
			// if v is the head, do nothing.
			// do nothing
		}
		return v.vale
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity < 1 {
		return
	}
	v := this.Get(key)
	if v != -1 {
		// that means the key exists, then replace the key
		this.ht[key].vale = value
		return
	}
	newNode := &LRUCacheNode{key: key, vale: value}
	this.ht[key] = newNode
	if this.size == 0 {
		// first element
		this.head = newNode
		this.tail = newNode
		this.size++
		return
	}
	newNode.next = this.head
	this.head.pre = newNode
	this.head = this.head.pre
	this.size++

	// pruning
	if this.size > this.capacity {
		delete(this.ht, this.tail.key)
		this.tail = this.tail.pre
		this.tail.next = nil
		this.size--
	}
}

func (this *LRUCache) Size() int {
	return this.size
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
