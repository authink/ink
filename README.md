# ink.go

## schema

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

# 配置环境变量
source .env

# run
air run
```

## 升级 golang 1.22

## 单元测试/go test/CI

```bash
go test ./src/...
```

## 部署/go build/CD

## API swagger 文档

## 搭建 Markdown Docs

## i18n
