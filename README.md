# ihome-愛家租房網 #
> 這是一個基於go語言的web實作練習項目

## 使用技術與練習重點 ##

- gin : 利用框架搭建web服務、RESTful風格API
- gorm : 基於orm架構，物件關係對映的練習
- MVC : 傳統設計模型，專案分層架構
- postman : web開發中用來測試API的好幫手
- mysql : sql語句使用、workbench資料匯出/匯入等操作
- redis : 緩存庫的觀念
- minio : 分散式儲存系統練習
- docker : 利用容器化的方式啟動mysql/redis/minio等服務
- git : 版本控制的基礎概念

特色tw本地化版本



### 基於版本 ###

golang : 1.17

mysql : 8

redis : 6

docker desktop : 4.1

github.com/jinzhu/gorm : 1.9.16 (這裡使用的是舊版，最新為gorm.io/gorm下的v2版本，兩者有較大差異且不兼容)



### 啟動 ###

1. 先在docker啟用3個資料庫服務

   ```
   docker run -d -it --name mysql8 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=1234 -e MYSQL_DATABASE=mydb mysql:8 --default-authentication-plugin=mysql_native_password
   ```

   ```
   docker run -d --name minio -e "MINIO_ROOT_USER=root" -e "MINIO_ROOT_PASSWORD=root1234" -p 9000:9000 -p 9001:9001 minio/minio server /data --console-address ":9001" 
   ```

   ```
   docker run --name redis6 -p 6379:6379 -d redis:6.0 redis-server --appendonly yes
   ```

   

2. go run ./web/main.go，此時會完成資料庫中表格的初始化

3. 透過mysql workbench手動將./home_tw.sql導入mydb

4. 開啟minio的控制台(localhost:9001)，將兩個預創立的bucket : avatar&house(存放頭像與房屋圖片用)權限設為public

5. 至首頁開始測試 http://127.0.0.1:22222/home/

6. 註冊帳戶->簡訊驗證碼功能，由於沒有實際串接業者服務，固定為1234

   

20211011-yoziming
