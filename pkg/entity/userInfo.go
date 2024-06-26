package entity

type UserInfo struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	PhoneNumber  string `json:"phone_number"`
	ActivityName string `json:"activity_name"`
	Date         string `json:"date"`
	Room         string `json:"class_room"`
}

func (stu UserInfo) TableName() string {
	return "user"
}
