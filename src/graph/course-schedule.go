package graph

import "aha-algorithm/src/util"

// https://leetcode-cn.com/problems/course-schedule/
// Graph Topology Sort + bfs
func canFinish(numCourses int, prerequisites [][]int) bool {
	// edge: from current to another
	edges := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, pair := range prerequisites {
		from, to := pair[1], pair[0]
		arr := edges[from] // pair is <post course, pre course>
		arr = append(arr, to)
		edges[pair[1]] = arr
		indegree[pair[0]]++ // every in-degree of node
	}

	// init add all nodes whose indegree is zero to the queue.
	var queue []int
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// bfs
	visited := 0
	for len(queue) > 0 {
		visited++
		course := queue[0]
		queue = queue[1:]
		for _, c := range edges[course] {
			indegree[c]--
			if indegree[c] == 0 {
				queue = append(queue, c)
			}
		}
	}
	return visited == numCourses
}

// Graph Topology Sort + dfs
func canFinishByDfs(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, pair := range prerequisites {
		from, to := pair[1], pair[0]
		arr := edges[from] // pair is <post course, pre course>
		arr = append(arr, to)
		edges[pair[1]] = arr
		indegree[pair[0]]++ // every in-degree of node
	}

	// init add all nodes whose indegree is zero to the queue.
	var stack []int
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			stack = append(stack, i)
		}
	}

	// dfs
	visited := make([]int, numCourses)
	for len(stack) > 0 {
		course := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[course] = 1

		for _, c := range edges[course] {
			indegree[c]--
			if indegree[c] == 0 && visited[c] != 1 {
				stack = append(stack, c)
			}
		}
	}
	return util.ArraySum(visited) == numCourses
}
