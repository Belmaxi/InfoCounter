package entity

type UserInfo struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	PhoneNumber  string `json:"phone_number"`
	UserNumber   int    `json:"user_number"`
	ActivityName string `json:"activity_name"`
	Date         string `json:"date"`
}

func (stu UserInfo) TableName() string {
	return "user"
}
