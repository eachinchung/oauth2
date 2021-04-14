# Oauth2 · [![GitHub license](https://img.shields.io/github/license/EachinChung/oauth2)](https://github.com/EachinChung/oauth2/blob/main/LICENSE) [![test](https://img.shields.io/badge/platform-Gin-9cf)]() [![GitHub stars](https://img.shields.io/github/stars/EachinChung/oauth2)](https://github.com/EachinChung/oauth2/stargazers) [![Release](https://img.shields.io/github/release/EachinChung/oauth2)](https://github.com/EachinChung/oauth2/releases)

一个基于Gin的Oauth2项目，通过一些简单的配置，就可以快速完成Oauth2登录以及生成JWT Token。

## 🚀 功能

- [x] 微信登录

## 📋 使用指南

1. 在根目录 conf 中，按 config_example.yaml 的信息新建一份 config.yaml 文件，qiniu用于七牛云的对象存储存储用户头像。
2. 按 init.sql 生成数据库表。

配置表中记得配置登录后跳转的系统与URL

```sql
INSERT INTO configs (config_key, version, config)
VALUES ('oauth2_system', 1, '{
  "oauth2": "http://127.0.0.1:3000/oauth2"
}');
```

3. 这时我们只要初始化go mod，并且运行air就可以把项目跑起来啦
4. 当系统运行起来时，我们也可以查看 [接口文档](http://127.0.0.1:8000/docs/index.html )
5. 微信登录示例：`https://open.weixin.qq.com/connect/oauth2/authorize?appid=你的appid&redirect_uri=http%3A%2F%2F127.0.0.1%3A8000%2Foauth2%2Fwechat%2Frequired%2Fsubscribe%3Fredirect%3D%2F%26service%3oauth2&response_type=code&scope=snsapi_base&state=state#wechat_redirect`



## LICENSE

本项目基于 [MIT](https://zh.wikipedia.org/wiki/MIT%E8%A8%B1%E5%8F%AF%E8%AD%89) 协议。