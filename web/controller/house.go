package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/yoziming/iHome/web/model"
	"github.com/yoziming/iHome/web/utils"
)

// 發布房源
func PostHouses(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)
	fmt.Println("===發布")

	s := sessions.Default(ctx)
	userName := s.Get("userName")

	// 從前端拿House
	var house model.HouseStu
	err := ctx.Bind(&house)
	//校驗數據
	if err != nil {
		fmt.Println("獲取數據錯誤", err)
		return
	}

	// 更新mysql中的資料
	houseId, err := model.UpdateHouse(userName.(string), house)
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		fmt.Println("===更新mysql中的UpdateHouse失敗")
		return // 如果出錯，那就報錯&退出
	}

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	respData := make(map[string]interface{})
	respData["house_id"] = houseId
	resp["data"] = respData

}

// 獲取用户發布的房源訊息
func GetUserHouses(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	s := sessions.Default(ctx)
	// 從session拿訊息
	userName := s.Get("userName")
	// 根據用户名從mysql中獲取用户ID

	houseInfo, err := model.GetUserHouse(userName.(string))
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		fmt.Println("===根據UserId從mysql中獲取用户房屋訊息err")
		return // 如果出錯，那就報錯&退出
	}

	respData := make(map[string]interface{})
	respData["houses"] = houseInfo
	resp["data"] = respData
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
}

// 上傳房屋圖片
func PostHousesImage(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	//獲取數據
	houseId := ctx.Param("id")
	fileHeader, err := ctx.FormFile("house_image")
	//校驗數據
	if houseId == "" || err != nil {
		fmt.Println("傳入數據不完整", err)
		return
	}
	//校驗 大小,類型
	if fileHeader.Size > 50000000 {
		fmt.Println("文件過大,請重新選擇")
		return
	}
	fileExt := path.Ext(fileHeader.Filename)
	if fileExt != ".png" && fileExt != ".jpg" {
		fmt.Println("文件類型錯誤,請重新選擇")
		return
	}

	//獲取文件字節切片
	file, _ := fileHeader.Open()
	/*
		buf := make([]byte, fileHeader.Size)
		file.Read(buf)
		f, _ := fileHeader.Open()
		var size int64 = fileHeader.Size
		buffer := make([]byte, size)
		// 讀取文件內容到buffer
		f.Read(buffer)
	*/
	// 上傳到minio (這步驟其實應該放在Model，不過比較繁瑣且不熟練，先放在這測試)
	// 新增一個哈希命名規則防止圖片覆蓋
	m5 := md5.New()
	m5.Write([]byte(houseId + fileHeader.Filename))
	fileName_hash := hex.EncodeToString(m5.Sum(nil))
	info, err := model.MinioClient.PutObject(ctx, "house", fileName_hash, file, -1, minio.PutObjectOptions{ContentType: "house"})
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Printf("Successfully uploaded %s of size %d\n", fileHeader.Filename, info.Size)
	// 拼接圖片成功上傳至Minio庫後的地址
	dstForUrl := "http://127.0.0.1:9000" + "/house" + "/" + fileName_hash
	// 向mysql存入房屋圖片路徑
	err = model.SaveHouseImg(houseId, dstForUrl)
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		fmt.Println("===向mysql存入頭像路徑失敗", err)
		return
	} else {
		// 房屋圖片上傳成功
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

		temp := make(map[string]interface{})
		temp["url"] = dstForUrl
		resp["data"] = temp
	}

}

// 獲取房屋詳細資訊服務
func GetHouseInfo(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)
	// 從請求網址獲取houseID
	houseId := ctx.Param("id")
	// fmt.Println("===houseId=", houseId)

	// 根據houseId從mysql中獲取房屋訊息
	houseInfo, err := model.GetHouseDetailWithId(houseId)
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		fmt.Println("===根據houseId從mysql中獲取房屋訊息err")
		return // 如果出錯，那就報錯&退出
	}

	respData := make(map[string]interface{})
	respData["house"] = houseInfo.House
	resp["data"] = respData
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
}

// 獲取首頁輪播圖片服務
func GetIndex(ctx *gin.Context) {

	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	// 獲取首頁輪播房屋資訊
	houses, err := model.GetIndexHouse()
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		fmt.Println("===獲取首頁輪播房屋資訊err", err)
		return // 如果出錯，那就報錯&退出
	}
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	temp := make(map[string]interface{})
	temp["houses"] = houses
	resp["data"] = temp

}

// 獲取（搜索）房源服務
func GetHouses(ctx *gin.Context) {
	// 回應http狀態用
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)
	// 根據請求網址獲得參數
	/*
		•	adi表示地區編號
		•	sd表示起始日期
		•	ed表示結束日期
		•	sk表示查詢方式
	*/
	areaId := ctx.Query("aid")
	sd := ctx.Query("sd")
	ed := ctx.Query("ed")
	sk := ctx.Query("sk")
	// if areaId == "" {
	// 	fmt.Println("===欲搜尋的資料不完整")
	// 	resp["errno"] = utils.RECODE_PARAMERR
	// 	resp["errmsg"] = utils.RecodeText(utils.RECODE_PARAMERR)
	// 	return
	// }
	// 調用model的搜尋房屋服務
	houses, err := model.SearchHouse(areaId, sd, ed, sk)
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		fmt.Println("===獲取（搜索）房源服務錯誤", err)
		return // 如果出錯，那就報錯&退出
	}
	// 回應http狀態
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	temp := make(map[string]interface{})
	temp["houses"] = houses
	temp["current_page"] = 1
	resp["data"] = temp
}
