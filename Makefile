.PHONY: mysql_init redis_init migrate_install migrate_init_db migrate_up sqlc test server swag
mysql_init:
	docker run --name mysql57 --privileged=true -p 3306:3306 -v /home/raja/workspace/go/src/Road/database/mysql/data:/var/lib/mysql -v /home/raja/workspace/go/src/Road/database/mysql/conf:/etc/mysql/ -e MYSQL_ROOT_PASSWORD=WW876001 -e MYSQL_DATABASE=road -e LANG=C.UTF-8 -e character_set_database=utf8 -d mysql:5.7
redis_init:
	docker run --name redis62 --privileged=true -p 7963:7963 -v /home/raja/workspace/go/src/Road/database/redis/data:/redis/data -v /home/raja/workspace/go/src/Road/database/redis/conf/redis.conf:/redis/redis.conf -d redis:6.2 redis-server /redis/redis.conf
migrate_install:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz & sudo mv migrate /usr/bin/migrate
migrate_init_db:
	migrate create -ext sql -dir internal/dao/mysql/db/migration -seq init_schema
migrate_up:
	./migrate -addr=localhost:3306 -source=internal/dao/mysql/migration -dbName=road -auth=root:WW876001
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
swag:
	swag init
docker_net:
	docker network create road
docker_connect:
	docker network connect road mysql57 & docker network connect road redis62
docker_build:
	docker build -t road:test .
docker_run:
	docker run --name road -p 8080:8080 --net road -it -d road:test
docker-compose_up:
	docker-compose up
docker-compose_down:
	docker-compose down
docker-compose_stop:
	docker-compose stop
#-e TZ=Asia/Shanghai -e default-time_zone='+8:00'
