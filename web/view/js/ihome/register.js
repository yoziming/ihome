function getCookie(name) {
    var r = document.cookie.match("\\b" + name + "=([^;]*)\\b");
    return r ? r[1] : undefined;
}

function generateUUID() {
    var d = new Date().getTime();
    if(window.performance && typeof window.performance.now === "function"){
        d += performance.now(); //use high-precision timer if available
    }
    var uuid = 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
        var r = (d + Math.random()*16)%16 | 0;
        d = Math.floor(d/16);
        return (c=='x' ? r : (r&0x3|0x8)).toString(16);
    });
    return uuid;
}

// 生成一個圖片驗證碼的編號，並設置頁面中圖片驗證碼img標籤的src屬性
function generateImageCode() {
    // 生成一個編號
    // 嚴格一點的使用uuid保證編號唯一， 不是很嚴謹的情況下，也可以使用時間戳
    imageCodeId = generateUUID();

    // 設置頁面中圖片驗證碼img標籤的src屬性
    var imageCodeUrl = "/api/v1.0/imagecode/" + imageCodeId;
    $(".image-code>img").attr("src", imageCodeUrl);
}

function sendSMSCode() {
    // 校驗參數，保證輸入框有數據填寫
    $(".phonecode-a").removeAttr("onclick");
    var mobile = $("#mobile").val();
    if (!mobile) {
        $("#mobile-err span").html("請填寫正確的手機號！");
        $("#mobile-err").show();
        $(".phonecode-a").attr("onclick", "sendSMSCode();");
        return;
    } 
    var imageCode = $("#imagecode").val();
    if (!imageCode) {
        $("#image-code-err span").html("請填寫驗證碼！");
        $("#image-code-err").show();
        $(".phonecode-a").attr("onclick", "sendSMSCode();");
        return;
    }

    // 通過ajax方式向後端接口發送請求，讓後端發送短信驗證碼
    var req = {
        text: imageCode, // 用户填寫的圖片驗證碼
        id: imageCodeId // 圖片驗證碼的編號
    }
    $.get("/api/v1.0/smscode/"+mobile, req, function (resp) {
        // 表示後端發送短信成功
        if (resp.errno == "0") {
            // 倒計時60秒，60秒後允許用户再次點擊發送短信驗證碼的按鈕
            var num = 60;
            // 設置一個計時器
            var t = setInterval(function () {
                if (num == 1) {
                    // 如果計時器到最後, 清除計時器對象
                    clearInterval(t);
                    // 將點擊獲取驗證碼的按鈕展示的文本回覆成原始文本
                    $(".phonecode-a").html("獲取驗證碼");
                    // 將點擊按鈕的onclick事件函數恢復回去
                    $(".phonecode-a").attr("onclick", "sendSMSCode();");
                } else {
                    num -= 1;
                    // 展示倒計時信息
                    $(".phonecode-a").html(num+"秒");
                }
            }, 1000, 60)
        } else {
            // 表示後端出現了錯誤，可以將錯誤信息展示到前端頁面中
            $("#phone-code-err span").html(resp.errmsg);
            $("#phone-code-err").show();
            // 將點擊按鈕的onclick事件函數恢復回去
            $(".phonecode-a").attr("onclick", "sendSMSCode();");
        }

    }, "json");

}

$(document).ready(function() {
    generateImageCode();  // 生成一個圖片驗證碼的編號，並設置頁面中圖片驗證碼img標籤的src屬性
    $("#mobile").focus(function(){
        $("#mobile-err").hide();
    });
    $("#imagecode").focus(function(){
        $("#image-code-err").hide();
    });
    $("#phonecode").focus(function(){
        $("#phone-code-err").hide();
    });
    $("#password").focus(function(){
        $("#password-err").hide();
        $("#password2-err").hide();
    });
    $("#password2").focus(function(){
        $("#password2-err").hide();
    });
    $(".form-register").submit(function(e){
        // 阻止瀏覽器對於表單的默認行為，即阻止瀏覽器把表單的數據轉換為表單格式kye=val&key=val的字符串發送到後端
        e.preventDefault();
        var mobile = $("#mobile").val();
        var phoneCode = $("#phonecode").val();
        var passwd = $("#password").val();
        var passwd2 = $("#password2").val();
        if (!mobile) {
            $("#mobile-err span").html("請填寫正確的手機號！");
            $("#mobile-err").show();
            return;
        } 
        if (!phoneCode) {
            $("#phone-code-err span").html("請填寫短信驗證碼！");
            $("#phone-code-err").show();
            return;
        }
        if (!passwd) {
            $("#password-err span").html("請填寫密碼!");
            $("#password-err").show();
            return;
        }
        if (passwd != passwd2) {
            $("#password2-err span").html("兩次密碼不一致!");
            $("#password2-err").show();
            return;
        }

        // 構造發送到後端的數據 方式一
        var req = {
            "mobile": mobile,
            "password": passwd,
            "sms_code": phoneCode
        };

        // // 方式二
        // var req = {};
        // // 將表單中的全部字段值保存到req對象中
        // $(".form-register").serializeArray().map(function(x){req[x.name]=x.value});
        // req == {mobile: "18511111111", imagecode: "3qbv", phonecode: "246810", password: "123456", password2: "123456"}


        // 向後端發送註冊請求
        $.ajax({
            url: "/api/v1.0/users",
            type: "POST",
            contentType: "application/json",  // 指明發送到後端的數據格式是json
            data: JSON.stringify(req),
            headers: {
                "X-CSRFToken": getCookie("csrf_token") // 後端開啓了csrf防護，所以前端發送json數據的時候，需要包含這個請求頭
            },
            dataType: "json", // 指明後端返回到前端的數據是json格式的
            success: function(resp){
                if (resp.errno == "0") {
                    // 表示註冊成功,跳轉到主頁
                    location.href = "/home/index.html";
                } else if (resp.errno == "4101") {
                    // 表示用户註冊成功，但是用户的登錄狀態後端未保存，所以跳轉到登錄頁面
                    location.href = "/home/login.html";
                } else {
                    // 在頁面中展示錯誤信息
                    $("#password2-err span").html(resp.errmsg);
                    $("#password2-err").show();
                }
            }
        });

    });
})