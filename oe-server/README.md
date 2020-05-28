项目结构说明
```javascript
    conf        配置
    db          数据库
    doa         数据与存储库交互层
    helpers     各种帮助类，封装
        |---middleware      中间件
        |---setting         获取配置块
        |---utils           各种常用封装(比如：加密，日志)
    logs        日志存储内容
    models      模型类层
    routers     路由层
    service     业务逻辑及服务层
```