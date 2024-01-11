# ink.go

## schema

```sql
CREATE DATABASE `ink` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

CREATE DATABASE `ink_test` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
```

```bash
migrate -database "mysql://username:password@tcp(localhost:3306)/ink" -path db/migrations up
```

```bash
migrate -database "mysql://username:password@tcp(localhost:3306)/ink" -path db/migrations down
```

## seed

```sql
CREATE TABLE users (
    id INT PRIMARY KEY,
    username VARCHAR(50) UNIQUE,
    email VARCHAR(50)
);

-- 使用INSERT IGNORE插入数据，如果用户名已存在则忽略
INSERT IGNORE INTO users (id, username, email)
VALUES (1, 'john_doe', 'john@example.com');
```

## hot reload

```bash
# install
go install github.com/cosmtrek/air@latest

# create .air.toml
air init
```

```toml
# .air.toml 增加如下一行
[envfile]
  path = ".env"
```

```bash
# run
air run
```

## 升级 golang 1.22

## 单元测试/go test/CI

设置 -> 搜(Test Env File)

`值: ${workspaceFolder}/.env.test`

可以通过 vscode 点击 Run Test，且加载 .env。

如果手动运行测试可安装 dotenv

```bash
go install github.com/joho/godotenv/cmd/godotenv@latest

godotenv -f .env.test go test -v -cover ./src/...

#手动运行程序
godotenv go run ./src/... run
```

- todo 单元测试开始执行，先在另一个测试库执行 migrate up，seed 测试数据，执行单元测试，完成后 migrate down，清理所有表。

```json
{
  "name": "Launch Ink",
  "type": "go",
  "request": "launch",
  "mode": "auto",
  "program": "${workspaceFolder}/src",
  "args": ["run"],
  "envFile": "${workspaceFolder}/.env"
}
```

## 部署/go build/CD

## API swagger 文档

## 搭建 Markdown Docs

## i18n
