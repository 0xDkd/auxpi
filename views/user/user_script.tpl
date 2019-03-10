<script src="/static/clipboard.js/2.0.1/clipboard.min.js"></script>
<script src="/static/app/js/app.js"></script>
<script src="/static/bootstrap/js/bootstrap.min.js"></script>
<script src="/static/app/js/jq-paginator.min.js"></script>
<script src="https://cdn.jsdelivr.net/combine/npm/mdui@0.4.2/dist/js/mdui.min.js,npm/jquery@3,npm/sweetalert2@7/dist/sweetalert2.min.js"></script>

<script>
    // 监听复制操作
    var clipboard = new ClipboardJS('#copy-url');
    clipboard.on('success', function (e) {
        app.msg(true, '复制成功！');
        e.clearSelection();
    });

    clipboard.on('error', function (e) {
        console.error('Action:', e.action);
        console.error('Trigger:', e.trigger);
        app.msg(false, '复制失败！');
    });







    {{/*$("#page").bootstrapPaginator({*/}}
        {{/*currentPage: '{{.Page.PageNo}}',*/}}
        {{/*totalPages: '{{.Page.TotalPage}}',*/}}
        {{/*bootstrapMajorVersion: 3,*/}}
        {{/*size: "small",*/}}
        {{/*onPageClicked: function (e, originalEvent, type, page) {*/}}
            {{/*window.location.href = "/?page=" + page*/}}
        {{/*}*/}}
    {{/*});*/}}


</script>