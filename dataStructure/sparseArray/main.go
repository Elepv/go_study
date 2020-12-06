package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {

	//先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子

	//输出看看原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}

	//用稀疏数组sparseArray来实现
	//1. 遍历chessMap，如果我们发现有一个元素的值不为0，创造一个node节点
	//2. 将node节点放入到对应的切片中
	var sparseArr []ValNode

	//标准的一个稀疏数组应该还有一个记录元素的二维数组的规模（行、列和默认值）

	//初始值设置
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个ValNode值的节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	//输出稀疏数组
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	//将稀疏数组存盘

	//将稀疏数组读取并回复成原始数组
	// for i, valNode := range sparseArr {
	// 	if i == 0 {
	// 		i1 := valNode.col
	// 		i2 := valNode.row
	// 		var ChessMapReset [i1][i2]int
	// 	} else {
	// 		chessMapReset[valNode.row][valNode.col] = valNode.val
	// 	}
	// }
}
