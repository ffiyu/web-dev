# 身份认证和权限校验（Authentication and Authorization）

## 概述

用户身份和权限是所有具有用户体系的系统最基本的功能。

身份认证指的是辨识本次操作的用户是谁。

权限指的是当用户执行某个操作时，验证当前用户是否具有对应的操作权限。身份认证是权限校验的前提。

## 需求

功能性需求：

- 用户注册
- 登入和登出
- 会话管理
- 密码管理

非功能性需求：

- 安全性
- 性能
  - 支持大量用户(Scalable)
  - 认证和鉴权速度快(Performance)

## 认证方式

### Session-based vs Token-based Authentication

- Session-based: 服务端维护一个 session 表，为每个已登录的用户创建一个 session-id，并设置到 cookie 中。这样，客户端每次请求都会自动带上 session-id，服务端再根据 session 表查出请求对应的用户身份。
- Token-based: 用户登录后，服务端创建一个**含有用户信息和有效时间的 token**，设置到 cookie。客户端请求到来时，服务端解密 token，即可获得用户身份。

Session 和 Token 最大的区别是：session 是一段无意义的字符串 id，服务端拿到 session-id 后需要查找 session 表，而 token 是带有用户信息，服务器解密后即可获得用户身份。换句话说，session 是有状态的，token 是无状态的。

Token 的性能比 Session 要好。最流行的 Token 方式是 JWT（Json Web Token）。

### 认证方法

[认证方法](./authentication-strategies/README.md)

## 概念

### Single-Factor Authentication vs Multi-Factor Authentication

SFA（1FA）指的是通过一种方式认证，比如用户名/密码。

1FA 被认为是不安全的。对一些安全性较高的系统，需要结合验证码、指纹等方式。
采用多种方式结合认证，则称为 MFA。现在最常用的是 2FA。

系统可以结合用户行为判断是采用 1FA 还是 MFA。
比如，在平时使用 1FA，在用户使用新设备或异地登录时，采用 MFA。
或者，在用户进行更敏感的操作时，要求用户进行面部识别等二次认证。

###
