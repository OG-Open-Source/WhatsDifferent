package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/OG-Open-Source/diffutil/pkg/diff"
)

func main() {
	// 定義命令行參數
	file1Ptr := flag.String("file1", "", "第一個文件的路徑")
	file2Ptr := flag.String("file2", "", "第二個文件的路徑")
	outputPtr := flag.String("output", "", "輸出差異結果的文件路徑 (可選)")
	flag.Parse()

	// 檢查必要參數
	if *file1Ptr == "" || *file2Ptr == "" {
		fmt.Println("錯誤: 必須提供兩個文件路徑")
		flag.Usage()
		os.Exit(1)
	}

	// 讀取文件
	text1, err := os.ReadFile(*file1Ptr)
	if err != nil {
		fmt.Printf("讀取文件 %s 錯誤: %v\n", *file1Ptr, err)
		os.Exit(1)
	}

	text2, err := os.ReadFile(*file2Ptr)
	if err != nil {
		fmt.Printf("讀取文件 %s 錯誤: %v\n", *file2Ptr, err)
		os.Exit(1)
	}

	// 計算差異
	diffs := diff.SimpleDiff(string(text1), string(text2))

	// 輸出差異
	output := os.Stdout
	if *outputPtr != "" {
		f, err := os.Create(*outputPtr)
		if err != nil {
			fmt.Printf("創建輸出文件錯誤: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		output = f
	}

	// 輸出差異
	fmt.Fprintf(output, "--- %s\n", *file1Ptr)
	fmt.Fprintf(output, "+++ %s\n", *file2Ptr)
	
	for _, d := range diffs {
		switch d.Op {
		case ' ':
			fmt.Fprintf(output, " %s\n", d.Text)
		case '+':
			fmt.Fprintf(output, "+%s\n", d.Text)
		case '-':
			fmt.Fprintf(output, "-%s\n", d.Text)
		}
	}
} 