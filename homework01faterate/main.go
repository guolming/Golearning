package main

import (
	"fmt"
)

func main() {
	names := [3]string{}
	weights := [3]float64{}
	talls := [3]float64{}
	ages := [3]int{}
	sexes := [3]string{}
	BMI := [3]float64{}
	fateRate := [3]float64{}
	results := [3]string{}
	var totalFateRate float64
	var totalPeople int
	for i := 0; i < len(names); i++ {

		fmt.Print("姓名：")
		fmt.Scanln(&names[i])

		fmt.Print("体重（千克）：")
		fmt.Scanln(&weights[i])

		fmt.Print("身高（米）：")
		fmt.Scanln(&talls[i])

		BMI[i] = weights[i] / (talls[i] * talls[i])

		fmt.Print("年龄：")
		fmt.Scanln(&ages[i])

		var sexWeight int
		fmt.Print("性别（男/女）：")
		fmt.Scanln(&sexes[i])

		if sexes[i] == "男" {
			sexWeight = 1
		} else {
			sexWeight = 0
		}
		fateRate[i] = (1.2*BMI[i] + 0.23*float64(ages[i]) - 5.4 - 10.8*float64(sexWeight)) / 100

		//fmt.Println(i+1,".",names[i]," ",BMI[i]," ",fateRate[i]," ")
		if sexes[i] == "男" {

			if ages[i] >= 18 && ages[i] <= 39 {
				if fateRate[i] >= 0 && fateRate[i] <= 0.1 {
					results[i] = "目前是：偏瘦，要多吃多锻炼，增强体质。"
					fmt.Println("目前是：偏瘦，要多吃多锻炼，增强体质。")
				} else if fateRate[i] > 0.1 && fateRate[i] <= 0.16 {
					results[i] = "目前是：标准，太棒了，要继续保持哦！"
					fmt.Println("目前是：标准，太棒了，要继续保持哦！")
				} else if fateRate[i] > 0.16 && fateRate[i] <= 0.21 {
					results[i] = "目前是：偏重，吃完饭要多散散步，消化消化。"
					fmt.Println("目前是：偏重，吃完饭要多散散步，消化消化。")
				} else if fateRate[i] > 0.21 && fateRate[i] <= 0.26 {
					results[i] = "目前是：偏胖，少吃点，多运动运动。"
					fmt.Println("目前是：偏胖，少吃点，多运动运动。")
				} else if fateRate[i] > 0.26 {
					results[i] = "目前是：非常肥胖，健身游泳跑步，即刻开始。"
					fmt.Println("目前是：非常肥胖，健身游泳跑步，即刻开始。")
				}
			} else if ages[i] > 39 && ages[i] <= 59 {
				if fateRate[i] >= 0 && fateRate[i] <= 0.11 {
					results[i] = "目前是：偏瘦，要多吃多锻炼，增强体质。"
					fmt.Println("目前是：偏瘦，要多吃多锻炼，增强体质。")
				} else if fateRate[i] > 0.11 && fateRate[i] <= 0.17 {
					results[i] = "目前是：标准，太棒了，要继续保持哦！"
					fmt.Println("目前是：标准，太棒了，要继续保持哦！")
				} else if fateRate[i] > 0.17 && fateRate[i] <= 0.22 {
					results[i] = "目前是：偏重，吃完饭要多散散步，消化消化。"
					fmt.Println("目前是：偏重，吃完饭要多散散步，消化消化。")
				} else if fateRate[i] > 0.22 && fateRate[i] <= 0.27 {
					results[i] = "目前是：偏胖，少吃点，多运动运动。"
					fmt.Println("目前是：偏胖，少吃点，多运动运动。")
				} else if fateRate[i] > 0.27 {
					results[i] = "目前是：非常肥胖，健身游泳跑步，即刻开始。"
					fmt.Println("目前是：非常肥胖，健身游泳跑步，即刻开始。")
				}
			} else if ages[i] > 59 {
				if fateRate[i] >= 0 && fateRate[i] <= 0.11 {
					results[i] = "目前是：偏瘦，要多吃多锻炼，增强体质。"
					fmt.Println("目前是：偏瘦，要多吃多锻炼，增强体质。")
				} else if fateRate[i] > 0.11 && fateRate[i] <= 0.17 {
					results[i] = "目前是：标准，太棒了，要继续保持哦！"
					fmt.Println("目前是：标准，太棒了，要继续保持哦！")
				} else if fateRate[i] > 0.17 && fateRate[i] <= 0.22 {
					results[i] = "目前是：偏重，吃完饭要多散散步，消化消化。"
					fmt.Println("目前是：偏重，吃完饭要多散散步，消化消化。")
				} else if fateRate[i] > 0.22 && fateRate[i] <= 0.27 {
					results[i] = "目前是：偏胖，少吃点，多运动运动。"
					fmt.Println("目前是：偏胖，少吃点，多运动运动。")
				} else if fateRate[i] > 0.27 {
					results[i] = "目前是：非常肥胖，健身游泳跑步，即刻开始。"
					fmt.Println("目前是：非常肥胖，健身游泳跑步，即刻开始。")
				}
			} else {
				fmt.Println("我们不看未成年人的体脂率，变化太大，无法评判")
			}
		} else {
			if ages[i] >= 18 && ages[i] <= 39 {
				if fateRate[i] >= 0 && fateRate[i] <= 0.2 {
					results[i] = "目前是：偏瘦，要多吃多锻炼，增强体质。"
					fmt.Println("目前是：偏瘦，要多吃多锻炼，增强体质。")
				} else if fateRate[i] > 0.2 && fateRate[i] <= 0.27 {
					results[i] = "目前是：标准，太棒了，要继续保持哦！"
					fmt.Println("目前是：标准，太棒了，要继续保持哦！")
				} else if fateRate[i] > 0.27 && fateRate[i] <= 0.34 {
					results[i] = "目前是：偏重，吃完饭要多散散步，消化消化。"
					fmt.Println("目前是：偏重，吃完饭要多散散步，消化消化。")
				} else if fateRate[i] > 0.34 && fateRate[i] <= 0.39 {
					results[i] = "目前是：偏胖，少吃点，多运动运动。"
					fmt.Println("目前是：偏胖，少吃点，多运动运动。")
				} else if fateRate[i] > 0.39 {
					results[i] = "目前是：非常肥胖，健身游泳跑步，即刻开始。"
					fmt.Println("目前是：非常肥胖，健身游泳跑步，即刻开始。")
				}
			} else if ages[i] > 39 && ages[i] <= 59 {
				if fateRate[i] >= 0 && fateRate[i] <= 0.21 {
					results[i] = "目前是：偏瘦，要多吃多锻炼，增强体质。"
					fmt.Println("目前是：偏瘦，要多吃多锻炼，增强体质。")
				} else if fateRate[i] > 0.21 && fateRate[i] <= 0.28 {
					results[i] = "目前是：标准，太棒了，要继续保持哦！"
					fmt.Println("目前是：标准，太棒了，要继续保持哦！")
				} else if fateRate[i] > 0.28 && fateRate[i] <= 0.35 {
					results[i] = "目前是：偏重，吃完饭要多散散步，消化消化。"
					fmt.Println("目前是：偏重，吃完饭要多散散步，消化消化。")
				} else if fateRate[i] > 0.35 && fateRate[i] <= 0.40 {
					results[i] = "目前是：偏胖，少吃点，多运动运动。"
					fmt.Println("目前是：偏胖，少吃点，多运动运动。")
				} else if fateRate[i] > 0.40 {
					results[i] = "目前是：非常肥胖，健身游泳跑步，即刻开始。"
					fmt.Println("目前是：非常肥胖，健身游泳跑步，即刻开始。")
				}
			} else if ages[i] > 59 {
				if fateRate[i] >= 0 && fateRate[i] <= 0.22 {
					results[i] = "目前是：偏瘦，要多吃多锻炼，增强体质。"
					fmt.Println("目前是：偏瘦，要多吃多锻炼，增强体质。")
				} else if fateRate[i] > 0.22 && fateRate[i] <= 0.29 {
					results[i] = "目前是：标准，太棒了，要继续保持哦！"
					fmt.Println("目前是：标准，太棒了，要继续保持哦！")
				} else if fateRate[i] > 0.29 && fateRate[i] <= 0.36 {
					results[i] = "目前是：偏重，吃完饭要多散散步，消化消化。"
					fmt.Println("目前是：偏重，吃完饭要多散散步，消化消化。")
				} else if fateRate[i] > 0.36 && fateRate[i] <= 0.41 {
					results[i] = "目前是：偏胖，少吃点，多运动运动。"
					fmt.Println("目前是：偏胖，少吃点，多运动运动。")
				} else if fateRate[i] > 0.41 {
					results[i] = "目前是：非常肥胖，健身游泳跑步，即刻开始。"
					fmt.Println("目前是：非常肥胖，健身游泳跑步，即刻开始。")
				}
			} else {
				fmt.Println("我们不看未成年人的体脂率，变化太大，无法评判")
			}
		}

	}

	for i := 0; i < len(names); i++ {
		totalFateRate += fateRate[i]
		totalPeople = 1 + i
		fmt.Println(names[i], weights[i], talls[i], ages[i], sexes[i], BMI[i], fateRate[i], results[i])
	}

	fmt.Println("总人数：", totalPeople, "平均体脂率：", totalFateRate/3)
}
