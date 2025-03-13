package whatsdifferent

import (
	"github.com/OG-Open-Source/whatsdifferent/pkg/diff"
	"github.com/OG-Open-Source/whatsdifferent/pkg/formatter"
	"github.com/OG-Open-Source/whatsdifferent/pkg/types"
)

// WhatsDifferent 是比较文本差异的主对象
type WhatsDifferent struct {
	differ    diff.Differ
	formatter formatter.Formatter
}

// New 创建一个新的 WhatsDifferent 实例
func New(formatType types.FormatterType) *WhatsDifferent {
	return &WhatsDifferent{
		differ:    diff.NewLineDiffer(),
		formatter: formatter.GetFormatter(formatType),
	}
}

// SetFormatter 设置格式化器
func (wd *WhatsDifferent) SetFormatter(formatType types.FormatterType) {
	wd.formatter = formatter.GetFormatter(formatType)
}

// Compare 比较两个字符串并返回差异结果
func (wd *WhatsDifferent) Compare(a, b string) types.CompareResult {
	return wd.differ.Compare(a, b)
}

// FormatDiff 格式化差异结果
func (wd *WhatsDifferent) FormatDiff(result types.CompareResult) string {
	return wd.formatter.Format(result)
}

// CompareAndFormat 比较并格式化差异
func (wd *WhatsDifferent) CompareAndFormat(a, b string) string {
	result := wd.Compare(a, b)
	return wd.FormatDiff(result)
} 