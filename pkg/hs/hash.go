// Copyright 2025 TimeWtr
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/spaolacci/murmur3"
)

const (
	charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base    = 62
)

type Murmur3 struct{}

func NewMurmur3() Hasher {
	return &Murmur3{}
}

// 将64位整数转换为Base62字符串
func (m *Murmur3) encodeBase62(n uint64) string {
	if n == 0 {
		return string(charset[0])
	}
	var s string
	for n > 0 {
		s = string(charset[n%base]) + s
		n = n / base
	}
	return s
}

// ShortenURL 对URL进行哈希缩短计算
func (m *Murmur3) ShortenURL(url string) (string, error) {
	// 创建一个32位的MurmurHash3哈希器
	h := murmur3.New32()
	// 写入URL字节
	_, err := h.Write([]byte(url))
	if err != nil {
		return "", err
	}
	// 获取32位哈希值
	hash := h.Sum32()
	// 将哈希值转换为Base62字符串
	shortCode := m.encodeBase62(uint64(hash))
	return shortCode, nil
}
