<script src="/static/clipboard.js/2.0.1/clipboard.min.js"></script>
<script src="/static/bootstrap-fileinput/4.5.1/js/fileinput.min.js"></script>
<script src="/static/bootstrap-fileinput/4.5.1/js/locales/zh.js"></script>
<script src="/static/bootstrap/3.3.7/js/bootstrap.min.js"></script>
<script>

    var clipboard = new ClipboardJS('.copy', {
        text: function (trigger) {
            return $(trigger).parent('li').text();
        }
    });

    clipboard.on('success', function (e) {
        app.msg(true, '复制成功');
    });

    clipboard.on('error', function (e) {
        app.msg(false, '复制失败');
    });

    $("#image").fileinput({
        uploadUrl: "{{ urlfor "WebUpLoadController.UpLoadHandle"}}",
        language: "zh",
        uploadAsync: true,
        overwriteInitial: false,
        //browseClass: "btn btn-file",
        maxFileSize: "{{.maxPicSize}}",// kb
        maxFileCount: "{{.maxNumber}}",
        showCaption: true,
        dropZoneEnabled: true,
        browseIcon: "<i class=\"glyphicon glyphicon-picture\"></i> ",
        allowedFileExtensions: JSON.parse('["jpg","jpeg","gif","png","ico"]'),
        uploadExtraData: {
            "_xsrf":{{.xsrf_token}},
            "apiSelect": "{{ .apiSelect}}"
        }
    }).on("fileuploaded", function (event, data, previewId, index) {
        var form = data.form, files = data.files, extra = data.extra, response = data.response, reader = data.reader;
        if (200 === response.code) {
            app.msg(true,response.data.name+"上传完成");
            $("#code-url ul").prepend("<li>" + response.data.url + "<i class=\"copy iconfont icon-copy\"></i></li>");
            $("#code-html ul").prepend("<li>&lt;img src=\"" + response.data.url + "\" alt=\"" + response.data.name + "\" title=\"" + response.data.name + "\" /&gt;<i class=\"copy iconfont icon-copy\"></i></li>");
            $("#code-bbcode ul").prepend("<li>[img]" + response.data.url + "[/img]<i class=\"copy iconfont icon-copy\"></i></li>");
            $("#code-markdown ul").prepend("<li>![" + response.data.name + "](" + response.data.url + ")<i class=\"copy iconfont icon-copy\"></i></li>");
            $("#code-markdown-with-link ul").prepend("<li>[![" + response.data.name + "](" + response.data.url + ")](" + response.data.url + ")<i class=\"copy iconfont icon-copy\"></i></li>");
            $(".success-info").css("width", "inherit");
            if (response.data.quota && response.data.use_quota) {
                $('.quota-container progress').attr('max', response.data.quota);
                $('.quota-container progress').val(response.data.use_quota);
                $('.quota-container span.quota').text(app.bytesToSize(response.data.quota));
                $('.quota-container span.use-quota').text(app.bytesToSize(response.data.use_quota));
            }
        } else if (500 === response.code) {
            mdui.alert(response.msg, '发生异常');
        } else {
            mdui.alert(response.msg);
        }
    });
    $('#image').css("display", "block");
</script>

<script>
    $(function () {
        var toTop = $("#to-top"), toTopHide = function () {
            if ($(window).scrollTop() > 50) {
                toTop.removeClass('mdui-fab-hide');
            } else {
                toTop.addClass('mdui-fab-hide');
            }
        };
        if ($(window).scrollTop() > 50) {
            toTopHide();
        }
        $(window).scroll(function () {
            toTopHide();
        });

        toTop.click(function () {
            $('body,html').animate({scrollTop: 0}, 500);
        });

        $('.open-menu').click(function () {
            if ($(window).width() > 1024) {
                app.cookie.set('menu', $('body').hasClass('mdui-drawer-body-left') ? 'open' : 'close', 10, '/');
            }
        });
    });
</script>