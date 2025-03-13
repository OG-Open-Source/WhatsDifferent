package types

// Change 表示文档中的一个差异
type Change struct {
	Type    ChangeType
	Content string
	Line    int
}

// ChangeType 表示差异的类型
type ChangeType int

const (
	// 相同的内容
	Same ChangeType = iota
	// 添加的内容
	Added
	// 删除的内容
	Deleted
	// 修改的内容
	Modified
)

// CompareResult 保存比较的结果
type CompareResult struct {
	Changes []Change
}

// FormatterType 定义输出格式类型
type FormatterType int

const (
	// 简单格式
	SimpleFormat FormatterType = iota
	// 详细格式
	DetailedFormat
	// 紧凑格式
	CompactFormat
) 