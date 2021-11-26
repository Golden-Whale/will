package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"will/common"
	"will/model"
	"will/util"
)

// GetWill 获取Will
func GetWill(ctx *gin.Context) {
	db := common.GetDB()

	// 获取参数
	limit, _ := ctx.GetQuery("limit")
	offset, _ := ctx.GetQuery("offset")
	month, _ := ctx.GetQuery("date")

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		intLimit = 0
	}
	intOffset, err := strconv.Atoi(offset)
	if err != nil {
		intOffset = 0
	}
	startTime, err := util.StringMonthToTime(month)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "data": nil, "message": "时间格式错误"})
		return
	}

	// 获取数据
	var data []model.Will
	var totalDataCount int64
	db.Order("-created_at").Limit(intLimit).Offset(intOffset).Where("created_at BETWEEN ? AND ?", startTime, startTime.AddDate(0, 1, 0)).Find(&data)
	db.Find(&model.Will{}).Where("created_at BETWEEN ? AND ?", startTime, startTime.AddDate(0, 1, 0)).Count(&totalDataCount)
	// 返回响应
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"totalDataCount": totalDataCount, "data": data}, "message": "Success"})
}

// PutWill 新增Will
func PutWill(ctx *gin.Context) {
	db := common.GetDB()
	// 增加Will
	newWill := model.Will{}
	var totalDataCount int64

	db.Create(&newWill)
	startTime, _ := util.StringMonthToTime(newWill.CreatedAt.GetYearAndMonth())

	db.Find(&model.Will{}).Where("created_at BETWEEN ? AND ?", startTime, startTime.AddDate(0, 1, 0)).Count(&totalDataCount)
	// 返回响应
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"totalDataCount": totalDataCount, "data": newWill}, "message": "Success"})
}

// GetWillHourRange 获取Will时间范围 {"data": [0: {count: 100, "data": {}}]}
func GetWillHourRange(ctx *gin.Context) {
	db := common.GetDB()

	// 获取参数
	// 获取数据
	var data []model.Will
	var totalDataCount int64
	listMap := make(gin.H)
	for i := 0; i < 24; i++ {
		db.Where("CAST(strftime('%H', created_at) AS INT) >= ?  AND CAST(strftime('%H', created_at) AS INT) < ?", i, i+1).Find(&data)
		db.Find(&model.Will{}).Where("CAST(strftime('%H', created_at) AS INT) >= ? AND CAST(strftime('%H', created_at) AS INT) < ?", i, i+1).Count(&totalDataCount)
		listMap[strconv.Itoa(i)] = gin.H{"DataCount": totalDataCount, "data": data}
	}

	// 返回响应
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": listMap, "message": "Success"})
}
