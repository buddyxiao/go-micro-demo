# 微服务学习

使用的技术：go-micro、gorm、gin、consul、sentinel

采用go1.18 的 go work 建立项目工作区

微服务框架 go-micro : https://github.com/go-micro/go-micro

ORM框架 gorm: https://gorm.io/zh_CN/docs/index.html

HTTP框架 gin: https://gin-gonic.com/zh-cn/

服务发现组件 consul: https://www.consul.io/

限流熔断组件 sentinel: https://sentinelguard.io/zh-cn/docs/golang/quick-start.html

链路追踪组件 jaeger: https://www.jaegertracing.io  
追踪API规范 opentracing:https://opentracing.io/

项目结构：

    cmd    
    common
    gin-api-gateway
    rand-service
    user-service
    
## 学习资料
https://space.bilibili.com/478093818
https://www.bilibili.com/video/BV1zz411v7ye

## 下一步学习

seata 分布式事务: https://seata.io/zh-cn/