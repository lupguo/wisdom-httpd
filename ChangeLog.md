# ChangeLog

## 2025/1/20

[x] 优化Webpack，支持执行`package.json`中`script`命令实现不同环境配置前端项目构建 -> 👍
[x] 本地基于Nginx配置快速部署
[x] 开发`PopClip插件`，支持通过PopClip请求`/api/save`保存信息
[x] 优化下`/api/save`接口，支持`Skey`密钥限制，防止被人恶意POST信息请求

## 2025/1/19

[x] 基于Caddy部署 `wisdom-httpd` 服务，支持每30s更换一次wisdom内容
[x] 创建一个`Cron`定时任务，支持定期备份到指定目录

## 2025/1/17

1. `webpack.config.js`重新配置，参考前后端分离方案，让Go渲染`wisdom.tmpl`模版时候，仅用包含`main.js`即可
2. wisdom-http服务响应，支持`CORS`处理

## 2025/1/15

1. Web日志、DB日志规范和重构
2. 重构`GetOneWisdom`、`SaveWisdom`接口，基于DB来存储和查询金句
3. 新增`ctx.GetHTTPReqEntry()`、`crp.xxReq`、`crp.xxRsp`来处理请求和响应参数
4. 新增`tool`工具模块
5. 解决了一些Bug

## 2025/1/2

1. 重构了wisdom的框架处理流程，使之更切合WebHandler模式

## 2024/12/24

1. 重构了一部分代码，使之更容易理解