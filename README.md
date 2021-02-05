# dobo | RESTful API 






## 更新记录 

v1.1 [2020-02-04]
- 支持 **PostgreSQL** 、 **MySQL** 和 **SQLite3** 数据库，驱动并内置
- 使用 **swagger** 生成接口文档
- 鉴权使用 **go-jwt** 
- **CORS** 跨域设置为 *
- 日志使用 **logrus** 丰富日志打印
- web框架为gin



## 支持的数据库

dobo 使用的是 [GORM] 作为 ORM. GORM 支持 **SQLite3**, **MySQL**,
**PostgreSQL** 和 **Microsoft SQL Server**.

在dobo中, **MySQL**, **PostgreSQL** and **SQLite3** 驱动已经内置，**Microsoft SQL Server**后期考虑加入。

**SQLite3需要注意:**
- `DBUSER`, `DBPASS`, `DBHOST` and `DBPORT` 保持默认
- `DBNAME` 需要设置数据库存放路径和名字; i.e,
```
./database.db
```
**本地运行**
- 本地运行需要开启go  mod
- 执行 go get -u github.com/swaggo/swag/cmd/swag  安装swag
- 执行 **go mod tidy**  下载依赖包
- 进入到 dobo目录下 将.env.sample 修改成.env
- 修改.env里面的系统配置信息
- 执行 swag init 安装swagger页面
- 执行go run main.go
- 访问 http://localhost:8083/swagger/index.html  访问接口页面


**初始化用户**
```
email:    admin@gibbon.com
password: 123456
```
**接口示例:**
- http://localhost:8083/api/v1/register 
  - `POST` [创建用户 不对外提供注册接口，需要使用Token进行注册]
```
{
    "Email":"example@example.com",
    "Password":"123456"
}
```
- http://localhost:8083/api/v1/login
  - `POST` [登录用户获取Token]
```
{
    "Email":"example@example.com",
    "Password":"123456"
}
```

### 文件列表

```
dobo
│---README.md
│---LICENSE
│---.gitignore
│---.env.sample
│---go.mod
│---go.sum
│---main.go
│
└───config
│    └---configMain.go
│    └---database.go
│    └---server.go
│
│───controller
│    └---auth.go
│    └---login.go
│    └---dockerController.go
│    └---imageController.go
│    └---xxxxx.go
│
└───docs
│    └─--docks.go
│    └─--swagger.json
│    └─--swagger.yml
│
└───database
│    └─--dbConnect.go
│    │
│    └───model
│         └---auth.go
│         └---docker.go
│         └---image.go
│         └---hobby.go
│         └---xxxxx.go
│
└───middleware
│         └---cors.go
│         └---jwt.go
│
└───service
│     └---auth.go
│
└───utils
      └---log.go
      └---RespUtils.go
      └---context.go
```

## License
MIT