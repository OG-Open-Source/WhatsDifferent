package diff

import (
	"strings"

	"github.com/OG-Open-Source/whatsdifferent/pkg/types"
)

// Differ 是计算差异的接口
type Differ interface {
	Compare(a, b string) types.CompareResult
}

// LineDiffer 按行比较文本
type LineDiffer struct{}

// NewLineDiffer 创建一个按行比较的对象
func NewLineDiffer() *LineDiffer {
	return &LineDiffer{}
}

// Compare 比较两个字符串，并返回行差异
func (d *LineDiffer) Compare(a, b string) types.CompareResult {
	aLines := strings.Split(a, "\n")
	bLines := strings.Split(b, "\n")
	
	// 简化实现 - 在实际使用中，您可能需要更复杂的差异算法，如 Myers 算法
	result := types.CompareResult{
		Changes: make([]types.Change, 0),
	}
	
	// 找出最大行数
	maxLines := len(aLines)
	if len(bLines) > maxLines {
		maxLines = len(bLines)
	}
	
	for i := 0; i < maxLines; i++ {
		// 检查是否超出任一文档范围
		if i >= len(aLines) {
			// B 中有额外行
			result.Changes = append(result.Changes, types.Change{
				Type:    types.Added,
				Content: bLines[i],
				Line:    i,
			})
		} else if i >= len(bLines) {
			// A 中有额外行
			result.Changes = append(result.Changes, types.Change{
				Type:    types.Deleted,
				Content: aLines[i],
				Line:    i,
			})
		} else if aLines[i] != bLines[i] {
			// 行已修改
			result.Changes = append(result.Changes, types.Change{
				Type:    types.Modified,
				Content: bLines[i],
				Line:    i,
			})
		} else {
			// 行相同
			result.Changes = append(result.Changes, types.Change{
				Type:    types.Same,
				Content: aLines[i],
				Line:    i,
			})
		}
	}
	
	return result
} 