package formatter

import (
	"fmt"
	"strings"

	"github.com/OG-Open-Source/whatsdifferent/pkg/types"
)

// CompactFormatter 将差异格式化为紧凑文本
type CompactFormatter struct{}

// NewCompactFormatter 创建一个紧凑格式化器
func NewCompactFormatter() *CompactFormatter {
	return &CompactFormatter{}
}

// Format 将差异结果格式化为紧凑文本
func (f *CompactFormatter) Format(result types.CompareResult) string {
	var builder strings.Builder
	
	isInDiffBlock := false
	for _, change := range result.Changes {
		if change.Type != types.Same {
			if !isInDiffBlock {
				builder.WriteString(fmt.Sprintf("--- Change at line %d ---\n", change.Line+1))
				isInDiffBlock = true
			}
			
			switch change.Type {
			case types.Added:
				builder.WriteString(fmt.Sprintf("+ %s\n", change.Content))
			case types.Deleted:
				builder.WriteString(fmt.Sprintf("- %s\n", change.Content))
			case types.Modified:
				builder.WriteString(fmt.Sprintf("* %s\n", change.Content))
			}
		} else if isInDiffBlock {
			builder.WriteString("---\n")
			isInDiffBlock = false
		}
	}
	
	return builder.String()
} 