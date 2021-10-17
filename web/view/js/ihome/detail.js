function hrefBack() {
    history.go(-1);
}

// 解析提取url中的查詢字符串參數
function decodeQuery(){
    var search = decodeURI(document.location.search);
    return search.replace(/(^\?)/, '').split('&').reduce(function(result, item){
        values = item.split('=');
        result[values[0]] = values[1];
        return result;
    }, {});
}

$(document).ready(function(){
    // 獲取詳情頁面要展示的房屋編號
    var queryData = decodeQuery();
    var houseId = queryData["id"];

    // 獲取該房屋的詳細信息
    $.get("/api/v1.0/houses/" + houseId, function(resp){
        if ("0" == resp.errno) {
            $(".swiper-container").html(template("house-image-tmpl", {img_urls:resp.data.house.img_urls, price:resp.data.house.price}));
            $(".detail-con").html(template("house-detail-tmpl", {house:resp.data.house}));

            // resp.user_id為訪問頁面用户,resp.data.user_id為房東
            if (resp.data.user_id != resp.data.house.user_id) {
                $(".book-house").attr("href", "/home/booking.html?hid="+resp.data.house.hid);
                $(".book-house").show();
            }
            var mySwiper = new Swiper ('.swiper-container', {
                loop: true,
                autoplay: 2000,
                autoplayDisableOnInteraction: false,
                pagination: '.swiper-pagination',
                paginationType: 'fraction'
            });
        }
    })
})