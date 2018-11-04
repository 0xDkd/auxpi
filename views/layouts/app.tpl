<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0,maximum-scale=1.0, user-scalable=no"/>
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>
    <meta name="renderer" content="webkit">
    <meta http-equiv="Cache-Control" content="no-siteapp"/>
    <meta name="keywords" content="{{.siteName}}"/>
    <meta name="description" content="{{.siteName}} sina Pictures & SouGou Pirtures & pictures upload"/>
    <title>{{ .siteName}}</title>

    <link href="/static/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/bootstrap-fileinput/4.5.1/css/fileinput.min.css" rel="stylesheet">

    <link rel="shortcut icon" href="/favicon.ico">
    <link rel="stylesheet" href="/static/app/iconfont/iconfont.css">
    <link rel="stylesheet" href="/static/app/newicon/iconfont.css">
    <link rel="stylesheet" href="/static/mdui/0.4.1/css/mdui.min.css">
    <link rel="stylesheet" href="/static/app/css/app.css">
    <!--[if IE]>
    <script>window.location.href = '/compatibility.html';</script>
    <![endif]-->
</head>
<body class="mdui-theme-accent-indigo mdui-appbar-with-toolbar ">
{{/*header*/}}
{{ .Header}}
{{/*left*/}}
{{ .Left}}
{{/*content*/}}
<div class="mdui-container">
    <main>
    {{.LayoutContent}}
    </main>
</div>
{{/*footer*/}}
{{ .Footer}}

</body>
<script src="/static/jquery/3.3.1/jquery.min.js"></script>
<script src="/static/mdui/0.4.1/js/mdui.min.js"></script>
<script src="/static/app/js/app.js"></script>

{{ .Scripts}}
</html>