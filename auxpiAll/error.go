// Copyright (c) 2019 aimerforreimu. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//  GNU GENERAL PUBLIC LICENSE
//                        Version 3, 29 June 2007
//
//  Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
//  Everyone is permitted to copy and distribute verbatim copies
// of this license document, but changing it is not allowed.
//
// repo: https://github.com/aimerforreimu/auxpi

package auxpi

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type Error struct {
	code  uint32
	msg   string
	where string
}

func (e *Error) Error() string {
	return fmt.Sprintf("code = %d ; msg = %s", e.code, e.msg)
}

func NewCoder(code uint32, msg string) *Error {
	where := caller(1, false)
	return &Error{code: code, msg: msg, where: where}
}

func Wrap(err error, extMsg ...string) *Error {
	msg := err.Error()
	if len(extMsg) != 0 {
		msg = strings.Join(extMsg, ":") + ":" + msg
	}
	return &Error{msg: msg}
}

func caller(calldepth int, short bool) string {
	_, file, line, ok := runtime.Caller(calldepth + 1)
	if !ok {
		file = "???"
		line = 0
	} else {
		file = filepath.Base(file)
	}

	return fmt.Sprintf("%s:%d", file, line)
}

//统一化处理为 并且转为 string 类型
//[xxxxx] : "xxxxxxxxxxxxxx"  ==>[file:line]
func ErrorToString(err error, tip ...string) string {
	if err != nil {
		content := fmt.Sprintf("%v", err)
		prefix := ""
		if len(tip) != 0 {
			prefix = "[" + tip[0] + "]"
		}
		e := prefix + content + `  ===>[` + caller(1, false) + "]"

		return e
	}
	return ""
}

func FormatError(err error, tip ...string) (fErr error) {
	fErr = fmt.Errorf("%v", ErrorToString(err))
	return
}
