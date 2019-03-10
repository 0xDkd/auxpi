<!DOCTYPE html>
<html>
<head>
    <title>Email</title>
    <meta http-equiv="content-type" content="text/html; charset=UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0">
    <style>
        * {
            box-sizing: border-box;
        }
        body {
            font-family: 'PingFang SC','Helvetica Neue','Helvetica','STHeitiSC-Light','Arial','Microsoft yahei','\005fae\008f6f\0096c5\009ed1',Verdana,sans-serif;
            font-weight: 400;
            font-size: 14px;
            line-height: 1.6;
            color: #333333;
            background: #f2f5f8;
        }
        a {
            color: #3292ff;
            text-decoration: none;
            border-bottom-style: dotted;
            border-bottom-width: 1px;
        }
        a:hover{
            text-decoration:none !important;
            border-bottom-style: solid;
        }
        h1,h2,h3,h4,h5,h6{
            font-weight: 500;
            line-height: 1.5;
            margin-bottom: 20px;
        }
        p {
            padding: 10px 0;
            /*margin: 0 20px;*/
            margin-bottom: 10px;
        }
        .btn {
            font-weight: 500;
            border-radius: 4px;
            padding: 16px 30px;
            text-align: center;
            background: #3885ff;
            text-decoration: none;
            color: #fff;
            margin: 15px auto;
            font-size: 16px;
            display: inline-block;
            border: none;
        }
        .btn:hover{
            border: none;
        }
        .btn-success {
            background: #2ecc71
        }
        blockquote {
            border-radius: 3px;
            background: #f5f5f5;
            margin: 10px 0;
            padding: 15px 20px;
            color: #455667;
        }
        .center{
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="wrapper" style="margin:0;padding:0;">
    <table width="100%" border="0" cellspacing="0" cellpadding="0" align="center">
        <tbody>
        <tr>
            <td class="inner" style="padding:30px 25px 40px 25px;background:#f2f5f8;" bgcolor="#f2f5f8" align="center">
                <table width="750" cellspacing="0" cellpadding="0" border="0">
                    <tbody>
                    <tr>
                        <td width="100%" style="border-radius:4px;background:#ffffff;" bgcolor="#ffffff" align="center">
                            <!-- Header -->
                            <table width="100%" border="0" cellspacing="0" cellpadding="0">
                                <tbody>
                                <tr>
                                    <td width="50%" align="left"><div style="width:150px;height:60px;padding: 15px 0 0 20px;"><img src="{{ .}}"  title="{{ .SiteName}}" style="display:inline;margin:0;max-height:50px;width: auto;" border="0"></div></td>
                                    <td width="50%" align="right"><div style="width:250px;height:60px;line-height:60px;padding-right:20px;">
                                            <a href="{{ .SiteLink}}" title="{{ .SiteName}}" style="font-size:12px;line-height:60px;color:#222222;text-decoration:none;border:none;padding:0 6px;">首页</a>
                                                                                        <a href="{{ .UserIndex}}" title="" style="font-size:12px;line-height:60px;color:#222222;text-decoration:none;border:none;padding:0 6px;">我的图片</a>
                                                                                        <a href="{{ .UserCenter}}" title="" style="font-size:12px;line-height:60px;color:#222222;text-decoration:none;border:none;padding:0 6px;">个人中心</a>
                                        </div></td>
                                </tr>
                                </tbody>
                            </table>
                            <!-- Main Body -->
                            <table width="100%" border="0" cellspacing="0" cellpadding="0">
                                <tbody>
                                <tr>
                                    <td width="100%">
                                        <div style="padding:20px 20px 40px;font-size:15px;line-height:1.5;color:#3d464d;">
                                            <!-- Content -->

<style>
    img{max-width:100%;}
</style>
<p>{{ .Content | html}}</p>
<p style="padding:10px 15px;background-color:#f4f4f4;margin-top:10px;color:#000;border-radius:3px;"><a href="{{ .Active}}" target="_blank">{{ .Active}}</a></p>                                            <!-- ./ Content -->
                                        </div>
                                    </td>
                                </tr>
                                </tbody>
                            </table>
                            <!-- Footer -->
                        </td>
                    </tr>
                    <tr>
                        <!-- Outer Footer -->
                        <td width="100%" align="center" style="font-size:10px;line-height: 1.5;color: #999999;padding: 5px 0;text-align:center;">
                            <p style="margin:10px 0 0;">此为系统自动发送邮件, 请勿直接回复.</p>
                            <p style="margin:10px 0 0;">版权所有 © {{ .SiteName}}</p>
                        </td>
                    </tr>
                    </tbody>
                </table>
            </td>
        </tr>
        </tbody>
    </table>
</div>
</body>
</html>
