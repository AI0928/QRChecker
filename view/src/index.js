const http = require('http');
const fs = require('fs');
const url = require('url');

const indexPage = fs.readFileSync('./index.html', 'UTF-8');
const otherPage1 = fs.readFileSync('./seito.html', 'UTF-8');
const otherPage2 = fs.readFileSync('./sensei.html', 'UTF-8');
const loginPage = fs.readFileSync('./login.html', 'UTF-8');
const testPage = fs.readFileSync('./test.html', 'UTF-8');
const styleCss = fs.readFileSync('./index.css', 'UTF-8');
const tableCss = fs.readFileSync('./table.css', 'UTF-8');
const scriptJs = fs.readFileSync('./js/test.js', 'UTF-8');
const login = fs.readFileSync('./js/login.js', 'UTF-8');
const home = fs.readFileSync('./home.html', 'UTF-8');
const test1 = fs.readFileSync('./test1.html', 'UTF-8');
const attend = fs.readFileSync('./attend.html', 'UTF-8');

const hostname = '0.0.0.0'; //ローカル なら "127.0.0.1"
const port = 8000;

const server = http.createServer(RouteSetting);

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});

function RouteSetting(req, res) {
  const url_parts = url.parse(req.url);
  switch (url_parts.pathname) {
    case '/':
    case '/login':
      res.writeHead(200, {'Content-Type': 'text/html'});
      res.write(loginPage);
      res.end();
      break;
    
    case '/test1':
      res.writeHead(200, {'Content-Type': 'text/html'});
      res.write(test1);
      res.end();
      break;  
    
    case '/attend':
      res.writeHead(200, {'Content-Type': 'text/html'});
      res.write(attend);
      res.end();
      break;  

    case '/index.html':
      res.writeHead(200, {'Content-Type': 'text/html'});
      res.write(indexPage);
      res.end();
      break;
    
    case '/seito.html':
    	res.writeHead(200, {'Content-Type': 'text/html'});
    	res.write(otherPage1);
      res.end();
      break;

    case '/sensei.html':
    	res.writeHead(200, {'Content-Type': 'text/html'});
    	res.write(otherPage2);
      res.end();
      break;
    
    case '/home.html':
      res.writeHead(200, {'Content-Type': 'text/html'});
      res.write(home);
      res.end();
      break;

    case '/test.html':
      res.writeHead(200, {'Content-Type': 'text/html'});
      res.write(testPage);
      res.end();
      break;

    case '/index.css':
    	res.writeHead(200, {'Content-Type': 'text/css'});
   	 	res.write(styleCss);
    	res.end();
    	break;
    
      case '/table.css':
        res.writeHead(200, {'Content-Type': 'text/css'});
          res.write(tableCss);
        res.end();
        break;

    case '/js/test.js':
    	res.writeHead(200, {'Content-Type': 'text/plain'});
    	res.write(scriptJs);
    	res.end();
    	break;

    case '/js/login.js':
      res.writeHead(200, {'Content-Type': 'text/plain'});
      res.write(login);
      res.end();
      break;

    case '/photo/logo.png':  //　←　アドレスは任意。本当はuuidとか使うのがいいのかもしれませんが。
      res.writeHead(200, {
      'Content-Type': `image/png; charset=utf-8`  //　← ここがキモ！
      });
      var image = fs.readFileSync("./photo/logo.png", "binary"); // ← ファイルpathはその環境に合わせてください
      res.end(image, "binary");
      break;
    
    case '/photo/back.png':  //　←　アドレスは任意。本当はuuidとか使うのがいいのかもしれませんが。
      res.writeHead(200, {
      'Content-Type': `image/png; charset=utf-8`  //　← ここがキモ！
      });
      var image = fs.readFileSync("./photo/back.png", "binary"); // ← ファイルpathはその環境に合わせてください
      res.end(image, "binary");
      break;
    
    case '/photo/home.png':  //　←　アドレスは任意。本当はuuidとか使うのがいいのかもしれませんが。
      res.writeHead(200, {
      'Content-Type': `image/png; charset=utf-8`  //　← ここがキモ！
      });
      var image = fs.readFileSync("./photo/home.png", "binary"); // ← ファイルpathはその環境に合わせてください
      res.end(image, "binary");
      break;

    default:
      res.writeHead(200, {'Content-Type': 'text/plain'});
      res.end('お探しのページは見つかりません。');
      break;
  }
}
