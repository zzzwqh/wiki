#!/bin/bash
DATE=`date "+%Y%m%d%H%M%S"`
DUMP_ADDR="127.0.0.1"
DUMP_USER="root"
DUMP_PASSWORD="root"
EXPORT_DIR="./dumpfiles/$DATE"
LEN=60
[[ -d ${EXPORT_DIR} ]] || mkdir -pv ${EXPORT_DIR}
DATABASES=$(mysql -u${DUMP_USER} -p${DUMP_PASSWORD} -h${DUMP_ADDR} -e "select distinct table_schema from information_schema.tables where table_schema not in ('mysql','sys','information_schema','performance_schema')" | grep -Ev "TABLE_SCHEMA")
echo ${DATABASES}
echo "===== 开始导出数据库: ${DATE} ===== "
num=0
for DB in $DATABASES; do
  let num++
  echo $num:$DB
  FILENAME="${EXPORT_DIR}/${DB}.sql"
  mysqldump -u${DUMP_USER} -p${DUMP_PASSWORD} -h${DUMP_ADDR} --set-gtid-purged=off --opt -R ${DB} > ${FILENAME}
done

echo "导出完成！文件保存在${EXPORT_DIR}"

