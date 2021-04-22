# ScanForLogin
## 一、redis的下载、安装和配置

1.下载redis：
```text
/home/SoftWareForCoding$ wget https://download.redis.io/releases/redis-6.2.1.tar.gz
```

2.解压
```text
/home/SoftWareForCoding$ tar xzf redis-6.2.1.tar.gz
/home/SoftWareForCoding$ cd redis-6.2.1
```
3.编译（但是redis-cli和redis-server命令不能全局运行）
```text
/home/SoftWareForCoding/redis-6.2.1$ make
```

4.使redis-cli和redis-server命令能全局运行
```text
/home/SoftWareForCoding/redis-6.2.1$ make install
```

5.编辑redis.conf
```text
/home/SoftWareForCoding/redis-6.2.1$ vim redis.conf

daemon no #该选项不需要修改，默认就好（网上很多教程说吧no改成yes，没啥用）

# 将该行注释，或者改为bind 0.0.0.0 
# (这个bind设置，不是设置什么ip什么ip就能访问redis，而是当绑定具有多个网卡的计算机时，确定绑定的是哪块网卡，访问者只能通过绑定的网卡进行访问)
bind 127.0.0.1 -::1

# 将下面这行改为 protected-mode no
protected-mode yes

# 自动设置密码（每次重启redis，只要带上conf，就会自动设置）
requirepass xxxxxx

# 最后一步尤其重要，如果你远程连接redis的时候，一直报io timeout，很有可能是自己的服务器没有开启安全组
# 如果是租的阿里云服务器，需要在阿里云安全组中配置6379端口的访问权限

redis.conf默认权限是666，需要修改一下，添加文件所有者的执行权限
chmod 766 (文件所有者、用户组内成员、其他成员)
```

6.启动redis
```text
# 服务器上启动redis报错：Creating Server TCP listening socket *.*.*.*:6379: bind: Cannot assign requested address
是由于bind配置项错误（参考第5节）

# 根据默认配置启动redis
redis-server

# 如果修改了redis.conf,需要在后面带上配置文件的路径，就会加载修改的配置项
redis-server redis.conf
```

7.关闭redis
```text
redis-cli shutdown
```

8.设置密码
```text
>config set requirepass 124541

在conf文件中设置密码参考第五节
```

9.go get下载时指定branch或tag
```text
-u表示更新包到最新版本，-v表示输出详细信息

指定branch:
go get -u -v <path-to-repo>@<branch_name>

指定tag：
go get -u -v <path-to-repo>/tag

```

## 二、Go的encoding/json
1.结构体无法直接存入redis，如果结构体中的字段是小写的，则转换为json之后得到的结果只是一个"{}"大括号，不包含任何字段，因为结构体中的字段为小写则为不可导

2.json tag用于设置结构体转换为json后字段的名称,如果没有json tag，则默认等于字段名
```text
	a := struct {
		C string `json:"c"`
		B int `json:"b"`
	}{
		C: "robert lu",
		B: 124,
	}
	// aj为[]byte类型
	aj, err := json.Marshal(&a)
	fmt.Println(aj)
```

3.设置数据的时候一般后面要带上Result(),因为这样如果设置失败可以抛出error；
如果从redis中获取值时，可以按需考虑，如果key不存在时需要抛出异常，则使用Result，如果key不存在也要正常运行，则使用String()等函数

4.在使用json.Marshal时，参数何时使用&val,何时使用val，两种有区别吗？