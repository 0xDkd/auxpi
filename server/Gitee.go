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
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	auxpi "github.com/auxpi/auxpiAll"
	"github.com/auxpi/bootstrap"
	auxpiLog "github.com/auxpi/log"
	"github.com/auxpi/models"
)

type Gitee struct {
	FileLimit []string
	MaxSize   int
}

func (g *Gitee) Upload(image *ImageParam) (ImageReturn, error) {
	//必须要一下信息,才可以工作，否者返回对应的错误
	//access_token
	//owner
	//repo
	var giteeAccount = auxpi.GiteeAccount{}

	err := giteeAccount.UnmarshalJSON([]byte(models.GetOption("gitee", "conf")))
	if err != nil {
		auxpiLog.SetAWarningLog("SERVER", err)
	}

	if giteeAccount.AccessToken == "" {
		err := errors.New("gitee:access_token is none")
		return ImageReturn{}, err
	}
	if giteeAccount.Owner == "" {
		err := errors.New("gitee:owner is none")
		return ImageReturn{}, err
	}
	if giteeAccount.Repo == "" {
		err := errors.New("gitee:repo is none")
		return ImageReturn{}, err
	}

	name := bootstrap.GenerateImageName(image.Name)
	url := g.formatUrl(giteeAccount.Owner, giteeAccount.Repo, name)
	//上传逻辑
	benc := base64.StdEncoding
	content := benc.EncodeToString(*image.Content)
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	err = bodyWriter.WriteField("access_token", giteeAccount.AccessToken)
	if err != nil {
		return ImageReturn{}, err
	}

	err = bodyWriter.WriteField("content", content)
	if err != nil {
		return ImageReturn{}, err
	}

	msg := "commit by auxpi -> file original name: " + image.Name
	err = bodyWriter.WriteField("message", msg)
	if err != nil {
		return ImageReturn{}, err
	}
	contentType := bodyWriter.FormDataContentType()

	err = bodyWriter.Close()
	if err != nil {
		return ImageReturn{}, err
	}
	resp, err := http.Post(url, contentType, bodyBuf)

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ImageReturn{}, err
	}

	var gitteResp = auxpi.GiteeResp{}
	err = gitteResp.UnmarshalJSON(respBody)
	if err != nil {
		var maps = make(map[string]string)
		err1 := json.Unmarshal(respBody, &maps)
		if err1 != nil {
			return ImageReturn{}, err1
		}
		err = errors.New(maps["message"])
		return ImageReturn{}, err
	}

	return ImageReturn{
		Url:    gitteResp.Content.DownloadURL,
		Delete: gitteResp.Commit.Sha,
		ID:     20,
	}, nil

}

// https://gitee.com/api/v5/repos/Aimerfor/auxpi-store/contents/233adas3sa3.txt/?access_token=
// abe033927a8c7b6d1a0be2a74d665393
func (g *Gitee) formatUrl(owner, repo, name string) string {
	return "https://gitee.com/api/v5/repos/" + owner + "/" + repo + "/contents/" + name
}

//func (g *Gitee) createHttpClient {
//
//}
