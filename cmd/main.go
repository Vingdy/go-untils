package main

import (
	"fmt"
	"go_untils/file"
	"go_untils/random"
	"time"
)

// TODO 命令行化
// TODO 参数化
func main() {
	res, err := RandomXlsxSelectTarget("/Users/RedYuzi/Desktop/target.xlsx")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf(res)
	res, err = RandomXlsxSelectTarget("/Users/RedYuzi/Desktop/target.xlsx")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf(res)
}

func RandomXlsxSelectTarget(fileName string) (string, error) {
	sheetData, err := file.GetXlsxFileExtract(fileName)
	if err != nil {
		return "", err
	}

	var randomRes string
	var randomSheet, randomRow, randomCol int

	for randomRes == "" {
		randomSheet = random.RandomNumberSelectOne(time.Now().UnixNano(), len(sheetData))
		randomRow = random.RandomNumberSelectOne(time.Now().UnixNano()+1, len(sheetData[randomSheet].Data)-1) //第一行为标题
		randomCol = random.RandomNumberSelectOne(time.Now().UnixNano()-1, len(sheetData[randomSheet].Data[randomRow]))
		randomRes = sheetData[randomSheet].Data[randomRow+1][randomCol]
	}

	randomRes = sheetData[randomSheet].Data[0][randomCol] + randomRes
	return randomRes, nil
}
