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

package hs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenURL(t *testing.T) {
	url := "https://www.gog.com/test/appId=1234"
	m := NewMurmur3()
	sc, err := m.ShortenURL(url)
	assert.Nil(t, err)
	t.Logf("shorten url: %s", sc)
}

func BenchmarkShortenURL(b *testing.B) {
	url := "https://www.gog.com/test/appId=1234"
	m := NewMurmur3()
	for i := 0; i < b.N; i++ {
		sc, err := m.ShortenURL(url)
		assert.Nil(b, err)
		b.Logf("shorten url: %s", sc)
	}
}

func BenchmarkShortenURL_NoLog(b *testing.B) {
	url := "https://www.gog.com/test/appId=1234"
	m := NewMurmur3()
	for i := 0; i < b.N; i++ {
		_, err := m.ShortenURL(url)
		assert.Nil(b, err)
	}
}
