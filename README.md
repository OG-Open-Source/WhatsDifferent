# WhatsDifferent 使用手冊

## 概述

WhatsDifferent 是一個強大的文件差異分析工具，可以幫助您詳細追蹤和顯示文件之間的變化。它支持多種格式的輸出，能夠處理二進制文件，並提供豐富的差異分析功能。

## 主要功能

1. **文件比較**：比較兩個文件之間的差異
2. **字符串比較**：比較兩個字符串之間的差異
3. **行比較**：比較兩個文本行之間的差異
4. **多種輸出格式**：支持簡單、可讀性強和標記語言三種輸出格式
5. **二進制文件檢測**：自動檢測並適當處理二進制文件
6. **語言感知**：對 Python 和其他語言提供不同的重要變更檢測

## 安裝

```bash
pip install whatsdifferent
```

## 基本用法

### 比較兩個文件

```python
from whatsdifferent import WhatsDifferent

wd = WhatsDifferent()
changes = wd.compare_files("old_file.txt", "new_file.txt")
print(wd.format_changes(changes))
```

### 比較兩個字符串

```python
from whatsdifferent import WhatsDifferent

wd = WhatsDifferent()
old_content = "Hello, world!"
new_content = "Hello, Python world!"
changes = wd.compare_strings(old_content, new_content)
print(wd.format_changes(changes))
```

### 比較行

```python
from whatsdifferent import WhatsDifferent

wd = WhatsDifferent()
old_line = "This is a very long line with some content that might be changed."
new_line = "This is a very long line with some content that has been changed."
changes = wd.compare_lines(old_line, new_line)
print(wd.format_changes(changes))
```

## 輸出格式

WhatsDifferent 支持三種輸出格式：

### 1. 簡單格式 (simple)

簡單格式提供基本的差異信息，顯示每行的添加或刪除內容。

```python
wd = WhatsDifferent(format='simple')  # 或者
wd = WhatsDifferent().set_format('simple')
```

輸出示例：

```
@@ 'old_file.txt' & 'new_file.txt' @@ modified
[1,-]This is the old content
[1,+]This is the new content
```

### 2. 可讀性強格式 (readable)

可讀性強格式提供更詳細的差異信息，包括上下文和變更摘要。

```python
wd = WhatsDifferent(format='readable')  # 或者
wd = WhatsDifferent().set_format('readable')
```

輸出示例：

```
@@ 'old_file.txt' & 'new_file.txt' @@ modified

Change Summary:
- Total changes: 2
- Added content: 1 lines
- Removed content: 1 lines

Changes at line 1:
Old file context:
  > 1: This is the old content
  2: Some unchanged content
New file context:
  > 1: This is the new content
  2: Some unchanged content
[1,-]This is the old content
[1,+]This is the new content
```

### 3. 標記語言格式 (markup)

標記語言格式使用 Markdown 風格的輸出，按語義單元（如類、函數等）組織變更。

```python
wd = WhatsDifferent(format='markup')  # 或者
wd = WhatsDifferent().set_format('markup')
```

輸出示例：

```
# Changes between 'old_file.txt' and 'new_file.txt' (modified)
Total changes: 2 (1 added, 1 removed)

## Global scope
- `This is the old content`
+ `This is the new content`

## Key Changes
- Removed: `This is the old content`
+ Added: `This is the new content`
```

## 高級功能

### 全局設置格式

您可以為所有 WhatsDifferent 實例設置默認格式：

```python
WhatsDifferent.set_format('markup')
```

### 二進制文件處理

WhatsDifferent 會自動檢測二進制文件，並提供適當的提示而不是嘗試比較內容：

```python
changes = wd.compare_files("image.png", "image_modified.png")
print(wd.format_changes(changes, "image.png", "image_modified.png"))
# 輸出: @@ 'image.png' & 'image_modified.png' @@ modified
#      Binary file detected - cannot display differences
```

### 語言感知功能

WhatsDifferent 會根據文件類型提供不同的重要變更檢測：

- **Python 文件**：檢測函數和類定義、導入語句、返回語句和函數調用
- **其他語言文件**：檢測函數定義、變量聲明、訪問修飾符等

## 類和方法詳解

### WhatsDifferent 類

主要類，提供所有差異比較和格式化功能。

#### 初始化參數

- `format`：輸出格式（'simple', 'readable', 或 'markup'，默認為 'simple'）

#### 主要方法

- `compare_files(old_file_path, new_file_path)`：比較兩個文件
- `compare_strings(old_content, new_content, file_path="unknown")`：比較兩個字符串
- `compare_lines(old_line, new_line, file_path="unknown")`：比較兩個行
- `format_changes(changes, old_file_path=None, new_file_path=None, format_type=None)`：格式化變更
- `format_for_markup(changes, old_file_path=None, new_file_path=None)`：使用標記語言格式格式化變更
- `set_format(format_type)`：設置輸出格式

### Change 類

表示單個變更的數據類。

#### 屬性

- `content`：變更的內容
- `change_type`：變更類型（'added', 'removed', 或 'info'）
- `location`：變更位置（ChangeLocation 對象）

### ChangeLocation 類

表示變更位置的數據類。

#### 屬性

- `file_path`：文件路徑
- `line_number`：行號
- `char_position`：字符位置（可選）

## 使用場景

1. **代碼審查**：比較代碼變更，識別重要的修改
2. **文檔比較**：比較文檔版本之間的差異
3. **配置文件分析**：分析配置文件的變更
4. **自動化測試**：在測試中驗證輸出與預期結果的差異
5. **版本控制集成**：與版本控制系統集成，提供更詳細的差異視圖

## 最佳實踐

1. 對於代碼文件，使用 'markup' 格式以獲得按語義組織的差異視圖
2. 對於簡單文本文件，使用 'simple' 或 'readable' 格式
3. 處理大型文件時，考慮使用 'simple' 格式以提高性能
4. 在處理可能包含二進制內容的文件時，WhatsDifferent 會自動檢測並適當處理
5. 比較文件時，請確保提供文件路徑參數給 `format_changes` 方法

## 限制

1. 對於非常大的文件，比較可能會消耗大量內存
2. 二進制文件只能檢測是否為二進制，無法顯示具體差異
3. 某些特殊編碼的文件可能需要預先處理
4. 'markup' 格式的語義分組功能在非代碼文件上效果有限

## 示例腳本

### 完整的文件比較示例

```python
from whatsdifferent import WhatsDifferent
import os

def compare_directories(dir1, dir2, output_dir):
	"""比較兩個目錄中的所有文件"""
	wd = WhatsDifferent(format='markup')

	if not os.path.exists(output_dir):
		os.makedirs(output_dir)

	files1 = set(os.listdir(dir1))
	files2 = set(os.listdir(dir2))

	all_files = files1.union(files2)

	for file in all_files:
		file1_path = os.path.join(dir1, file)
		file2_path = os.path.join(dir2, file)

		if file in files1 and file in files2:
			# 文件在兩個目錄中都存在
			changes = wd.compare_files(file1_path, file2_path)
			result = wd.format_changes(changes, file1_path, file2_path)
		elif file in files1:
			# 文件只在第一個目錄中存在
			changes = wd.compare_files(file1_path, "")
			result = wd.format_changes(changes, file1_path, "")
		else:
			# 文件只在第二個目錄中存在
			changes = wd.compare_files("", file2_path)
			result = wd.format_changes(changes, "", file2_path)

		output_file = os.path.join(output_dir, f"{file}_diff.md")
		with open(output_file, 'w', encoding='utf-8') as f:
			f.write(result)

		print(f"Comparison for {file} saved to {output_file}")

# 使用示例
compare_directories("old_version", "new_version", "diff_results")
```

### 與版本控制系統集成

```python
from whatsdifferent import WhatsDifferent
import subprocess
import os

def git_diff_with_details(repo_path, commit1, commit2):
	"""使用 WhatsDifferent 提供 Git 差異的詳細視圖"""
	os.chdir(repo_path)

	# 獲取兩個提交之間更改的文件列表
	cmd = ["git", "diff", "--name-only", commit1, commit2]
	result = subprocess.run(cmd, capture_output=True, text=True)
	changed_files = result.stdout.strip().split('\n')

	wd = WhatsDifferent(format='markup')

	for file in changed_files:
		if not file:  # 跳過空行
			continue

		# 從兩個提交中獲取文件內容
		cmd1 = ["git", "show", f"{commit1}:{file}"]
		cmd2 = ["git", "show", f"{commit2}:{file}"]

		try:
			content1 = subprocess.run(cmd1, capture_output=True, text=True).stdout
			content2 = subprocess.run(cmd2, capture_output=True, text=True).stdout

			# 比較內容
			changes = wd.compare_strings(content1, content2, file)
			result = wd.format_changes(changes)

			print(f"\n{'='*50}\n{file}\n{'='*50}")
			print(result)

		except subprocess.CalledProcessError:
			print(f"Error processing file: {file}")

# 使用示例
git_diff_with_details("/path/to/repo", "HEAD~5", "HEAD")
```

## 總結

WhatsDifferent 是一個功能豐富的文件差異分析工具，提供多種輸出格式和詳細的變更追蹤。它可以處理文本文件和二進制文件，並能夠識別重要的代碼變更。無論是用於代碼審查、文檔比較還是與版本控制系統集成，WhatsDifferent 都能提供清晰、詳細的差異視圖。
