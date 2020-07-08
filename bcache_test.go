package bcache

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(expect, v) {
		t.Errorf("callback failed")
	}
}

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

var loadCount = make(map[string]int, len(db))

func TestGet(t *testing.T) {
	// 实例化一个group，并写入callback
	bcache := NewGroup("scores", 2<<10, GetterFunc(getter))

	// 循环获取key为db中key的值（调用callback+命中缓存
	for k, v := range db {
		// load from callback function
		if value, err := bcache.Get(k); err != nil || value.String() != v {
			t.Fatalf("fail to get value of %s", k)
		}

		// cache hit
		if _, err := bcache.Get(k); err != nil || loadCount[k] > 1 {
			t.Fatalf("cache %s miss", k)
		}
	}

	// 查找一个不存在的值
	if v, err := bcache.Get("Unknown"); err == nil {
		t.Fatalf("the value of Unknown should be empty, but %s got", v)
	}
}

func getter(key string) ([]byte, error) {
	log.Printf("[SlowDB] search key: %s", key)

	if score, ok := db[key]; ok {
		if _, ok := loadCount[key]; !ok {
			loadCount[key] = 0
		}

		loadCount[key]++
		return []byte(score), nil
	}

	return nil, fmt.Errorf("%s not exists", key)
}
