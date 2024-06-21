# minimark
小型超市商品管理系统

## 构建
1. 克隆仓库
```shell
git clone git@github.com:zhanglp0129/minimark.git
```

2. 进入项目目录
```shell
cd minimark
```

3. 编译
```shell
go build -o minimark minimark/cmd/main
```

4. 创建数据库
```sql
-- 以下为sql语句，不是shell命令，请在mysql或sqlserver终端运行
create database minimark;
```

5. 编写配置文件
```shell
cp config_sample.toml config.toml
vim config.toml
```

6. 运行
```shell
./minimark
```