package model

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// 初始化minio
var MinioClient *minio.Client

func InitMinio() {
	ctx := context.Background()
	endpoint := "localhost:9000"
	accessKeyID := "root"
	secretAccessKey := "root1234"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	// 方便全局使用
	MinioClient = minioClient

	// 創建預設的bucket
	// Make a new bucket called mymusic.
	bucketName := "avatar"
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	bucketName2 := "house"
	err = minioClient.MakeBucket(ctx, bucketName2, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName2)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName2)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName2)
	}
}

// 初始化mysql
var GlobalConn *gorm.DB

func InitDb() (*gorm.DB, error) {
	db, err := gorm.Open("mysql",
		"root:1234@tcp(localhost:3306)/mydb?parseTime=true&loc=Local")

	if err == nil {
		// 初始化 全局連接池句柄
		GlobalConn = db
		GlobalConn.DB().SetMaxIdleConns(10)
		GlobalConn.DB().SetConnMaxLifetime(100)

		db.SingularTable(true)
		db.AutoMigrate(new(User), new(House), new(Area), new(Facility), new(HouseImage), new(OrderHouse))
		return db, nil
	}
	return nil, err
}

// 初始化redis
var Rdb *redis.Client
var Ctx = context.Background()

func InitRdb() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	Rdb = rdb
}

// ===以下為定義的資料結構與常量等===

/* 用户 table_name = user */
type User struct {
	ID            int           //用户編號
	Name          string        `gorm:"size:32;unique"`  //用户名
	Password_hash string        `gorm:"size:128" `       //使用者密碼加密的
	Mobile        string        `gorm:"size:11;unique" ` //手機號
	Real_name     string        `gorm:"size:32" `        //真實姓名  實名認證
	Id_card       string        `gorm:"size:20" `        //身份證號  實名認證
	Avatar_url    string        `gorm:"size:256" `       //使用者頭像路徑       通過fastdfs進行圖片存儲
	Houses        []*House      //使用者發佈的房屋資訊  一個人多套房
	Orders        []*OrderHouse //使用者下的訂單       一個人多次訂單
}

/* 房屋信息 table_name = house */
type House struct {
	ID              uint      `gorm:"primary_key" json:"house_id"` //房屋編號
	CreatedAt       time.Time `json:"ctime"`
	UpdatedAt       time.Time
	DeletedAt       *time.Time    `sql:"index"`
	UserId          uint          `json:"user_id"`                     //房屋主人的用户編號  與用户進行關聯
	AreaId          uint          `json:"area_id"`                     //歸屬地的區域編號   和地區表進行關聯
	Title           string        `gorm:"size:64" json:"title"`        //房屋標題
	Address         string        `gorm:"size:512" json:"adress"`      //地址
	Room_count      int           `gorm:"default:1" json:"room_count"` //房間數目
	Acreage         int           `gorm:"default:0" json:"acreage"`    //房屋總面積
	Price           int           `json:"price"`
	Unit            string        `gorm:"size:32;default:''" json:"unit"`               //房屋單元,如 幾室幾廳
	Capacity        int           `gorm:"default:1" json:"capacity"`                    //房屋容納的總人數
	Beds            string        `gorm:"size:64;default:''" json:"beds"`               //房屋牀鋪的配置
	Deposit         int           `gorm:"default:0" json:"deposit"`                     //押金
	Min_days        int           `gorm:"default:1" json:"min_days"`                    //最少入住的天數
	Max_days        int           `gorm:"default:0" json:"max_days"`                    //最多入住的天數 0表示不限制
	Order_count     int           `gorm:"default:0" json:"order_count"`                 //預定完成的該房屋的訂單數
	Index_image_url string        `gorm:"size:256;default:''" json:"index_image_url"`   //房屋主圖片路徑
	Facilities      []*Facility   `gorm:"many2many:house_facilities" json:"facilities"` //房屋設施   與設施表進行關聯
	Images          []*HouseImage `json:"img_urls"`                                     //房屋的圖片   除主要圖片之外的其他圖片位址
	Orders          []*OrderHouse `json:"orders"`                                       //房屋的訂單    與房屋表進行管理
}

/* 區域資訊 table_name = area */ //區域資訊是需要我們手動添加到資料庫中的
type Area struct {
	Id     int      `json:"aid"`                  //區域編號     1    2
	Name   string   `gorm:"size:32" json:"aname"` //區域名字     昌平 海澱
	Houses []*House `json:"houses"`               //區域所有的房屋   與房屋表進行關聯
}

/* 設施資訊 table_name = "facility"*/ //設施資訊 需要我們提前手動添加的
type Facility struct {
	Id     int      `json:"fid"`     //設施編號
	Name   string   `gorm:"size:32"` //設施名字
	Houses []*House //都有哪些房屋有此設施  與房屋表進行關聯的
}

/* 房屋圖片 table_name = "house_image"*/
type HouseImage struct {
	Id      int    `json:"house_image_id"`      //圖片id
	Url     string `gorm:"size:256" json:"url"` //圖片url     存放我們房屋的圖片
	HouseId uint   `json:"house_id"`            //圖片所屬房屋編號
}

/* 訂單 table_name = order */
type OrderHouse struct {
	gorm.Model            //訂單編號
	UserId      uint      `json:"user_id"`       //下單的用户編號   //與用户表進行關聯
	HouseId     uint      `json:"house_id"`      //預定的房間編號   //與房屋資訊進行關聯
	Begin_date  time.Time `gorm:"type:datetime"` //預定的起始時間
	End_date    time.Time `gorm:"type:datetime"` //預定的結束時間
	Days        int       //預定總天數
	House_price int       //房屋的單價
	Amount      int       //訂單總金額
	Status      string    `gorm:"default:'WAIT_ACCEPT'"` //訂單狀態
	Comment     string    `gorm:"size:512"`              //訂單評論
	Credit      bool      //表示個人徵信情況 true表示良好
}

/*
這裡是gorm/v2的練習，由於太多錯誤(主要是結構體中有切片就報錯)跟不上教學
只好暫時停用，用回舊版

var GlobalConn *gorm.DB

func InitDb() (*gorm.DB, error) {
	// 連接mysql數據庫，?parseTime=true&loc=Local使用本地時區
	dsn := "root:root@tcp(127.0.0.1:3306)/ihome?parseTime=true&loc=Local"

	GlobalConn.SingularTable(true)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用單數表名，啓用該選項後，`User` 表將是`user`
		},
	})
	if err != nil {
		fmt.Println("gorm.open err", err)
		return nil, err
	}
	GlobalConn = db
	// 藉助gorm創建數據庫表
	err = db.AutoMigrate(new(User), new(House), new(Area), new(Facility), new(HouseImage), new(OrderHouse))
	if err != nil {
		fmt.Println("AutoMigrate err", err)
		return nil, err
	}
	return nil, err
}
*/
