$(document).ready(function(){
    // 對於發佈房源，只有認證後的用户才可以，所以先判斷用户的實名認證狀態
    $.get("/api/v1.0/user/auth", function(resp){
        if ("4101" == resp.errno) {
            // 用户未登錄
            location.href = "/home/login.html";
        } else if ("0" == resp.errno) {
            // 未認證的用户，在頁面中展示 "去認證"的按鈕
            if (!(resp.data.real_name && resp.data.id_card)) {
                $(".auth-warn").show();
                return;
            }
            // 已認證的用户，請求其之前發佈的房源信息
            $.get("/api/v1.0/user/houses", function(resp){
                if ("0" == resp.errno) {
                    $("#houses-list").html(template("houses-list-tmpl", {houses:resp.data.houses}));
                } else {
                    $("#houses-list").html(template("houses-list-tmpl", {houses:[]}));
                }
            });
        }
    });
})