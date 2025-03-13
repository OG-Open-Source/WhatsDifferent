package formatter

import (
	"fmt"
	"strings"

	"github.com/OG-Open-Source/whatsdifferent/pkg/types"
)

// SimpleFormatter 将差异格式化为简单文本
type SimpleFormatter struct{}

// NewSimpleFormatter 创建一个简单格式化器
func NewSimpleFormatter() *SimpleFormatter {
	return &SimpleFormatter{}
}

// Format 将差异结果格式化为简单文本
func (f *SimpleFormatter) Format(result types.CompareResult) string {
	var builder strings.Builder
	
	for _, change := range result.Changes {
		prefix := ""
		switch change.Type {
		case types.Same:
			prefix = "  "
		case types.Added:
			prefix = "+ "
		case types.Deleted:
			prefix = "- "
		case types.Modified:
			prefix = "* "
		}
		
		builder.WriteString(fmt.Sprintf("%s%s\n", prefix, change.Content))
	}
	
	return builder.String()
} 