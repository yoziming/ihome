package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/yoziming/iHome/web/controller"
	"github.com/yoziming/iHome/web/model"
)

// 添加gin開發基礎

func main() {
	// 初始化
	router := gin.Default()

	// 初始化mysql跟redis
	model.InitDb()
	model.InitRdb()
	model.InitMinio()
	// 初始化session容器
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	// 使用容器，中間件必須放在路由之前
	router.Use(sessions.Sessions("mysession", store))

	// 路由匹配到首頁 (dev:http://127.0.0.1:22222/home/)
	router.Static("/home", "view")
	// 匹配favicon
	router.StaticFile("/favicon.ico", "view/favicon.ico")

	// 獲取（搜索）房源服務 *************TODO
	// router.GET("/search.html", controller.GetHouses)

	// 優化過的router.Group
	r1 := router.Group("api/v1.0/")
	{
		// 匹配登陸&註冊頁面
		r1.GET("session", controller.GetSession)
		// 匹配註冊驗證碼頁面
		r1.GET("imagecode/:uuid", controller.GetImageCd)
		// sms驗證頁面
		r1.GET("smscode/:phone", controller.GetSmscd)
		// 註冊頁面
		r1.POST("users", controller.PostRet)
		// 區域頁面
		r1.GET("areas", controller.GetArea)
		// 登入頁面
		r1.POST("sessions", controller.PostLogin)

		// 獲取（搜索）房源服務
		r1.GET("houses", controller.GetHouses)
		// 獲取首頁輪播圖片服務
		r1.GET("house/index", controller.GetIndex)
		// 獲取房屋詳細資訊服務
		r1.GET("houses/:id", controller.GetHouseInfo)

		// 從此插入中間件檢驗session，往下的路由都不用檢驗了
		r1.Use(LoginFilter())

		// 登出
		r1.DELETE("session", controller.DeleteSession)
		// 用户訊息
		r1.GET("user", controller.GetUserInfo)
		// 更新用户名
		r1.PUT("user/name", controller.PutUserInfo)
		// 上傳頭像
		r1.POST("user/avatar", controller.PostAvatar)
		// 實名認證上傳
		r1.POST("user/auth", controller.PostUserAuth)
		// 獲得已上傳的實名認證資訊
		r1.GET("user/auth", controller.GetUserAuth)
		// 發布房源
		r1.POST("houses", controller.PostHouses)
		// 獲取用户發布的房源訊息
		r1.GET("user/houses", controller.GetUserHouses)
		// 上傳房屋圖片
		r1.POST("houses/:id/images", controller.PostHousesImage)

		// 發送（發佈）訂單服務
		r1.POST("orders", controller.PostOrders)
		// 獲取房東/租户訂單資訊服務
		r1.GET("user/orders", controller.GetUserOrder)
		// 更新房東同意/拒絕訂單
		r1.PUT("orders/:id/status", controller.PutOrders)
		// 更新使用者評價訂單資訊
		r1.PUT("orders/:id/comment", controller.PutComment)

	}

	// 啟用
	router.Run(":22222")
	// 路由匹配到首頁 (dev:http://127.0.0.1:22222/home/)
}

// 以中間件形式來檢驗session
func LoginFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		// 從session拿訊息
		userName := s.Get("userName")
		if userName == nil {
			// 發生異常，沒登錄但進入了用户頁面
			c.Abort()
		} else {
			c.Next() // 繼續
		}
	}
}
