#### 基于gin的api脚手架
0、支持指定配置文件启动     
1、不破坏gin的原有特性，基本不影响原有gin的运行速度    
2、wire负责管理依赖注入  
3、集成常用组件gorm/zap/viper/go-redis等  
4、方便集成第三放组件  
5、可扩展解耦方便，只需要替换data层数据源即可（biz负责定义repo/usercase，data层负责实现repo）  
6、输入输出结构体单独提到api文档方便管理，约定对应具体的service模块   
说明：实际开发只需要在api目录定义输入输出结构体，router添加路由，internal目录分别实现service/biz/data层即可    
（数据校验等原有gin的方法照旧使用即可）

```
├── api  // 定义输入/输出的结构体
│   ├── user //对应user模块
│       ├── user.go //user模块相关的结构体
│   └── order //对应order模块
│       └── order.go //user模块相关的结构体
├── cmd  // 应用入口文件
│   ├── main.go //入口文件
│   ├── wire.go //依赖注入文件
│   ├── wire_gen.go //wire生成
├── configs  // 定义配置文件
├── internal  // 应用业务代码
|       └── biz //定义repo接口，解耦data实现
|       └── conf 
|       └── data //repo的具体实现
|           └── model //实体
|           └── data.go //注册各种组件
|       └── pkg //供其它目录使用的包
|       └── service //业务入口，处理业务校验等
├── middleware  // 中间件
├── pkg  // 公共使用的包
├── pkg  // 路由定义文件
    ├── app.go //初始化gin.Engine
    ├── router.go //自定义用户路由
└── third_party //第三方类库
└── tool //公共函数目录
└── docs //说明文档目录
└── go.mod
```
待增加功能：错误处理、平滑启动、多服务端管理