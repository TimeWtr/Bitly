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

package logger

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestZapLogger_Debugf(t *testing.T) {
	zl, err := zap.NewDevelopment()
	assert.Nil(t, err)
	l := NewZapLogger(zl)
	l.Debugf("test debug", Field{
		Key: "key",
		Val: "value",
	})
}

func TestZapLogger_Infof(t *testing.T) {
	zl, err := zap.NewDevelopment()
	assert.Nil(t, err)
	l := NewZapLogger(zl)
	l.Infof("test info", Field{
		Key: "key",
		Val: "value",
	})
}

func TestZapLogger_Warnf(t *testing.T) {
	zl, err := zap.NewDevelopment()
	assert.Nil(t, err)
	l := NewZapLogger(zl)
	l.Warnf("test warn", Field{
		Key: "key",
		Val: "value",
	})
}

func TestZapLogger_Errorf(t *testing.T) {
	zl, err := zap.NewDevelopment()
	assert.Nil(t, err)
	l := NewZapLogger(zl)
	l.Errorf("test error", Field{
		Key: "key",
		Val: "value",
	})
}
