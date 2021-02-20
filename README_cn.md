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
* 归并排序
* 堆排序
* ...

## 高级数据结构
* [双向链表](src/list/double-linked-list.go)
* [并查集](src/tree/disjoint-set.go)
* [堆](src/heap/heap.go)
* 跳跃表 (TODO)
* [Trie树](src/tree/trie.go)
* Double Array Tire (TODO)
* B-树 (TODO)
* B+树 (TODO)
* 线段树 (TODO)
* 三叉搜索树(Ternary Search Tree) (TODO)
* 有限状态转换机(FST, Finite State Transducer) (TODO)
* 红黑树 (TODO)
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
* [斐波那契数列](src/array/fibonacci_number.go)
* [任务调度](src/optimalization/task-scheduler.go)
* [不同的二叉搜索树](src/tree/unique-binary-search-trees.go)
* [前k个高频元素](src/array/top-k-frequent-elements.go)
* [单词拆分](src/string/word-break.go)

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fpsc0606%2Faha-algorithm.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fpsc0606%2Faha-algorithm?ref=badge_large)
