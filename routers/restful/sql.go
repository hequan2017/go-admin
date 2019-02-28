package restful

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/hequan2017/go-admin/pkg/app"
	"github.com/hequan2017/go-admin/pkg/e"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func GetAll(c *gin.Context) {
	appG := app.Gin{C: c}
	offset := com.StrTo(c.DefaultQuery("offset", "0")).MustInt()
	count := com.StrTo(c.DefaultQuery("count", "10")).MustInt()
	queryData, err := SQLQueryByMap(
		c.Query("columnname"),
		c.Query("feilds"),
		c.Param("tablename"),
		c.Query("where"),
		c.Query("order"),
		offset,
		count)
	if err != nil {
		log.Fatal(err.Error())

	} else {
		appG.Response(http.StatusOK, e.SUCCESS, queryData)
	}
}

func GetId(c *gin.Context) {
	appG := app.Gin{C: c}

	queryData, err := SQLQueryByMap(
		"",
		c.Query("feilds"),
		c.Param("tablename"),
		"id="+c.Param("id"),
		"", 0, 1)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, queryData)
	}
}

func Post(c *gin.Context) {
	appG := app.Gin{C: c}
	body, _ := ioutil.ReadAll(c.Request.Body)
	affect, err := SQLInsert(
		c.Param("tablename"),
		body)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, strconv.FormatInt(affect, 10))
	}
}

func Put(c *gin.Context) {
	appG := app.Gin{C: c}
	body, _ := ioutil.ReadAll(c.Request.Body)

	affect, err := SQLUpdate(
		c.Param("tablename"),
		"id="+c.Param("id"),
		body)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, strconv.FormatInt(affect, 10))
	}
}

func Delete(c *gin.Context) {
	appG := app.Gin{C: c}
	affect, err := SQLDelete(
		c.Param("tablename"),
		"id="+c.Param("id"))
	if err != nil {
		log.Fatal(err.Error())
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, strconv.FormatInt(affect, 10))
	}
}
