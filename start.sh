#!/bin/sh

set -e #确保脚本在指令返回非零直接返回

echo "run db migrate"
/app/migrate -addr=mysql57:3306 -source=/app/migration -dbName=road -auth=root:WW876001

echo "start the app"
exec "$@" # 执行传递给脚本的所有参数
