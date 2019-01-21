<div align="center">

## AUXPI

**基于 API 的简单图床**

[![GitHub issues](https://img.shields.io/github/issues/aimerforreimu/AUXPI.svg)](https://github.com/aimerforreimu/AUXPI/issues)
[![GitHub forks](https://img.shields.io/github/forks/aimerforreimu/AUXPI.svg)](https://github.com/aimerforreimu/AUXPI/network)
[![GitHub stars](https://img.shields.io/github/stars/aimerforreimu/AUXPI.svg)](https://github.com/aimerforreimu/AUXPI/stargazers)
[![GitHub license](https://img.shields.io/github/license/aimerforreimu/AUXPI.svg)](https://github.com/aimerforreimu/AUXPI)
</div>
<br>

## 功能 特色

* 支持 web 上传图片
* 支持 API 上传图片
* 支持图床:
    
    * 搜狗
    * 新浪 (私有+公共)
    * SMMS 
    * 奇虎 (360)
    * 百度  
    * 阿里
    * 京东
    * Upload.cc
    * Flickr
    * 网易
    * 掘金
    * 本地 (由于需要对接后台，正在紧张开发中)

    
* API 可以设置 token 可以私用，也可以选择关闭 API，只保留 web 上传
* 可以设置是否开启新浪图床上传（因为新浪图床需要登录自己的账号）
* 可以设置允许上传的图片最大大小 和 一次性上传的最多张数

* 简单部署即可使用，从 Release 中获取，开箱即用
* 轻量级不使用任何数据库 (暂时)
* 后台控制，用户和管理员双后台(开发中)
* Go 语言编写速度加成(滑稽)
    


## 项目截图

![首页](https://ws4.sinaimg.cn/large/ed24e93ely1fzeamnpm6yj21hc0o7n31.jpg)

![无上传图片样式](https://ws2.sinaimg.cn/large/ed24e93ely1fzeao21i4wj21hb0odad1.jpg)



## 部署使用

**首先前排不要脸求Star ❤ (/ω＼)**

**有问题请到 Issue 中提出**

**演示站点为 [AuXpI 图床](https://imgx.0w0.tn/)**

**因为服务器在国外，可能会上传稍微慢一些 ~**



查看 WIKI 中的 [安装教程 ](https://github.com/aimerforreimu/AUXPI/wiki/%E5%AE%89%E8%A3%85%E9%83%A8%E7%BD%B2)



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

