# HTTP Basic Authentication

> 📌 **适用场景**
>
> Basic auth 适用于在开放的 HTTP API 限制资源（接口）的访问，要求访问者必须提供用户名/密码进行认证。
> Basic auth 是 HTTP 标准的一部分，浏览器、Postman、服务端 web 框架等都有相应的支持，开发成本非常低，只需要实现用户名和密码的验证。
>
> **非常不推荐在具有登录图形界面的系统中使用 basic auth**。
> 这类系统几乎都需要实现登录注册界面、会话管理、权限体系，无法享受 basic auth 的简单、低成本的好处，反而在安全、拓展性上更差。

## 概述

Basic auth 是 HTTP 标准中的简单认证方案，在 [RFC 7617](https://datatracker.ietf.org/doc/html/rfc7617) 中定义。
Basic auth 要求客户端提供 `Authorization: Basic base64_encoded(username:password)` 请求头进行认证。

## 验证流程

1. 如果客户端请求不包含 `Authorization` 或验证失败，服务端回复 `401 Unauthorized` 状态码 和 `WWW-Authenticate: Basic realm="User Visible Realm"` 请求头。
2. (如果客户端是浏览器)浏览器弹框提示用户输入用户名和用户码，用户输入提交后，用 base64 编码，再次发送请求。
3. 服务端对请求的 `Authorization` 验证通过，认证通过，响应资源。

其中，`realm` 表示身份认证的域，由服务端指定，可以是任意格式的文本，客户端可见。浏览器等客户端会缓存用户名和密码，匹配到相同的 realm 时，自动填写认证内容，避免每次都提示用户填写。实践中经常用域名或资源地址前缀作为 realm 的值。

## 安全问题

- 必须和 HTTPS/LTS 加密传输协议配合使用。因为 base64 是明文编码，攻击者可以解码获得用户名和密码。
- 无法防范重放攻击，攻击者可以抓取请求后，原样发送，在不需要知道用户名和密码的情况下获取资源。

## 实现

在实现上，只需要在服务侧按规范对 `Authorization` 请求头进行校验，校验失败时，设置 `401 Unauthorized` 和 `WWW-Authenticate` 请求头即可。

### 手动实现

[example/basic-auth](../examples/basic-auth/server.go) 展示了如何手动实现 basic auth 的服务端逻辑。

### 使用中间件

几乎所有的 web server 框架，都会提供 basic auth 的中间件，我们只需要补充用户名和密码的校验逻辑即可。

[example/basic-auth-with-gin](../examples/basic-auth-with-gin/server.go) 展示了如何使用 gin 框架的 basic auth 中间件进行基本的身份验证。

### 登入与登出

使用浏览器时，访问 url，会自动弹出框，填入用户名和密码即可。
登录后，浏览器会保存用户名和密码，下次访问时，会自动填充。

浏览器重启后，下次需要重新输入用户名和密码。
如果想要手动清除用户名和密码，可以输入 `https://wrong_user:wrong_psw@example.com/path`

使用 postman 时，选择 `Authorization` 选项卡，选择 `Basic Auth`，输入用户名和密码即可。

## 【拓展】Digest Authentication

Digest Authentication 是为了解决 Basic Authentication 传递明文用户名和密码的问题而提出的（Base64 编码属于明文编码）。

Digest 的基本流程和 Basic 类似，但是，服务端的 `WWW-Authenticate` 和客户端的 `Authentication` 会包含更多的字段信息。
同时，凭据会被 MD5 或 SHA 加密，即使请求被拦截，也无法获得用户密码。
客户端请求可以在 nonce 字段携带一个随机数，服务端可以对 nonce 进行校验，防止重放攻击。

Digest 解决了密码明文和重放攻击，但是仍然无法解决中间人攻击，不能脱离 HTTPS 加密传输协议。
加上实现复杂，很少在实际中使用。

实际中，如果要求安全性，会使用 HTTPS 和 JWT 等更高级的认证方案。
如果要求成本低，就使用 Basic Authentication。

See [HTTP 摘要认证 - 维基百科，自由的百科全书](https://zh.wikipedia.org/wiki/HTTP%E6%91%98%E8%A6%81%E8%AE%A4%E8%AF%81)
