// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fsnotify from examples-go
// Author ixyxj 729104978@qq.com

package fsnotify

import (
	"os"
	"fmt"
	"io"
	"strings"
	"bufio"
	"io/ioutil"
)


// GetAppendText 获取改变的文本
func GetAppendText(path string, offset int64) (string, int64) {
	// 排除负数
	if offset < 0 {
		offset = 0
	}
	f, err := os.Open(path)
	if err != nil {
		fmt.Errorf("can not open file path %s, error: %v\n", path, err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		fmt.Errorf("can not open file stat:%v\n", err)
	}

	current := stat.Size()

	// 如果文件字符减少返回字符空，并将偏移量回传
	if current <= offset {
		return "", current
	}

	// 文件偏移量seek
	if _, err := f.Seek(offset, io.SeekStart); err != nil {
		fmt.Errorf("can not seek file :%v\n", err)
	}

	c, _ := ioutil.ReadAll(bufio.NewReader(f))

	// 取最小偏移，判断是删除还是增加
	return strings.TrimSpace(string(c)), current
}
