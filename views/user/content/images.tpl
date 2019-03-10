<div class="mdui-row">
    {{range $index, $image := .Images}}
    <div class="mdui-col-sm-6 mdui-col-md-3" style="margin:10px 0px 0px 0px;">
    <div class="mdui-card">
            <div class="mdui-card-media">
                <img src="{{ $image.Link}}" style="height:221px;width: 400px;"/>
                <div class="mdui-card-media-covered">
                    <div class="mdui-card-primary">
                        <div class="mdui-card-primary-title">上传于</div>
                        <div class="mdui-card-primary-subtitle">{{$image.CreatedDay}}</div>
                    </div>
                </div>
            </div>
            <div class="mdui-card-actions">
                {{/*<button class="mdui-btn mdui-btn-raised mdui-btn-dense mdui-color-theme-accent mdui-ripple">查看信息</button>*/}}
                <button class="mdui-btn mdui-ripple mdui-color-theme-accent  mdui-btn-raised mdui-btn-dense" data-clipboard-action="copy" data-clipboard-text="{{ $image.Link}}" id="copy-url">复制链接</button>
                {{/*<button class="mdui-btn mdui-ripple mdui-color-red mdui-btn-raised">删除</button>*/}}

            </div>
        </div>
    </div>
    {{end}}
</div>

<div class="mdui-row" style="margin: 30px 0px 0px 10px">
    {{str2html .Page}}
</div>


<script src="/static/jquery/3.3.1/jquery.min.js"></script>
