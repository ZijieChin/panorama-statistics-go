package controller

import (
	"Panorama-Statistics/database"
	"Panorama-Statistics/util"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Response struct {
	MSG  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var res = &Response{}

// @Summary Insert visit info
// @Description Insert visit information using post method
// @Accept application/x-www-form-urlencoded
// @Param user body string true "user's wechat_work id"
// @Param page body string true "page name"
// @Success 200 {string} json "{"msg": "success", "data": ""}"
// @Router /stat [post]
func InsertInfo(c *gin.Context) {
	user := c.PostForm("user")
	page := c.PostForm("page")
	result := database.InsertStat(database.Statistics{
		User: user,
		Page: page,
	})
	if result != nil {
		res.MSG = "fail"
		res.Data = result.Error()
		c.JSON(500, res)
		return
	}
	res.MSG = "success"
	c.JSON(200, "")
}

// @Summary Get page view
// @Description Get page view data with certain condition
// @Param user body string false "user's wechat_work id"
// @Param page body string false "page name"
// @Param start body string false "start date"
// @Param end body string false "end date"
// @Success 200 {string} json "{"msg": "success", "data": ""}"
// @Router /stat [get]
func GetInfo(c *gin.Context) {
	user := c.Query("user")
	page := c.Query("page")
	start := c.Query("start")
	end := c.Query("end")
	stime := util.TimeHandler(start, 0)
	etime := util.TimeHandler(end, 1)
	result := database.GetStat(user, page, stime, etime)
	res.MSG = "success"
	res.Data = util.ResFormatter(result)
	c.JSON(200, res)
}

// @Summary Get page view file
// @Description Get page view data with certain condition and return with Excel file
// @Param user body string false "user's wechat_work id"
// @Param page body string false "page name"
// @Param start body string false "start date"
// @Param end body string false "end date"
// @Router /file [get]
func GetFile(c *gin.Context) {
	user := c.Query("user")
	page := c.Query("page")
	start := c.Query("start")
	end := c.Query("end")
	stime := util.TimeHandler(start, 0)
	etime := util.TimeHandler(end, 1)
	result := database.GetStat(user, page, stime, etime)
	resdata := util.ResFormatter(result)
	filename := util.XLSXGen(resdata)
	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "attachment; filename=PageView.xlsx")
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.File("./temp/" + filename)
	defer os.RemoveAll("./temp/" + filename)
}
