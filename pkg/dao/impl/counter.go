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

	result, err := impl.db.Exec("insert into user values(?,?,?,?,?,?,?)", info.Id, info.Name, info.Class, info.PhoneNumber, info.ActivityName, info.Date, info.Room)
	fmt.Println(info)
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
	err := result.Scan(&info.Id, &info.Name, &info.Class, &info.PhoneNumber, &info.ActivityName, &info.Date, &info.Room)
	print(err.Error())
	if err != nil {
		fmt.Println("Invalid query")
	}
	return info
}

func (impl StatisticsDAOImpl) GetTables() []entity.UserInfo {
	result, err := impl.db.Query("select * from user")
	var infos = make([]entity.UserInfo, 0, 100)
	if err == nil {
		print(1)
		for result.Next() {
			var info entity.UserInfo
			err = result.Scan(&info.Id, &info.Name, &info.Class, &info.PhoneNumber, &info.ActivityName, &info.Date, &info.Room)
			infos = append(infos, info)
		}
	}
	return infos
}

func (impl StatisticsDAOImpl) CreateUser(id string, password string, permission int) bool {
	_, err := impl.db.Exec("insert into admin values(?,?,?)", id, password, permission)
	if err != nil {
		fmt.Println("增加账号错误")
		return false
	}
	return true
}

func (impl StatisticsDAOImpl) GetPassword(id string) string {
	result := impl.db.QueryRow("select * from admin where id = ?", id)
	var id_ string
	var psw string
	var pem int
	err := result.Scan(&id_, &psw, &pem)
	if err != nil {
		fmt.Println("无法获取账户信息")
	}
	return psw
}
