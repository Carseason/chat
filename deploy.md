#   部署教程
    程序默认为 8082端口启动,
    默认 websocket 地址为 /v1/websocket
    直接执行 ./chat即可

#   nginx代理
    location ^~/v1/websocket {
        proxy_pass http://127.0.0.1:8082;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_read_timeout 360s;
    }
    使用nginx部署需添加以上配置