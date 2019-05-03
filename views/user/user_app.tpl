<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <meta name="_xsrf" content="{{.xsrf_token}}" />
    <title>{{ .siteName}} 管理</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/mdui@0.4.2/dist/css/mdui.min.css">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/sweetalert2@7/dist/sweetalert2.min.css">
    <link rel="stylesheet" type="text/css" href="/static/bootstrap/css/bootstrap.min.css">
    <link rel="shortcut icon" href="/static/favicon.ico" type="image/x-icon">
    <style type="text/css">
        a:hover{
            text-decoration:none;
        }

    </style>

</head>
<body class="mdui-drawer-body-left  mdui-appbar-with-toolbar mdui-theme-primary-indigo mdui-theme-accent-blue">
{{ .Header}}
{{ .SiderBar}}
<a id="anchor-top"></a>
{{/*{{ .Content}}*/}}
<div class="mdui-container-fluid">
    <div class="mdui-container-fluid mdui-m-t-2 mdui-m-b-2">
{{ .Content}}
    </div>
</div>
{{ .Script}}
</body>

</html>
