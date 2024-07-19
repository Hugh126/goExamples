> 由于对go的不熟悉，本项目是为了试验Go的运用，及一些第三方模块的使用  

## Go配置及入门：

参考[微软文档：](https://learn.microsoft.com/zh-cn/training/modules/go-get-started/2-install-go?pivots=linux)

1. 下载

    [https://go.dev/dl/](https://go.dev/dl/)
2. 解压到local

    ```shell
    sudo tar -zxvf go1.20.14.linux-amd64.tar.gz -C /usr/local/
    ```
3. 配置环境变量

    1. sudo vim /etc/profile

        ```shell
        export GOROOT\=/usr/local/go
        export PATH\="\$PATH:\$GOROOT/bin"
        # 如果没有这个目录，可以创建或者指定其他
        export GOPATH\=\$HOME/goWorkSpace
        ```
    2. source /etc/profile
    3. 新开终端输入`go env`查看配置变量
4. [Go代理配置](https://goproxy.cn/)

    ```shell
    $ go env -w GO111MODULE=on
    $ go env -w GOPROXY=https://goproxy.cn,direct
    ```
5. 运行

    ```go
    package main
    import "fmt"
    func main() {
    	fmt.Println("Hello World!")
    }

    ```

    在代码目录下运行 `go run main.go`
6. Go工作区

    传统每个 Go 工作区都包含三个基本文件夹：

    * *bin*：包含应用程序中的可执行文件。
    * *src*：包括位于工作站中的所有应用程序源代码。
    * *pkg*：包含可用库的已编译版本。 编译器可以链接这些库，而无需重新编译它们

    > 自从 Go 1.11 引入 Go Modules 之后，Go 的工作区结构变得更加灵活。使用 Go Modules 时，你不再需要严格按照 src、bin 和 pkg 的结构组织代码
    >
7. Go Modules（包管理）

    GO111MODULE 有三个值：off, on和auto（默认值）。

    * GO111MODULE\=off，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。
    * GO111MODULE\=on，go命令行会使用modules，而一点也不会去GOPATH目录下查找。
    * GO111MODULE\=auto，默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：

      1. 当前目录在\$GOPATH/src之外且该目录包含go.mod文件
      2. 当前文件在包含go.mod文件的目录下面。
8. Go work（工作区管理，多个包）

    ```shell
    # 初始化工作区文件，用于生成go.work工作区文件
    go work init {目录}
    # 添加新的模块到工作区
    go work use {模块目录1} ...
    ```