<!DOCTYPE html>
<html lang="en">
{{ .Header}}
<body class="my-login-page">
<section class="h-100">
    <div class="container h-100">
        <div class="row justify-content-md-center align-items-center h-100">
            <div class="card-wrapper">
                <div class="brand">
                    {{/*/static/app/images/logo.jpg*/}}
                    <a href="{{ .SiteLink}}" target="_blank"><img src="{{ .Logo}}" alt="{{ .SiteName}}"></a>
                </div>
                <div class="card fat">
                    {{ .Content}}
                </div>
                <div class="footer">
                    Copyright &copy; {{ .Time}} &mdash; <a href="{{ .SiteLink}}" target="_blank">{{ .SiteName}}</a>
                </div>
            </div>
        </div>
    </div>
</section>
{{ .Footer}}
</body>
</html>