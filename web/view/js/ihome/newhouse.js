function getCookie(name) {
    var r = document.cookie.match("\\b" + name + "=([^;]*)\\b");
    return r ? r[1] : undefined;
}

$(document).ready(function(){
    // $('.popup_con').fadeIn('fast');
    // $('.popup_con').fadeOut('fast');
    $.get("/api/v1.0/areas", function (resp) {
        if ("0" == resp.errno) {
            // // 表示查詢到了數據,修改前端頁面
            // for (var i=0; i<resp.data.length; i++) {
            //     // 向頁面中追加標籤
            //     var areaId = resp.data[i].aid;
            //     var areaName = resp.data[i].aname;
            //     $("#area-id").append('<option value="'+ areaId +'">'+ areaName +'</option>');
            // }

            // 使用前端模板
            rendered_html = template("areas-tmpl", {areas: resp.data});
            $("#area-id").html(rendered_html);
        } else {
            alert(resp.errmsg);
        }
    }, "json");

    // 處理房屋基本信息的表單數據
    $("#form-house-info").submit(function (e) {
        e.preventDefault();
        // 檢驗表單數據是否完整
        // 將表單的數據形成json，向後端發送請求
        var formData = {};
        $(this).serializeArray().map(function (x) { formData[x.name] = x.value });

        // 對於房屋設施的checkbox需要特殊處理
        var facility = [];
        // $("input:checkbox:checked[name=facility]").each(function(i, x){ facility[i]=x.value });
        $(":checked[name=facility]").each(function(i, x){ facility[i]=x.value });

        formData.facility = facility;

        // 使用ajax向後端發送請求
        $.ajax({
            url: "/api/v1.0/houses",
            type: "post",
            data: JSON.stringify(formData),
            contentType: "application/json",
            dataType: "json",
            headers: {
                "X-CSRFToken": getCookie("csrf_token")
            },
            success: function(resp){
                if ("4101" == resp.errno) {
                    location.href = "/home/login.html";
                } else if ("0" == resp.errno) {
                    // 後端保存數據成功
                    // 隱藏基本信息的表單
                    $("#form-house-info").hide();
                    // 顯示上傳圖片的表單
                    $("#form-house-image").show();
                    // 設置圖片表單對應的房屋編號那個隱藏字段
                    $("#house-id").val(resp.data.house_id);
                } else {
                    alert(resp.errmsg);
                }
            }
        });
    })

    // 處理圖片表單的數據
    $("#form-house-image").submit(function (e) {
        e.preventDefault();
        var house_id = $("#house-id").val();
        // 使用jquery.form插件，對錶單進行異步提交，通過這樣的方式，可以添加自定義的回調函數
        $(this).ajaxSubmit({
            url: "/api/v1.0/houses/"+house_id+"/images",
            type: "post",
            headers: {
                "X-CSRFToken": getCookie("csrf_token")
            },
            dataType: "json",
            success: function (resp) {
                if ("4101" == resp.errno) {
                    location.href = "/home/login.html";
                } else if ("0" == resp.errno) {
                    // 在前端中添加一個img標籤，展示上傳的圖片
                    $(".house-image-cons").append('<img src="'+ resp.data.url+'">');
                } else {
                    alert(resp.errmsg);
                }
            }
        })
    })


})