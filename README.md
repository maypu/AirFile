# AirFile

#### 介绍
AirFile 文件快速分享系统

#### 软件架构
`golang`
`gin`

#### 部署
https://hub.docker.com/r/maypu/airfile

Simple startup command:
```
docker run -d -p 8086:8086 --name=airfile --restart=always maypu/airfile:latest
```

More custom configuration startup commands:
```
docker run -d -p 8086:8086 -v /path/to/host/database:/app/database -v /path/to/host/files:/app/files --name=airfile --restart=always maypu/airfile:latest
```
Replace `/path/to/host/files` with the host path you wish to place

#### 开发
```shell
git clone https://github.com/maypu/AirFile.git
cd AirFile
go mod tidy
```

#### 截图
![AirFile首页](https://pictg.teahot.top/file/96c25c0d0f8853175df55.png)