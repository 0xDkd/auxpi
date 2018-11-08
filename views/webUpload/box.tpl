<div class="upload-container">
    <div class="title">
        <h1>Image Upload <i
                class="mdui-list-item-icon mdui-icon material-icons iconfont icon-{{.iconStyle}} mdui-text-color-{{ .iconColor}}"></i>
        </h1>
        <p>最大可上传 {{ .maxPicSize}} KB的图片，单次同时可选择 {{ .maxNumber}} 张。</p>
    </div>
    <form action="" method="post" enctype="multipart/form-data">
        <input id="image" style="display: none;" type="file" multiple name="image" accept="image/*">
    </form>
    <div class="success-info">
        <div class="mdui-tab mdui-tab-scrollable" mdui-tab>
            <a href="#code-url" class="mdui-ripple mdui-tab-active">URL</a>
            <a href="#code-html" class="mdui-ripple">HTML</a>
            <a href="#code-bbcode" class="mdui-ripple">BBCode</a>
            <a href="#code-markdown" class="mdui-ripple">Markdown</a>
            <a href="#code-markdown-with-link" class="mdui-ripple">Markdown with link</a>
        </div>
        <div id="code-url">
            <ul></ul>
        </div>
        <div id="code-html">
            <ul></ul>
        </div>
        <div id="code-bbcode">
            <ul></ul>
        </div>
        <div id="code-markdown">
            <ul></ul>
        </div>
        <div id="code-markdown-with-link">
            <ul></ul>
        </div>
    </div>
</div>