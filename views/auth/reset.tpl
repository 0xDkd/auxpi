<div class="card-body">
    <h4 class="card-title">重置密码</h4>
    <form method="POST" class="my-login-validation" novalidate="">
        <div class="form-group">
            <label for="new-password">新密码</label>
            <input id="new-password" type="password" class="form-control" name="password" required autofocus data-eye>
            <input type="hidden" name="reset" value="{{.resetToken}}">
            <div class="invalid-feedback">
                必须填写新密码
            </div>
            <div class="form-text text-muted">
                请确保您的密码足够健壮并请牢记
            </div>
        </div>

        <div class="form-group m-0">
            <button type="submit" class="btn btn-primary btn-block">
                确认重置
            </button>
        </div>
    </form>
</div>