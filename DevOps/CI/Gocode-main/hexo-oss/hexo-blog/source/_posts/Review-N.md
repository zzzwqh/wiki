---
title: Review-N
date: 2021-01-16 16:17:49
tags: Review
categories: Review
---

<!--MORE-->



## 编译安装脚本

> 编译安装配置模块可自行定义

```bash
wget http://nginx.org/download/nginx-1.22.0.tar.gz  && tar xf nginx-1.22.0.tar.gz -C /usr/local/
mv /usr/local/nginx-1.22.0/ /usr/local/nginx
useradd -M -s /sbin/nologin nginx
# 配置模块
cd /usr/local/nginx  && ./configure --prefix=/usr/local/nginx --user=nginx --group=nginx --with-http_stub_status_module
# 编译安装
make && make install
ln -s /usr/local/nginx/sbin/nginx /usr/local/sbin/
mkdir /usr/local/nginx/logs
touch /usr/local/nginx/logs/access.log
touch /usr/local/nginx/logs/error.log
chown nginx.nginx /usr/local/nginx -R 
# 启动
nginx
```

[Nginx 配置文件 nginx.conf 详细说明](https://www.jb51.net/article/103968.htm)

## 冲突指令以谁为准？

> Nginx 中两种指令，值指令可以覆盖，动作指令不可以

![image-20220816153813879](image-20220816153813879.png)

![image-20220816153852342](image-20220816153852342.png)





## Alias 和 Root 的区别

root和alias都可以定义在location模块中，都是用来 <font color="red">**指定请求资源的真实路径**</font> ，比如：

1、 alias 只能作用在 location 中，而 root 可以存在 server、http 和 location 中。

2、 alias 后面必须要用 “/” 结束，否则会找不到文件，而 root 则对 ”/” 可有可无。

```BASH
location /i/ {  
    root /data/w3;
}
```

请求 `http://foofish.net/i/top.gif ` 这个地址时，那么在服务器里面对应的真正的资源是` /data/w3/i/top.gif`文件

![img](https://img-blog.csdnimg.cn/img_convert/20413acc2c8ea3324c2b344457a0dd42.png)

而 alias 正如其名，alias 指定的路径是 location 的别名，不管 location 的值怎么写，资源的真实路径都是 alias 指定的路径 ，比如：

```bash
location /i/ {  
  alias /data/w3/;
}
```

![img](https://img-blog.csdnimg.cn/img_convert/4ac86f824c26eb31e0d5a944adc14d33.png)

## Listen 指令

本地通信可以用 unix 通信，性能更好

![image-20220816161545376](image-20220816161545376.png)

## 接受请求的 Event 模块

请求先到操作系统内核，Nginx 的 Event 模块一直有一个 epoll_wait 

![image-20220816162508770](image-20220816162508770.png)

## 接受请求的 Http 模块

![image-20220816181928420](image-20220816181928420.png)

## Nginx 的正则表达式

>[nginx 进行正则匹配（常见正则匹配符号表示） - 腾讯云开发者社区-腾讯云 (tencent.com)](https://cloud.tencent.com/developer/article/1500808)

![image-20220817114724634](image-20220817114724634.png)

> 一款用于测试正则表达式的工具  pcretest ，将你的正则用 /^ 和 $/  包裹起来，nginx 中 / 不需要用 \ 转义，这个测试工具命令行中需要用 \ 做转义

```BASH
[root@master01 ~]# yum -y install pcre-tools
[root@master01 ~]# pcretest
PCRE version 8.32 2012-11-30

  re> /^\/admin\/(\d+)\/images\/(\w+)\.(\w+)$/
data> /admin/1/images/icon.jpg
 0: /admin/1/images/icon.jpg
 1: 1
 2: icon
 3: jpg


# 0 就是 data 本身
# 1 是匹配到的第一项、2 是匹配到的第二项、以此类推
# Nginx 中的写法，就变成了/admin/(\d+)/images/(\w+)\.(\w+) 
```

## 处理请求的 Server 块（支持正则）

> 主域名的作用？当 server_name_in_redirect 的值配置为 off 时，无作用
>
> 当配置为 on 时，请求其他的域名时，会跳转（302）重定向到主域名的地址

![image-20220817145817994](image-20220817145817994.png)



> Server 模块中， server_name 指令中的正则

![image-20220817150235400](image-20220817150235400.png)

 

![image-20220817151337183](image-20220817151337183.png)

## Nginx 的 11 个处理阶段



![image-20220817151454701](image-20220817151454701.png)

![image-20220817151941694](image-20220817151941694.png)

## Postread 阶段 — 拿到 Client 真实 IP

> 需要用到 realip 模块，需要重新编译# 重新编译时候，使用–add-module=/root/nginx-push-stream-module指定添加模块
>
> 只使用make进行编译，把编译好的在objs的nginx替换掉原来的/usr/local/nginx/sbin/nginx

```bash
# 获取之前的配置参数
[root@master01 nginx]# nginx -V
nginx version: nginx/1.22.0
built by gcc 4.8.5 20150623 (Red Hat 4.8.5-44) (GCC)
configure arguments: --prefix=/usr/local/nginx --user=nginx --group=nginx --with-http_stub_status_module

# 添加模块
[root@master01 nginx]# ./configure --prefix=/usr/local/nginx --user=nginx --group=nginx --with-http_stub_status_module --with-http_realip_module
# 只需要编译
[root@master01 nginx]# make

# 把 objs 中的 nginx （新编译的）移到 sbin 命令路径下
[root@master01 nginx]# cp /usr/local/nginx/objs/nginx /usr/local/nginx/sbin/
```

![image-20220817152345238](image-20220817152345238.png)![image-20220817152351432](image-20220817152351432.png)

> set_real_ip_from 指定可以获取真实 IP 地址来源的网段，默认是全都获取（不用管）
>
> real_ip_header 的值需要我们设定为 X-Forwarded-For 才能追溯到最源头的客户端 IP 地址
>
> real_ip_recursive  默认为 off （不用管）
>
> 怎么写 Server 指令块？

```bash
server {
    listen 80;
    server_name return.test.com;
    root html/;
    index index.html index.htm;
    real_ip_header X-Forwarded-For;
    error_page  404 /403.html;
    location / {
        return 200 "Client Real IP : $remote_addr\n";
    }
}
```

> 在另外一台机器上 Curl 该地址

```bash
[root@master01 ~]# curl http://42.192.150.241/
Client Real IP : 121.89.244.58
```



## Rewrite 模块

### Return 指令

> 关于 302 和 301 的对比，302 是浏览器不会缓存这个重定向，每次访问该地址，还是会经过原始地址 URL 重定向到跳转的地址 URL
>
> 301 是浏览器会缓存这个永久重定向，直接缓存访问这个地址会跳转到的 URL
>
> 上述内容就是 permanent 301 和 redirect 302 的区别

![image-20220817160304623](image-20220817160304623.png)

> 303、307、308 中的改变方法指的是 POST 、GET 方法



### Return 指令和 error_page 之间的联系

![image-20220817160854156](image-20220817160854156.png)

### Rewrite 指令

![image-20220817173126693](image-20220817173126693.png)

![image-20220817173326325](image-20220817173326325.png)

> 上述代码如下，last 和 location 的区别如下

```bash
server {
    listen 80;
    server_name return.test.com;
    root html/;
    index index.html index.htm;
    real_ip_header X-Forwarded-For;
    real_ip_recursive on;
    error_page  404 /403.html;
    location /first {
        rewrite /first(.*) /second$1 last;	# last 意味着当前 rewrite 后的地址，还可以在当前 Server 里匹配其他 location
        return 200 'first_return\n';
    }
    location /second {
        rewrite /second(.*) /third$1 break; 
        return 200 'second_return\n';
    }
    location /third {
        return 200 'third_return\n';
    }
}
server {
    listen 80;
    server_name return.test.com;
    root html/;
    rewrite_log on;
    index index.html index.htm;
    real_ip_header X-Forwarded-For;
    real_ip_recursive on;
    error_page  404 /403.html;
    location /first {
        # last 意味着当前 rewrite 后的地址，还可以在当前 Server 里匹配其他 location
        rewrite /first(.*) /second$1 last;
        return 200 'first_return\n';
    }
    location /second {
        # 如果不带 break，不会去寻找 $ROOT/third/ 真实路径下的文件
        #rewrite /second(.*) /third$1;
        # 因为带了 break，后面的 return 语句也不会执行（更不会匹配其他 location），会到服务器上寻找 $ROOT/third/ 路径下的 $1 文件
        rewrite /second(.*) /third$1 break;
        return 200 'second_return\n';
    }
    location /third {
        return 200 'third_return\n';
    }
}
```

![image-20220817180912735](image-20220817180912735.png)

![image-20220817181251576](image-20220817181251576.png)

```bash
# 加入上面的 location 指令块
server {
    listen 80;
    server_name rewrite.test.com;
    root html/;
    rewrite_log on;
    index index.html index.htm;
    real_ip_header X-Forwarded-For;
    real_ip_recursive on;
    error_page  404 /403.html;
    location /first {
        rewrite /first(.*) /second$1 last;
        return 200 'first_return\n';
    }
    location /second {
        # 如果不带 break，不会去寻找 $ROOT/third/ 真实路径下的文件
        #rewrite /second(.*) /third$1;
        # 因为带了 break，后面的 return 语句也不会执行，会到服务器上寻找 $ROOT/third/ 路径下的 $1 文件
        rewrite /second(.*) /third$1 break;
        return 200 'second_return\n';
    }
    location /third {
        return 200 'third_return\n';
    }
    location /redirect1 {
        rewrite /redirect1(.*) $1 permanent;	# 301
    }
    location /redirect2 {
        rewrite /redirect2(.*) $1 redirect;	# 302
    }
    location /redirect3 {
        rewrite /redirect3(.*) http://42.192.150.241$1;	# 302
    }
    location /redirect4 {
        rewrite /redirect4(.*) http://42.192.150.241$1 permanent;	# 301
    }
}
```

![image-20220818102900144](image-20220818102900144.png)

### If 指令

![image-20220818103321609](image-20220818103321609.png)

![image-20220818103344802](image-20220818103344802.png)

![image-20220818103450144](image-20220818103450144.png)

## 找到处理请求的 Location（ Location 匹配顺序优先级）

![image-20220818103636370](image-20220818103636370.png)

![image-20220818103648308](image-20220818103648308.png)

![image-20220818103732082](image-20220818103732082.png)

```bash
# 在 Server 指令块下加入如下 location 指令块
    location ~ /Test1/$ {
         return 200 "first regular expressions match!\n";
     }
     location ~* /Test1(\w+)$ {
         return 200 'longest regular expressions match!\n';
     }
     location ^~ /Test1/ {
         return 200 'stop regular expressions match!\n';
     }
     location /Test1/Test2 {
         return 200 'longest prefix expression match!\n';
     }
     location /Test1 {
         return 200 'short prefix expression match!\n';
     }
     location = /Test1 {
         return 200 'exact match!\n';
     }
```

![ ](image-20220818112008310.png)

根据上图中的规则，我们尝试验证下 Curl 得到的结果

![image-20220818115805764](image-20220818115805764.png)

> 下面的例子，就是将 location ^~ /Test1/  改成了  location ^~ /Test1 ，将  location /Test1 改成了  location ^~ /Test1/  会有如下效果

![image-20220818120710043](image-20220818120710043.png) 

## Access 阶段 — 如何限制每个客户端的并发连接数 limit_conn 指令

> 模块不需要自己加，默认就有

![image-20220818121011654](image-20220818121011654.png)

![image-20220818124819736](image-20220818124819736.png)

![image-20220818121226773](image-20220818121226773.png)

```bash
limit_conn_zone $binary_remote_addr zone=addr:10m;

server {
    listen 80;
    server_name rewrite.test.com;
    root html/;
    rewrite_log on;
    index index.html index.htm;
    real_ip_header X-Forwarded-For;
    real_ip_recursive on;
    location / {
        # 当并发超出定义的最大数量时，返回的状态码
        limit_conn_status 500;
        # 并发连接日志的输出等级
        limit_conn_log_level warn;
        # 每个并发输出每秒输出的字节数是 50 Byte，传输速度慢下来，能够模拟 2 个并发时，返回 500 状态码的情况
        limit_rate 50;
        # 根据 addr 这个 Key 限制客户端的最大并发数，是 1
        limit_conn addr 1;
    }
}
```

> 客户端执行 Curl 请求

![image-20220818125410236](image-20220818125410236.png)

> 再开启一个窗口，模拟客户端并发，再次执行 Curl

![image-20220818125550747](image-20220818125550747.png)

## Access 阶段 — 对用户名密码做限制的 auth_basic 模块

[可参考链接](http://t.zoukankan.com/rdchenxi-p-11159821.html) 

![image-20220818143926534](D:\myhexo\source\_posts\Review-N\image-20220818143926534.png)![image-20220818143938485](D:\myhexo\source\_posts\Review-N\image-20220818143938485.png)

```bash
# 先生成用户名密码文件
[root@ethanz conf.d]# htpasswd -cb /usr/local/nginx/pass/passwd ethan 123

[root@ethanz conf.d]# cat server.conf
limit_conn_zone $binary_remote_addr zone=addr:10m;
server {
    listen 80;
    server_name rewrite.test.com;
    root html/;
    rewrite_log on;
    index index.html index.htm;
    real_ip_header X-Forwarded-For;
    real_ip_recursive on;
    location /{
        satisfy any;
        auth_basic "Some description";
        auth_basic_user_file /usr/local/nginx/pass/passwd;      # 密码文件的路径
        deny all;
    }
}
```

