

### badger基本用法

参考： https://pkg.go.dev/github.com/dgraph-io/badger#section-readme

1. 安装badger
```
go get github.com/dgraph-io/badger/v3
```


2. 创建数据库
```
db, err := badger.Open(badger.DefaultOptions("./data"))
// 设置值大小阈值，默认1MB，超过阈值则写入磁盘，而不是内存LSM树中
// db, err := badger.Open(badger.DefaultOptions("./data").WithValueThreshold(1<<20))
defer db.Close()
if err != nil {
    log.Fatal(err)
}
// 读写事务的封装；设置值，无则新增，有则修改
db.Update(func(txn *badger.Txn) error {
    err := txn.Set([]byte("hello"), []byte("world"))
    return err
})
// 只读事务的封装；查看值
db.View(func(txn *badger.Txn) error {
    item, err := txn.Get([]byte("hello"))
    if err != nil {
        return err
    }
    // 必须copy后使用
    val, err := item.ValueCopy(nil)
    fmt.Printf("hello=%s\n", val)   // hello=world
})
```


### 本demo中封装了badger的批量增删查改
```
go run badger/main.go
```

