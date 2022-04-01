<a href="https://github.com/0RAJA/Road">A Road of Code 简易博客系统</a>

gin+mysql+redis+docker-compose部署,数据挂载在本地 已经完成了后端的基础建设，

# 整体结构

```go
├── cmd //可执行程序入口
│   ├── migrate //数据库迁移工具
│   └── test //Main测试
├── configs //配置文件
│   ├── app //实际配置文件
│   └── model //配置文件模型
├── database //挂载数据库和项目的本地资源
│   ├── mysql
│   │   ├── conf
│   │   └── data
│   ├── redis
│   │   ├── conf
│   │   └── data
│   └── road
│       ├── configs
│       │   ├── app
│       │   └── model
│       ├── storage //持久化层
│       │   └── Applogs //日志持久化
│       └── uploads //文件的保存
├── docs //接口文档
├── internal //具体代码
    ├── controller //接口层 对外路由
    ├── dao //与数据库交互
    │   ├── mysql
    │   │   ├── migration //数据库的迁移脚本
    │   │   ├── query //sql语句
    │   │   └── sqlc //对应sqlc生成的go代码
    │   └── redis
    ├── global //全局变量（设置，日志，项目根路径，生成token等等辅助的变量被初始化在此层）
    ├── logic //逻辑处理层，接收接口层的合法数据进行操作（可能与dao数据库层操作）
    ├── middleware //中间件(限流，鉴权，recover，等等。。。)
    ├── pkg //工具包
    │   ├── app //格式化输入输出(错误码标准化，分页处理，格式化响应等与格式相关的操作)
    │   ├── bind //利用gin+validator+translator绑定数据并通过翻译中间件进行翻译错误信息
    │   ├── conversion //常用数据格式转换
    │   ├── email //发送邮件
    │   ├── limiter //限流（令牌桶）
    │   ├── logger //zap+lumberjack实现日志的持久化
    │   ├── password //使用bcrypt对密码进行加密
    │   ├── setting //读取配置文件
    │   ├── singleflight //防止缓存击穿和缓存穿透
    │   ├── snowflake //雪花算法实现分布式ID生成器
    │   ├── struct //常用数据结构（前缀树路由规则匹配）
    │   ├── times //常用时间格式化工具
    │   ├── token //使用passto生成token
    │   ├── upload //对上传文件的检验与保存的封装
    │   └── utils //常用工具（随机数，hash等）
    ├── routing //注册路由和中间件
    └── settions //配置文件对应的数据结构
```

数据库表

![image-20220401115735145](https://gitee.com/ORaja/picture/raw/master/img/image-20220401115735145.png)

![image-20220401115939049](https://gitee.com/ORaja/picture/raw/master/img/image-20220401115939049.png)

# 接口文档

https://humraja.xyz/road/swagger/index.html#

# 总体总结

1. 因为一些原因没有先和前端讨论接口，就自己先思考罗列一下需求，

    因为是个小博客，所以就设置一个管理员权限就可以，mysql里存一些必要的数据即可，用户信息就采取github授权，通过管理员信息来判断权限来决定一些接口的访问权限，redis就放个缓存，以及像点赞信息和访问量等信息就先存储在redis，然后再定期入库（因为怕redis崩了，所以采取新增访问量的形式。）部署因为嫌麻烦就全部容器化，最后直接使用docker-compose一键部署了。

2. 对于mysql的部署一定注意时区，`timestamp` 会存储时区信息，把你的时间先转换为UTC时间，取的时候再转换为你的时区，而且可以自动更新修改时间和创建时间，但是使用时要注意把容器的时区调整对，还有就是中文的支持，在建表语句的时候记得加上(utf8)，最后mysql的数据库迁移工具我采取的是<a href="https://github.com/golang-migrate/migrate">golang-migrate</a> 但是命令行总是不成功，就直接写成一个程序写在Dockerfile里进行迁移也可以。

3. 基础的sql语句

    使用<a href="https://github.com/kyleconroy/sqlc">sqlc</a>进行对sql语句的转换，感觉也比较方便，有些地方虽然也用得不方便，但是也可以接受。对于单元测试我使用<a href="github.com/stretchr/testify/require">require</a>库就十分简洁

4. 然后是接口文档，使用<a href="https://swagger.io/">swagger</a> 自动生成接口文档就很方便。然后使用gin中封装好的`validator`库来进行入参校验，对于像时间之类的可以自定义参数校验，对于像分页，回复等使用封装好的包来统一规范即可。

5. 之后是逻辑处理层，我是将系统错误封装，原错误入日志，返回封装好的错误信息即可，也是将错误及时处理而不是返回给接口层处理。

6. 然后是中间件就可以直接用之前写好的包就可以。在原有基础上进行简单的修改，如限流器的匹配规则就使用前缀树进行改进。

7. 然后就是对整个项目的功能进行测试，接口测试我使用的是`apifox`也很方便。测试过程中也是遇到些问题，但都不是很难解决，毕竟整个逻辑也很清晰，对于数据库的单元测试也很完善。

8. 最后就是部署上线，书写docker-compose来一键部署。其中对于宿主挂载点一定要把权限开开，还有就是记得先把配置文件先放到挂载点。服务器上就用nginx进行反向代理即可。

9. 整体逻辑也不麻烦，大概一个星期就全部完成。之后等待前端将界面写好然后正式上线启用，再有什么问题就之后再慢慢解决。

# 启动方式

1. 先创建configs/app/app.yml(按照configs/model/app.yml为模板)，然后根据自己进行设置。
2. 修改docker-compose.yml中挂载在本地的路径为自己设备的路径
3. 然后`docker-compose up` 一键部署启动

# 注意事项

1. 数据库的数据和配置文件挂载在本地，需要查看接口文档只需要更改`Dockerfile`中`RUN go build -o main main.go`为`RUN go build -tags "doc" -o main main.go`
   然后浏览器访问 `http//http://127.0.0.1:8080/swagger/index.html#/` 即可
2. 默认管理员账号 用户名`0RAJA` 密码`123456`
3. 需要将配置文件先放到宿主机挂载的路径下再启动项目，否则配置文件会被覆盖为空。注意给宿主机挂载的路径开启权限，否则可能出现权限报错
4. 升级https `https://certbot.eff.org/instructions?ws=nginx&os=centosrhel7`
