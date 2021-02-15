package tree

// https://leetcode-cn.com/problems/implement-trie-prefix-tree/
// Use fixed 26 alpha, may not be perfect. you can use dynamic structure instead.
const R int = 26

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	links [R]*TrieNode
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: &TrieNode{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root
	for _, c := range word {
		if node.links[c-'a'] == nil {
			node.links[c-'a'] = new(TrieNode)
		}
		node = node.links[c-'a']
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.root
	for _, c := range word {
		if node.links[c-'a'] == nil {
			return false
		}
		node = node.links[c-'a']
	}
	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, c := range prefix {
		if node.links[c-'a'] == nil {
			return false
		}
		node = node.links[c-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
