<div class="card-body">
    <h4 class="card-title">密码找回</h4>
    <form method="POST" class="my-login-validation" novalidate="">
        <div class="form-group">
            <label for="email">Email 地址</label>
            <input id="email" type="email" class="form-control" name="email" value="" required autofocus>
            <div class="invalid-feedback">
                email 不能为空
            </div>
            <div class="form-text text-muted">
                点击"重置密码"的按钮，我们会发送一封带有重置链接的邮件到您的邮箱
            </div>
        </div>

        <div class="form-group m-0">
            <button type="submit" class="btn btn-primary btn-block">
                重置密码
            </button>
        </div>
    </form>
</div>