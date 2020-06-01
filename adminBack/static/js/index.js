//判断访问哪个功能区
function adminOnclick() {
    document.getElementById("one").classList.add("active");
    document.getElementById("two").classList.remove("active");
    document.getElementById("indexIframe").setAttribute("src","table.html")
}
function nextOnclick() {
    document.getElementById("one").classList.remove("active");
    document.getElementById("two").classList.add("active");
    document.getElementById("indexIframe").setAttribute("src","new.html")
}
//新增修改时的判断
function checkName() {
    var name = $("#name").val();
    //正则表达式，不是字符串
    var reg = /^[A-Za-z0-9_\\-\u4e00-\u9fa5]{1,20}$/;
    if (reg.test(name)){
        $("#name").removeClass("err");
        return true;
    }else {
        $("#name").addClass("err");
        return false;
    }
}
function checkPasswd() {
    var Passwd = $("#passwd").val();
    //正则表达式，不是字符串
    var reg = /^[A-Za-z]\w{7,17}$/;
    if (reg.test(Passwd)){
        $("#passwd").removeClass("err");
        return true;
    }else {
        $("#passwd").addClass("err");
        return false;
    }
}
function checkBirthday() {
    var Birthday = $("#birthday").val();
    //正则表达式，不是字符串
    var reg = /^\d{4}-\d{1,2}-\d{1,2}$/;
    if (reg.test(Birthday)){
        $("#birthday").removeClass("err");
        return true;
    }else {
        $("#birthday").addClass("err");
        return false;
    }
}
function checkPhone() {
    var Phone = $("#phone").val();
    //正则表达式，不是字符串
    var reg = /^0?(13|14|15|17|18|19)[0-9]{9}$/;
    if (reg.test(Phone)){
        $("#phone").removeClass("err");
        return true;
    }else {
        $("#phone").addClass("err");
        return false;
    }
}
function checkEmail() {
    var Email = $("#email").val();
    //正则表达式，不是字符串
    var reg = /^\w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\.)+[A-Za-z]{2,14}$/;
    if (reg.test(Email)){
        $("#email").removeClass("err");
        return true;
    }else {
        $("#email").addClass("err");
        return false;
    }
}
function checkAddr() {
    var Addr = $("#addr").val();
    var reg = /^.{0,150}$/;
    // if (Addr!=null&&Addr!=''){
    if(reg.test(Addr)){
        $("#addr").removeClass("err");
        return true;
    }else {
        $("#addr").addClass("err");
        return false;
    }
}
//添加用户
function subAdd() {
    if (checkName() && checkPasswd() && checkBirthday() && checkPhone() && checkEmail() && checkAddr()) {
        $("#msg").html("");
        var nullflag = false
        //判断表单数据是否为空
        var t = $('#form').serializeArray();
        //序列化表单数据
        var jsonObj = serializeForm($('#form'))
        //将序列化好的数据字符串化
        var jsonStr = JSON.stringify(jsonObj)
        $.each(t, function (i, item) {
            if (item['value'] === '') {
                document.getElementById("msg").innerText = "请输入完整信息！"
                nullflag = false
            } else {
                nullflag = true
            }
        })
        if (nullflag) {
            $.ajax({
                type: "POST",//方法类型
                dataType: "JSON",//预期服务器返回的数据类型
                url: "/admin/addAdmin",//url
                contentType: "application/json;charset=utf-8",//(可以)//请求需要发送的处理数据  编码格式
                //后台只接收json格式，因此要JSON.stringify格式化
                data: jsonStr,
                crossDomain: true, //跨域问题
                async: false,
                //post 提交的数据
                success: function (result) {
                    if (result.status) {
                        window.location.href = "/table.html";
                    } else {
                        document.getElementById("msg").innerText = result.message;
                    }
                },
                error: function () {
                    alert("请求失败")
                }
            });
        }
    }
    else {
        $("#msg").html("请验证输入的格式是否正确");
    }
}
//修改用户
function subModify() {
    if (checkName() && checkPasswd() && checkBirthday() && checkPhone() && checkEmail() && checkAddr()) {
        $("#msg").html("");
        var nullflag = false
        //判断表单数据是否为空
        var t = $('#modifyfrom').serializeArray();
        //序列化表单数据
        var jsonObj = serializeForm($('#modifyfrom'))
        //将序列化好的数据字符串化
        var jsonStr = JSON.stringify(jsonObj)
        //表单获取不到id值，只能用传参方法进行
        var url = "/admin/modifyAdmin?id="+$("#id").val()
        $.each(t, function (i, item) {
            if (item['value'] === '') {
                document.getElementById("msg").innerText = "请输入完整信息！"
                nullflag = false
            } else {
                nullflag = true
            }
        })
        if (nullflag) {
            $.ajax({
                type: "POST",//方法类型
                dataType: "JSON",//预期服务器返回的数据类型
                url: url,//url
                contentType: "application/json;charset=utf-8",//(可以)//请求需要发送的处理数据  编码格式
                //后台只接收json格式，因此要JSON.stringify格式化
                data: jsonStr,
                crossDomain: true, //跨域问题
                async: false,
                //post 提交的数据
                success: function (result) {
                    if (result.status) {
                        window.location.href = "/table.html";
                    } else {
                        document.getElementById("msg").innerText = result.message;
                    }
                },
                error: function () {
                    alert("请求失败")
                }
            });
        }
    }
    else {
        $("#msg").html("请验证输入的格式是否正确");
    }
}
var limit = 2;
//模糊查询名称
function queryByName(curr) {
    var name = $("#findName").val();
    if(name!=""&&name!=null){
        var url = "/admin/queryByName?name="+name+"and"+curr+"and"+limit;
        // url = encodeURI(url);
        $.ajax({
            type:"GET",
            dataType: "json",//预期服务器返回的数据类型
            url: url ,//url
            crossDomain : true, //跨域问题
            success: function (result) {
                if (result.status) {
                    makelist(result.admin)
                    var count = result.total;
                    laypage(count,limit,2,curr)
                }else {
                    alert("未匹配到相关信息，请重试!")
                }
            },
            error : function() {
                alert("异常！");
            }
        })
    }else {
        alert("请输入要查询的名称！")
    }

}
//根据性别查询
function queryBySex(curr) {
    var sex = $("#findSex").val();
    var url = "/admin/queryBySex?sex="+sex+"and"+curr+"and"+limit;
    $.ajax({
        type:"GET",
        dataType:"JSON",
        url:url,
        success:function (result) {
            if (result.status){
                makelist(result.admin)
                var count = result.total;
                laypage(count,limit,3,curr)
            }else {
                alert("未匹配到相关信息，请重试!")
            }
        },
        error:function () {
            alert("异常")
        }

    })

}
//根据生日日期查询
function queryByBirth(curr) {
    if($("#findBirth1").val()!=""||$("#findBirth2").val()!=""){
        var date = $("#findBirth1").val()+"and"+$("#findBirth2").val();
        var url = "/admin/queryByDate?date="+date+"and"+curr+"and"+limit;;
        $.ajax({
            type:"GET",
            dataType:"JSON",
            url:url,
            success:function (result) {
                if (result.status){
                    makelist(result.admin)
                    var count = result.total;
                    laypage(count,limit,4,curr)
                }else {
                    alert("未匹配到相关信息，请重试!")
                }
            },
            error:function () {
                alert("异常")
            }
        })
    }else {
        alert("请选择日期期限！")
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
//分页的条数

//table获取数据库列表
function queryAll(curr) {
    var url = "/admin/index?curr="+curr+"and"+limit;
    $.ajax({
        type: "GET",//方法类型
        dataType: "json",//预期服务器返回的数据类型
        url: url ,//url
        crossDomain : true, //跨域问题
        //异步请求
        async:false,
        //post 提交的数据
        success: function (result) {
            if (result.status) {
                makelist(result.admin);
                var count = result.total;
                laypage(count,limit,1,curr)
            }
        },
        error : function() {
            alert("异常！");
        }
    });
}
//按创建时间排序
function queryAllByCreat(curr){
    var url ="/admin/queryAllByCreat?curr="+curr+"and"+limit;
    $.ajax({
        type: "GET",//方法类型
        dataType: "json",//预期服务器返回的数据类型
        url: url ,//url
        crossDomain : true, //跨域问题
        //post 提交的数据
        success: function (result) {
            if (result.status) {
                makelist(result.admin)
                var count = result.total;
                laypage(count,limit,5,curr)
            }
        },
        error : function() {
            alert("异常！");
        }
    });
}
//按修改时间排序
function queryAllByModify(curr){
    var url ="/admin/queryAllByModify?curr="+curr+"and"+limit;
    $.ajax({
        type: "GET",//方法类型
        dataType: "json",//预期服务器返回的数据类型
        url: url ,//url
        crossDomain : true, //跨域问题
        //post 提交的数据
        success: function (result) {
            if (result.status) {
                makelist(result.admin)
                var count = result.total;
                laypage(count,limit,6,curr)
            }
        },
        error : function() {
            alert("异常！");
        }
    });
}
//写表单
function makelist(admins) {
    var lis = ''
    for (i=0;i<admins.length;i++){
        var admin = admins[i]
        var li = '<tr>\n            <td style="width: 3%">'+(i+1)+'</td>\n        ' +
            '    <td style="width: 7%">'+admin.name+'</td>\n            <td style="width: 14%">'+admin.passwd+'</td>\n        ' +
            '    <td style="width: 8%">'+admin.birthday+'</td>\n            <td style="width: 3%">'+admin.sex+'</td>\n       ' +
            '     <td style="width: 10%">'+admin.phone+'</td>\n            <td style="width: 10%">'+admin.email+'</td>\n   ' +
            '         <td style="width: 16%">'+admin.addr+'</td>\n            <td style="width: 3%">'+admin.education+'</td>\n           ' +
            ' <td style="width: 9%">'+admin.create_time+'</td>\n            <td style="width: 9%">'+admin.last_time+'</td>\n      ' +
            '      <td style="width: 5%">\n  <input class="tableButton" type="button" value="修改" onclick="location=\'modify.html?id='+admin.id+'\'">\n             ' +
            '   <input class="tableButton" type="button" value="删除" onclick="deleteAdminByid('+admin.id+')">\n            </td>\n        </tr>'
        lis +=li;
    }
    $("#tbody").html(lis);
    var src = "location='add.html?id="+(admins.length+1)+"'";
    $("#adds").attr("onclick",src);
}

//根据id删除信息
function deleteAdminByid(id) {
    var url = "/admin/delete?id="+id
    $.ajax({
                type: "GET",//方法类型
                dataType: "json",//预期服务器返回的数据类型
                url: url,//url
                contentType:"application/x-www-form-urlencoded",
                async :true,
                success: function (result) {
                    if (result.status) {
                        location.reload();
                    }else {
                        alert(result.message);
                    }
                },
                error : function() {
                    alert("异常！");
                }
    });
}

//获取modify的信息
function modifyInfo() {
    //根据url获取id字段
    var ids = getQueryString("id")
    var url = "/admin/modifyInfo?id="+ids
    $.ajax({
        type:"GET",
        dataType:"json",
        url : url,
        async :true,
        success:function (result) {
            var admin = result.admin;
            var li = '<h2 >修改用户</h2>\n  ' +
                '  <div><label>I&nbsp;&nbsp;&nbsp;D：<input class="add-input" id="id" name="id" type="text" placeholder="'+ids+'" disabled value="'+ids+'"></label></div>\n  ' +
                '  <div><label>名称：<input class="add-input" name="name" id="name" type="text" placeholder="请输入名称" value="'+admin.name+'"></label></div>\n  ' +
                '  <div><label>密码：<input class="add-input" name="passwd" id="passwd"  type="passwd" placeholder="请输入密码" value="'+admin.passwd+'"></label></div>\n   ' +
                ' <div><label>生日：<input class="add-input" name="birthday" id="birthday" type="date" value="'+admin.birthday+'"></label></div>\n    ' +
                '<div class="add-sellects"><label>性别：<select name="sex" id="sex">\n     ' +
                '           <option value="男" >男</option>\n   <option value="女" >女</option>\n            </select></label>\n        &nbsp;&nbsp;&nbsp;&nbsp;\n       ' +
                ' <label>学历：<select name="education" id="education">\n              ' +
                '  <option value="博士">博士</option>\n   <option value="硕士">硕士</option>\n     <option  value="本科">本科</option>\n    <option value="专科">专科</option>\n        <option value="高中">高中</option>\n            </select></label></div>\n  ' +
                '  <div><label>电话：<input class="add-input" name="phone" id="phone"  type="text" placeholder="请输入手机号" value="'+admin.phone+'"></label></div>\n   ' +
                ' <div><label>邮箱：<input class="add-input" name="email" id="email" type="text" placeholder="请输入邮箱" value="'+admin.email+'"></label></div>\n ' +
                '   <div><label>地址：<input class="add-input" name="addr" id="addr" type="text" placeholder="请输入地址" value="'+admin.addr+'"></label></div>\n   ' +
                ' <div class="submit"><input type="button"  class="login-submit" onclick="location=\'table.html\'" value="返回">\n       ' +
                ' <input type="button" id="sub" class="login-submit"  value="提交"></div>\n    ' +
                '<div class="msg"><span id="msg"></span></div>';
            $("#modifyfrom").html(li);
            $("#sex").val(admin.sex);
            $("#education").val(admin.education);
            $("#name").blur(checkName);
            $("#passwd").blur(checkPasswd);
            $("#birthday").blur(checkBirthday);
            $("#phone").blur(checkPhone);
            $("#email").blur(checkEmail);
            $("#addr").blur(checkAddr);
            $("#sub").click(subModify)
        },
        error:function () {
            alert("异常")
        }
    })
}
//获取url后面的字段信息，传入参数键值对key
function getQueryString(name) {
    //区分大小写匹配
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    //从url第一位开始进行正则匹配
    var r = window.location.search.substr(1).match(reg);

    //r:id=5,,5,
    if (r != null) return unescape(r[2]);
    return null;
}

//获取addIndex
function addIndex() {
    //根据url获取id字段
    var id = getQueryString("id");
    $("#id").val(id);
}

