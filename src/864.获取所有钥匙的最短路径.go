package main

/*
 * @lc app=leetcode.cn id=864 lang=golang
 *
 * [864] 获取所有钥匙的最短路径
 */
// dynamic-programming|breadth-first-search

import "fmt"

// @lc code=start
type Element interface{}

type Queue struct {
	sli []Element
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Size() int {
	return len(q.sli)
}

func (q *Queue) Push(e Element) {
	q.sli = append(q.sli, e)
}

func (q *Queue) Pop() Element {
	if q.Size() == 0 {
		return nil
	}
	e := q.sli[0]
	q.sli = q.sli[1:]
	return e
}

func shortestPathAllKeys(grid []string) int {
	mp := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	m, n := len(grid), len(grid[0])
	// 1. 初始化染色数组
	vis := make([][][]int, m)
	for i := 0; i < len(vis); i++ {
		vis[i] = make([][]int, n)
		for j := 0; j < len(vis[i]); j++ {
			vis[i][j] = make([]int, 64)
		}
	}
	// 2. 找到起点，计算钥匙数
	si, sj, keycnt := 0, 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] >= 'a' && grid[i][j] <= 'f' {
				keycnt++
			} else if grid[i][j] == '@' {
				si, sj = i, j
			}
		}
	}
	keys := 0
	for i := 0; i < keycnt; i++ {
		keys |= 1 << i
	}
	// 3. BFS
	vis[si][sj][0] = 1 // value表示当前的步数
	q := NewQueue()
	q.Push([]int{si, sj, 0})
	for q.Size() != 0 {
		e := q.Pop().([]int)
		i, j, state := e[0], e[1], e[2] // state表示手中的钥匙数
		step := vis[i][j][state]
		if state == keys {
			return step - 1
		}
		for k := 0; k < 4; k++ {
			x, y := i+mp[k][0], j+mp[k][1]
			if x < 0 || y < 0 || x >= m || y >= n {
				continue
			}
			if grid[x][y] == '.' || grid[x][y] == '@' ||
				(grid[x][y] >= 'a' && grid[x][y] <= 'f') ||
				(grid[x][y] >= 'A' && grid[x][y] <= 'F' && state&(1<<(grid[x][y]-'A')) != 0) {
				state_new := state
				if grid[x][y] >= 'a' && grid[x][y] <= 'f' {
					state_new |= 1 << (grid[x][y] - 'a')
				}
				if vis[x][y][state_new] == 0 {
					q.Push([]int{x, y, state_new})
					// fmt.Printf("(%d, %d) -> (%d, %d): %d %d\n", i, j, x, y, state_new, step+1)
					vis[x][y][state_new] = step + 1
				}
			}
		}
	}
	return -1
}

// @lc code=end
func main() {
	fmt.Println(shortestPathAllKeys([]string{"@.a.#", "###.#", "b.A.B"}))
}
