# 修改

```
const (
port                = ":6002"  //端口
upFolder            = "./up"  //存放目录
accessPassword      = "123456"  //访问密码
fileRetentionPeriod = 3 * time.Hour  //文件存放小时，到时间自动删除
)
```

# up.py

```
Usage: python up.py <host:port> <file_path> <access_password>
```

# /etc/systemd/system/dateupload.service

```
[Unit]
Description=up Service
After=network.target

[Service]
WorkingDirectory=/big/upload/   //启动目录
ExecStart=/big/upload/uploads  //启动命令
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

# 重载

```
systemctl daemon-reload
```

# 启动

```
sudo systemctl enable dateupload.service
sudo systemctl start dateupload.service
```


