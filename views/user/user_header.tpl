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