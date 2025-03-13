package formatter

import (
	"github.com/OG-Open-Source/whatsdifferent/pkg/types"
)

// Formatter 是格式化差异输出的接口
type Formatter interface {
	Format(result types.CompareResult) string
}

// 根据格式类型获取对应的格式化器
func GetFormatter(format types.FormatterType) Formatter {
	switch format {
	case types.SimpleFormat:
		return NewSimpleFormatter()
	case types.DetailedFormat:
		return NewDetailedFormatter()
	case types.CompactFormat:
		return NewCompactFormatter()
	default:
		return NewSimpleFormatter()
	}
} 