<div class="mdui-drawer" id="main-drawer">
    <div class="mdui-list" mdui-collapse="{accordion: true}">
        <a href="/" class="mdui-list-item mdui-ripple">
            <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-purple">home</i>
            <div class="mdui-list-item-content">首页</div>
        </a>
        <a href="/users/index" class="mdui-list-item mdui-ripple">
            <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-pink">image</i>
            <div class="mdui-list-item-content">图片管理</div>
        </a>
        <a href="/users/edit" class="mdui-list-item mdui-ripple">
            <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-blue">settings</i>
            <div class="mdui-list-item-content">个人设置</div>
        </a>



        {{ if .IsAdmin}}
        <a href="/admin" class="mdui-list-item mdui-ripple">
            <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-green">poll</i>
            <div class="mdui-list-item-content">后台管理</div>
        </a>
        {{ end}}


    </div>
</div>