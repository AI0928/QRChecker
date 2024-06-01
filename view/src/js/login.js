async function loginFunc() {
  const uid=document.getElementById('id');
  const pass=document.getElementById('pass');
  // alert("" + uid + pass);
  ///home?name=Management_Information_Systems=3=10000=IyodaAkira
  const url = 'http://localhost:1323/users'; //usersのJSON
  const response = await fetch(url);
  const json = await response.json();
  for (var i = 0; i < json.length; i++) {
    var email = json[i].email;
    var password = json[i].password;
    var id = json[i].user_id;
    if(email==uid.value&&password==pass.value){
        const url='./home?id=' + id;
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
