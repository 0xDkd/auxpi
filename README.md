<div align="center">

## AUXPI

**基于 API 的简单图床**

[![GitHub issues](https://img.shields.io/github/issues/aimerforreimu/AUXPI.svg)](https://github.com/aimerforreimu/AUXPI/issues)
[![GitHub forks](https://img.shields.io/github/forks/aimerforreimu/AUXPI.svg)](https://github.com/aimerforreimu/AUXPI/network)
[![GitHub stars](https://img.shields.io/github/stars/aimerforreimu/AUXPI.svg)](https://github.com/aimerforreimu/AUXPI/stargazers)
[![GitHub license](https://img.shields.io/github/license/aimerforreimu/AUXPI.svg)](https://github.com/aimerforreimu/AUXPI)
</div>
<br>

## 项目截图

![Snipaste_2018-11-05_21-09-45.png](https://ws4.sinaimg.cn/large/006A66c0ly1fwxhzushiuj31hc0m5q4v.jpg)

![Snipaste_2018-11-05_21-10-02.png](https://ws3.sinaimg.cn/large/006A66c0ly1fwxhzuskkhj31h50o0myt.jpg)

![Snipaste_2018-11-05_21-10-11.png](https://ws1.sinaimg.cn/large/006A66c0ly1fwxhzup2k6j31ha0o8my3.jpg)

![Snipaste_2018-11-05_21-17-01.png](https://ws1.sinaimg.cn/large/006A66c0ly1fwxhzve3syj313j0lojup.jpg)

![Snipaste_2018-11-05_21-11-31.png](https://ws2.sinaimg.cn/large/006A66c0ly1fwxhzuvyx9j31gw0ny0wv.jpg)




## 部署使用

**首先前排不要脸求Star ❤ (/ω＼)**

**有问题请到 Issue 中提出**

**演示站点为 [AuXpI 图床](http://img.0w0.tn)**

**因为服务器在国外，可能会上传稍微慢一些 ~**



使用该图床有两种选择，一种选择是使用 Release 中已经编译好的文件，另外一种选择是自己编译代码。

### 从 Release 中获取编译完成的文件
首先需要到 [Release](https://github.com/aimerforreimu/AUXPI/releases) 中下载符合您服务器需要的最新版本的压缩包，然后在您的服务器端进行解压，解压完成以后的目录为 (以 Linux/Mac 系统为例):

```text
├── LICENSE
├── README.md
├── conf #配置文件目录
├── static #静态文件存放目录
├── auxpi # 编译以后的可执行程序
└── views # 视图文件
```

**基础配置**

您可以在 `conf/app.conf` 下面进行基础站点的配置，配置示例如下


```text
appname = auxpi # 程序的名称
httpport = 2333 # 程序所运行的端口
runmode = dev   # 程序的运行环境 dev 为开发模式 pro 为正常模式
enablexsrf = true # 是否开启 CSRF 攻击防御（必须开启）
xsrfkey = ads093jmfas93j*3sd-212df923 #CSRF key 这里请随意填写字符串
xsrfexpire = 3600 # CSRF token 过期时间
```
请注意，`dev` 模式 需要设置 $GOPATH 所以如果没有 Go 环境，请直接使用 pro 模式，调试请使用源代码进行调试

设置完成以后，请执行 如下命令赋予文件执行权限

```bash
# 赋予运行权限
chmod u+x auxpi

#运行程序
./auxpi

#后台运行程序
nohup ./auxpi
```


在程序第一次运行的时候会在 `conf/` 目录下生成 `install.lock` 和 `siteConfig.json` 如果删除 `install.lock` 的话 `siteConfig.json` 就会被初始化为最初的值，所以请不要轻易的删除 `install.lock`

**站点配置**

接下来您可以在 `siteConfig.json` 配置您的站点设置

```text

{
  "site_name": "BusterApi 图床", #站点名称
  "site_footer": "你好世界", #footer 输出的内容
  "site_url": "/", # 站点 url
  "site_upload_max_number": 10,#一次性最多可以上传多少张图片
  "site_up_load_max_size": 5,#最大允许上传的图片大小，单位 MB
  "open_api_up_load": true,# 是否开启 API 上传 
  "api_token": "",# API token 空为不设置
  "api_default": "SouGou", # API 默认上传上去的图床
  "cache_config": true, #是否对配置进行缓存(建议开启)
  "site_upload_way": {
    "local_store": false, # 是否开启本地储存(此功能等待开发)
    "open_sina_pic_store": false, #是否启用新浪图床
    "sina_account": {
      "user_name": "", #若开启微博图床，请填写您的微博登录用户名
      "pass_word": "", #填写您的 微博登录密码
      "reset_sina_cookie_time": 3600,  # 微博 cookie 缓存时间 s
      "defult_pic_size": "large" # 默认返回的微博图片的大小
    }
  }
}

```

其中需要说明的是

`reset_sina_cookie_time` 这一项最好不要更改，更改的话不要让其大于 3600s

`defult_pic_size` 可选的参数为 

```text
square
thumb150
orj360
orj480
mw690
mw1024
mw2048
small
bmiddle
large
```

请根据自己的需要进行配置，**配置以后需要重新启动程序才能生效**

您现在可以通过访问 http://yourip:yourport 进行访问

（yourport 为 conf 中设置的端口）


**绑定域名**

现在程序已经在您的后台中跑了起来，如果您想要绑定域名可以通过 Nginx 进行反代，Nginx 的配置如下

```conf
server
{
    listen 80;
    server_name  你的域名;
    error_page 404 /404.html;
    error_page 502 /502.html;

    location /
    {
        proxy_pass http://yourip:yourport;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header REMOTE-HOST $remote_addr;
        add_header X-Cache $upstream_cache_status;
        expires 12h;
    }

    location ~ ^/(\.user.ini|\.htaccess|\.git|\.svn|\.project|LICENSE|README.md)
    {
        return 404;
    }

}
```

替换 `你的域名` 和 `http://yourip:yourport` 为你实际的值即可，保存以后，需要检查 nginx 的配置,然后重启 ，访问您的域名即可看到您的网站


### 从源码中构建

#### 使用 Bee 工具运行程序
```bash
cd $GOPATH/src
git clone git@github.com:aimerforreimu/AUXPI.git
cd AUXPI/
bee run auxpi
```

#### 交叉编译
 
```bash
#Mac/Windows环境编译成 Linux 程序 
GOOS=linux GOARCH=amd64 bee pack 
#Mac/Linux 环境编译 Windows 程序
GOOS=windows GOARCH=amd64 bee pack
#Windows/Linux 编译 Mac 程序
GOOS=darwin GOARCH=amd64 bee pack
``` 
更多交叉编译请参考 [Go 交叉编译](https://www.jianshu.com/p/4b345a9e768e)

## API 上传

其实当时写这个程序的时候没有想要写前端的页面，是看到了另一位大佬的图床，感觉这个前端页面很好看才写网页版上传，本来想直接写个 API 服务.

### API 上传实例

**图片上传 V1 接口**

| 功能 | 图片上传接口 |
| --- | --- |
| HTTP 请求方式 |  POST |
| URL  | http://yourname/api/v1/upload |

**请求参数**

| 参数名称 | 类型 | 是否必须|描述|
| --- | --- |---| --- |
| image | File | 是 | 表单名称,上传图片|
| token  | String | 是 | 认证所必须的 token ，如果站在没有开启则留空即可 |
| apiSelect  | String | 是 | 所选择的 API 类型 |

**apiSelect可选参数**

| apiSelect 可选参数 | 参数说明
| --- | ---|
| SouGou| 搜狗图床|
|Sina|新浪图床|
|Smms|SMMS 图床|


**成功上传返回**

```json
{
    "code": 200,
    "msg": "上传成功",
    "data": {
        "name": "Snipaste_2018-08-28_01-17-58.png",
        "url": "https://img04.sogoucdn.com/app/a/100520146/0dcb98aadb59c6b29dc0832eb7cc094a"
    }
}

```

```json
{
    "code": 200,
    "msg": "上传成功",
    "data": {
        "name": "Snipaste_2018-08-28_01-17-58.png",
        "url": "https://i.loli.net/2018/11/05/5be038b1b4af6.png"
    }
}
```

**失败返回值**

上传出错返回值

```json
{
    "code": 500,
    "msg": "上传失败"
}
```

API 未开启返回值

```json
{
    "code": 405,
    "msg": "Method not allowed"
}
```
 Token 验证失败返回值
 
```json
{
    "code": 403,
    "msg": "Forbidden"
}
```

选择文件为空返回值

```json

{
    "code": 500,
    "msg": "No files were uploaded."
}

```

文件太大返回值

```json

{
    "code": 500,
    "msg": "File is too large."
}

```

## TODO 

* [x] API 上传

* [ ] API 自动文档

* [ ] API v2 版本分发上传,返回所有图床储存链接 

* [ ] 用户系统

* [ ] 前后端分离,Vue 驱动前端

* [ ] 后台控制

* [ ] 本地上传，各大平台对接储存

* [ ] 使用 MySQL 而不是 JSON


## 说明

本项目是学习 Go 的过程中，边学边写出来的程序，可能存在 bug 连篇 ，逻辑让人无法接受，南辕北辙，代码无法让人直视等副作用。

## 致敬

[wisp-x](https://github.com/wisp-x)

[astaxie](https://github.com/astaxie)

