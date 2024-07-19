package main

import (
	"log"
	"fmt"
	badger "github.com/dgraph-io/badger"
)
// 批量添加或修改键值对
func batchAddOrUpdateKeyValue(db *badger.DB, kvs map[string]string) error {
	return db.Update(func(txn *badger.Txn) error {
		for key, value := range kvs {
			if err := txn.Set([]byte(key), []byte(value)); err != nil {
				return err
			}
		}
		return nil
	})
}

// 查看键值对
func getKeyValue(db *badger.DB, key string) (string, error) {
	var value []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(nil)
		return err
	})
	return string(value), err
}

// 批量查看键值对
func batchGetKeyValue(db *badger.DB, keys []string) (map[string]string, error) {
	result := make(map[string]string)
	err := db.View(func(txn *badger.Txn) error {
		for _, key := range keys {
			item, err := txn.Get([]byte(key))
			if err != nil {
				return err
			}
			value, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			result[key] = string(value)
		}
		return nil
	})
	return result, err
}

// 批量删除键值对
func batchDeleteKeyValue(db *badger.DB, keys []string) error {
	return db.Update(func(txn *badger.Txn) error {
		for _, key := range keys {
			if err := txn.Delete([]byte(key)); err != nil {
				return err
			}
		}
		return nil
	})
}


func main() {
	// 打开一个Badger数据库
	// options := badger.DefaultOptions("badger_db")
	options := badger.DefaultOptions("/tmp/badger")
	options.Logger = nil // 关闭日志输出
	db, err := badger.Open(options)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 批量添加键值对
	kvs := map[string]string{
		"name": "Lily",
		"age":  "18",
	}
	err = batchAddOrUpdateKeyValue(db, kvs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("批量新增两个键值对")

	// 批量查看键值对
	keys := []string{"name", "age"}
	values, err := batchGetKeyValue(db, keys)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("批量查看两个键值对Value: %v\n", values)

	// 修改键值对
	kvs = map[string]string{
		"name": "Lily Updated",
		"age": "20",
	}
	err = batchAddOrUpdateKeyValue(db, kvs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("批量修改键值对")

	// 查看修改后的键值对
	values, err = batchGetKeyValue(db, []string{"name", "age"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("批量修改后的值为: %v\n", values)

	// 批量删除键值对
	err = batchDeleteKeyValue(db, keys)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("批量删除键值对")

	// 尝试查看已删除的键值对
	values, err = batchGetKeyValue(db, []string{"name", "age"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("批量删除后的值为: %v\n", values)
}