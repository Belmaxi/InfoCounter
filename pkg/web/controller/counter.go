package controller

import (
	"InfoCounter/pkg/dao/impl"
	"InfoCounter/pkg/entity"
	"github.com/gin-gonic/gin"
	"strconv"
)

type InfoControllerImpl struct {
	dao *impl.StatisticsDAOImpl
}

type NumInfoController interface {
	AddTable(c *gin.Context)
	GetTableByID(c *gin.Context)
}

func NewInfoControllerImpl() *InfoControllerImpl {
	return &InfoControllerImpl{dao: impl.NewStatisticsDAOImpl()}
}

func (impl InfoControllerImpl) AddTable(c *gin.Context) {
	var info entity.UserInfo
	info.Id = c.PostForm("id")
	info.Name = c.PostForm("name")
	info.PhoneNumber = c.PostForm("phone_number")
	info.Class = c.PostForm("class")
	info.ActivityName = c.PostForm("activity_name")
	info.Date = c.PostForm("date")
	ok := impl.dao.AddTable(c, info)
	if ok == "ok" {
		c.JSON(200, map[string]interface{}{"message": "ok"})
	} else if ok == "duplicate key" {
		c.JSON(200, map[string]interface{}{"message": "Duplicate entry " + info.Id + " for key"})
	} else {
		c.JSON(200, map[string]interface{}{"message": "other error"})
	}
}

func (impl InfoControllerImpl) GetTableByID(c *gin.Context) {
	id := c.Param("key")

	info := impl.dao.GetTableByID(c, id)
	c.JSON(200, gin.H{
		"message": "ok",
		"table": gin.H{
			"name":          info.Name,
			"phone_number":  info.PhoneNumber,
			"class":         info.Class,
			"activity_name": info.ActivityName,
			"date":          info.Date,
		},
	})
}

func (impl InfoControllerImpl) CreateUser(c *gin.Context) {
	id := c.PostForm("id")
	psw := c.PostForm("password")
	permission, _ := strconv.Atoi(c.PostForm("permission"))

	if impl.dao.CreateUser(id, psw, permission) {
		c.JSON(200, map[string]interface{}{"message": "error"})
	} else {
		c.JSON(200, map[string]interface{}{"message": "ok"})
	}
}

func (impl InfoControllerImpl) GetPassword(c *gin.Context) {
	id := c.PostForm("id")
	psw := impl.dao.GetPassword(id)
	c.JSON(200, map[string]interface{}{"password": psw})
}
