package main

import "fmt"

func exchangeNumber(arr *[4][4]int) {
	//交换第1行和第4行
	for j := 0; j < len(arr[0]); j++ {
		arr[0][j], arr[len(arr)-1][j] = arr[len(arr)-1][j], arr[0][j]
	}

	//交换第2行和第3行
	for j := 0; j < len(arr[0]); j++ {
		arr[1][j], arr[len(arr)-2][j] = arr[len(arr)-2][j], arr[1][j]
	}

	//输出交换后的数组
	for _, v := range arr {
		for _, w := range v {
			fmt.Printf("%d\t", w)
		}
		fmt.Println()
	}
}

func main() {
	//定义一个4行4列的二维数组,逐个从键盘输入值,然后将第1行和第4行的数据进行交换,将第2行和第3行的数据进行交换
	//定义数组,并从键盘输入值
	var arr [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("请输入第%d行第%d列的值\n", i+1, j+1)
			fmt.Scan(&arr[i][j])
		}
	}
	exchangeNumber(&arr)

}
