package formatter

import (
	"fmt"
	"strings"

	"github.com/OG-Open-Source/whatsdifferent/pkg/types"
)

// DetailedFormatter 将差异格式化为详细文本
type DetailedFormatter struct{}

// NewDetailedFormatter 创建一个详细格式化器
func NewDetailedFormatter() *DetailedFormatter {
	return &DetailedFormatter{}
}

// Format 将差异结果格式化为详细文本
func (f *DetailedFormatter) Format(result types.CompareResult) string {
	var builder strings.Builder
	
	for _, change := range result.Changes {
		line := change.Line + 1
		switch change.Type {
		case types.Same:
			builder.WriteString(fmt.Sprintf("[%4d]     %s\n", line, change.Content))
		case types.Added:
			builder.WriteString(fmt.Sprintf("[%4d] +++ %s\n", line, change.Content))
		case types.Deleted:
			builder.WriteString(fmt.Sprintf("[%4d] --- %s\n", line, change.Content))
		case types.Modified:
			builder.WriteString(fmt.Sprintf("[%4d] *** %s\n", line, change.Content))
		}
	}
	
	return builder.String()
} 