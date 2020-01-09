# infoweb

infoweb-api: go语言搭建的基础后端

使用[singo框架](https://singo.gourouting.com/)（Simple Single Golang Web Service）

## singo

Singo采用了一系列Golang中比较流行的组件，以此为基础快速搭建Restful Web API

Singo已经整合了许多开发API所必要的组件：

1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的 
2. [GORM](http://gorm.io/docs/index.html): ORM工具。本项目需要配合Mysql使用 
3. [Gin-Session](https://github.com/gin-contrib/sessions): Gin框架提供的Session操作工具
4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端
5. [godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具，方便使用环境变量
6. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件
7. 自行实现了国际化i18n的一些基本功能
8. 本项目是使用基于cookie实现的session来保存登录状态的，如果需要可以自行修改为token验证

## interface

1. 实现了```/api/v1/user/register```用户注册接口
2. 实现了```/api/v1/user/login```用户登录接口
3. 实现了```/api/v1/user/me```用户资料接口(需要登录后获取session)
4. 实现了```/api/v1/user/logout```用户登出接口(需要登录后获取session)
5. 实现了```/api/v1/user/info```信息查询接口

## 结构

本项目已经预先创建了一系列文件夹划分出下列模块:

1. api文件夹：MVC框架的controller，负责协调各部件完成任务
2. model文件夹：负责存储数据库模型和数据库操作相关的代码
3. service：负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
4. serializer：储存通用的json模型，把model得到的数据库模型转换成api需要的json对象
5. cache：负责redis缓存相关的代码
6. auth：权限控制文件夹
7. util：一些通用的小工具
8. conf：放一些静态存放的配置文件，其中locales内放置翻译相关的配置文件

## 运行

```shell
go run main.go
```

项目运行后启动在3000端口（可以修改，参考gin文档)
