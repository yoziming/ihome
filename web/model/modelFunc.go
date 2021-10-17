package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// 以下存放暫時使用(大多為接收資料用)的結構體

type Houses struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	AreaName             string   `protobuf:"bytes,2,opt,name=area_name,json=areaName,proto3" json:"area_name,omitempty"`
	Ctime                string   `protobuf:"bytes,3,opt,name=ctime,proto3" json:"ctime,omitempty"`
	HouseId              int32    `protobuf:"varint,4,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
	ImgUrl               string   `protobuf:"bytes,5,opt,name=img_url,json=imgUrl,proto3" json:"img_url,omitempty"`
	OrderCount           int32    `protobuf:"varint,6,opt,name=order_count,json=orderCount,proto3" json:"order_count,omitempty"`
	Price                int32    `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	RoomCount            int32    `protobuf:"varint,8,opt,name=room_count,json=roomCount,proto3" json:"room_count,omitempty"`
	Title                string   `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`
	UserAvatar           string   `protobuf:"bytes,10,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type HouseDetail struct {
	AreaId   uint   `json:"area_id"`
	Acreage  int32  `protobuf:"varint,1,opt,name=acreage,proto3" json:"acreage,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Beds     string `protobuf:"bytes,3,opt,name=beds,proto3" json:"beds,omitempty"`
	Capacity int32  `protobuf:"varint,4,opt,name=capacity,proto3" json:"capacity,omitempty"`
	//comment
	Comments []*CommentData `protobuf:"bytes,5,rep,name=comments,proto3" json:"comments,omitempty"`
	Deposit  int32          `protobuf:"varint,6,opt,name=deposit,proto3" json:"deposit,omitempty"`
	//展示所有的圖片 主圖片和副圖片
	Facilities           []int32  `protobuf:"varint,7,rep,packed,name=facilities,proto3" json:"facilities,omitempty"`
	Hid                  int32    `protobuf:"varint,8,opt,name=hid,proto3" json:"hid,omitempty"`
	ImgUrls              []string `protobuf:"bytes,9,rep,name=img_urls,json=imgUrls,proto3" json:"img_urls,omitempty"`
	MaxDays              int32    `protobuf:"varint,10,opt,name=max_days,json=maxDays,proto3" json:"max_days,omitempty"`
	MinDays              int32    `protobuf:"varint,11,opt,name=min_days,json=minDays,proto3" json:"min_days,omitempty"`
	Price                int32    `protobuf:"varint,12,opt,name=price,proto3" json:"price,omitempty"`
	RoomCount            int32    `protobuf:"varint,13,opt,name=room_count,json=roomCount,proto3" json:"room_count,omitempty"`
	Title                string   `protobuf:"bytes,14,opt,name=title,proto3" json:"title,omitempty"`
	Unit                 string   `protobuf:"bytes,15,opt,name=unit,proto3" json:"unit,omitempty"`
	UserAvatar           string   `protobuf:"bytes,16,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
	UserId               int32    `protobuf:"varint,17,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,18,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type DetailData struct {
	House                *HouseDetail `protobuf:"bytes,1,opt,name=house,proto3" json:"house,omitempty"`
	UserId               int32        `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

type CommentData struct {
	Comment              string   `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	Ctime                string   `protobuf:"bytes,2,opt,name=ctime,proto3" json:"ctime,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type GetData struct {
	Houses               []*Houses `protobuf:"bytes,1,rep,name=houses,proto3" json:"houses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

type OrderStu struct {
	EndDate   string `json:"end_date"`
	HouseId   string `json:"house_id"`
	StartDate string `json:"start_date"`
}

type UserData struct {
	Id int
}

type StatusStu struct {
	Action string `json:"action"`
	Reason string `json:"reason"`
}

type CommentStu struct {
	Order_id string `json:"order_id"`
	Comment  string `json:"comment"`
}
type UserOrder struct{}

type OrdersData struct {
	Amount               int32    `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Ctime                string   `protobuf:"bytes,3,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Days                 int32    `protobuf:"varint,4,opt,name=days,proto3" json:"days,omitempty"`
	EndDate              string   `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	ImgUrl               string   `protobuf:"bytes,6,opt,name=img_url,json=imgUrl,proto3" json:"img_url,omitempty"`
	OrderId              int32    `protobuf:"varint,7,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	StartDate            string   `protobuf:"bytes,8,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	Status               string   `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	Title                string   `protobuf:"bytes,10,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// 把圖片驗證碼存到redis
func SaveImgCode(code, uuid string) error {
	// 設定寫入的驗證碼60秒過期
	err := Rdb.Set(Ctx, uuid, code, 60*time.Second).Err()
	if err != nil {
		panic(err)
	}
	return err
}

// 比對圖片驗證碼
func CheckImgCode(uuid, imgCode string) bool {
	val, err := Rdb.Get(Ctx, uuid).Result()
	if err != nil {
		fmt.Println("rdb.Get err", err)

	}
	fmt.Println("uuid=", val)
	return val == imgCode
}

// 把簡訊驗證碼存到redis
func SaveSmsCode(code, phone string) error {
	// 設定寫入的驗證碼600秒過期
	err := Rdb.Set(Ctx, phone, code, 600*time.Second).Err()
	if err != nil {
		fmt.Println("簡訊驗證碼存到redis err", err)
	}
	return err
}

// 比對SMS驗證碼
func CheckSmsCode(phone, code string) bool {
	val, err := Rdb.Get(Ctx, phone).Result()
	if err != nil {
		fmt.Println("比對SMS驗證碼 err", err)
	}
	fmt.Println("phone=", val)
	return val == code
}

// 將用户註冊到mysql資料庫
func RegisterUser(mobile, pwd string) error {
	var user User
	user.Name = mobile // 暫時使用手機號當用户名稱
	user.Mobile = mobile
	// 使用md5對pwd加密
	m5 := md5.New()
	m5.Write([]byte(pwd))                       // 把pwd寫入緩衝
	pwd_hash := hex.EncodeToString(m5.Sum(nil)) // 不用額外密鑰
	user.Password_hash = pwd_hash
	// 將user寫入mysql
	return GlobalConn.Create(&user).Error
}

// 驗證登入帳號密碼，並獲取用户名稱
func Login(mobile, pwd string) (string, error) {
	var user User
	// 對參數pwd先轉換成md5 hash
	m5 := md5.New()
	m5.Write([]byte(pwd))
	pwd_hash := hex.EncodeToString(m5.Sum(nil))
	// 查找mysql資料庫
	err := GlobalConn.Where("mobile=?", mobile).Select("name").
		Where("password_hash = ?", pwd_hash).Find(&user).Error
	return user.Name, err
}

// 獲取用户訊息
func GetUserInfoFromSql(userName string) (User, error) {
	// 實現sql: select * from user where name = userName
	var user User
	err := GlobalConn.Where("name = ?", userName).First(&user).Error
	return user, err
}

// 更新用户名
func UpdateUserName(newName, oldName string) error {
	// update user set name = 'newName' where name = 'oldName'
	return GlobalConn.Model(new(User)).Where("name = ?", oldName).Update("name", newName).Error
}

// 更新用户頭像
func UpdateAvatar(userName, avatar string) error {
	// update user set avatar_url = avatar, where name = username
	return GlobalConn.Model(new(User)).Where("name = ?", userName).
		Update("avatar_url", avatar).Error
}

// 更新實名資料
func UpdateRealName(userName, realName, id string) error {
	return GlobalConn.Model(new(User)).Where("name = ?", userName).
		Updates(map[string]interface{}{"real_name": realName, "id_card": id}).Error
}

// 根據用户名從mysql中獲取用户訊息
func GetAuthFromSql(userName string) (User, error) {
	var user User
	err := GlobalConn.Where("name = ?", userName).First(&user).Error
	return user, err
}

// 根據Username從mysql中獲取用户房屋訊息
func GetUserHouse(userName string) ([]*Houses, error) {
	var houseInfos []*Houses
	//有用户名
	var user User
	if err := GlobalConn.Where("name = ?", userName).Find(&user).Error; err != nil {
		fmt.Println("獲取當前用户信息錯誤", err)
		return nil, err
	}
	//房源信息   一對多查詢
	var houses []House
	GlobalConn.Model(&user).Related(&houses)

	for _, v := range houses {
		var houseInfo Houses
		houseInfo.Title = v.Title
		houseInfo.Address = v.Address
		houseInfo.Ctime = v.CreatedAt.Format("2006-01-02 15:04:05")
		houseInfo.HouseId = int32(v.ID)
		houseInfo.ImgUrl = v.Index_image_url
		houseInfo.OrderCount = int32(v.Order_count)
		houseInfo.Price = int32(v.Price)
		houseInfo.RoomCount = int32(v.Room_count)
		houseInfo.UserAvatar = user.Avatar_url

		//獲取地域信息
		var area Area
		//related函數可以是以主表關聯從表,也可以是以從表關聯主表
		GlobalConn.Where("id = ?", v.AreaId).Find(&area)
		houseInfo.AreaName = area.Name

		houseInfos = append(houseInfos, &houseInfo)
	}
	return houseInfos, nil
}

type HouseStu struct {
	Acreage   string   `json:"acreage"`
	Address   string   `json:"address"`
	AreaId    string   `json:"area_id"`
	Beds      string   `json:"beds"`
	Capacity  string   `json:"capacity"`
	Deposit   string   `json:"deposit"`
	Facility  []string `json:"facility"`
	MaxDays   string   `json:"max_days"`
	MinDays   string   `json:"min_days"`
	Price     string   `json:"price"`
	RoomCount string   `json:"room_count"`
	Title     string   `json:"title"`
	Unit      string   `json:"unit"`
}

// 上傳房屋資料
func UpdateHouse(userName string, request HouseStu) (houseId int, err error) {
	var houseInfo House
	//給house賦值
	houseInfo.Address = request.Address
	//根據userName獲取userId
	var user User
	if err := GlobalConn.Where("name = ?", userName).Find(&user).Error; err != nil {
		fmt.Println("查詢當前用户失敗", err)
		return 0, err
	}
	//sql中一對多插入,只是給外鍵賦值
	houseInfo.UserId = uint(user.ID)
	houseInfo.Title = request.Title
	//類型轉換
	price, _ := strconv.Atoi(request.Price)
	roomCount, _ := strconv.Atoi(request.RoomCount)
	houseInfo.Price = price
	houseInfo.Room_count = roomCount
	houseInfo.Unit = request.Unit
	houseInfo.Capacity, _ = strconv.Atoi(request.Capacity)
	houseInfo.Beds = request.Beds
	houseInfo.Deposit, _ = strconv.Atoi(request.Deposit)
	houseInfo.Min_days, _ = strconv.Atoi(request.MinDays)
	houseInfo.Max_days, _ = strconv.Atoi(request.MaxDays)
	houseInfo.Acreage, _ = strconv.Atoi(request.Acreage)
	//一對多插入
	areaId, _ := strconv.Atoi(request.AreaId)
	houseInfo.AreaId = uint(areaId)

	//request.Facility    所有的傢俱  房屋
	for _, v := range request.Facility {
		id, _ := strconv.Atoi(v)
		var fac Facility
		if err := GlobalConn.Where("id = ?", id).First(&fac).Error; err != nil {
			fmt.Println("傢俱id錯誤", err)
			return 0, err
		}
		//查詢到了數據
		houseInfo.Facilities = append(houseInfo.Facilities, &fac)
	}

	// 創建該房屋資料到mysql中
	err = GlobalConn.Create(&houseInfo).Error
	if err != nil {
		fmt.Println("插入房屋信息失敗", err)
		return 0, err
	}

	return int(houseInfo.ID), err
}

// 把圖片的憑證存儲到資料中   更新   主圖,次圖  第一張圖片是主圖,剩下的圖片是副圖
func SaveHouseImg(houseId, imgPath string) error {
	/*return GlobalConn.Model(new(House)).Where("id = ?",houseId).
	Update("index_image_url",imgPath).Error*/
	// 如何判斷上傳的圖是當前房屋的第一張圖片
	var houseInfo House
	if err := GlobalConn.Where("id = ?", houseId).Find(&houseInfo).Error; err != nil {
		fmt.Println("查詢不到房屋資訊", err)
		return err
	}

	if houseInfo.Index_image_url == "" {
		// 説明沒有上傳過圖片  現在上傳的圖片是主圖
		return GlobalConn.Model(new(House)).Where("id = ?", houseId).
			Update("index_image_url", imgPath).Error
	}

	// 上傳副圖
	var houseImg HouseImage
	houseImg.Url = imgPath
	hId, _ := strconv.Atoi(houseId)
	houseImg.HouseId = uint(hId)
	return GlobalConn.Create(&houseImg).Error
}

// 根據userName & houseId從mysql中獲取房屋訊息
func GetHouseDetail(houseId, userName string) (DetailData, error) {
	var respData DetailData
	// 給houseDetail賦值
	var houseDetail HouseDetail

	var houseInfo House
	if err := GlobalConn.Where("id = ?", houseId).Find(&houseInfo).Error; err != nil {
		fmt.Println("===查詢房屋資訊錯誤", err)
		return respData, err
	}
	{
		houseDetail.Acreage = int32(houseInfo.Acreage)
		houseDetail.Address = houseInfo.Address
		houseDetail.Beds = houseInfo.Beds
		houseDetail.Capacity = int32(houseInfo.Capacity)
		houseDetail.Deposit = int32(houseInfo.Deposit)
		houseDetail.Hid = int32(houseInfo.ID)
		houseDetail.MaxDays = int32(houseInfo.Max_days)
		houseDetail.MinDays = int32(houseInfo.Min_days)
		houseDetail.Price = int32(houseInfo.Price)
		houseDetail.RoomCount = int32(houseInfo.Room_count)
		houseDetail.Title = houseInfo.Title
		houseDetail.Unit = houseInfo.Unit

		if houseInfo.Index_image_url != "" {
			houseDetail.ImgUrls = append(houseDetail.ImgUrls, ""+houseInfo.Index_image_url)
		}

	}

	//評論在order表
	var orders []OrderHouse
	if err := GlobalConn.Model(&houseInfo).Related(&orders).Error; err != nil {
		fmt.Println("查詢房屋評論資訊", err)
		return respData, err
	}
	//var comments []*house.CommentData
	for _, v := range orders {
		var commentTemp CommentData
		commentTemp.Comment = v.Comment
		commentTemp.Ctime = v.CreatedAt.Format("2006-01-02 15:04:05")
		var tempUser User
		GlobalConn.Model(&v).Related(&tempUser)
		commentTemp.UserName = tempUser.Name

		houseDetail.Comments = append(houseDetail.Comments, &commentTemp)
	}

	//獲取房屋的傢俱資訊  多對多查詢
	var facs []Facility
	if err := GlobalConn.Model(&houseInfo).Related(&facs, "Facilities").Error; err != nil {
		fmt.Println("查詢房屋傢俱資訊錯誤", err)
		return respData, err
	}
	for _, v := range facs {
		houseDetail.Facilities = append(houseDetail.Facilities, int32(v.Id))
	}

	//獲取副圖片  幅圖找不到算不算錯
	var imgs []HouseImage
	if err := GlobalConn.Model(&houseInfo).Related(&imgs).Error; err != nil {
		fmt.Println("該房屋只有主圖", err)
	}

	// for _, v := range imgs {
	//  if len(imgs) != 0 {
	//      houseDetail.ImgUrls = append(houseDetail.ImgUrls, "http://192.168.137.81:8888/"+v.Url)
	//  }
	// }

	//獲取房屋所有者資訊
	var user User
	if err := GlobalConn.Model(&houseInfo).Related(&user).Error; err != nil {
		fmt.Println("查詢房屋所有者資訊錯誤", err)
		return respData, err
	}
	houseDetail.UserName = user.Name
	houseDetail.UserAvatar = user.Avatar_url
	houseDetail.UserId = int32(user.ID)

	respData.House = &houseDetail

	//獲取當前流覽人資訊
	var nowUser User
	if err := GlobalConn.Where("name = ?", userName).Find(&nowUser).Error; err != nil {
		fmt.Println("查詢當前流覽人資訊錯誤", err)
		return respData, err
	}
	respData.UserId = int32(nowUser.ID)
	return respData, nil
}

// 根據houseId從mysql中獲取房屋訊息
func GetHouseDetailWithId(houseId string) (DetailData, error) {
	var respData DetailData
	// 給houseDetail賦值
	var houseDetail HouseDetail

	var houseInfo House
	if err := GlobalConn.Where("id = ?", houseId).First(&houseInfo).Error; err != nil {
		fmt.Println("===查詢房屋資訊錯誤", err)
		return respData, err
	}
	{
		houseDetail.Acreage = int32(houseInfo.Acreage)
		houseDetail.Address = houseInfo.Address
		houseDetail.Beds = houseInfo.Beds
		houseDetail.Capacity = int32(houseInfo.Capacity)
		houseDetail.Deposit = int32(houseInfo.Deposit)
		houseDetail.Hid = int32(houseInfo.ID)
		houseDetail.MaxDays = int32(houseInfo.Max_days)
		houseDetail.MinDays = int32(houseInfo.Min_days)
		houseDetail.Price = int32(houseInfo.Price)
		houseDetail.RoomCount = int32(houseInfo.Room_count)
		houseDetail.Title = houseInfo.Title
		houseDetail.Unit = houseInfo.Unit

		if houseInfo.Index_image_url != "" {
			houseDetail.ImgUrls = append(houseDetail.ImgUrls, ""+houseInfo.Index_image_url)
		}

		// 用預設圖片代替
		// str1 := "images/home01.jpg"
		// str2 := "images/home02.jpg"
		// str3 := "images/home03.jpg"
		// houseDetail.ImgUrls = append(houseDetail.ImgUrls, str1, str2, str3)

	}

	//評論在order表
	var orders []OrderHouse
	if err := GlobalConn.Model(&houseInfo).Related(&orders).Error; err != nil {
		fmt.Println("查詢房屋評論資訊", err)
		return respData, err
	}
	//var comments []*house.CommentData
	for _, v := range orders {
		var commentTemp CommentData
		commentTemp.Comment = v.Comment
		commentTemp.Ctime = v.CreatedAt.Format("2006-01-02 15:04:05")
		var tempUser User
		GlobalConn.Model(&v).Related(&tempUser)
		commentTemp.UserName = tempUser.Name

		houseDetail.Comments = append(houseDetail.Comments, &commentTemp)
	}

	//獲取房屋的傢俱資訊  多對多查詢
	var facs []Facility
	if err := GlobalConn.Model(&houseInfo).Related(&facs, "Facilities").Error; err != nil {
		fmt.Println("查詢房屋傢俱資訊錯誤", err)
		return respData, err
	}
	for _, v := range facs {
		houseDetail.Facilities = append(houseDetail.Facilities, int32(v.Id))
	}

	//獲取副圖片  幅圖找不到算不算錯
	var imgs []HouseImage
	if err := GlobalConn.Model(&houseInfo).Related(&imgs).Error; err != nil {
		fmt.Println("該房屋只有主圖", err)
	}

	for _, v := range imgs {
		if len(imgs) != 0 {
			houseDetail.ImgUrls = append(houseDetail.ImgUrls, ""+v.Url)
		}
	}

	//獲取房屋所有者資訊
	var user User
	if err := GlobalConn.Model(&houseInfo).Related(&user).Error; err != nil {
		fmt.Println("查詢房屋所有者資訊錯誤", err)
		return respData, err
	}
	houseDetail.UserName = user.Name
	houseDetail.UserAvatar = user.Avatar_url
	houseDetail.UserId = int32(user.ID)

	respData.House = &houseDetail

	return respData, nil
}

// 獲取首頁輪播房屋資訊
func GetIndexHouse() ([]*Houses, error) {

	var housesResp []*Houses

	var houses []House
	if err := GlobalConn.Limit(5).Find(&houses).Error; err != nil {
		fmt.Println("獲取房屋資訊失敗", err)
		return nil, err
	}

	for _, v := range houses {
		var houseTemp Houses
		houseTemp.Address = v.Address
		//根據房屋資訊獲取地域資訊
		var area Area
		var user User

		GlobalConn.Model(&v).Related(&area).Related(&user)

		houseTemp.AreaName = area.Name
		houseTemp.Ctime = v.CreatedAt.Format("2006-01-02 15:04:05")
		houseTemp.HouseId = int32(v.ID)
		houseTemp.ImgUrl = v.Index_image_url
		houseTemp.OrderCount = int32(v.Order_count)
		houseTemp.Price = int32(v.Price)
		houseTemp.RoomCount = int32(v.Room_count)
		houseTemp.Title = v.Title
		houseTemp.UserAvatar = user.Avatar_url

		housesResp = append(housesResp, &houseTemp)
	}

	return housesResp, nil
}

// 搜尋房屋
func SearchHouse(areaId, sd, ed, sk string) ([]*Houses, error) {
	var houseInfos []House
	if areaId == "" {
		// 表示沒選區域就搜索，顯示所有房源
		err := GlobalConn.Model(new(House)).Where("area_id > ?", 0).Find(&houseInfos).Error
		if err != nil {
			fmt.Println("搜索房屋失敗", err)
			return nil, err
		}
	} else {
		// 前端設定了搜尋條件
		//   minDays  <  (結束時間  -  開始時間) <  max_days
		//計算一個差值  先把string類型轉為time類型
		sdTime, _ := time.Parse("2006-01-02", sd)
		edTime, _ := time.Parse("2006-01-02", ed)
		dur := edTime.Sub(sdTime)
		fmt.Println("***傳入的搜尋房屋條件(areaId,durtion)", areaId, dur.Hours()/24)
		err := GlobalConn.Model(new(House)).Where("area_id = ?", areaId).
			// 計算時間要選的限制很多很煩，先無視
			// Where("min_days < ?", dur.Hours()/24).
			// Where("max_days > ?", dur.Hours()/24).
			Find(&houseInfos).Error
		if err != nil {
			fmt.Println("搜索房屋失敗", err)
			return nil, err
		}
	}

	//獲取[]*Houses
	var housesResp []*Houses
	for _, v := range houseInfos {
		var houseTemp Houses
		houseTemp.Address = v.Address
		//根據房屋資訊獲取地域資訊
		var area Area
		var user User

		GlobalConn.Model(&v).Related(&area).Related(&user)

		houseTemp.AreaName = area.Name
		houseTemp.Ctime = v.CreatedAt.Format("2006-01-02 15:04:05")
		houseTemp.HouseId = int32(v.ID)
		houseTemp.ImgUrl = v.Index_image_url
		houseTemp.OrderCount = int32(v.Order_count)
		houseTemp.Price = int32(v.Price)
		houseTemp.RoomCount = int32(v.Room_count)
		houseTemp.Title = v.Title
		houseTemp.UserAvatar = user.Avatar_url

		housesResp = append(housesResp, &houseTemp)
	}
	return housesResp, nil
}

func InsertOrder(houseId, beginDate, endDate, userName string) (int, error) {
	//獲取插入對象
	var order OrderHouse

	//給對象賦值
	hid, _ := strconv.Atoi(houseId)
	order.HouseId = uint(hid)

	//把string類型的時間轉換為time類型
	bDate, _ := time.Parse("2006-01-02", beginDate)
	order.Begin_date = bDate

	eDate, _ := time.Parse("2006-01-02", endDate)
	order.End_date = eDate

	//需要userId
	/*var user User
	GlobalConn.Where("name = ?",userName).Find(&user)*/
	//select id form user where name = userName

	var userData UserData
	if err := GlobalConn.Raw("select id from user where name = ?", userName).Scan(&userData).Error; err != nil {
		fmt.Println("獲取用户數據錯誤", err)
		return 0, err
	}

	//獲取days
	dur := eDate.Sub(bDate)
	order.Days = (int(dur.Hours()) / 24) + 1
	order.Status = "WAIT_ACCEPT"

	//房屋的單價和總價
	var house House
	GlobalConn.Where("id = ?", hid).Find(&house).Select("price")
	order.House_price = house.Price
	order.Amount = house.Price * order.Days

	order.UserId = uint(userData.Id)
	if err := GlobalConn.Create(&order).Error; err != nil {
		fmt.Println("插入訂單失敗", err)
		return 0, err
	}
	return int(order.ID), nil
}

//獲取房東訂單如何實現?
func GetOrderInfo(userName, role string) ([]*OrdersData, error) {
	//最終需要的數據
	var orderResp []*OrdersData
	//獲取當前用户的所有訂單
	var orders []OrderHouse

	var userData UserData
	//用原生查詢的時候,查詢的字段必須跟數據庫中的字段保持一直
	GlobalConn.Raw("select id from user where name = ?", userName).Scan(&userData)

	//查詢租户的所有的訂單
	if role == "custom" {
		if err := GlobalConn.Where("user_id = ?", userData.Id).Find(&orders).Error; err != nil {
			fmt.Println("獲取當前用户所有訂單失敗")
			return nil, err
		}
	} else {
		//查詢房東的訂單  以房東視角來查看訂單
		var houses []House
		GlobalConn.Where("user_id = ?", userData.Id).Find(&houses)

		for _, v := range houses {
			var tempOrders []OrderHouse
			GlobalConn.Model(&v).Related(&tempOrders)

			orders = append(orders, tempOrders...)
		}
	}

	//循環遍歷一下orders
	for _, v := range orders {
		var orderTemp OrdersData
		orderTemp.OrderId = int32(v.ID)
		orderTemp.EndDate = v.End_date.Format("2006-01-02")
		orderTemp.StartDate = v.Begin_date.Format("2006-01-02")
		orderTemp.Ctime = v.CreatedAt.Format("2006-01-02")
		orderTemp.Amount = int32(v.Amount)
		orderTemp.Comment = v.Comment
		orderTemp.Days = int32(v.Days)
		orderTemp.Status = v.Status

		//關聯house表
		var house House
		GlobalConn.Model(&v).Related(&house).Select("index_image_url", "title")
		orderTemp.ImgUrl = house.Index_image_url
		orderTemp.Title = house.Title

		orderResp = append(orderResp, &orderTemp)
	}
	return orderResp, nil
}

//更新訂單狀態
func UpdateStatus(action, id, reason string) error {
	db := GlobalConn.Model(new(OrderHouse)).Where("id = ?", id)

	if action == "accept" {
		//標示房東同意訂單
		return db.Update("status", "WAIT_COMMENT").Error
	} else {
		//表示房東不同意訂單  如果拒單把拒絕的原因寫到comment中
		return db.Updates(map[string]interface{}{"status": "REJECTED", "comment": reason}).Error
	}
}

// 更新評價
func UpdateComment(order_id, comment string) error {
	db := GlobalConn.Model(new(OrderHouse)).Where("id = ?", order_id)
	err := db.Updates(map[string]interface{}{"status": "COMPLETE", "comment": comment}).Error
	if err != nil {
		fmt.Println("更新OrderHouse失敗")
		return err
	}
	// 獲取houseId
	var order OrderHouse
	err = GlobalConn.Model(new(OrderHouse)).Where("id = ?", order_id).First(&order).Error
	if err != nil {
		fmt.Println("從order_id獲取整張order資料失敗")
		return err
	}
	houseId := order.HouseId
	var house House
	err = GlobalConn.Model(new(House)).Where("id = ?", houseId).First(&house).Error
	if err != nil {
		fmt.Println("從houseId獲取整個house資料失敗")
		return err
	}
	newCount := house.Order_count + 1
	err = GlobalConn.Model(new(House)).Where("id = ?", houseId).Update("order_count", newCount).Error
	if err != nil {
		fmt.Println("從houseId更新計數失敗")
		return err
	}
	return err
}
