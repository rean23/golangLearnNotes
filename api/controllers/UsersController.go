package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	m "api/models"
)

/**
获取用户信息
*/
func UserInfo(context *gin.Context) {
	context.String(http.StatusOK, context.Param("username"))
}

func CreateUser(context *gin.Context) {
	user := m.User{Name: "Rean", Age: 28}

	db := m.GetDB()

	//fmt.Println(user)
	db.Create(&user)
}
