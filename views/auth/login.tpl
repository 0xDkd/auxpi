<div class="card-body">
    <h4 class="card-title">登录</h4>
    <form method="POST" class="my-login-validation" novalidate="">
        <div class="form-group">
            <label for="email">Email 地址</label>
            <input id="email" type="email" class="form-control" name="email" value="" required autofocus>
            <div class="invalid-feedback">
                邮件不合法
            </div>
        </div>

        <div class="form-group">
            <label for="password">密码
                <a href="/forgot" class="float-right">
                    忘记密码?
                </a>
            </label>
            <input id="password" type="password" class="form-control" name="password" required data-eye>
            <div class="invalid-feedback">
                密码是必须的
            </div>
        </div>

        <div class="form-group">
            <div class="custom-checkbox custom-control">
                <input type="checkbox" name="remember" id="remember" class="custom-control-input">
                <label for="remember" class="custom-control-label">记住我</label>
            </div>
        </div>

        <div class="form-group m-0">
            <button type="submit" class="btn btn-primary btn-block">
                登录
            </button>
        </div>
        <div class="mt-4 text-center">
            没有账号? <a href="/register">创建一个</a>
        </div>
    </form>
</div>