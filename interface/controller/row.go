package row

import (
	"database/sql"

	usertable "../../domain/table"
)

func RowToArrayMap(rows *sql.Rows) map[int]*usertable.Column {
	var usersMap = map[int]*usertable.Column{}
	var count = 0
	for rows.Next() {
		var user usertable.Column
		rows.Scan(&user.Id)

		usersMap[count] = &usertable.Column{Id: user.Id}
		count++
	}
	return usersMap
}
