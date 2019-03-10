<div class="mdui-appbar mdui-appbar-fixed">
    <div class="mdui-toolbar mdui-color-black">
        <span mdui-drawer="{target: '#menu'}" class="mdui-btn mdui-btn-icon open-menu"><i
                class="mdui-icon material-icons">&#xe5d2;</i></span>
        <a href="/" class="mdui-typo-headline">{{ .siteName}}</a>
        <div class="mdui-toolbar-spacer"></div>
        {{if .IsLogin}}
        <a href="/users/edit" class="mdui-btn"><i class="mdui-icon material-icons">face</i>{{ .UserName}}</a>
        <a  href="/logout" class="mdui-btn mdui-btn-icon"><i class="mdui-icon material-icons">exit_to_app</i></a>
        {{else}}
        <a href="/login" class="mdui-btn"><i class="mdui-icon material-icons">account_circle</i>登录</a>
        <a href="/register" class="mdui-btn"><i class="mdui-icon material-icons">group_add</i>注册</a>
         {{end}}
    </div>
</div>
