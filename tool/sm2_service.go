package tool

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
)

// 一下为tool.go，辅助函数
/*
	保存系统参数
*/
func saveOneValue(filename string,save *big.Int)  {
	outputFile, outputError := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)//os.O_APPEND追加
	if outputError != nil {
		fmt.Println(outputError)
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString(save.String()+"\n")
	// 一定得记得将缓冲区内容刷新到磁盘文件
	outputWriter.Flush()
}

/*
	保存各个成员的参数情况——用于测试
*/
func saveValue(filename string,save []*big.Int)  {
	//写文件 保存p值
	k := len(save)

	outputFile, outputError := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)//os.O_APPEND追加
	if outputError != nil {
		fmt.Println(outputError)
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	for i := 0; i < k; i++ {
		outputWriter.WriteString(save[i].String()+"\n")
	}
	// 一定得记得将缓冲区内容刷新到磁盘文件
	outputWriter.Flush()
}

/*
	读取各个成员的参数情况——用于测试
*/
func ReadValue(filename string) (re []*big.Int) {
	temp := new(big.Int)

	inputFile, inputError := os.Open(filename)
	if inputError != nil {
		fmt.Println(inputError)
		return nil
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		temp.SetString(inputString,10)
		re = append(re,new(big.Int).Set(temp))
		if readerError == io.EOF {
			break
		}
	}
	return re
}

/*
	读取系统参数情况
*/
func ReadOneValue(filename string) (*big.Int) {
	re := new(big.Int)
	inputFile, inputError := os.Open(filename)
	if inputError != nil {
		fmt.Println(inputError)
		return nil
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		re.SetString(inputString,10)

		if readerError == io.EOF {
			break
		}
	}
	return re
}
