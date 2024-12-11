package medium

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/number-of-islands/
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。
//
// 示例 1：
// 输入：grid = [
//
//	["1","1","1","1","0"],
//	["1","1","0","1","0"],
//	["1","1","0","0","0"],
//	["0","0","0","0","0"]
//
// ]
// 输出：1
// 示例 2：
// 输入：grid = [
//
//	["1","1","0","0","0"],
//	["1","1","0","0","0"],
//	["0","0","1","0","0"],
//	["0","0","0","1","1"]
//
// ]
// 输出：3

func TestNumIslands(t *testing.T) {
	g := [][]byte{{1, 0, 1, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 0}, {0, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1}, {1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0}, {0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 1, 1}, {1, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1}, {0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0}, {1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0}, {0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1}, {0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0}, {0, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1}, {0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1}, {1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0}, {0, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0}, {0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1}, {0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1}, {0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1}, {0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1}, {0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0}}
	fmt.Println(numIslands(g) == 23)
	fmt.Println(numIslands2(g) == 23)
}

// numIslands 个人思路总结 二维数组遍历，独立的岛屿用id标记, 当(x,y)的上一层(x-1,y)不存在时递增id并记mark[x][y]=id,
// 当(x,y)的上一层(x-1,y)存在时，与当前curMarkId比较，id不一致时取小的id，并记录id指向idMap[oid] = id
// 注意：！！比较(x-1,y)的preId时，该坐标的id可能已经指向其他id，所以需要preId = getTailValue，并且preId小于curMarkId时，preId也需要修改id指向idMap[preId] = curMarkId
func numIslands(grid [][]byte) int {
	xlen := len(grid)
	ylen := len(grid[0])
	markId := make([][]int, xlen)
	for i := 0; i < xlen; i++ {
		markId[i] = make([]int, ylen)
	}
	id := 0
	idMap := make(map[int]int)
	for x := 0; x < xlen; x++ {
		curMarkId := 0
		for y := 0; y < ylen; y++ {
			if grid[x][y] == 0 {
				curMarkId = 0
			} else if x == 0 {
				if curMarkId == 0 {
					id++
					idMap[id] = id
					curMarkId = id
				}
				markId[x][y] = curMarkId
			} else {
				preId := markId[x-1][y]
				if curMarkId == 0 && preId != 0 {
					curMarkId = getTailValue(idMap, preId)
				} else if preId != 0 {
					preVal := getTailValue(idMap, preId)
					// fix: preVal!=marked时，与下（坐标轴为左）层独立mark链接上，如下情况
					// ...
					// . .
					// .
					// ...
					if preVal > curMarkId {
						idMap[preId] = curMarkId
						idMap[preVal] = curMarkId
					} else if preVal < curMarkId {
						idMap[curMarkId] = preVal
						curMarkId = preVal
					}
				}
				if curMarkId == 0 {
					id++
					idMap[id] = id
					curMarkId = id
				}
				markId[x][y] = curMarkId
			}
		}
	}

	count := 0
	for k, v := range idMap {
		if k == v {
			count++
		}
	}

	return count
}

func getTailValue(idMap map[int]int, id int) int {
	for {
		nid := idMap[id]
		if nid == 0 || nid == id {
			return nid
		} else {
			id = nid
		}
	}
}

// 标准题解:dfs
func numIslands2(grid [][]byte) int {
	xlen := len(grid)
	ylen := len(grid[0])

	count := 0
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			if grid[x][y] == 1 {
				count++
				dfs(grid, xlen, ylen, x, y)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, xlen, ylen, x, y int) {
	if x < 0 || y < 0 || x >= xlen || y >= ylen || grid[x][y] == 0 {
		return
	}
	grid[x][y] = 0
	dfs(grid, xlen, ylen, x+1, y)
	dfs(grid, xlen, ylen, x-1, y)
	dfs(grid, xlen, ylen, x, y+1)
	dfs(grid, xlen, ylen, x, y-1)
}
