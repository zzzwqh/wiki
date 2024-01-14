
## MySQL 实例备份

```bash
#!/bin/bash
DATE=`date "+%Y%m%d%H%M%S"`
DUMP_USER="dump"
DUMP_PASSWORD="dump@123456"
DUMP_ADDRESS="127.0.0.1"
EXPORT_DIR="./dumpfiles/$DATE"
LEN=60
[[ -d $EXPORT_DIR ]] || mkdir -pv $EXPORT_DIR
DATABASES=$(mysql -u$DUMP_USER -p$DUMP_PASSWORD -h$DUMP_ADDRESS -e "select distinct table_schema from information_schema.tables where table_schema not in ('mysql','sys','information_schema','performance_schema')" | grep -Ev "TABLE_SCHEMA")
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


```

## 备份后导入
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

