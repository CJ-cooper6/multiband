# 多波段视频监控系统


## 启动nginx

```
cd /h/nginx
./anginx -c conf/winnginx.conf
```

## 一起推流

```
ffmpeg -f dshow -i video="Chicony USB2.0 Camera" -f dshow -i audio="麦克风 (Realtek High Definition Audio)" -vcodec libx264 -preset ultrafast -tune:v zerolatency -f flv rtmp://127.0.0.1/live2/stream
```

<br>

```
go run main.go
```

# todo
错误处理
