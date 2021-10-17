package utils

const (
	RECODE_OK        = "0"
	RECODE_DBERR     = "4001"
	RECODE_NODATA    = "4002"
	RECODE_DATAEXIST = "4003"
	RECODE_DATAERR   = "4004"

	RECODE_SESSIONERR = "4101"
	RECODE_LOGINERR   = "4102"
	RECODE_PARAMERR   = "4103"
	RECODE_USERONERR  = "4104"
	RECODE_ROLEERR    = "4105"
	RECODE_PWDERR     = "4106"
	RECODE_USERERR    = "4107"
	RECODE_SMSERR     = "4108"
	RECODE_MOBILEERR  = "4109"

	RECODE_REQERR    = "4201"
	RECODE_IPERR     = "4202"
	RECODE_THIRDERR  = "4301"
	RECODE_IOERR     = "4302"
	RECODE_SERVERERR = "4500"
	RECODE_UNKNOWERR = "4501"
)

var recodeText = map[string]string{
	RECODE_OK:         "成功",
	RECODE_DBERR:      "資料庫查詢錯誤",
	RECODE_NODATA:     "無數據",
	RECODE_DATAEXIST:  "資料已存在",
	RECODE_DATAERR:    "資料錯誤",
	RECODE_SESSIONERR: "用戶未登錄",
	RECODE_LOGINERR:   "用戶登錄失敗",
	RECODE_PARAMERR:   "參數錯誤",
	RECODE_USERERR:    "用戶不存在或未啟動",
	RECODE_USERONERR:  "用戶已經註冊",
	RECODE_ROLEERR:    "使用者身份錯誤",
	RECODE_PWDERR:     "密碼錯誤",
	RECODE_REQERR:     "非法請求或請求次數受限",
	RECODE_IPERR:      "IP受限",
	RECODE_THIRDERR:   "協力廠商系統錯誤",
	RECODE_IOERR:      "檔讀寫錯誤",
	RECODE_SERVERERR:  "內部錯誤",
	RECODE_UNKNOWERR:  "未知錯誤",
	RECODE_SMSERR:     "短信失敗",
	RECODE_MOBILEERR:  "手機號錯誤",
}

func RecodeText(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}
