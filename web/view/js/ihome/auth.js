function showSuccessMsg() {
    $('.popup_con').fadeIn('fast', function() {
        setTimeout(function(){
            $('.popup_con').fadeOut('fast',function(){}); 
        },1000) 
    });
}


function getCookie(name) {
    var r = document.cookie.match("\\b" + name + "=([^;]*)\\b");
    return r ? r[1] : undefined;
}

$(document).ready(function(){
    // 查詢用户的實名認證信息
    $.get("/api/v1.0/user/auth", function(resp){
        // 4101代表用户未登錄
        if ("4101" == resp.errno) {
            location.href = "/home/login.html";
        }
        else if ("0" == resp.errno) {
            // 如果返回的數據中real_name與id_card不為null，表示用户有填寫實名信息
            if (resp.data.real_name && resp.data.id_card) {
                $("#real-name").val(resp.data.real_name);
                $("#id-card").val(resp.data.id_card);
                // 給input添加disabled屬性，禁止用户修改
                $("#real-name").prop("disabled", true);
                $("#id-card").prop("disabled", true);
                // 隱藏提交保存按鈕
                $("#form-auth>input[type=submit]").hide();
            }
        }
    }, "json");

    // 管理實名信息表單的提交行為
    $("#form-auth").submit(function(e){
        e.preventDefault();
        // 如果用户沒有填寫完整，展示錯誤信息
        if ($("#real-name").val()=="" || $("#id-card").val() == "") {
            $(".error-msg").show();
            return
        }

        // 將表單的數據轉換為json字符串
        var data = {};
        $(this).serializeArray().map(function(x){data[x.name] = x.value;});
        var jsonData = JSON.stringify(data);

        // 向後端發送請求
        $.ajax({
            url:"/api/v1.0/user/auth",
            type:"post",
            data: jsonData,
            contentType: "application/json",
            dataType: "json",
            headers: {
                "X-CSRFTOKEN": getCookie("csrf_token")
            },
            success: function (resp) {
                if (0 == resp.errno) {
                    $(".error-msg").hide();
                    // 顯示保存成功的提示信息
                    showSuccessMsg();
                    $("#real-name").prop("disabled", true);
                    $("#id-card").prop("disabled", true);
                    $("#form-auth>input[type=submit]").hide();
                }else {
                    // $(".error-msg").show();
                    $(".error-msg").html('<i class="fa fa-exclamation-circle"></i>信息不符,驗證失敗,請重新填寫!</div>').show();
                }
            }
        });
    })

})