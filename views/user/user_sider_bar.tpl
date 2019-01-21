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