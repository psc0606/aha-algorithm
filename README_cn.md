<a title="Hit" target="_blank" href="https://github.com/psc0606/aha-algorithm"><img src="https://hits.b3log.org/psc0606/aha-algorithm.svg"></a>
[![Build Status](https://travis-ci.com/psc0606/aha-algorithm.svg?branch=main)](https://travis-ci.com/psc0606/aha-algorithm)
[![codecov](https://codecov.io/gh/psc0606/aha-algorithm/branch/main/graph/badge.svg)](https://codecov.io/gh/psc0606/aha-algorithm)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fpsc0606%2Faha-algorithm.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fpsc0606%2Faha-algorithm?ref=badge_shield)

# aha-algorithm
啊哈, 算法 golang 实现. 英文版本 [README.md](./README.md)

# 目录
## 字符串匹配算法
* 多模匹配算法 Wu-Manber(WM) 算法
* 单模 Shift-And 算法

## 排序算法
* [插入排序](src/sort/insert-sort.go)
* [冒泡排序](src/sort/bubble-sort.go)
* [选择排序](src/sort/select-sort.go)
* [希尔排序](src/sort/shell-sort.go)
* [快速排序](src/sort/quick-sort.go)
* [归并排序](src/sort/merge-sort.go)
* 堆排序
* [计数排序](src/sort/counting-sort.go)
* 基数排序
* 桶排序
* ...

## 高级数据结构
* [双向链表](src/list/double-linked-list.go)
* [并查集](src/tree/disjoint-set.go)
* [堆](src/heap/heap.go)
* 树堆 (TODO)
* 跳跃表 (TODO)
* 红黑树 (TODO)
* [Trie树](src/tree/trie.go)
* Double Array Tire (TODO)
* B-树 (TODO)
* B+树 (TODO)
* B*树 (TODO)  
* 线段树 (TODO)
* 三叉搜索树(Ternary Search Tree) (TODO)
* 有限状态转换机(FST, Finite State Transducer) (TODO)
* B* (Branch Star分支寻路算法) (TODO)
* ...

## 地域结构
* GeoHash (TODO)
* 墨卡托 (TODO)
* ...

## 概率型数据结构
* HyperLogLog (TODO)
* 布隆过滤器 (TODO)
* Count-Min Sketch (TODO)
* ...

## 缓存
* [最近最少使用缓存淘汰](src/cache/lru.go)
* LRU with timestamp (TODO)
* LFU (TODO)
* W-TintyLFU (TODO)
* ...

## 选举算法
* [摩尔投票法](src/election/majority-element.go) - [论文链接](paper/A%20Fast%20Majority%20Vote%20Algorithm.pdf)

## 趣味算法
**[E]** means **简单难度**, **[M]** means **中等难度**, **[H]** means **困难**.
* [汉明距离](src/bit/hamming-distance.go) [E]
* [比特位计数](src/bit/counting-bits.go) [M]
* [斐波那契数列](src/array/fibonacci_number.go) [E]
* [两数之和](src/array/twosum.go) [E]
* [三数之和](src/array/threesum.go) [M]
* [买卖股票的最佳时机](src/optimalization/best-time-to-buy-and-sell-stock.go) [E]
* [最佳买卖股票时机含冷冻期](src/optimalization/best-time-to-buy-and-sell-stock-with-cooldown.go) [M]  
* [任务调度](src/optimalization/task-scheduler.go) [M]
* [零钱兑换](src/optimalization/coin-change.go) [M]
* [完全平方数](src/optimalization/perfect-squares.go) [M]
* [跳跃游戏](src/optimalization/jump-game.go) [M]
* [乘积最大子数组](src/optimalization/maximum-product-subarray.go) [M]  
* [不同的二叉搜索树](src/tree/unique-binary-search-trees.go) [M]
* [前k个高频元素](src/array/top-k-frequent-elements.go) [M]
* [单词拆分](src/string/word-break.go) [M]
* [树层序遍历](src/tree/binary-tree-level-order-traversal.go) [M]
* [二叉树中序遍历(非递归)](src/tree/binary-tree-inorder-traversal.go) [M]  
* [旋转有序数组搜索](src/array/search-in-rotated-sorted-array.go) [M]
* [二叉树展开成单链表](src/tree/flatten-binary-tree-to-linked-list.go) [M]
* [最近公共祖先](src/tree/lowest-common-ancestor-of-a-binary-tree.go) [M]
* [删除链表的倒数第N个结点](src/list/remove-nth-node-from-end-of-list.go) [M]
* [环形链表 II](src/list/linked-list-cycle-ii.go) [M]
* [电话号码字母组合](src/backtrace/letter-combinations-of-a-phone-number.go) [M]
* [组合总和](src/backtrace/combination-sum.go) [M]  
* [生成括号](src/backtrace/generate-parentheses.go) [M]
* [全排列](src/backtrace/permutations.go) [M]

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fpsc0606%2Faha-algorithm.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fpsc0606%2Faha-algorithm?ref=badge_large)
