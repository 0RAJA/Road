A Road of Code
go+mysql+redis 使用docker-compose部署,数据库挂载在本地
已经完成了后端的建设，
# 整体结构
```
├── cmd 辅助可执行文件(数据库迁移文件，测试文件)
├── configs (配置文件路径)
├── database (容器持久化路径)
├── docker-compose.yml (docker-compose配置文件，按需修改)
├── Dockerfile (用于启动road项目的配置文件)
├── docs (接口文档)
├── go.mod
├── go.sum
├── internal (代码区)
├── main.go (启动文件)
├── Makefile (常用的指令)
├── README.md
├── sqlc.yaml (sqlc配置文件)
├── start.sh (用于在docker-compose项目部署成功后初始化数据库)
├── storage (持久化区，日志文件所在地)
├── test.md (redis的默认配置文件)
├── uploads (上传文件所在区)
└── wait-for.sh (用于等待redis，mysql启动后再启动road)
```
# 启动方式
1. 先创建configs/app/app.yml(按照configs/model/app.yml为模板)，然后根据自己设备的配置进行设置。
2. 修改docker-compose.yml中挂载在本地的路径为自己设备的路径
3. 然后`docker-compose up` 一键部署启动
# 须知
数据库的数据和配置文件挂载在本地，需要查看接口文档只需要更改`Dockerfile`中`RUN go build -o main main.go`为`RUN go build -tags "doc" -o main main.go`
然后浏览器访问 `http//http://127.0.0.1:8080/swagger/index.html#/` 即可
