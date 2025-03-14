package diff

import (
	"strings"
)

// DiffOp 表示差異操作類型
type DiffOp struct {
	Op   byte   // '+', '-', 或 ' '
	Text string
}

// SimpleDiff 實現一個簡單的行級差異算法
func SimpleDiff(text1, text2 string) []DiffOp {
	lines1 := strings.Split(text1, "\n")
	lines2 := strings.Split(text2, "\n")
	
	// 移除可能的空行
	if len(lines1) > 0 && lines1[len(lines1)-1] == "" {
		lines1 = lines1[:len(lines1)-1]
	}
	if len(lines2) > 0 && lines2[len(lines2)-1] == "" {
		lines2 = lines2[:len(lines2)-1]
	}
	
	var result []DiffOp
	
	// 使用簡單的行比較
	i, j := 0, 0
	for i < len(lines1) || j < len(lines2) {
		// 兩個文件都還有行
		if i < len(lines1) && j < len(lines2) {
			if lines1[i] == lines2[j] {
				// 相同行
				result = append(result, DiffOp{' ', lines1[i]})
				i++
				j++
			} else {
				// 尋找下一個匹配點
				matchFound := false
				
				// 嘗試在第二個文件中尋找當前第一個文件的行
				for k := j; k < min(j+10, len(lines2)); k++ {
					if lines1[i] == lines2[k] {
						// 找到匹配，添加第二個文件中的新行
						for m := j; m < k; m++ {
							result = append(result, DiffOp{'+', lines2[m]})
						}
						result = append(result, DiffOp{' ', lines1[i]})
						i++
						j = k + 1
						matchFound = true
						break
					}
				}
				
				// 如果沒找到，嘗試在第一個文件中尋找當前第二個文件的行
				if !matchFound {
					for k := i; k < min(i+10, len(lines1)); k++ {
						if lines2[j] == lines1[k] {
							// 找到匹配，添加第一個文件中的刪除行
							for m := i; m < k; m++ {
								result = append(result, DiffOp{'-', lines1[m]})
							}
							result = append(result, DiffOp{' ', lines2[j]})
							i = k + 1
							j++
							matchFound = true
							break
						}
					}
				}
				
				// 如果仍然沒找到匹配，則認為當前行被修改
				if !matchFound {
					result = append(result, DiffOp{'-', lines1[i]})
					result = append(result, DiffOp{'+', lines2[j]})
					i++
					j++
				}
			}
		} else if i < len(lines1) {
			// 第二個文件已結束，第一個文件還有行
			result = append(result, DiffOp{'-', lines1[i]})
			i++
		} else {
			// 第一個文件已結束，第二個文件還有行
			result = append(result, DiffOp{'+', lines2[j]})
			j++
		}
	}
	
	return result
}

// min 返回兩個整數中的較小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
} 