package storage

import (
	"database/sql"
	"fmt"
	"goclip/model"
	"goclip/pkg"
	_ "modernc.org/sqlite"
	"os"
	"time"
)

var DB *sql.DB

func NewDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	//链接测试
	if pingErr := db.Ping(); pingErr != nil {
		return nil, pingErr
	}
	if _, CreateErr := db.Exec(`create table if not exists clip_history(
    id integer primary key autoincrement,
    content text not null,
    create_at datetime 
);`); CreateErr != nil {
		db.Close()
		return nil, CreateErr
	}
	DB = db
	return DB, nil
}

func AddDate(content string) error {
	_, err := DB.Exec(`insert into clip_history (content,create_at) values (?,?)`, content, time.Now())
	if err != nil {
		return err
	}
	pkg.WriteInfoLog("service.log", "数据添加成功")
	return nil
}

func List(limit string) {
	// ToDo 如何将查询结果放入切片中
	var history model.History
	//返回一个查询结果集的游标
	rows, err := DB.Query("select * from clip_history limit ?", limit)
	if err != nil {
		return
	}
	defer rows.Close()
	//遍历游标直至为空
	for rows.Next() {
		err := rows.Scan(&history.Id, &history.Content, &history.Time)
		if err != nil {
			return
		}
		fmt.Fprintf(os.Stdout, "ID:%d,Content:%s,Time:%s\n", history.Id, history.Content, history.Time)
	}

	fmt.Fprintf(os.Stdout, "输出成功,总共查询了%d条记录", limit)
}

func GetCount() (int64, error) {
	result, err := DB.Exec(`select count(*) from clip_history`)
	if err != nil {
		return 0, err
	}
	count, _ := result.RowsAffected()
	return count, nil
}

func DeleteData(count int64) error {
	//使用事务确保数据
	tx, TErr := DB.Begin()
	if TErr != nil {
	}

	//如果事务还未提交，保证回滚，维持数据安全
	defer func() {
		if TErr != nil {
			tx.Rollback()
		}
	}()

	//删除操作
	_, TErr = tx.Exec(`delete from clip_history where id in (
    select id from clip_history order by create_at asc limit ?
)`, count)
	//如果错误不为空，那么就返回错误，此时会触发defer从而回滚
	if TErr != nil {
		return TErr
	}
	if TErr = tx.Commit(); TErr != nil {
		return TErr
	}
	return nil
}

func CleanData() error {
	_, err := DB.Exec("delete from clip_history")
	if err != nil {
		return err
	}
	return nil
}

func GetData() []model.History {
	rows, err := DB.Query(`select * from clip_history`)
	var historyList []model.History
	index := 0
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(historyList[index].Id, historyList[index].Content, historyList[index].Time)
		index++
	}
	return historyList
}

func UpdateData(id int, content string, newtime time.Time) error {
	_, err := DB.Exec(`update clip_history set content=? and create_at=? where id =? `, content, newtime, id)
	if err != nil {
		return err
	}
	return nil
}
