# 微服务
> 微服务是一种设计模式，网上流行的说法是一种区别于MVC的设计模式，本人并不这样认为。
- 本人认为微服务是一种基于MVC发展出来的一种设计模式。
    - mvc 
    - oop
        - model
        - 方法设定重于范围
        - 接口设定重于结果
    - 数据库选型
        - mongo map
        - mysql struct
    - 分布一致性 
        - redis
            - 缓存
            - 分布锁
- golang与微服务
    - golang语言特性 （业务逻辑简单map型）
    - 建议golang按微服务来设计
- 容器化
    - 容器化范围
    - 状态机设计
- 容器件通信交流
    - 消息中间件（异步）
        >用于流量消峰
        - rabbitmq
        - kafka
        - redis(极其不推荐)
    - 链式中间件
        - 通信数据结构
            - json
            - probuffer(grpc)
        - 通信协议
            - rpc
                - grpc
                - 基于redis协议的rpc
            - restful
    - golang中两种中间件模式
        - 函数式
        - handler
- 任务
    - select 
    - time
- websocket
    - 与异步消息协同使用
- 流操作
    - golang的接口思想
    - bufio包
- 容器中的信息统一
    - context
    - sso单点登录

# net/http包
> The net/http package provides all the features we need to write HTTP clients and servers.