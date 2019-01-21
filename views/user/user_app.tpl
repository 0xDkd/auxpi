<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{{ .siteName}} 管理</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/mdui@0.4.2/dist/css/mdui.min.css">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/sweetalert2@7/dist/sweetalert2.min.css">

</head>
<body class="mdui-drawer-body-left mdui-appbar-with-toolbar mdui-theme-primary-indigo mdui-theme-accent-blue">
{{/*{{ .Header}}*/}}
<header class="mdui-appbar mdui-appbar-fixed">
    <div class="mdui-toolbar mdui-color-theme">
        <span class="mdui-btn mdui-btn-icon mdui-ripple mdui-ripple-white"
              mdui-drawer="{target: '#main-drawer', swipe: true}"><i class="mdui-icon material-icons">menu</i></span>
        <a href="{{urlfor "PagesController.UserIndexShow"}}" target="_blank"
           class="mdui-typo-headline mdui-hidden-xs">{{ .siteName}} 用户中心</a>
        <div class="mdui-toolbar-spacer"></div>
        <a href="https://pan.0w0.tn/admin/bind" class="mdui-btn"><i class="mdui-icon material-icons">face</i>这里填写用户名</a>
        <a onclick="event.preventDefault();document.getElementById('logout-form').submit();"
           href="javascript:void(0)" class="mdui-btn mdui-btn-icon"><i class="mdui-icon material-icons">exit_to_app</i></a>
        <form id="logout-form" action="https://pan.0w0.tn/logout" method="POST"
              class="mdui-hidden">
        </form>
    </div>
</header>
{{/*{{ .SiderBar}}*/}}
<div class="mdui-drawer" id="main-drawer">
    <div class="mdui-list" mdui-collapse="{accordion: true}">
        <div
                class="mdui-collapse-item  mdui-collapse-item-open ">
            <div class="mdui-collapse-item-header mdui-list-item mdui-ripple">
                <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-blue">settings</i>
                <div class="mdui-list-item-content">设置</div>
                <i class="mdui-collapse-item-arrow mdui-icon material-icons">keyboard_arrow_down</i>
            </div>
            <div class="mdui-collapse-item-body mdui-list">
                <a class="mdui-list-item mdui-ripple  mdui-list-item-active "
                   href="https://pan.0w0.tn/admin">基础设置 </a>
                <a class="mdui-list-item mdui-ripple "
                   href="https://pan.0w0.tn/admin/show">显示设置 </a>
                <a class="mdui-list-item mdui-ripple "
                   href="https://pan.0w0.tn/admin/profile">密码设置 </a>
                <a class="mdui-list-item mdui-ripple "
                   href="https://pan.0w0.tn/admin/bind">绑定设置 </a>
            </div>
        </div>
        <div
                class="mdui-collapse-item ">
            <div class="mdui-collapse-item-header mdui-list-item mdui-ripple">
                <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-green">widgets</i>
                <div class="mdui-list-item-content">文件操作</div>
                <i class="mdui-collapse-item-arrow mdui-icon material-icons">keyboard_arrow_down</i>
            </div>
            <div class="mdui-collapse-item-body mdui-list">
                <a class="mdui-list-item mdui-ripple "
                   href="https://pan.0w0.tn/admin/file">普通文件上传 </a>
                <a class="mdui-list-item mdui-ripple "
                   href="https://pan.0w0.tn/admin/file/other">其它操作 </a>
            </div>
        </div>

        <div
                class="mdui-collapse-item">
            <div class="mdui-collapse-item-header mdui-list-item mdui-ripple">
                <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-red">clear_all</i>
                <div class="mdui-list-item-content">缓存</div>
                <i class="mdui-collapse-item-arrow mdui-icon material-icons">keyboard_arrow_down</i>
            </div>
            <div class="mdui-collapse-item-body mdui-list">
                <a class="mdui-list-item mdui-ripple"
                   href="https://pan.0w0.tn/admin/clear">缓存清理 </a>
                <a class="mdui-list-item mdui-ripple"
                   href="https://pan.0w0.tn/admin/refresh"
                   onclick="mdui.snackbar({ message: '正在刷新缓存，请稍等', position: 'right-top' });">缓存刷新 </a>
            </div>
        </div>
        <a href="https://pan.0w0.tn/image" class="mdui-list-item mdui-ripple">
            <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-pink">image</i>
            <div class="mdui-list-item-content">图床</div>
        </a>
        <a href="https://pan.0w0.tn/admin/debug" class="mdui-list-item mdui-ripple">
            <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-grey">bug_report</i>
            <div class="mdui-list-item-content">调试日志</div>
        </a>
        <a href="https://onedrive.live.com" class="mdui-list-item mdui-ripple">
            <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-blue">cloud</i>
            <div class="mdui-list-item-content">OneDrive管理</div>
        </a>
    </div>
</div>

<a id="anchor-top"></a>

{{/*{{ .Content}}*/}}
<div class="mdui-container">
    <div class="mdui-container-fluid mdui-m-t-2 mdui-m-b-2">

        <div class="mdui-typo">
            <h1>基本设置
                <small>参数设置</small>
            </h1>
        </div>
        <form action="" method="post">
            <input type="hidden" name="_token" value="OZZVKkIkl78rMfZ8BveOfnNhgjCyVCY2FxzhvVXx">
            <div class="mdui-textfield mdui-textfield-floating-label">
                <label class="mdui-textfield-label" for="name">站点名称</label>
                <input type="text" class="mdui-textfield-input" id="name" name="name"
                       value="OLAINDEX">
            </div>
            <br>
            <div class="mdui-textfield mdui-textfield-floating-label">
                <label class="mdui-textfield-label" for="root">OneDrive根目录</label>
                <input type="text" class="mdui-textfield-input" id="root" name="root"
                       value="/" required>
                <div class="mdui-textfield-helper">目录索引起始文件夹地址，文件或文件夹名不能以点开始或结束，且不能包含以下任意字符: " * : <>? / \ | 否则无法索引。
                </div>
            </div>
            <br>
            <div class="mdui-textfield mdui-textfield-floating-label">
                <label class="mdui-textfield-label" for="expires">缓存时间(分钟)</label>
                <input type="text" class="mdui-textfield-input" id="expires" name="expires"
                       value="10">
                <div class="mdui-textfield-helper">建议缓存时间小于60分钟，否则会导致缓存失效</div>
            </div>
            <br>
            <div class="mdui-textfield mdui-textfield-floating-label">
                <label class="mdui-textfield-label" for="encrypt_path">加密</label>
                <textarea name="encrypt_path" id="encrypt_path" class="mdui-textfield-input"
                          rows="3"></textarea>
                <div class="mdui-textfield-helper">填写需要加密的文件或文件夹路径，格式如： /path1/xxx/ /path2/xxx/ password1,/path3/xxx/
                    /path4/xxx/ password2 (以OneDrive根目录为基础)
                </div>
            </div>
            <br>
            <label for="image_hosting" class="mdui-textfield-label">加密选项</label> &nbsp; &nbsp;
            <br>
            <label class="mdui-checkbox">
                <input type="checkbox" name="encrypt_option[]" value="list"
                       checked/>
                <i class="mdui-checkbox-icon"></i>
                加密目录列表
            </label> &nbsp; &nbsp;
            <label class="mdui-checkbox">
                <input type="checkbox" name="encrypt_option[]" value="show"
                       checked/>
                <i class="mdui-checkbox-icon"></i>
                加密文件预览
            </label> &nbsp; &nbsp;
            <label class="mdui-checkbox">
                <input type="checkbox" name="encrypt_option[]" value="download"
                       checked/>
                <i class="mdui-checkbox-icon"></i>
                加密文件下载
            </label> &nbsp; &nbsp;
            <label class="mdui-checkbox">
                <input type="checkbox" name="encrypt_option[]" value="view"
                       checked/>
                <i class="mdui-checkbox-icon"></i>
                加密图片查看页
            </label> &nbsp; &nbsp;
            <br>
            <br>
            <label for="image_hosting" class="mdui-textfield-label">是否开启图床</label> &nbsp; &nbsp;
            <br>
            <select name="image_hosting" id="image_hosting" class="mdui-select" mdui-select="{position: 'bottom'}">
                <option value="1" selected>开启</option>
                <option value="0">关闭</option>
                <option value="2">仅管理员开启</option>
            </select>
            <br>
            <br>
            <label for="image_home" class="mdui-textfield-label">是否将图床设为首页</label> &nbsp; &nbsp;
            <br>
            <select name="image_home" id="image_home" class="mdui-select" mdui-select="{position: 'bottom'}">
                <option value="">开启图床设为首页</option>
                <option value="1" selected>开启</option>
                <option value="0">关闭</option>
            </select>
            <div class="mdui-textfield mdui-textfield-floating-label">
                <label class="mdui-textfield-label" for="image_hosting_path">OneDrive中图床保存地址</label>
                <input type="text" class="mdui-textfield-input" id="image_hosting_path" name="image_hosting_path"
                       value="Hosting">
                <div class="mdui-textfield-helper">文件或文件夹名不能以点开始或结束，且不能包含以下任意字符: " * : <>? / \ |</div>
            </div>

            <br>
            <div class="mdui-textfield mdui-textfield-floating-label">
                <label class="mdui-textfield-label" for="hotlink_protection">防盗链</label>
                <input type="text" class="mdui-textfield-input" id="hotlink_protection" name="hotlink_protection"
                       value="">
                <div class="mdui-textfield-helper">留空则不开启。白名单链接以空格分开（此处采用 Http Referer 防盗链机制，如需加强请自行从服务器层面配置）</div>

            </div>
            <br>
            <div class="mdui-textfield mdui-textfield-floating-label">
                <label class="mdui-textfield-label" for="copyright">自定义版权显示</label>
                <input type="text" class="mdui-textfield-input" id="copyright" name="copyright"
                       value="">
                <div class="mdui-textfield-helper">留空则不显示。使用markdown格式表示 如：Made by [xxx](https://xxx)</div>
            </div>

            <br>
            <div class="mdui-textfield mdui-textfield-floating-label">
                <label class="mdui-textfield-label" for="statistics">统计代码</label>
                <input type="text" class="mdui-textfield-input" id="statistics" name="statistics"
                       value="">
                <div class="mdui-textfield-helper">js 统计代码</div>
            </div>

            <br>

            <button class="mdui-btn mdui-color-theme-accent mdui-ripple mdui-float-right" type="submit"><i
                        class="mdui-icon material-icons">check</i> 保存
            </button>
        </form>
    </div>
</div>

{{/*{{.Script}}*/}}
<script src="https://cdn.jsdelivr.net/combine/npm/mdui@0.4.2/dist/js/mdui.min.js,npm/jquery@3,npm/sweetalert2@7/dist/sweetalert2.min.js"></script>
</body>

</html>
