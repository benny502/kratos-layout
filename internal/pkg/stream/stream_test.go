package stream

import (
	"fmt"
	"strconv"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestStream(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3306)/hy_flexbi_report?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("%s", err)
	}
	data := make([]map[string]interface{}, 0)
	db.Table("m_source_table_126").Limit(100).Find(&data)

	result := make([]map[string]interface{}, 0)

	From(func(source chan<- interface{}) {
		for _, row := range data {
			source <- row
		}
	}).Group(func(item interface{}) interface{} {
		row := item.(map[string]interface{})
		return row["field_959"]
	}).ForEach(func(item interface{}) {
		group := item.([]interface{})
		row := make(map[string]interface{})
		_, err := Just(group...).Reduce(func(pipe <-chan interface{}) (interface{}, error) {

			var value float64
			for item := range pipe {
				row["name"] = item.(map[string]interface{})["field_959"]
				num, _ := strconv.ParseFloat(item.(map[string]interface{})["field_962"].(string), 64)
				value += num
				row["value"] = num
			}
			return row, nil
		})
		if err != nil {
			return
		}

		result = append(result, row)
	})

	fmt.Println(result)
}
