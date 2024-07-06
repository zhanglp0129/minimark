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

4. 打包前端代码
```shell
cd web
npm i
npm run build
cd ..
```

5. 编译后端代码
    - Linux
    ```shell
    go build -o minimark minimark/cmd/main
    ```
    - Windows
    ```shell
    go build -o minimark.exe minimark/cmd/main
    ```

6. 创建数据库
```sql
-- 以下为sql语句，请在mysql或sqlserver终端运行
create database minimark;
```

7. 编写配置文件
    - Linux
    ```shell
    cp config_sample.toml config.toml
    vim config.toml
    ```
    - Windows
    ```shell
    copy config_sample.toml config.toml
    notepad config.toml
    ```

8. 运行

   - Linux
    ```shell
    ./minimark
    ```
    - Windows的cmd终端
    ```shell
    minimark
    ```
    - Windows的PowerShell终端
    ```shell
    ./minimark.exe
    ```