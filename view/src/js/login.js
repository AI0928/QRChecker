alert("login!!");
async function loginFunc() {
  alert("確認");
  const uid=document.getElementById('id');
  const pass=document.getElementById('pass');
  // alert("" + uid + pass);
  ///home?name=Management_Information_Systems=3=10000=IyodaAkira
  const url = 'http://localhost:1323/users'; //usersのJSON
  // fetch(url)
  // .then(function (data) {
  //   return data.json(); // 読み込むデータをJSONに設定
  // })
  // .then(function (json) {
  //   // alert(json)
  //   for (var i = 0; i < json.length; i++) {
  //     var major = json[i].major;
  //     var grade= json[i].grade;
  //     var user_id= json[i].user_id;
  //     var name = json[i].name;
  //     var email = json[i].email;
  //     var password = json[i].password;

  //     if(email==uid.value&&password==pass.value){
  //         const url='./home?name='+major+'='+grade+'='+user_id+'='+name;
  //         window.location.href = url;
  //         return;
  //     }
  //   }
  //   alert("ログイン失敗");
  // });
  const response = await fetch(url);
  const json = await response.json();
  //console.log(json, json.length)
  for (var i = 0; i < json.length; i++) {
    var major = json[i].major;
    var grade= json[i].grade;
    var user_id= json[i].user_id;
    var name = json[i].name;
    var email = json[i].email;
    var password = json[i].password;

    if(email==uid.value&&password==pass.value){
        const url='./home?name='+major+'='+grade+'='+user_id+'='+name;
        window.location.href = url;
        return;
    }
    else{
        alert("ログイン失敗");
    }
  }
}

document.getElementById('form').addEventListener('submit', function() {
  event.preventDefault(); // フォームのデフォルトの送信動作をキャンセル  
  loginFunc();
});
