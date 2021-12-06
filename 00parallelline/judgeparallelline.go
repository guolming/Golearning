package main

import (
	"fmt"
)

func main() {
	pointOne := [2]float64{}
	pointTwo := [2]float64{}
	pointThree := [2]float64{}
	pointFour := [2]float64{}
	var k1, k2, b1, b2 float64
	var yesOrNo string

	for j := 0; j < 10; j++ {
		fmt.Println("round:", j+1, "请输入两条直线四个点的坐标：（X1，Y1），（X2，Y2），（X3，Y3），（X4，Y4）")
		fmt.Println("输入:X1：")
		fmt.Scanln(&pointOne[0])
		fmt.Println("输入:Y1：")
		fmt.Scanln(&pointOne[1])
		fmt.Println("输入:X2：")
		fmt.Scanln(&pointTwo[0])
		fmt.Println("输入:Y2：")
		fmt.Scanln(&pointTwo[1])

		fmt.Println("输入:X3：")
		fmt.Scanln(&pointThree[0])
		fmt.Println("输入:Y3：")
		fmt.Scanln(&pointThree[1])
		fmt.Println("输入:X4：")
		fmt.Scanln(&pointFour[0])
		fmt.Println("输入:Y4：")
		fmt.Scanln(&pointFour[1])

		if (pointTwo[0]-pointOne[0] == 0.0) || (pointFour[0]-pointThree[0] == 0.0) == true {
			if (pointTwo[0]-pointOne[0] != 0.0) || (pointFour[0]-pointThree[0] != 0.0) == true {
				fmt.Println("两条直线不平行。")
			} else if pointOne[0] == pointThree[0] {
				fmt.Println("两条直线重合。")
			} else {
				fmt.Println("两条直线平行。")
			}

		} else {
			k1 = (pointTwo[1] - pointOne[1]) / (pointTwo[0] - pointOne[0])
			k2 = (pointFour[1] - pointThree[1]) / (pointFour[0] - pointThree[0])
			b1 = pointOne[1] - (pointTwo[0] * ((pointTwo[1] - pointOne[1]) / (pointTwo[0] - pointOne[0])))          //第一条直线的截距
			b2 = pointThree[1] - (pointFour[0] * ((pointFour[1] - pointThree[1]) / (pointFour[0] - pointThree[0]))) //第二条直线的截距
			if k1 == k2 {
				if b1 == b2 {
					fmt.Println("两条直线重合。")
				} else {
					fmt.Println("两条直线平行。")
				}
			} else {
				fmt.Println("两条直线不平行。")
			}
		}
		fmt.Println("是否继续？（y/n）")
		fmt.Scanln(&yesOrNo)
		if yesOrNo != "y" {
			break
		}
	}

}
