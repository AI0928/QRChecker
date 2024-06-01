window.addEventListener('load', () => {
    startQR();
});
const user_id = location.search.split('=')[1];
const syusseki = "出席";
const video = document.createElement('video');
const canvasElement = document.getElementById('canvas');
const canvas = canvasElement.getContext('2d');
const loadingMessage = document.getElementById('loadingMessage');
const outputContainer = document.getElementById('output');
const outputMessage = document.getElementById('outputMessage');
const myHeaders = new Headers();
myHeaders.append('Content-Type', 'application/x-www-form-urlencoded; charset=utf-8');
myHeaders.append('Access-Control-Allow-Origin', 'http://localhost:1323');

const startQR = () => {
    // console.log("startQR");
    navigator.mediaDevices.getUserMedia({
        video: {
            audio: false,
            facingMode: 'environment'//'user' でインカメを使う
        }
    }).then((stream) => {
        video.srcObject = stream;
        video.setAttribute('playsinline', true);
        video.play();
        requestAnimationFrame(tick);
    }).catch((err) => {
        alert(err.message)
    })
};


//QRの解析
function tick() {
    // console.log("tick");
    loadingMessage.innerHTML = 'Loading video...';
    if (video.readyState === video.HAVE_ENOUGH_DATA) {
        loadingMessage.hidden = true;
        canvasElement.hidden = false;
        outputContainer.hidden = false;

        canvasElement.height = video.videoHeight;
        canvasElement.width = video.videoWidth;
        canvas.drawImage(video, 0, 0, canvasElement.width, canvasElement.height);
        //CanvasのgetImageDataメソッドで指定範囲のImageDataオブジェクトを取得する
        const imageData = canvas.getImageData(0, 0, canvasElement.width, canvasElement.height);
        //jsQRのメソッド
        const code = jsQR(imageData.data, imageData.width, imageData.height, {
            inversionAttempts: 'dontInvert',
        });

        //QRが読み込めた時の挙動
        if (code) {
            drawLine(code.location.topLeftCorner, code.location.topRightCorner, "#FF3B58");
            drawLine(code.location.topRightCorner, code.location.bottomRightCorner, "#FF3B58");
            drawLine(code.location.bottomRightCorner, code.location.bottomLeftCorner, "#FF3B58");
            drawLine(code.location.bottomLeftCorner, code.location.topLeftCorner, "#FF3B58");
            outputMessage.hidden = true;
            video.style.display = 'none';
            video.pause();
            
            // time_check(code.data);
            put(code.data, syusseki)
        } else {
            outputMessage.hidden = false;
        }
    }
    requestAnimationFrame(tick);
};


//QRを囲むライン
const drawLine = (begin, end, color) => {
    // console.log("drawLine");
    canvas.beginPath();
    canvas.moveTo(begin.x, begin.y);
    canvas.lineTo(end.x, end.y);
    canvas.lineWidth = 4;
    canvas.strokeStyle = color;
    canvas.stroke();
};

//off
const videoOff = () => {
    video.pause();
    video.src = "";
    video.srcObject.getTracks()[0].stop();
};

async function time_check(data){  
    var now = new Date();
    var HourValue = now.getHours();
    var MinValue = now.getMinutes();
    var SecValue = now.getSeconds();
    // var now_time = HourValue + ":" + MinValue + ":" + SecValue;
    var code_texts = data.split('=')
    const rurl_lecture = 'http://localhost:1323/lecture/'+code_texts[0];
    const getparameter = {
        method: 'GET',
        headers: myHeaders,
    }
    const lecture = await fetch(rurl_lecture, getparameter).then((response) => {
        return response.json();
    });
    start_time = lecture.start_time.split(':');
    lecture_time = 90
    n_time = HourValue * 60 + MinValue;
    console.log(start_time, HourValue, MinValue, SecValue);
    s_time = parseInt(start_time[0]) * 60 + parseInt(start_time[1]);
    e_time = parseInt(end_time[0]) * 60 + parseInt(end_time[1]);
    if (s_time <= n_time && s_time + 3 >= n_time){
            atenndanceStatus = '出席';
    } else if (s_time <= n_time && s_time + lecture_time / 2 >= n_time){
            atenndanceStatus = '遅刻';
    } else {
        atenndanceStatus = '欠席';
    }
}

async function put(a, atenndanceStatus) {
    // console.log("put");
    const value = a.split('=');
    const lecture_id = value[0];
    const count = value[1];
    // e.PUT("/attendances/:lecture_id/:user_id/:user_name/:count", updateAttendance)
    const rurl = 'http://localhost:1323/attendances/'+lecture_id+'/'+user_id+'/'+count + '/' + atenndanceStatus;
    const rurl_lecture = 'http://localhost:1323/lecture/'+lecture_id;
    console.log(rurl)

    const putparameter = {
        method: 'PUT',
        headers: myHeaders,
    }

    const result = await fetch(rurl, putparameter).then((response) => {
        return response.json();
    });
    console.log(result)
    const getparameter = {
        method: 'GET',
        headers: myHeaders,
    }

    const lecture = await fetch(rurl_lecture, getparameter).then((response) => {
        return response.json();
    });

    const url='./check?id='+'='+user_id+'='+lecture.name+'='+count+'='+atenndanceStatus;
    window.location.href = url;
}

