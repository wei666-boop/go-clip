package internal

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"goclip/storage"
	"os"
)

func ExportJson(name string) error {
	result := storage.GetData()
	if result == nil {
	}

	//格式化输出为json
	jsonData, err := json.MarshalIndent(result, "", "")
	if err == nil {
	}

	//写入到json文件
	err = os.WriteFile(name, jsonData, 0644)
	if err != nil {
	}
	return nil
}

func ExportCsv(name string) error {
	result := storage.GetData()
	if result == nil {
	}

	file, err := os.Create(name)
	if err != nil {
	}
	defer file.Close()

	//写入表头
	writer := csv.NewWriter(file)
	//刷新缓冲区
	defer writer.Flush()

	header := []string{"id", "content", "create_at"}
	if err = writer.Write(header); err != nil {
	}

	//写入数据行
	for _, row := range result {
		//将结构体转换为字符串切片，因为Write只接受字符串切片
		record := []string{fmt.Sprint(row.Id), row.Content, fmt.Sprint(row.Content)}
		if err = writer.Write(record); err != nil {
		}
	}
	return nil
}
