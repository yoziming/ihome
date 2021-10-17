package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yoziming/iHome/web/model"
	"github.com/yoziming/iHome/web/utils"
)

// TODO

type OrderStu struct {
	EndDate   string `json:"end_date"`
	HouseId   string `json:"house_id"`
	StartDate string `json:"start_date"`
}

// 發送（發佈）訂單服務
func PostOrders(ctx *gin.Context) {
	// 回應http狀態用
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)
	//獲取資料
	var order OrderStu
	err := ctx.Bind(&order)

	//校驗資料
	if err != nil {
		fmt.Println("獲取資料錯誤", err)
		return
	}
	//獲取用户名
	userName := sessions.Default(ctx).Get("userName")

	// 調用插入訂單服務
	order_id, err := model.InsertOrder(order.HouseId, order.StartDate, order.EndDate, userName.(string))
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		fmt.Println("===發送（發佈）訂單服務", err)
		return // 如果出錯，那就報錯&退出
	}
	// 回應http狀態
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	temp := make(map[string]interface{})
	temp["order_id"] = order_id
	resp["data"] = temp

}

// 獲取房東/租户訂單資訊服務
func GetUserOrder(ctx *gin.Context) {

	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	//獲取get請求傳參
	role := ctx.Query("role")
	//校驗數據
	if role == "" {
		fmt.Println("獲取數據失敗")
		return
	}
	//獲取用户名
	userName := sessions.Default(ctx).Get("userName")

	orderInfo, err := model.GetOrderInfo(userName.(string), role)
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		fmt.Println("===獲取房東/租户訂單資訊服務", err)
		return // 如果出錯，那就報錯&退出
	}
	// 回應http狀態
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	temp := make(map[string]interface{})
	temp["orders"] = orderInfo
	resp["data"] = temp
}

// 更新房東同意/拒絕訂單
func PutOrders(ctx *gin.Context) {

	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	//獲取數據，id=訂單id
	id := ctx.Param("id")
	var statusStu model.StatusStu
	err := ctx.Bind(&statusStu)
	//校驗數據
	if err != nil || id == "" {
		fmt.Println("獲取數據錯誤", err)
		return
	}

	err = model.UpdateStatus(statusStu.Action, id, statusStu.Reason)
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
	}
	// 回應http狀態
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
}

// 更新使用者評價訂單資訊
func PutComment(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)
	//獲取數據，id=訂單id
	id := ctx.Param("id")
	var comStu model.CommentStu
	err := ctx.Bind(&comStu)
	//校驗數據
	if err != nil || id == "" || comStu.Comment == "" {
		fmt.Println("評論錯誤", err)
		resp["errno"] = utils.RECODE_PARAMERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_PARAMERR)
		return
	}
	err = model.UpdateComment(id, comStu.Comment)
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
	}
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
}
