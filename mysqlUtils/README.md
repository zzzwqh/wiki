# MySQL 实例简易备份

```bash
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

```
# 备份后导入
```bash
# root @ ECS-project-v-alpha-online-public in /data/mysql [14:31:50]
$ cd dumpfiles

# root @ ECS-project-v-alpha-online-public in /data/mysql/dumpfiles [14:31:59]
$ ls
20240110123214  import_data.sh

# root @ ECS-project-v-alpha-online-public in /data/mysql [14:31:50]
$ cat import_data.sh
#!/bin/bash
# 建库
#for i in $(ls ./20240110123214 |  awk -F'.sql' '{print$1}' | egrep -v "schema|meta");do mysql -uroot -pRov-alpha -h127.0.0.1 -e "create database if not exists $i";done

# 导入数据
#for  i in $(ls ./20240110123214 |  awk -F'.sql' '{print$1}' | egrep -v "schema|meta"); do mysql -uroot -pRov-alpha -h 127.0.0.1 -D $i < ./20240110123214/$i.sql ;done
```
# MySQL 其他备份方式
```bash
# 位点 -> 表结构
mysqldump -uroot -pabc123 -hlocalhost -P3306 --master-data=2 --single-transaction --max-allowed-packet=1G --triggers --routines --events --no-data --databases mayfly > /data/mayfly_20230902.sql
# gtid -> 表结构
mysqldump -uroot -pabc123 -hlocalhost -P3306 --set-gtid-purged=COMMENTED --single-transaction --max-allowed-packet=1G --triggers --routines --events --no-data --databases mayfly > /data/mayfly_20230902.sql
```