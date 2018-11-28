function isPoneAvailable(str) {
    var myreg=/^[1][3,4,5,7,8][0-9]{9}$/;
    if (!myreg.test(str)) {
        return false;
    } else {
        return true;
    }
}


function login() {
    if ($("#uniq").val()==""){
        alert("请填写识别码！");
        return;
    }
    $.ajax({
        //几个参数需要注意一下
        cache: false,
        type: "POST",//方法类型
        dataType: "json",//预期服务器返回的数据类型
        url: "/login" ,
        data: $('#form-login').serialize(),
        success: function (result) {
            if (result.Result == true){
                alert("预约成功！");
            }else{
                alert("提交失败！"+ result.Info);
                $("#uniq").val("");
            }
        },
        error : function() {
            alert("提交异常！");
        }
    });
}




function xsSubmit() {
    alert($('#form1').serialize()+"&"+$('#form2').serialize()+"&"+$('#form3').serialize());
    // if ($("#contact_message").val()==""){
    //     //alert("请填写内容！");
    //     return;
    // }
    $.ajax({
        //几个参数需要注意一下
        cache: false,
        type: "POST",//方法类型
        dataType: "json",//预期服务器返回的数据类型
        url: "/xs" ,
        data: $('#form1').serialize()+"&"+$('#form2').serialize()+"&"+$('#form3').serialize(),
        success: function (result) {
            if (result.Result == true){
                $("#contact_phone").val("");
                $("#contact_message").val("");
                alert("预约成功！");
            }else{
                alert("提交失败！"+ result.Info);
            }
        },
        error : function() {
            alert("提交异常！");
        }
    });
}