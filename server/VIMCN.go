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

package server

import (
	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/tools"
)

type VimCN struct {
	FileLimit []string
	MaxSize   int
}

func (s *VimCN) Upload(image *ImageParam) (ImageReturn, error) {
	url := "https://img.vim-cn.com/"

	file := &auxpi.FormFile{
		Name:  image.Name,
		Key:   "image",
		Value: *image.Content,
		Type:  image.Type,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	return ImageReturn{
		Url: data,
		ID:  15,
	}, nil
}

func (s *VimCN) UploadToVimCN(img []byte, imgInfo string, imgType string) string {
	url := "https://img.vim-cn.com/"
	name := tools.GetFileNameByMimeType(imgInfo)

	file := &auxpi.FormFile{
		Name:  name,
		Key:   "image",
		Value: img,
		Type:  imgType,
	}
	var header map[string]string
	data := tools.FormPost(file, url, header)
	return data
}
