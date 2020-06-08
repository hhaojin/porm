package Helper

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//把数据结构 映射成 map切片
func DBMap(columns []string, rows *sql.Rows) ([]interface{}, error) {
	allRows := make([]interface{}, 0) //所有行  大切片
	for rows.Next() {
		oneRow := make([]interface{}, len(columns)) //定义一行切片
		scanRow := make([]interface{}, len(columns))
		fieldMap := make(map[string]interface{})
		for i, _ := range oneRow {
			scanRow[i] = &oneRow[i]
		}
		err := rows.Scan(scanRow...)
		if err != nil {
			return nil, err
		}
		for i, val := range oneRow {
			v, ok := val.([]byte) //断言
			if ok {
				fieldMap[columns[i]] = string(v)
			}
		}
		allRows = append(allRows, fieldMap)
	}
	return allRows, nil
}
