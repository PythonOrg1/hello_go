package search
//

import (
	"encoding/json"
	"os"
)

// 数据文件
const dataFile = "data/data.json"

// Feed 包含我们需要处理的数据源的信息
// 结构体
// `json："xxx"` 中的称为 'tag' 标记
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// 获取数据源列表
// RetrieveFeeds reads and unmarshals the feed data file.
func RetrieveFeeds() ([]*Feed, error) {
	// Open the file.
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once
	// the function returns.
	// defer 在函数执行返回之后才会执行， 关闭文件
	defer file.Close()

	// Decode the file into a slice of pointers 解码文件至一个切片里
	// to Feed values.
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// We don't need to check for errors, the caller can do this.
	return feeds, err
}