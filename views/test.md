# bind 127.0.0.1 #限制redis只能本地访问，有其他需求可自行修改

protected-mode no #默认yes，开启保护模式，自选

daemonize no#默认no，改为yes意为以守护进程方式启动，可后台运行，除非kill进程，改为yes会使配置文件方式启动redis失败

logfile /opt/redis.log //对应log日志地址，挂载再宿主机上

dir /opt/data #输入本地redis数据库存放文件夹（可选）

# appendonly yes #redis持久化（可选）
