# Go-Web-Study

项目高版本: `go version 1.20+`

执行`go mod tidy`拉取项目依赖



<br>
`sqlc` : sqlc https://github.com/kyleconroy/sqlc



`tables/queries`目录下: 

- 如果有新表加入,请修改tables
- 如果有新的业务sql语句加入, 请修改对应业务文件的sql\

在项目根目录下执行  sqlc generate, 自动生成sql业务逻辑