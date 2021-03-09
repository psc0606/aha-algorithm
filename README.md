<a title="Hit" target="_blank" href="https://github.com/psc0606/aha-algorithm"><img src="https://hits.b3log.org/psc0606/aha-algorithm.svg"></a>
[![Build Status](https://travis-ci.com/psc0606/aha-algorithm.svg?branch=main)](https://travis-ci.com/psc0606/aha-algorithm)
[![codecov](https://codecov.io/gh/psc0606/aha-algorithm/branch/main/graph/badge.svg)](https://codecov.io/gh/psc0606/aha-algorithm)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fpsc0606%2Faha-algorithm.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fpsc0606%2Faha-algorithm?ref=badge_shield)

# aha-algorithm
Aha, algorithm by go. See Chinese [README.md](./README_cn.md)

# Table
## String Matching
* Wu-Manber(WM) Algorithm
* Shift-And Algorithm

## Sort Algorithm
* [Insertion Sort](src/sort/insert-sort.go)
* [Bubble Sort](src/sort/bubble-sort.go)
* [Selection Sort](src/sort/select-sort.go)  
* [Shell Sort](src/sort/shell-sort.go)
* [Quick Sort](src/sort/quick-sort.go)
* [Merge Sort](src/sort/merge-sort.go)
* Heap Sort
* [Counting Sort](src/sort/counting-sort.go)
* Radix Sort
* Bucket Sort
* ...

## Advanced Structure
* (Double Linked List)(src/list/double-linked-list.go)
* [Disjoint Set (Union Find Set)](src/tree/disjoint-set.go)
* [Heap](src/heap/heap.go)
* Treap (TODO)
* SkipList (TODO)
* RB-Tree (TODO)
* [Trie](src/tree/trie.go)
* Double Array Tire (TODO)
* B-Tree (TODO)
* B+Tree (TODO)
* B*Tree (TODO)
* R Tree (TODO)
* Segment Tree (TODO)
* Ternary Search Tree (TODO)
* FST (TODO)
* B* (Branch Star) (TODO)
* ...

## Geo
* GeoHash (TODO)
* Mercator (TODO)
* ...

## Probabilistic Data Structure
* HyperLogLog (TODO)
* Bloom Filter (TODO)
* Count-Min Sketch (TODO)
* ...

## Cache
* [LRU cache](src/cache/lru.go)
* LRU with timestamp (TODO)
* LFU (TODO)
* W-TintyLFU (TODO)
* ...

## Election Algorithm
*  [Boyer–Moore majority vote algorithm](src/election/majority-element.go) - [paper link](paper/A%20Fast%20Majority%20Vote%20Algorithm.pdf)

## Funny Problem
**[E]** means **EASY**, **[M]** means **MEDIUM**, **[H]** means **HARD**, **[I]** means **INTERVIEW**

---
### Bits Operation
* [Hamming Distance](src/bit/hamming-distance.go) [E]
* [Counting Bits](src/bit/counting-bits.go) [M]
* [Subsets](src/backtrace/subsets.go) [M]

---
### Array
* [Remove Duplicates From Sorted Array](src/array/remove_duplicates.go) [E] [I]
* [Move Zero](src/array/move-zero.go) [E] [I]
* [Plus One](src/array/plus-one.go) [E] [I]
* [Merge Sorted Array](src/array/merge-sorted-array.go) [E] [I]  
* [Fibonacci Number](src/array/fibonacci_number.go) [E]
* [Two Sum](src/array/twosum.go) [E] [I]
* [Three Sum](src/array/threesum.go) [M] [I]
* [Top-k frequent elements](src/array/top-k-frequent-elements.go) [M]
* [Search In Rotated Sorted Array](src/array/search-in-rotated-sorted-array.go) [M]
* [Find The Duplicate Number](src/array/find-the-duplicate-number.go) [M]
* [Container With Most Water](src/array/container-with-most-water.go) [M] [I]
* [Rotate Array](src/array/rotate-array.go) [M] [I]
* [Group Anagrams](src/array/group-anagrams.go) [M] [I]
* [K-th Largest Element In An Array](src/array/kth-largest-element-in-an-array.go) [M] [I]

---
### List
* [Merge Two Sorted Lists](src/list/merge-two-sorted-lists.go) [M] [I]
* [Remove n-th Node From End Of List](src/list/remove-nth-node-from-end-of-list.go) [M]
* [Linked List Cycle](src/list/linked-list-cycle.go) [E] [I]
* [Linked List Cycle II](src/list/linked-list-cycle-ii.go) [M] [I]
* [Swap Nodes In Pairs](src/list/swap-nodes-in-pairs.go) [M] [I]

---
### Stack
* [Minimum Stack](src/stack/min-stack.go) [E] [I]
* [Valid Parentheses](src/stack/valid-parentheses.go) [E] [I]
* [Decode String](src/stack/decode-string.go) [M]

---
### Dynamic Programming And Greedy
* [Climbing Stairs](src/optimalization/climbing-stairs.go) [E] [I]
* [Best Time To Buy And Sell Stock](src/optimalization/best-time-to-buy-and-sell-stock.go) [E]
* [Best Time To Buy And Sell Stock With Cool Down](src/optimalization/best-time-to-buy-and-sell-stock-with-cooldown.go) [M]  
* [Task scheduler](src/optimalization/task-scheduler.go) [M]
* [Jump Game](src/optimalization/jump-game.go) [M]
* [Maximum Product Subarray](src/optimalization/maximum-product-subarray.go) [M]
* [word break](src/string/word-break.go) [M]

---
#### Knapsack problem
* [Coin Change](src/optimalization/coin-change.go) [M]
* [Perfect Squares](src/optimalization/perfect-squares.go) [M]

---
### Tree
* [Unique binary search trees](src/tree/unique-binary-search-trees.go) [M]
* [Tree Level Traversal](src/tree/binary-tree-level-order-traversal.go) [M]
* [Tree Inorder Traversal(non-recursive)](src/tree/binary-tree-inorder-traversal.go) [M] [I]
* [Flatten Binary Tree To Linked List](src/tree/flatten-binary-tree-to-linked-list.go) [M]
* [LCA](src/tree/lowest-common-ancestor-of-a-binary-tree.go) [M]

---
### Backtrace
* [Letter Combinations Of A Phone Number](src/backtrace/letter-combinations-of-a-phone-number.go) [M]
* [Generate Parenthesis](src/backtrace/generate-parentheses.go) [M] [I]
* [Combination Sum](src/backtrace/combination-sum.go) [M]
* [Permute](src/backtrace/permutations.go) [M]
* [Number of Islands](src/backtrace/number-of-islands.go) [M]
* [Subsets](src/backtrace/subsets.go) [M]

### Disjoint Set (Union Find Set)
* [Accounts merge](src/tree/accounts-merge.go) [M]
* [Number of Islands](src/backtrace/number-of-islands.go) [M]

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fpsc0606%2Faha-algorithm.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fpsc0606%2Faha-algorithm?ref=badge_large)
