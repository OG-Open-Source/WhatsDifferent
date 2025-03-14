package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/OG-Open-Source/diffutil/pkg/diff"
)

func main() {
	// 定義命令行參數
	testCasePtr := flag.String("test", "basic", "要運行的測試案例: basic, complex, file")
	flag.Parse()

	switch *testCasePtr {
	case "basic":
		runBasicTest()
	case "complex":
		runComplexTest()
	case "file":
		runFileTest()
	default:
		fmt.Println("未知的測試案例:", *testCasePtr)
		os.Exit(1)
	}
}

func runBasicTest() {
	fmt.Println("運行基本測試...")
	
	text1 := "line1\nline2\nline3\nline4\n"
	text2 := "line1\nline2 modified\nline3\nline4\nline5\n"
	
	diffs := diff.SimpleDiff(text1, text2)
	
	fmt.Println("--- 原始文本")
	fmt.Println("+++ 修改後文本")
	
	for _, d := range diffs {
		switch d.Op {
		case ' ':
			fmt.Printf(" %s\n", d.Text)
		case '+':
			fmt.Printf("+%s\n", d.Text)
		case '-':
			fmt.Printf("-%s\n", d.Text)
		}
	}
}

func runComplexTest() {
	fmt.Println("運行複雜測試...")
	
	text1 := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// 這是一個註釋
	fmt.Println("Goodbye!")
}
`
	
	text2 := `package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, " + os.Args[1] + "!")
	// 這是一個修改過的註釋
	fmt.Println("Goodbye!")
	fmt.Println("See you later!")
}
`
	
	diffs := diff.SimpleDiff(text1, text2)
	
	fmt.Println("--- 原始代碼")
	fmt.Println("+++ 修改後代碼")
	
	for _, d := range diffs {
		switch d.Op {
		case ' ':
			fmt.Printf(" %s\n", d.Text)
		case '+':
			fmt.Printf("+%s\n", d.Text)
		case '-':
			fmt.Printf("-%s\n", d.Text)
		}
	}
}

func runFileTest() {
	fmt.Println("運行文件測試...")
	
	// 創建臨時文件
	tmpFile1 := "temp_file1.txt"
	tmpFile2 := "temp_file2.txt"
	
	content1 := "This is a test file.\nIt has multiple lines.\nSome lines will be changed.\nOthers will remain the same.\n"
	content2 := "This is a test file.\nIt has multiple lines with changes.\nSome lines will be changed.\nOthers will remain the same.\nAnd some new lines will be added.\n"
	
	err := os.WriteFile(tmpFile1, []byte(content1), 0644)
	if err != nil {
		fmt.Printf("創建臨時文件錯誤: %v\n", err)
		os.Exit(1)
	}
	defer os.Remove(tmpFile1)
	
	err = os.WriteFile(tmpFile2, []byte(content2), 0644)
	if err != nil {
		fmt.Printf("創建臨時文件錯誤: %v\n", err)
		os.Exit(1)
	}
	defer os.Remove(tmpFile2)
	
	// 讀取文件
	text1, err := os.ReadFile(tmpFile1)
	if err != nil {
		fmt.Printf("讀取文件錯誤: %v\n", err)
		os.Exit(1)
	}
	
	text2, err := os.ReadFile(tmpFile2)
	if err != nil {
		fmt.Printf("讀取文件錯誤: %v\n", err)
		os.Exit(1)
	}
	
	// 計算差異
	diffs := diff.SimpleDiff(string(text1), string(text2))
	
	fmt.Println("--- " + tmpFile1)
	fmt.Println("+++ " + tmpFile2)
	
	for _, d := range diffs {
		switch d.Op {
		case ' ':
			fmt.Printf(" %s\n", d.Text)
		case '+':
			fmt.Printf("+%s\n", d.Text)
		case '-':
			fmt.Printf("-%s\n", d.Text)
		}
	}
} 