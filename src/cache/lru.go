package cache

// https://leetcode-cn.com/problems/lru-cache/
// Implement by double linked list.
// A better double-linked list implementing trick: use dummy head and dummy tail to avoid if-condition.
type LRUCache struct {
	capacity   int
	size       int
	ht         map[int]*LRUCacheNode
	head, tail *LRUCacheNode
}

type LRUCacheNode struct {
	key, value int
	prev, next *LRUCacheNode
}

func initLRUCacheNode(key, value int) *LRUCacheNode {
	return &LRUCacheNode{key: key, value: value}
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		ht:       make(map[int]*LRUCacheNode, capacity),
		head:     initLRUCacheNode(0, 0), // A dummy head
		tail:     initLRUCacheNode(0, 0), // A dummy tail
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.ht[key]; !ok {
		return -1
	}
	node := this.ht[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.ht[key]; !ok {
		node := initLRUCacheNode(key, value)
		this.ht[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.ht, removed.key)
			this.size--
		}
	} else {
		node := this.ht[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) Size() int {
	return this.size
}

func (this *LRUCache) addToHead(node *LRUCacheNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *LRUCacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *LRUCacheNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *LRUCacheNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
