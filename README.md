#   C聊
    一款非侵入式前端即时聊天工具。
    
#   如何使用?
    在</body>前添加
        <div id="c-chat">
            <chat api="wss://服务器地址/v1/websocket" id="默认ID" title="聊天室标题" name="当前用户名称,可留空" avatar="当前用户头像,可留空" logo="当前图标LOGO,可留空"></chat>
        </div>
     前端需引用资源文件 
        /src/js/app.js
        /src/js/style.css

#   Server部署
<a href="deploy.md" target="_blank">Server部署</a>
#   演示站点
<a href="https://carseason.github.io/chat" target="_blank">演示网站</a>
