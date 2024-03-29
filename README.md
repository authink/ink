# ink

## Usage

```bash
$ ./bin/ink run
```

具体可以查看 help 信息

```bash
$ ./bin/ink -h
Usage:
  ink [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  gen         Generate DB Models
  help        Help about any command
  migrate     Migrate schema up or down
  run         Run server
  seed        Seed the database
  swag        Generate swagger docs

Flags:
  -h, --help   help for ink

Use "ink [command] --help" for more information about a command.
```

```bash
$ ./bin/ink help run
Run ink server

Usage:
  ink run [flags]

Flags:
  -h, --help          help for run
  -l, --live-reload   Enable live reload
```

## 前置条件

- 在工作目录克隆项目代码

```bash
$ cd {work dir}
$ git clone git@github.com:authink/ink.git
$ git clone git@github.com:authink/ink.schema.git
```

- 项目根目录中新建 .env.local

```conf
# .env.local
DB_USER={your_db_username}
DB_PASSWORD={your_db_password}
```

- 创建名为 ink 的数据库

```sql
CREATE DATABASE `ink` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
```

- make，方便执行 gen/swag/test/build/run 等任务

如果未安装 make，可参考 Makefile 中的任务命令，手动执行

## Schema

```bash
# up
$ APP_ENV=dev ./bin/ink migrate -d up

# down
$ APP_ENV=dev ./bin/ink migrate -d down
```

## Seed

```bash
$ APP_ENV=dev ./bin/ink seed
```

## Generate DB Models

```bash
$ make gen
```

## 生成 API swagger 文档

```bash
$ make swag
```

```go
// MainAPI 文件增加 import
_ "github.com/authink/ink/src/docs"
```

启动 Ink server

然后[打开 Swagger API 文档](http://localhost:8080/swagger/index.html)

## Quick Run

```bash
$ make run
```

## Live reload

此功能只能在 APP_ENV=dev 环境中开启

```bash
$ make run ARGS="-l"
```

## 单元测试/go test/CI

```bash
$ make test
```

## 部署/go build/CD

```bash
$ make build
```

## 搭建 Markdown Docs

## I18n

todo: ink translate
自动扫描所有未翻译字段，写入 locales 文件中

## Env

根据 APP_ENV 读取不同的 env，其中 .local 文件是在本地开发时用来覆盖默认配置，不会提交到 git 仓库

1. APP_ENV=dev

.env.dev.local > .env.local > .env.dev > .env

2. APP_ENV=test

.env.test.local > .env.local > .env.test > .env

3. APP_ENV=prod

.env.prod.local > .env.local > .env.prod > .env

提供 API 可以获取所有变量值
