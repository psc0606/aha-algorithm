package list

// A dummy head and tail help to simplify the operations on double linked list.
type DLinkedList struct {
	size       int
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	value      interface{}
	prev, next *DLinkedNode
}

func initDLinkedNode(value interface{}) *DLinkedNode {
	return &DLinkedNode{value: value}
}

func Constructor() *DLinkedList {
	dl := &DLinkedList{
		size: 0,
		head: initDLinkedNode(nil), // A dummy head
		tail: initDLinkedNode(nil), // A dummy tail
	}
	dl.head.next = dl.tail
	dl.tail.prev = dl.head
	return dl
}

func (this *DLinkedList) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *DLinkedList) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *DLinkedList) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *DLinkedList) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
