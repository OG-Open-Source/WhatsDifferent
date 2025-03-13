package main

import (
	"fmt"

	"github.com/OG-Open-Source/whatsdifferent"
	"github.com/OG-Open-Source/whatsdifferent/pkg/types"
)

func main() {
	// 创建示例文本
	text1 := `第一行
这是第二行
这是第三行
第四行是相同的
这是第五行，将被删除
这是原始文档的最后一行`

	text2 := `第一行已修改
这是第二行
这是新的第三行
第四行是相同的
这是新文档的最后一行
这是添加的新行`

	// 使用三种不同格式
	fmt.Println("=== 简单格式 ===")
	simpleCompare(text1, text2)
	
	fmt.Println("\n=== 详细格式 ===")
	detailedCompare(text1, text2)
	
	fmt.Println("\n=== 紧凑格式 ===")
	compactCompare(text1, text2)
}

func simpleCompare(a, b string) {
	wd := whatsdifferent.New(types.SimpleFormat)
	result := wd.CompareAndFormat(a, b)
	fmt.Println(result)
}

func detailedCompare(a, b string) {
	wd := whatsdifferent.New(types.DetailedFormat)
	result := wd.CompareAndFormat(a, b)
	fmt.Println(result)
}

func compactCompare(a, b string) {
	wd := whatsdifferent.New(types.CompactFormat)
	result := wd.CompareAndFormat(a, b)
	fmt.Println(result)
} 