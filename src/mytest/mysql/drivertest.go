package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func TestDriver() {
	db, err := sql.Open("mysql", "root:213456@/test?charset=utf8")
	CheckErr(err)

	//插入数据
	var sql_string string = "INSERT userinfo set username=?,departname=?,created=?"
	var id int = 2
	stmt, err := db.Prepare(sql_string)
	CheckErr(err)

	res, err := stmt.Exec("song", "研发部门", "2012-12-12")
	CheckErr(err)

	fmt.Println(id)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	CheckErr(err)

	res, err = stmt.Exec("songUpdate", id)
	CheckErr(err)

	affect, err := res.RowsAffected()
	CheckErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	CheckErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		CheckErr(err)
		fmt.Println("----------")
		fmt.Println("|", uid, "|")
		fmt.Println("|", username, "|")
		fmt.Println("|", department, "|")
		fmt.Println("|", created, "|")
	}

	//删除数据
	stmt, err = db.Prepare("DELETE FROM userinfo where uid=?")
	CheckErr(err)

	affect, err = res.RowsAffected()
	CheckErr(err)
	fmt.Println(affect)
	db.Close()
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
