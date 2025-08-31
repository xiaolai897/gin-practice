# gin-practice


此项目依赖数据库：postgresql



配置文件 /config/config.yaml

```yaml
domain: "localhost"

postgres:
  host: "127.0.0.1"
  port: "5432"
  user: "postgres"
  password: "postgres"
  dbname: "test"
  schema: "public"
  time-zone: "Asia/Shanghai"
  sslmode: "disable"
  max-idle-conns: 20
  max-open-conns: 100

zap:
  level: "info"
  path: "logs/blog.log"
  max-size: 5
  max-backup: 30
  max-age: 180
  compress: true

```



初始化数据库

```go
go run cmd/migrate/main.go -c config/config.yaml
```



运行项目

```go
go run cmd/main.go -c config/config.yaml
```



用户注册

```
curl -X POST -H -H "Content-Type: application/json" -d '{"name":"name","email":"email","username":"username","password":"password"}' http://127.0.0.1:7900/register
```



用户登录

```
curl -X POST -H -H "Content-Type: application/json" -d '{"username":"username", "password":"password"}' http://127.0.0.1:7900/login
```
