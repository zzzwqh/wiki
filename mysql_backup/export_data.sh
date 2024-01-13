## MySQL 实例备份
#!/bin/bash
DATE=`date "+%Y%m%d%H%M%S"`
DUMP_USER="dump"
DUMP_PASSWORD="dump@123456"
EXPORT_DIR="./dumpfiles/$DATE"
LEN=60
[[ -d $EXPORT_DIR ]] || mkdir -pv $EXPORT_DIR
DATABASES=$(mysql -u$DUMP_USER -p$DUMP_PASSWORD -e "select distinct table_schema from information_schema.tables where table_schema not in ('mysql','sys','information_schema','performance_schema')" | grep -Ev "TABLE_SCHEMA")
echo "===== 开始导出数据库: $DATE ===== "
num=0
for DB in $DATABASES; do
  let num++
  echo $num:$DB
  FILENAME="$EXPORT_DIR/$DB.sql.gz"
  mysqldump --set-gtid-purged=off --opt -R $DB | gzip -9 > $FILENAME
done

echo "导出完成！文件保存在$EXPORT_DIR"

#echo "开始清理 $LEN 天前的备份"
#find $EXPORT_DIR -type d -name "20*" -maxdepth 1 -mtime +$LEN -exec rm -rf {} \;
#echo "清理完成！
