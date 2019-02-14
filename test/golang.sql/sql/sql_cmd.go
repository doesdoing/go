package sql

import (
	"database/sql"
	"fmt"
)

func Check(a error) {
	if a != nil {
		fmt.Println(a)
	}
}

/*
GetSQL sdfdsfdsfdsfsdfsdfdsfsd
*/
func GetSQL(db *sql.DB) []map[string]string {
	rows, err := db.Query("select * from users_info")
	Check(err)

	defer rows.Close()
	cloumns, err := rows.Columns()
	Check(err)

	var info = []map[string]string{}
	values := make([]sql.RawBytes, len(cloumns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		MAP := make(map[string]string)
		err = rows.Scan(scanArgs...)
		Check(err)

		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			MAP[cloumns[i]] = value
		}
		info = append(info, MAP)
	}
	Check(err)

	return info
}

/* InSQL...sadsad

 */
func InSQL(db *sql.DB, v []*string) {
	stmt, err := db.Prepare("INSERT users_info SET username=?,password=?,level=?,who=?")
	Check(err)

	res, err := stmt.Exec(v[0], v[1], v[2], v[3])
	Check(err)

	id, err := res.LastInsertId()
	Check(err)

	fmt.Println(id)
}
