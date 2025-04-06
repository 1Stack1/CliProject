# 说明

**实现命令行调用FOFA**

生成cli.exe

```shell
go build -o cli.exe main.go
```

进行FOFA查询，并调用chrome截图

```shell
./cli.exe --key YOUR_FOFA_KEY  --query "百度"
```

输出格式

```shell
url  状态码  标题  浏览器截图存放路径
```

# 支持配置文件读取

```yml
number_concurrency: 10 #并发截图数量
count: 17 #FOFA查询数量
proxy_url: http://localhost:7890 #代理地址
```

