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
│    └---user.go
│    └---post.go
│    └---hobby.go
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
│         └---user.go
│         └---post.go
│         └---hobby.go
│         └---userHobby.go
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
```

## License
