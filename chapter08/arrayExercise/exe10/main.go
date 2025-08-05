package main

import (
	"fmt"
	"math"
)

func evaluateScores(scores [8]float64) {
	// 找最高分与最低分
	maxScore := scores[0]
	minScore := scores[0]
	maxIndex := 0
	minIndex := 0
	for i, v := range scores {
		if v > maxScore {
			maxScore = v
			maxIndex = i
		}
		if v <= minScore {
			minScore = v
			minIndex = i
		}
	}
	if maxScore == minScore {
		fmt.Println("注意：所有评委打分一致，无法区分最高和最低评委")
	}

	fmt.Printf("\n最高分是 %.2f，来自第 %d 位评委\n", maxScore, maxIndex+1)
	fmt.Printf("最低分是 %.2f，来自第 %d 位评委\n", minScore, minIndex+1)

	// 求最终得分（去掉最高和最低）
	sum := 0.0
	for i, _ := range scores {
		if i != minIndex && i != maxIndex {
			sum += scores[i]
		}
	}
	finalScore := sum / 6
	fmt.Printf("\n最终得分(去掉最高最低后)为%.2f\n", finalScore)

	//找到最佳评委和最差评委
	minDiff := math.MaxFloat64 // 最小偏差，初始设为无穷大
	maxDiff := 0.0             // 最大偏差，初始设为0
	bestIndex := -1
	worstIndex := -1
	for i, v := range scores {
		diff := math.Abs(v - finalScore) //当前评委的偏差
		if diff < minDiff {
			minDiff = diff
			bestIndex = i
		}
		if diff > maxDiff {
			maxDiff = diff
			worstIndex = i
		}
	}
	fmt.Printf("最佳评委是第 %d 位，打分 %.2f，偏差 %.2f\n", bestIndex+1, scores[bestIndex], minDiff)
	fmt.Printf("最差评委是第 %d 位，打分 %.2f，偏差 %.2f\n", worstIndex+1, scores[worstIndex], maxDiff)
}

func main() {
	/*
		跳水比赛，8个评委打分。
		运动员的成绩是8个成绩去掉一个最高分，去掉一个最低分，剩下的6个分数的平均分就是最后得分。
		使用一维数组实现如下功能:
		(1)请把打最高分的评委和最低分的评委找出来。
		(2)找出最佳评委和最差评委。最佳评委就是打分和最后得分最接近的评委。最差评委就是打分和最后得分相差最大的。
	*/
	var scores [8]float64
	for i := 0; i < len(scores); i++ {
		fmt.Printf("请第%d个评委打分:", i+1)
		_, err := fmt.Scan(&scores[i])
		if err != nil {
			fmt.Println("输入有误，请输入数字")
			return
		}
	}
	for i := 0; i < len(scores); i++ {
		fmt.Printf("%.2f\t", scores[i])
	}
	fmt.Println() //换行

	evaluateScores(scores)

}
