function getCookie(name) {
    var r = document.cookie.match("\\b" + name + "=([^;]*)\\b");
    return r ? r[1] : undefined;
}

$(document).ready(function() {
    $("#mobile").focus(function(){
        $("#mobile-err").hide();
    });
    $("#password").focus(function(){
        $("#password-err").hide();
    });
    $(".form-login").submit(function(e){
        // 阻止瀏覽器對於表單的默認提交行為
        e.preventDefault();
        var mobile = $("#mobile").val();
        var passwd = $("#password").val();
        if (!mobile) {
            $("#mobile-err span").html("請填寫正確的手機號！");
            $("#mobile-err").show();
            return;
        } 
        if (!passwd) {
            $("#password-err span").html("請填寫密碼!");
            $("#password-err").show();
            return;
        }
        // 將表單的數據存放到對象data中
        var data = {};
        $(this).serializeArray().map(function(x){data[x.name] = x.value;});
        // 將data轉為json字符串
        var jsonData = JSON.stringify(data);
        $.ajax({
            url:"/api/v1.0/sessions",
            type:"post",
            data: jsonData,
            contentType: "application/json",
            dataType: "json",
            headers:{
                "X-CSRFTOKEN":getCookie("csrf_token"),
            },
            success: function (data) {
                if ("0" == data.errno) {
                    // 登錄成功，跳轉到主頁
                    location.href = "/home/";
                    return;
                }
                else {
                    // 其他錯誤信息，在頁面中展示
                    $("#password-err span").html(data.errmsg);
                    $("#password-err").show();
                    return;
                }
            }
        });
    });
})