<div class="card-body">
    <h4 class="card-title">注册</h4>
    <form method="POST" class="my-login-validation" novalidate="">
        <div class="form-group">
            <label for="name">用户名</label>
            <input id="name" type="text" class="form-control" name="name" required autofocus>
            <div class="invalid-feedback">
                用户名必须填写
            </div>
        </div>

        <div class="form-group">
            <label for="email">Email 地址</label>
            <input id="email" type="email" class="form-control" name="email" required>
            <div class="invalid-feedback">
                邮箱地址不能为空
            </div>
        </div>

        <div class="form-group">
            <label for="password">密码</label>
            <input id="password" type="password" class="form-control" name="password" required data-eye>
            <div class="invalid-feedback">
                密码不能为空
            </div>
        </div>

        <div class="form-group">
            <div class="custom-checkbox custom-control">
                <input type="checkbox" name="agree" id="agree" class="custom-control-input" required="">
                <label for="agree" class="custom-control-label">我同意 <a href="/about">条款</a></label>
                <div class="invalid-feedback">
                    您必须同意我们的服务条款
                </div>
            </div>
        </div>

        <div class="form-group m-0">
            <button type="submit" class="btn btn-primary btn-block">
                注册
            </button>
        </div>
        <div class="mt-4 text-center">
            已经有账户啦? <a href="/login">登录</a>
        </div>
    </form>
</div>