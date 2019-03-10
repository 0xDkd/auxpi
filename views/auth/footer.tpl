<script src="/static/jquery/3.3.1/jquery.min.js"></script>
<script src="/static/bootstrap/js/bootstrap.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/sweetalert2@8.2.5/dist/sweetalert2.all.min.js"></script>

<script>
    'use strict';

    $(function () {

        // author badge :)


        $("input[type='password'][data-eye]").each(function (i) {
            var $this = $(this),
                id = 'eye-password-' + i,
                el = $('#' + id);

            $this.wrap($("<div/>", {
                style: 'position:relative',
                id: id
            }));

            $this.css({
                paddingRight: 60
            });
            $this.after($("<div/>", {
                html: 'Show',
                class: 'btn btn-primary btn-sm',
                id: 'passeye-toggle-' + i,
            }).css({
                position: 'absolute',
                right: 10,
                top: ($this.outerHeight() / 2) - 12,
                padding: '2px 7px',
                fontSize: 12,
                cursor: 'pointer',
            }));

            $this.after($("<input/>", {
                type: 'hidden',
                id: 'passeye-' + i
            }));

            var invalid_feedback = $this.parent().parent().find('.invalid-feedback');

            if (invalid_feedback.length) {
                $this.after(invalid_feedback.clone());
            }

            $this.on("keyup paste", function () {
                $("#passeye-" + i).val($(this).val());
            });
            $("#passeye-toggle-" + i).on("click", function () {
                if ($this.hasClass("show")) {
                    $this.attr('type', 'password');
                    $this.removeClass("show");
                    $(this).removeClass("btn-outline-primary");
                } else {
                    $this.attr('type', 'text');
                    $this.val($("#passeye-" + i).val());
                    $this.addClass("show");
                    $(this).addClass("btn-outline-primary");
                }
            });
        });

        var mod = "{{ .Action}}";


        var xsrftoken = $('meta[name=_xsrf]').attr('content');

        const ToastMixin = Swal.mixin({
            toast: true,
            position: 'top-end',
            showConfirmButton: false,
            timer: 700
        });

        var login = function (data) {
            $.ajax({
                url: "/login",
                type: 'post',
                data: data,
                dataType: 'json',
                success: function (resp) {
                    if (resp.code != 200) {
                        ToastMixin.fire({
                            type: 'warning',
                            title: resp.msg,
                            showCloseButton: true
                        });
                        return
                    }
                    ToastMixin.fire({
                        type: "success",
                        title: "登录成功,即将跳转",
                        showCloseButton: true,
                    }).then(() => {
                        location.href = "/users/index"
                    })
                },
                error: function (resp) {
                    ToastMixin.fire({
                        type: "danger",
                        title: "超时",
                        showCloseButton: true,
                    })
                },
            });
        };

        var register = function (data) {
            $.ajax({
                url: "/register",
                type: 'post',
                data: data,
                dataType: 'json',
                success: function (resp) {
                    if (resp.code != 200) {
                        Swal.fire({
                            type: 'warning',
                            title: resp.msg,
                            html: resp.data,
                            showCloseButton: true
                        });
                        return;

                    }
                    Swal.fire(
                        '注册成功',
                        '请去您的邮箱查看验证邮件',
                        'success'
                    ).then(() => {
                        location.href = "/login"
                    })
                },
                error: function (resp) {
                    Swal.fire({
                        type: "danger",
                        title: "超时",
                        showCloseButton: true,
                    })
                },
            });
        };

        var forgot = function (data) {
            $.ajax({
                url: "/forgot",
                type: 'post',
                data: data,
                dataType: 'json',
                success: function (resp) {
                    if (resp.code != 200) {
                        Swal.fire({
                            type: 'warning',
                            title: resp.msg,
                            showCloseButton: true
                        });
                        return
                    }
                    Swal.fire(
                        '提交成功',
                        '请去您的邮箱查看找回邮件',
                        'success'
                    ).then(() => {
                        location.href = "/login"
                    })
                },
                error: function (resp) {
                    Swal.fire({
                        type: "danger",
                        title: "超时",
                        showCloseButton: true,
                    })
                },
            });
        };

        var reset = function (data) {
            $.ajax({
                url: '/reset',
                type: 'post',
                data: data,
                dataType: 'json',
                success: function (resp) {
                    if (resp.code != 200) {
                        Swal.fire({
                            type: 'warning',
                            title: resp.msg,
                            showCloseButton: true
                        });
                        return
                    }
                    Swal.fire(
                        '重置成功',
                        '请您使用新的密码进行登录',
                        'success'
                    ).then(() => {
                        location.href = "/login"
                    })
                },
                error: function (resp) {
                    Swal.fire({
                        type: "danger",
                        title: "超时",
                        showCloseButton: true,
                    })
                },
            });
        };


        $(".my-login-validation").submit(function (e) {
            var form = $(this);
            e.preventDefault();
            e.stopPropagation();
            form.addClass('was-validated');
            if (form[0].checkValidity() === false) {
                switch (mod) {
                    case "login":
                        ToastMixin.fire({
                            type: 'warning',
                            title: '请输入正确的用户名和密码',
                            showCloseButton: true
                        });
                        break;
                    case "register":
                        ToastMixin.fire({
                            type: 'warning',
                            title: '请输入正确的用户名,密码和邮箱',
                            showCloseButton: true
                        });
                        break;
                    case "forgot":
                        ToastMixin.fire({
                            type: 'warning',
                            title: '请输入正确的邮箱',
                            showCloseButton: true
                        });
                        break;
                    default:
                        ToastMixin.fire({
                            type: 'warning',
                            title: '请输入合法的密码',
                            showCloseButton: true
                        });
                        break;
                }
                return
            }

            var data = form.serializeArray();
            var token = {"name": "_xsrf", "value": xsrftoken};
            data.push(token);

            switch (mod) {
                case "login":
                    login(data);
                    break;
                case "register":
                    register(data);
                    break;
                case "forgot":
                    forgot(data);
                    break;
                default:
                    reset(data);
                    break;
            }
        });
    });
</script>
