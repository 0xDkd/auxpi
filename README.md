<div align="center">


## AUXPI

**集合多家 API 的新一代图床**

[![GitHub issues](https://img.shields.io/github/issues/aimerforreimu/AUXPI.svg)](https://github.com/aimerforreimu/AUXPI/issues)
[![GitHub forks](https://img.shields.io/github/forks/aimerforreimu/AUXPI.svg)](https://github.com/aimerforreimu/AUXPI/network)
[![GitHub stars](https://img.shields.io/github/stars/aimerforreimu/AUXPI.svg)](https://github.com/aimerforreimu/AUXPI/stargazers)
[![GitHub license](https://img.shields.io/github/license/aimerforreimu/AUXPI.svg)](https://github.com/aimerforreimu/AUXPI)

本项目使用 beego + vue + element-ui 进行开发

```text
    _       __  __  ___ _____
    /_\  /\ /\ \/ / / _ \\_   \
   //_\\/ / \ \  / / /_)/ / /\/
/  _  \ \_/ /  \/ ___/\/ /_
\_/ \_/\___/_/\_\/   \____/

🍭 A NEW API IMAGES STORE TOOL 🍭


```
</div>
<br>

## 功能 特色

* 支持 web 上传图片
* 支持 API 上传图片
* 支持分发，控制反转
* 各种自由定制请看下面的截图

    


## 项目截图

![首页](https://ws2.sinaimg.cn/large/007DFgJwgy1g10ecblh1dj31hc0nw770.jpg)

![管理员后台](https://ws3.sinaimg.cn/large/007DFgJwgy1g10eavu2zqj31ha0obdii.jpg)

![用户后台](https://ws3.sinaimg.cn/large/007DFgJwgy1g10ejf74lzj31h70ogtey.jpg)

![设置页面](https://ws4.sinaimg.cn/large/007DFgJwgy1g10fkgqdsnj30u018cq9x.jpg)

---

**如果您感觉不错，请您点个 Star，您的 Star 是对我最大的鼓励 (认真脸)**

有问题请到 Issue 中提出



## 安装教程

如果您是空白的 vps，您可以直接使用一键脚本进行安装，安装教程 

注意此脚本会给您安装 Nginx 和 Mysql，请确保您的服务器是干净的

https://github.com/aimerforreimu/AUXPI/wiki/%E4%B8%80%E9%94%AE%E5%AE%89%E8%A3%85%E8%84%9A%E6%9C%AC

如果您想手动安装，可以参考安装 wiki

https://github.com/aimerforreimu/AUXPI/wiki/%E5%9B%BE%E5%BA%8A%E9%85%8D%E7%BD%AE




演示站点为 [https://test.0w0.tn/](https://test.0w0.tn/)

请勿滥用，滥用者将被加入 ip 黑名单并且放到 github 上公示

演示站点管理员账户:
```text
用户名:admin123
密码:admin123
邮箱:auxpi@0w0.tn
```

## 开发人员 && 折腾用户

如果您想要从源代码中构建 auxpi，您可以按照下面的教程进行构建，如果您仅仅想使用 auxpi，这一段您可以跳过

请注意你，请确保您的电脑上有 Go 环境。
### 从源码中构建

#### 使用脚本构建

在 [最近一次更新](https://github.com/aimerforreimu/AUXPI/commit/00e061273b6a74e0c85ed7feb744fc58d0e2c797) 中加入了批量编译打包的脚本，您可以使用此脚本进行代码的构建

```bash
bash build.sh all # 编译所有平台的程序
bash build.sh mac # 编译 mac 程序
bash build.sh windows # 编译 Windows 程序
bash build.sh linux #编译 linux 程序
bash build.sh all 2.2.0 clear # 编译程序，分别打包所有的平台的项目，2.2.0 为版本号，清空编译以后文件，只保留压缩包
bahs build.sh help #查看帮助
```

打包前端

```bash
cd resource
yarn install
yarn run build
```


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

### 二次开发规范

如果这个项目让您感觉还不错，您想二次开发一下或者想为这个项目提交 PR

##### 1.命名规范
* 函数，私有变量必须采用小驼峰方式进行命名 即 `sendMail()`,`userInfo` 

* 共有变量需要使用**大驼峰**方式命名

* 结构体,接口必须使用**大驼峰**方式进行命名

##### 2.代码组织

最好请按照这个下面说明的结构去组织您的代码

* 中间件请存放在 `middleware` 文件夹中

* api 请按照版本号放在 `controller/api/v(0-9)/`中，所有 api 除去 `auth` 都不能进行模板引擎的渲染和操作，只允许输出 `json`

* 如果需要渲染模板引擎，请直接在 `controller/` 下面建立对应的 controller

* `utils` 下面的各种工具文件不允许与 `models` 下面的文件耦合在一起，如果要为 `models` 编写工具，请放到 `tools` 下

* `all.go` 中只允许写入经常被引入，需要格式化成 `json` , `xml` 等格式的结构体  


```text
.
├── LICENSE
├── README.md
├── auxpiAll
├── bootstrap
├── build
├── build.sh
├── conf
├── controllers
├── install.sh
├── log
├── main.go
├── middleware
├── models
├── pem
├── resource
├── routers
├── server
├── static
├── tests
├── tools
├── utils
└── views

```

## API 上传



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

* [x] API v2 版本分发上传,返回所有图床储存链接 

* [x] 用户系统

* [x] 前后端分离,Vue 驱动前端

* [x] 后台控制

* [x] 本地上传，各大平台对接储存

* [x] 使用 MySQL 而不是 JSON

* [ ] API 自动文档

* [ ] 国际化

* [ ] 多缓存选择 Memory|redis|file|Memcached

* [ ] 用户后台重构

* [ ] 定时任务调度

## 说明

本项目是学习 Go 的过程中，边学边写出来的程序，可能存在 bug 连篇 ，逻辑让人无法接受，南辕北辙，代码无法让人直视等副作用。

## 致敬
[@ astaxie](https://github.com/astaxie) (beego)

[@ PanJiaChen](https://github.com/PanJiaChen) (vue-element-admin)

[@ metowolf](https://github.com/metowolf/upimg-cli) (upimg-cli)

[@ wisp-x](https://github.com/wisp-x) (lsky-pro)



## LICENSE

GNU General Public License v3.0