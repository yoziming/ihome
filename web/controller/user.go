package controller

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"net/http"

	"github.com/afocus/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/yoziming/iHome/web/model"
	"github.com/yoziming/iHome/web/utils"
)

// 獲取Session
func GetSession(ctx *gin.Context) {
	// 初始化返回錯誤的map
	resp := make(map[string]interface{})
	// 獲取session
	s := sessions.Default(ctx)
	userName := s.Get("userName")
	if userName == nil {
		// 用户沒登錄，沒存在mysql中也沒存在session中
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	} else {
		// 已經登錄過有session
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		var nameData struct {
			Name string `json:"name"`
		}
		nameData.Name = userName.(string) // 類型斷言
		resp["data"] = nameData           // 返回給前端(根據接口文檔)
	}
	ctx.JSON(http.StatusOK, resp)
}

// 獲取圖片驗證碼
func GetImageCd(ctx *gin.Context) {
	// 實際驗證碼的數字
	uuid := ctx.Param("uuid")
	// 設定驗證碼樣式
	cap := captcha.New()
	cap.SetFont("./conf/comic.ttf")
	cap.SetSize(128, 64)
	// cap.SetFrontColor(color.RGBA{255, 255, 255, 0})
	cap.SetDisturbance(captcha.NORMAL) // 幹擾強度
	// 生成驗證碼
	img, str := cap.Create(4, captcha.NUM)
	// 把驗證碼存到redis
	err := model.SaveImgCode(str, uuid)
	if err != nil {
		fmt.Println("===圖片驗證碼存到redis失敗", err)
		return
	}
	// 編碼將圖像寫入PNG格式的ctx.Writer中
	png.Encode(ctx.Writer, img)
	/* 檢驗用
	fmt.Println("str=", str)
	fmt.Println("uuid=", uuid)
	*/
}

// 獲取smscd
func GetSmscd(ctx *gin.Context) {
	// 存http回報代碼用
	resp := make(map[string]string)
	// 獲取用户前端填入的資料與驗證碼
	phone := ctx.Param("phone")
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")
	fmt.Println("===收到圖片驗證請求...", phone, imgCode, uuid)

	// 檢查圖片驗證碼是否相符
	res := model.CheckImgCode(uuid, imgCode)

	if res {
		fmt.Println("===圖片驗證成功!")
		resp["errno"] = utils.RECODE_OK

		// 發簡訊驗證碼，需使用簡訊業者提供的API，這邊只看教學示範

		// 假設調用簡訊API後發送成功，得到回覆
		var smsSend = true
		if smsSend {
			// 暫時把簡訊驗證碼固定為1234
			var code = "1234"

			// 將簡訊驗證碼存到redis庫
			err := model.SaveSmsCode(code, phone)
			if err != nil {
				fmt.Println("SMS驗證碼存到redis失敗", err)
				resp["errno"] = utils.RECODE_DBERR
				resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
				return
			}
			fmt.Println("===收到簡訊驗證請求...", phone, code)
		} else {
			// 假設調用簡訊驗相關API失敗
			fmt.Println("SMS簡訊發送失敗!")
			resp["errno"] = utils.RECODE_SMSERR
			resp["errmsg"] = utils.RecodeText(utils.RECODE_SMSERR)
		}

	} else {
		fmt.Println("圖片驗證失敗!")
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
	}
	ctx.JSON(http.StatusOK, resp)
}

// 註冊
func PostRet(ctx *gin.Context) {

	// 獲取數據
	var regData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
		SmsCode  string `json:"sms_code"`
	}
	ctx.Bind(&regData)
	fmt.Println("===收到註冊請求...", regData)
	resp := make(map[string]string)

	bool := model.CheckSmsCode(regData.Mobile, regData.SmsCode)
	if bool {
		// 如果SMS檢驗正確，就註冊並寫入mysql
		err := model.RegisterUser(regData.Mobile, regData.PassWord)
		if err != nil {
			// 寫入mysql失敗
			resp["errno"] = utils.RECODE_DBERR
			resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		} else {
			resp["errno"] = utils.RECODE_OK
			resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		}
	} else { // SMS檢驗err
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
	}
	ctx.JSON(http.StatusOK, resp)
}

// 獲取區域
func GetArea(ctx *gin.Context) {
	var areas []model.Area
	// 先從緩存拿數據
	getTemp, err := model.Rdb.Get(model.Ctx, "areaData").Bytes()
	if err != nil {
		// redis中沒有緩存數據
		// 先從mysql中拿數據
		model.GlobalConn.Find(&areas)
		// fmt.Println("從Mysql得到區域數據")

		// 再把數據寫到redis中
		// 先序列化
		data, _ := json.Marshal(areas)
		err := model.Rdb.Set(model.Ctx, "areaData", data, 0).Err()
		if err != nil {
			fmt.Println("把數據寫到redis err", err)
		}
	} else {
		// redis中有名為"areaData"的緩存數據
		// fmt.Println("從redis得到區域數據")
		// 反序列化並存到areas
		json.Unmarshal(getTemp, &areas)
	}

	resp := make(map[string]interface{})
	resp["errno"] = "0"
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = areas

	ctx.JSON(http.StatusOK, resp)
}

// 登入
func PostLogin(ctx *gin.Context) {
	// 獲取前端輸入的數據，用一個結構體來接收
	var loginData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
	}
	ctx.Bind(&loginData)
	resp := make(map[string]interface{})
	// 調用登入驗證功能
	userName, err := model.Login(loginData.Mobile, loginData.PassWord)
	if err != nil {
		//登入失敗
		resp["errno"] = utils.RECODE_LOGINERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
	} else {
		//登入成功
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		// 將登入狀態存到session
		s := sessions.Default(ctx)  // 初始化session
		s.Set("userName", userName) // 把用户名存到session中
		s.Save()
	}
	ctx.JSON(http.StatusOK, resp)
}

// 登出
func DeleteSession(ctx *gin.Context) {
	resp := make(map[string]interface{})
	s := sessions.Default(ctx)
	// 刪除session
	s.Delete("userName")
	err := s.Save()
	if err != nil {
		fmt.Println("登出失敗")
		resp["errno"] = utils.RECODE_IOERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_IOERR)
	} else {
		// 成功登出
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}
	ctx.JSON(http.StatusOK, resp)
}

// 用户訊息
func GetUserInfo(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	s := sessions.Default(ctx)
	// 從session拿訊息
	userName := s.Get("userName")
	if userName == nil {
		// 發生異常，沒登錄但進入了用户頁面
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		return // 如果出錯，那就報錯&退出
	} else {
		// 根據用户名獲取用户訊息
		user, err := model.GetUserInfoFromSql(userName.(string))
		if err != nil {
			resp["errno"] = utils.RECODE_DBERR
			resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
			return // 如果出錯，那就報錯&退出
		} else {
			resp["errno"] = utils.RECODE_OK
			resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
			temp := make(map[string]interface{})
			temp["user_id"] = user.ID
			temp["name"] = user.Name
			temp["mobile"] = user.Mobile
			temp["real_name"] = user.Real_name
			temp["id_card"] = user.Id_card
			temp["avatar_url"] = user.Avatar_url

			resp["data"] = temp
		}
	}
}

// 更新用户名
func PutUserInfo(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	s := sessions.Default(ctx)
	// 從session拿訊息
	userName := s.Get("userName")

	// 從前端拿新的用户名
	var newName struct {
		Name string `json:"name"`
	}
	ctx.Bind(&newName)
	// fmt.Println("===newName:", newName)

	// 更新mysql中的用户名
	err := model.UpdateUserName(newName.Name, userName.(string))
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		fmt.Println("===更新mysql中的用户名失敗")
		return // 如果出錯，那就報錯&退出
	}
	// 更新session
	s.Set("userName", newName.Name)
	err = s.Save()
	if err != nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		fmt.Println("===新session存失敗")
		return
	}
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = newName
}

func PostAvatar(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)
	s := sessions.Default(ctx)
	// 從session拿訊息
	userName := s.Get("userName")
	// 獲取圖片文件資訊
	file, _ := ctx.FormFile("avatar")
	// 這裡使用minioClient.PutObject（）方法上傳，入參是file.Open()。
	f, _ := file.Open()
	// 新增一個哈希命名規則防止圖片覆蓋
	m5 := md5.New()
	m5.Write([]byte(userName.(string) + file.Filename))
	fileName_hash := hex.EncodeToString(m5.Sum(nil))
	// 上傳到minio (這步驟其實應該放在Model，不過比較繁瑣且不熟練，先放在這測試)
	info, err := model.MinioClient.PutObject(ctx, "avatar", fileName_hash, f, -1, minio.PutObjectOptions{ContentType: "avatar"})
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Printf("Successfully uploaded %s of size %d\n", file.Filename, info.Size)
	// 拼接圖片成功上傳至Minio庫後的地址
	dstForUrl := "http://127.0.0.1:9000" + "/avatar" + "/" + fileName_hash
	// 向mysql存入頭像路徑
	err = model.UpdateAvatar(userName.(string), dstForUrl)
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		fmt.Println("===向mysql存入頭像路徑失敗", err)
		return
	} else {
		// 頭像上傳成功
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

		temp := make(map[string]interface{})
		temp["avatar_url"] = dstForUrl
		resp["data"] = temp
	}
}

// 上傳實名認證資料
func PostUserAuth(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	s := sessions.Default(ctx)
	// 從session拿訊息
	userName := s.Get("userName")

	// 從前端拿realName
	var realName struct {
		RealName string `json:"real_name"`
		IdCard   string `json:"id_card"`
	}
	ctx.Bind(&realName)
	fmt.Println("===realName:", realName)

	// 更新mysql中的資料
	err := model.UpdateRealName(userName.(string), realName.RealName, realName.IdCard)
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		fmt.Println("===更新mysql中的realName失敗")
		return // 如果出錯，那就報錯&退出
	}

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
}

// 獲得已上傳的實名認證資訊
func GetUserAuth(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	s := sessions.Default(ctx)
	// 從session拿訊息
	userName := s.Get("userName")

	// 根據用户名從mysql中獲取用户訊息
	user, err := model.GetAuthFromSql(userName.(string))
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return // 如果出錯，那就報錯&退出
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		temp := make(map[string]interface{})
		temp["user_id"] = user.ID
		temp["name"] = user.Name
		temp["mobile"] = user.Mobile
		temp["real_name"] = user.Real_name
		temp["id_card"] = user.Id_card
		temp["avatar_url"] = user.Avatar_url

		resp["data"] = temp
	}
}
