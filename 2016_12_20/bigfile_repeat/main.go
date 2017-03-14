// 判断字符串是否已存在
package main

import (
	"crypto/md5"
	"encoding/base64"
	"os"
	"strconv"
)

// 生成md5字符串
func makeMd5(i int) string {
	md5byteArr := md5.Sum([]byte(strconv.Itoa(i)))
	return base64.StdEncoding.EncodeToString(md5byteArr[:])
}

// 保存到文件
func saveMd5(f *os.File, md5str string) {
	_, err := f.WriteString(md5str + "\n")
	if err != nil {
		panic(err)
	}
}

// 初始化
func init() {
	f, err := os.OpenFile("md5str", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}()

	num := 100000
	for i := 0; i < num; i++ {
		md5str := makeMd5(i)
		saveMd5(f, md5str)
	}
}

// 在上面的基础上实现这样的一个功能: 往文件中再写入一个md5字符串时，怎么能够以更快的速度判断这个字符串是否已经存在
func readMd5() {

}

func findMd5() {

}

func main() {

}
