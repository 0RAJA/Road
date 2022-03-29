#!/bin/sh

set -e #确保脚本在指令返回非零直接返回

echo "run db migrate"
/app/migrate --host=mysql57 --source=/app/migration

echo "start the app"
exec "$@" # 执行传递给脚本的所有参数

