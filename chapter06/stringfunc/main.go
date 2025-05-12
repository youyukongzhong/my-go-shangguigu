package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//统计字符串的长度,按字节 len(str)
	str := "hello北"                   //golang的编码统一为utf-8 (ascii的字符(字母和数字) 占一个字符, 汉字占用3个字节)
	fmt.Println("1)str的长度", len(str)) //8

	//2)字符串遍历,同时处理有中文的问题 r := []rune(str)
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("2)字符=%c,", r[i])
	}
	fmt.Println()

	//3)字符串转整数:n,err := strconv.Atoi("123")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("3)转换的结果是", n)
	}

	//4)整数转字符串 str = strconv.Itoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("4)整数转字符串值str=%v,类型str=%T\n", str, str)

	//5)字符串转[]byte: var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("5)bytes=%v\n", bytes)

	//6)[]byte转字符串: str = string([]byte{97,98,99})
	str = string([]byte{97, 98, 99})
	fmt.Printf("6)bytes=%v\n", str)

	//7)10进制转2,8,16进制 str = strconv.FormatInt(123,2)返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("7)123对应的2进制是%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("  123对应的16进制是%v\n", str)

	//8)查找子串是否在制定的字符串中:strings.Contains("seafood","foo") //true
	b := strings.Contains("seafood", "foo")
	fmt.Printf("8) b=%v\n", b)

	//9)统计一个字符串中有几个指定的子串: strings.Count("chelsea","e") //4
	num := strings.Count("chelsea", "e")
	fmt.Printf("9) num=%v\n", num)

	//10)不区分大小写的字符串比较(==是区分字母大小写的): fmt.Println(strings.EqualFold("abc","ABC"))//true
	b = strings.EqualFold("abc", "ABC")
	fmt.Printf("10) b=%v\n", b)          //true
	fmt.Println("10)结果", "abc" == "ABC") //false 区分字母大小写

	//11)返回子串在字符串第一次出现的index值,如果没有则返回-1: strings.Index("NLT_abc","abc") //4
	index := strings.Index("NLT_abc", "c")
	fmt.Printf("11) index=%v\n", index)

	//12)返回子串在字符串最后一次出现的index，如没有返回-1:strings.LastIndex("go golang”,"go")
	index = strings.LastIndex("go golang", "go")
	fmt.Printf("12) index=%v\n", index)

	//13)将指定的子串替换成另外一个子串: strings.Replace("go go hello","go","北京",n)
	//n可以指定你希望替换几个，如果n=-1表示全部替换
	str2 = "go go hello"
	str = strings.Replace(str2, "go", "北京", 2)
	fmt.Printf("13) str=%v str2=%v\n", str, str2)

	//14)按照指定的某个字符，为分割标识，将一个字符串拆分成字符串切片:strings.Split("hello,world,ok",",")
	strArr := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("14) str[%v] =%v\n", i, strArr[i])
	}
	fmt.Printf("14) strArr=%v strArr=%T\n", strArr, strArr)

	//15)将字符串的字母进行大小写的转换: strings.ToLower("Go”)//go | strings.ToUpper("Go")//GO
	str = "golang Hello"
	str = strings.ToLower(str)
	str = strings.ToUpper(str)
	fmt.Printf("15) str=%v\n", str) //golang hello | GOLANG HELLO

	//16)将字符串左右两边的空格去掉:strings.TrimSpace(" happy gopher entry   ")
	str = " happy gopher entry   "
	str = strings.TrimSpace(str)
	fmt.Printf("16) str=%q\n", str)

	//17)将字符串左右两边指定的字符去掉 : strings.Trim("! hello! "," !")//["hello"]//将左右两边!和" "去掉
	str = strings.Trim("! hello! ", "o !")
	fmt.Printf("17) str=%q\n", str)

	//18)将字符串左边指定的字符去掉 : strings.TrimLeft("! hello!"," !")//["hello"]//将左边!和" "去掉
	str = strings.TrimLeft("! hello!", "! ")
	fmt.Printf("18) str=%q\n", str)

	//19)将字符串右边指定的字符去掉 : strings.TrimRight("! hello! ","! ") //["hello"]//将右边!和" "去掉
	str = strings.TrimRight("! hello! ", "! ")
	fmt.Printf("19) str=%q\n", str)

	//20)判断字符串是否以指定的字符串开头: strings.HasPrefix("ftp://192.168.10.1",“ftp”) // true
	b = strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Printf("20) b=%v\n", b)
	//21)判断字符串是否以指定的字符串结束: strings.HasSuffix("NLT_abc.jpg”，"abc”) //false
	b = strings.HasSuffix("NLT_abc.jpg", "abc")
	fmt.Printf("21) b=%v\n", b)
}
