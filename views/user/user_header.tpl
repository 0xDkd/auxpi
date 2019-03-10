<header class="mdui-appbar mdui-appbar-fixed">
    <div class="mdui-toolbar mdui-color-theme">
        <span class="mdui-btn mdui-btn-icon mdui-ripple mdui-ripple-white"
              mdui-drawer="{target: '#main-drawer', swipe: true}"><i class="mdui-icon material-icons">menu</i></span>
        <a href="{{urlfor "PagesController.UserIndexShow"}}" target="_blank"
           class="mdui-typo-headline mdui-hidden-xs" style="text-decoration:none">{{ .siteName}} 用户中心</a>
        <div class="mdui-toolbar-spacer"></div>
        <a href="/users/edit" class="mdui-btn"><i class="mdui-icon material-icons">face</i>{{ .User.Username}}</a>
        <a  href="/logout" class="mdui-btn mdui-btn-icon"><i class="mdui-icon material-icons">exit_to_app</i></a>
    </div>
</header>