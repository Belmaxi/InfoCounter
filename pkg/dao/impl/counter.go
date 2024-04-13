package impl

import (
	"InfoCounter/pkg/config"
	"InfoCounter/pkg/entity"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type StatisticsDAOImpl struct {
	db *sql.DB
}

func NewStatisticsDAOImpl() *StatisticsDAOImpl {
	return &StatisticsDAOImpl{db: config.DB}
}

func (impl StatisticsDAOImpl) AddTable(ctx context.Context, info entity.UserInfo) string {

	result, err := impl.db.Exec("insert into user values(?,?,?,?,?,?)", info.Id, info.Name, info.Class, info.PhoneNumber, info.ActivityName, info.Date)
	if err != nil {
		fmt.Println("新增数据错误")
		error_type := err.Error()
		if strings.HasPrefix(error_type, "Error 1062 (23000)") {
			fmt.Println("重复添加key")
			return "duplicate key"
		}
		return "other error"
	}
	newID, _ := result.LastInsertId() //新增数据的ID
	i, _ := result.RowsAffected()     //受影响行数
	fmt.Printf("新增的数据ID：%d , 受影响行数：%d \n", newID, i)
	return "ok"
}

func (impl StatisticsDAOImpl) GetTableByID(ctx context.Context, id string) entity.UserInfo {
	result := impl.db.QueryRow("select * from user where id = ?", id)
	var info entity.UserInfo
	err := result.Scan(&info.Id, &info.Name, &info.Class, &info.PhoneNumber, &info.ActivityName, &info.Date)
	print(err.Error())
	if err != nil {
		fmt.Println("Invalid query")
	}
	return info
}
