package main

// test create: 
// $ for i in {1..10000};do echo "this is a test" >> testfile;done

import (
	"fmt"
	"os"
	"bufio"
	// "io"
)

func main(){	
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s filename\n", os.Args[0][2:])
		return
	}

	// 一次性读取
	// byteData, err := os.ReadFile(os.Args[1])
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(byteData))

	// 测试结果
	// real    0m0.102s
	// user    0m0.006s
	// sys     0m0.024s

	// 分片读取
	// file, _ := os.Open(os.Args[1])
	// defer file.Close()
	// for {
	// 	buf := make([]byte, 1024)
	// 	_, err := file.Read(buf)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Printf("%s",buf)
	// }
	
	// 测试结果
	// real    0m0.110s
	// user    0m0.015s
	// sys     0m0.001s

	// 带缓冲读
	// 按行读
	// file, err := os.Open(os.Args[1])
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// defer file.Close()
	// buf := bufio.NewReader(file)
	// for {
	// 	line, _, err := buf.ReadLine()
	// 	fmt.Println(string(line))
	// 	if err != nil {
	// 		break
	// 	}
	// }
	// 测试结果
	// real    0m0.118s
	// user    0m0.037s
	// sys     0m0.036s
	
	// 指定分隔符，按照单词读取
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	// 测试结果
	// real    0m0.459s
	// user    0m0.071s
	// sys     0m0.200s
}