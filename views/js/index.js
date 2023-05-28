var comein = document.getElementById('comein');
var regist = document.getElementById('regist');
var zhuce = document.getElementsByClassName('zhuce')[0];
regist.onclick = function () {
    document.getElementsByClassName('login')[0].style.display = "none";
    document.getElementsByClassName('register')[0].style.display = "block";
}
comein.onclick = function () {
    document.getElementsByClassName('register')[0].style.display = "none";
    document.getElementsByClassName('login')[0].style.display = "block";
}

var form2 = document.getElementsByClassName('form2');
zhuce.onclick = function () {
    var password = document.getElementById('zcpassword').value;
    var password2 = document.getElementById('zcpassword-again').value;
    var zcusername = document.getElementById('zcusername').value;
    if (zcusername == '' || zcusername == null) {
        alert('注册名不能为空，请重新输入！');
        form2.reset();
        return;
    } else if (password == '' || password == null || password2 == '' || password2 == null) {
        alert('密码不能为空，请重新输入！');
        form2.reset();
        return;
    } else if (password !== password2) {
        alert('请确认输入的密码是否相同！');
        form2.reset();
        return;
    } else {
        var zcusername = document.getElementById('zcusername').value;
        var zcpassword = document.getElementById('zcpassword').value;
        var zcpassword2 = document.getElementById('zcpassword-again').value;
        var obj = {};
        obj.zcusername = zcusername;
        obj.zcpassword = zcpassword;
        obj.zcpassword2 = zcpassword2;
        var data = JSON.stringify(obj);

        // var formdata = new FormData();
        // formdata.append("zcusername", zcusername);
        // formdata.append("zcpassword", zcpassword);
        // formdata.append("zcpassword2", zcpassword2);
        // console.log(formdata.get("zcpassword"));
        
        var xhr = new XMLHttpRequest(); 
        xhr.open("post", "http://127.0.0.1:8080/regist", true);
        xhr.send(data);
        // xhr.send(formdata);//发送数据
        xhr.onreadystatechange = function () {
            if (xhr.readyState == 4 && xhr.status >= 200) {
                // if ((xhr.status >=200 && xhr.status < 300) || xhr.status == 304){
                console.log(xhr.responseText);
                var obj=JSON.parse(xhr.responseText);
                if (obj.flag == 1) {
                    alert('注册成功！');
                    document.getElementsByClassName('register')[0].style.display = "none";
                    document.getElementsByClassName('login')[0].style.display = "block";
                }else{
                    alert('注册失败');
                }
                // }else{
                //     for(var i=0; i<obj.length;i++){
                //         if(zcusername == obj[i].name){
                //             alert('该用户名已被注册过，请重新输入！');
                //         }
                //     }
                // }
                // }
            } else {
                console.log("请求不成功" + xhr.status);
            }
        }
    }
}
var denglv = document.querySelector('.denglv');
denglv.onclick = function () {
    console.log('111');
    // denglv.href = "jump.html";
    var dlusername = document.getElementById('dlusername').value;
    var dlpassword = document.getElementById('dlpassword').value;
    var obj = {};
    obj.dlusername = dlusername;
    obj.dlpassword = dlpassword;
    var data = JSON.stringify(obj);
    // var formdata = new FormData();
    // formdata.append("dlusername", dlusername);
    // formdata.append("dlpassword", dlpassword);
    var xhr = new XMLHttpRequest();
    xhr.open("post", "http://127.0.0.1:8080/login", true);
    xhr.send(data);
    // xhr.send(formdata);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4 && xhr.status >= 200) {
            console.log(xhr.responseText);
            var obj=JSON.parse(xhr.responseText);
            if (obj.flag == 1) {
                window.location.href="jump.html";
                // self.location.href = "";
            } else {
                alert('登录失败！');
            }

        } else {
            console.log("请求不成功" + xhr.status);
        }
    }
}