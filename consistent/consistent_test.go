package consistent

import (
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	// 实例化一个Map，hash的方法直接将string转为uint32
	hash := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))

		return uint32(i)
	})

	// 往Map中添加节点
	hash.Add("6", "4", "2")

	// 添加测试用例
	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	// 验证测试用例
	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

	hash.Add("8")

	testCases["27"] = "8"

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}
}
