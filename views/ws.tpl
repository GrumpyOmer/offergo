<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<div id="msg"></div>
<input type="text" id="text">
<input type="submit" value="发送数据" onclick="song()">
<p>adfssa</p>
</body>

<script>
    document.querySelector("body").style.fontSize = '28px';
    // var wsServer = 'ws://47.106.142.180:9501'
    //调用websocket对象建立连接：
    var wsServer = 'ws://127.0.0.1:5678/ws/server'
    // var wsServer = 'wss://api.woflat.com/wss'
    //参数：ws/wss(加密)：//ip:port （字符串）
    var websocket = new WebSocket(wsServer);
    //onopen监听连接打开
    websocket.onopen = function (evt) {
        //websocket.readyState 属性：
        /*
        CONNECTING    0    The connection is not yet open.
        OPEN    1    The connection is open and ready to communicate.
        CLOSING    2    The connection is in the process of closing.
        CLOSED    3    The connection is closed or couldn't be opened.
        */

    };

    function song(){

        var text = document.getElementById('text').value;
        var text = '{"msg":'+text+',"id":89}';
        //向服务器发送数据
        websocket.send(text);
    }
    //监听连接关闭
       websocket.onclose = function (evt) {
           console.log("Disconnected");
       };

    //onmessage 监听服务器数据推送
    websocket.onmessage = function (evt) {
        // console.log('Retrieved data from server: ' + evt.data);
        var json = JSON.parse(evt.data);
        msg.innerHTML += json.msg +'<br>';
    };
    //监听连接错误信息
    //    websocket.onerror = function (evt, e) {
    //        console.log('Error occured: ' + evt.data);
    //    };

</script>
</html>
