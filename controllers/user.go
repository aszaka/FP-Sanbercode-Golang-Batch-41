package controllers

import (
	"FP/database"
	"FP/repository"
	"FP/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllUser(c *gin.Context) {
	var (
		result gin.H
	)

	users, err := repository.GetAllUser(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

func AddUser(c *gin.Context) {
	var user structs.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		panic(err)
	}

	err = repository.AddUser(database.DbConnection, user)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil input User",
	})
}

func AddAdmin(c *gin.Context) {
	var user structs.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		panic(err)
	}

	err = repository.AddAdmin(database.DbConnection, user)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil input Admin",
	})
}

func UpdateUser(c *gin.Context) {
	var user structs.User

	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	user.IdUser = int16(id)

	err = repository.UpdateUser(database.DbConnection, user)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"results": "Sukses Update User",
	})
}

func DeleteUser(c *gin.Context) {
	var user structs.User
	id, err := strconv.Atoi(c.Param("id"))

	user.IdUser = int16(id)

	err = repository.DeleteUser(database.DbConnection, user)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Sukses Hapus User",
	})
}
