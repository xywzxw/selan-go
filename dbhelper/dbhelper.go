package dbhelper

import (
	"database/sql"
	"fmt"
	"selan/utils"
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func SQLJSONData(sqlStr string) string {

	path := beego.AppConfig.String("domainpath")
	dbuser := utils.GetValue(path, "db", "dbuser").(string)
	dbpass := utils.GetValue(path, "db", "dbpass").(string)
	dburl := utils.GetValue(path, "db", "dburl").(string)
	dbport := utils.GetValue(path, "db", "dbport").(string)
	dbname := utils.GetValue(path, "db", "dbname").(string)

	dbpath := dbuser + ":" + dbpass + "@tcp(" + dburl + ":" + dbport + ")/" + dbname + "?charset=utf8"
	db, err := sql.Open("mysql", dbpath)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	list := "["

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("log:", err)
			panic(err.Error())
		}

		row := "{"
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			columName := strings.ToLower(columns[i])

			cell := fmt.Sprintf(`"%v":"%v"`, columName, value)
			row = row + cell + ","
		}
		row = row[0 : len(row)-1]
		row += "}"
		list = list + row + ","

	}
	list = list[0 : len(list)-1]
	list += "]"
	fmt.Println("成功获取数据")
	return list

}
