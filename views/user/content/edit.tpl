<div class="mdui-row">
    <div class="mdui-col-sm-6 mdui-col-md-5">
        <div class="mdui-card mdui-hoverable">
            <!-- 卡片头部，包含头像、标题、副标题 -->

            <div class="mdui-card-header">
                <img class="mdui-card-header-avatar" src="https://ws3.sinaimg.cn/large/5d821655ly1g0y2bmcsfxj207g07gt90.jpg"/>
                <div class="mdui-card-header-title">{{ .User.Username}}</div>
                <div class="mdui-card-header-subtitle">注册于 {{ .User.CreatedDay }}</div>
            </div>

            <!-- 卡片的媒体内容，可以包含图片、视频等媒体内容，以及标题、副标题 -->
            <div class="mdui-card-media">
                <img src="https://ws3.sinaimg.cn/large/5d821655ly1g0y2bq08jbj20ia0bdaal.jpg"/>
            </div>


            <!-- 卡片的标题和副标题 -->
            <div class="mdui-card-primary">
                <div class="mdui-card-primary-title">信息总览</div>
            </div>

            <!-- 卡片的内容 -->
            <div class="mdui-card-content">
                <div class="mdui-table-fluid">
                    <table class="mdui-table mdui-table-hoverable">
                        <thead>
                        <tr>
                            <th>用户 ID</th>
                            <th>用户组</th>
                            <th>api Token</th>
                            <th>用户版本号</th>
                        </tr>
                        </thead>
                        <tbody>
                        <tr>
                            <td>{{ .User.ID}}</td>
                            <td>{{ .User.Role.Name}}</td>
                            <td>{{ .User.Token}}</td>
                            <td>{{ .User.Version}}</td>
                        </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>

    <div class="mdui-col-sm-6 mdui-col-md-7 mdui-hoverable">
        <div class="mdui-panel" mdui-panel>
            <div class="mdui-panel-item mdui-panel-item-open">
                <div class="mdui-panel-item-header mdui-text-color-blue"><i
                            class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-blue">settings</i>&nbsp;基本资料
                </div>
                <div class="mdui-panel-item-body">

                    <div class="mdui-textfield mdui-textfield-floating-label">
                        <label class="mdui-textfield-label">User Name</label>
                        <input class="mdui-textfield-input" value="{{ .User.Username}}" type="text" disabled/>
                    </div>

                    <div class="mdui-textfield mdui-textfield-floating-label">
                        <label class="mdui-textfield-label">Email</label>
                        <input class="mdui-textfield-input" value="{{ .User.Email}}" type="email" disabled/>
                    </div>
                </div>
            </div>

            <div class="mdui-panel-item mdui-panel-item-open">
                <div class="mdui-panel-item-header  mdui-text-color-green"><i
                            class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-green">lock</i>&nbsp;修改密码
                </div>
                <div class="mdui-panel-item-body">
                    <form action="" class="reset-validation" method="POST">
                        <div class="mdui-textfield mdui-textfield-floating-label">
                            <label class="mdui-textfield-label">现在密码</label>
                            <textarea class="mdui-textfield-input" name="old_password" required autofocus></textarea>
                        </div>

                        <div class="mdui-textfield mdui-textfield-floating-label">
                            <label class="mdui-textfield-label">新密码</label>
                            <textarea class="mdui-textfield-input" name="new_password" required autofocus></textarea>
                        </div>

                        <div class="mdui-textfield mdui-textfield-floating-label">
                            <label class="mdui-textfield-label">重复密码</label>
                            <textarea class="mdui-textfield-input" name="re_password" required autofocus></textarea>
                        </div>

                        <br>
                        <button class="mdui-btn mdui-btn-raised mdui-btn-dense mdui-color-theme-accent mdui-ripple"
                                type="submit">提交
                        </button>
                    </form>
                </div>
            </div>

            <div class="mdui-panel-item mdui-panel-item-open">
                <div class="mdui-panel-item-header mdui-text-color-purple"><i
                            class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-purple">computer</i>其它设置
                </div>
                <div class="mdui-panel-item-body">
                    <button class="mdui-btn mdui-btn-raised mdui-btn-dense mdui-color-theme-accent mdui-ripple">刷新 api
                        Token
                    </button>
                </div>
            </div>

        </div>
    </div>
</div>


<script src="/static/jquery/3.3.1/jquery.min.js"></script>

<script>


    $(".reset-validation").submit(function (e) {
        var form = $(this);
        e.preventDefault();
        e.stopPropagation();
        var xsrftoken = $('meta[name=_xsrf]').attr('content');
        var data = form.serializeArray();
        var token = {"name": "_xsrf", "value": xsrftoken};
        data.push(token);
        $.ajax({
            url: "/users/{{ .User.ID}}/edit",
            type: 'post',
            data: data,
            dataType: 'json',
            success: function (resp) {
                mdui.snackbar({
                    message: resp.msg,
                    position: 'right-top',
                });
            },
            error: function () {
                mdui.snackbar({
                    message: "修改失败",
                    position: 'right-top',
                });
            },
        });
    });
</script>