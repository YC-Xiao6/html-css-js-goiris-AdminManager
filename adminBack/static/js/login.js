var flag = false;
var zkzcode; //在全局 定义验证码
function createCode() {
    zkzcode = "";
    var codeLength = 4;//验证码的长度
    var checkCode = document.getElementById("checkCode");//此处是显示验证码的地方
    var selectChar = new Array(1,2, 3, 4, 5, 6, 7, 8, 9,'A','B','C','D','E','F','G','H','J','K','M','N','P','Q','R','S','T','U','V','W','X','Y','Z');//所有候选组成验证码的字符，当然也可以用中文的
    for (var i = 0; i < codeLength; i++) {
        var charIndex = Math.floor(Math.random() * 31);
        zkzcode += selectChar[charIndex];
    }
    if (checkCode) {
        // checkCode.className = "zkzcode";
        checkCode.value = zkzcode;
    }
}

function checkCodes() {
    var inputCode = document.getElementById("inputCode");
    // var submit = document.getElementById("sub");
    var value = inputCode.value;
    if(value===zkzcode){
        flag=true;
        inputCode.className="login-check";
        document.getElementById("msg").innerText=""
    }else {
        flag=false;
        inputCode.className="login-checkErr";
        // submit.setAttribute("type","button")
        document.getElementById("msg").innerText="验证码错误，请重试！"
    }
}
function loginSubmit() {
    var nullflag = false
    //判断表单数据是否为空
    var t = $('#form').serializeArray();
    //序列化表单数据
    var jsonObj = serializeForm($('#form'))
    //将序列化好的数据字符串化
    var jsonStr = JSON.stringify(jsonObj)
    $.each(t,function(i,item){
        if(item['value'] === '') {
            document.getElementById("msg").innerText="请输入完整信息！"
            nullflag = false
        }else {
            nullflag = true
        }
    })
    if(flag&&nullflag) {
        $.ajax({
            type: "POST",//方法类型
            dataType: "JSON",//预期服务器返回的数据类型
            url: "/admin/login" ,//url
            contentType: "application/json;charset=utf-8",//(可以)//请求需要发送的处理数据  编码格式
            //后台只接收json格式，因此要JSON.stringify格式化
            data: jsonStr,
            crossDomain : true, //跨域问题
            async: false,
            //post 提交的数据
            success: function (result) {
                if (result.status) {
                    window.location.href = "/index";
                }else {
                    document.getElementById("msg").innerText=result.message;
                }
            },
            error : function() {
                alert("请求失败")
            }
        });
    }
}
//序列化表单数据的函数
function serializeForm(form){
    var obj = {};
    $.each(form.serializeArray(),function(index){
        if(obj[this['name']]){
            obj[this['name']] = obj[this['name']] + ','+this['value'];
        } else {
            obj[this['name']] =this['value'];
        }
    });
    return obj;
}
