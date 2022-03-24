.PHONY: mysql_init
mysql_init:
	docker run --name mysql57 -v /home/raja/workspace/database/road/mysql/database:/var/lib/mysql -v /home/raja/workspace/database/road/mysql/conf:/etc/mysql/ -e MYSQL_ROOT_PASSWORD=WW876001 -e MYSQL_DATABASE=road -p 3306:3306 -d mysql:5.7
redis_init:
	docker run --name redis62 --privileged=true -p 7963:7963 -v /home/raja/workspace/database/road/redis/logs/redis.log:/redis/redis.log -v /home/raja/workspace/database/road/redis/database:/redis/data -v /home/raja/workspace/database/road/redis/conf/redis.conf:/redis/redis.conf -d redis:6.2 redis-server /redis/redis.conf
migrate_install:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz & sudo mv migrate /usr/bin/migrate
migrate_init_db:
	migrate create -ext sql -dir internal/dao/mysql/db/migration -seq init_schema
migrate_up:
	migrate -path internal/dao/mysql/db/migration -database 'mysql://root:WW876001@tcp(localhost:3306)/road' --verbose up
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
swag:
	swag init
#-e TZ=Asia/Shanghai -e default-time_zone='+8:00'
