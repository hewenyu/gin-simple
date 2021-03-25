## gin-simple
gin的一个样本的代码


## gin 为项目根目录
main.go 为入口文件
Router 为路由目录
Middlewares 为中间件目录
Controllers 为控制器目录（MVC）
Services 为服务层目录，这里把DAO逻辑也写入其中，如果分开也可以
Models 为模型目录
Databases 为数据库初始化目录
Sessions 为session初始化目录
文件 引用顺序 大致如下：
main.go(在main中关闭数据库) - router(Middlewares) - Controllers - Services(sessions) - Models - Databases