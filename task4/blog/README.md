个人博客系统设计
使用 Go 语言结合 Gin 框架和 GORM 库开发一个个人博客系统的后端，实现博客文章的基本管理功能，包括文章的创建、读取、更新和删除（CRUD）操作，同时支持用户认证和简单的评论功能，描述下大概设计框架，有那些模块，每个模块的作用。
1.	项目初始化
•	创建一个新的 Go 项目，使用 go mod init 初始化项目依赖管理。
•	安装必要的库，如 Gin 框架、GORM 以及数据库驱动（如 MySQL 或 SQLite）。
2.	数据库设计与模型定义
•	设计数据库表结构，至少包含以下几个表：
o	users 表：存储用户信息，包括 id 、 username 、 password 、 email 等字段。
o	posts 表：存储博客文章信息，包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段。
o	comments 表：存储文章评论信息，包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。
•	使用 GORM 定义对应的 Go 模型结构体。
3.	用户认证与授权
•	实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。
•	使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。
4.	文章管理功能
•	实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
•	实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
•	实现文章的更新功能，只有文章的作者才能更新自己的文章。
•	实现文章的删除功能，只有文章的作者才能删除自己的文章。
5.	评论功能
•	实现评论的创建功能，已认证的用户可以对文章发表评论。
•	实现评论的读取功能，支持获取某篇文章的所有评论列表。
6.	错误处理与日志记录
•	对可能出现的错误进行统一处理，如数据库连接错误、用户认证失败、文章或评论不存在等，返回合适的 HTTP 状态码和错误信息。
•	使用日志库记录系统的运行信息和错误信息，方便后续的调试和维护。
项目结构：
config 层：数据库、服务、Jwt配置、数据库初始化链接
controllers 层：控制器层，主要实现业务处理逻辑
global 层：全局变量
middleware 层:中间件主要提供鉴权、日志、异常统一处理
models 层:对象结构体
router 层：路由配置，初始化引擎实例
utils 层:实现鉴权相关功能
reponse 层：异常响应结构
下载需要的包：
go mod init  
go get -u github.com/gin-gonic/gin
go get github.com/spf13/viper
go get -u github.com/gin-gonic/gin  

go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u golang.org/x/crypto/bcrypt
go get github.com/golang-jwt/jwt/v5


不需要鉴权接口：
协议类型：POST
注册接口：/api/auth/register
参数：json类型：{
    "Username": "zk001",
    "Password": "123456"
}
登录接口：/api/auth/login
参数：json类型：{
    "Username": "zk",
    "Password": "123456"
}
查询文章列表接口：/api/getPostlist
参数：不需要

通过文章ID查询文章详情接口：/api/getPostById
参数：json类型：{ "id": 4}

通过文章ID获取文章评论接口：/api/getcommentsbypostID
参数：json类型：{"ID":4}

需要鉴权接口：
协议类型：POST
将登录获取到的token，放入下面每个接口Header中，如Authorization：Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTcyMzc1NTIsImlkIjoxLCJ0b2tlbnZlcnNpb24iOjcsInVzZXJuYW1lIjoiemsifQ.WxaMuvgjNIZJ62ezEMfo8fvhUPNy0TiVw7X0TtIY0yk

创建文章接口：/api/createpost
参数：json类型：{
"Title":"0906第五篇文章",
"Content":"如何使用go"}

更新文章接口：/api/updatePost
参数：json类型：id为文章id
{
    "id": 9,
    "title": "修改后的0906第五篇文章",
    "content": "如何使用go-修改后"
}
删除文章接口：/api/deletePost
参数：json类型：id为文章id
{"id": 4}

创建文章评论接口：/api/createcomment
{
"Content":"0906第五篇文章评论3",
"PostID":5
}