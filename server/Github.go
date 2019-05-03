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
	"fmt"
	"io/ioutil"
	"net/http"

	auxpi "github.com/auxpi/auxpiAll"
	"github.com/auxpi/bootstrap"
	auxpiLog "github.com/auxpi/log"
	"github.com/auxpi/models"
)

type Github struct {
	FileLimit []string
	MaxSize   int
}

func (g *Github) Upload(image *ImageParam) (ImageReturn, error) {

	var githubAccount = auxpi.GithubAccount{}
	err := githubAccount.UnmarshalJSON([]byte(models.GetOption("github", "conf")))
	if err != nil {
		auxpiLog.SetAWarningLog("SERVER", err)
	}

	user := githubAccount.Owner
	email := "img-pi@admin.com"
	if githubAccount.AccessToken == "" {
		err := errors.New("github:access_token is none")
		return ImageReturn{}, err
	}
	if githubAccount.Owner == "" {
		err := errors.New("github:owner is none")
		return ImageReturn{}, err
	}
	if githubAccount.Repo == "" {
		err := errors.New("github:repo is none")
		return ImageReturn{}, err
	}

	if githubAccount.Email != "" {
		email = githubAccount.Email
	}

	name := bootstrap.GenerateImageName(image.Name)
	url := g.formatUrl(githubAccount.Owner, githubAccount.Repo, name, githubAccount.AccessToken)
	//上传逻辑
	benc := base64.StdEncoding
	bodyBuf := &bytes.Buffer{}
	content := benc.EncodeToString(*image.Content)
	msg := "commit by auxpi -> file original name: " + image.Name

	j := auxpi.GithubRequest{
		Message: msg,
		Content: content,
	}
	j.Committer.Name = user
	j.Committer.Email = email
	//Make body
	body, err := j.MarshalJSON()
	if err != nil {
		return ImageReturn{}, err
	}
	bodyBuf.Write(body)

	req, err := http.NewRequest("PUT", url, bodyBuf)

	if err != nil {
		return ImageReturn{}, err
	}
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return ImageReturn{}, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return ImageReturn{}, err
	}

	var githubResp = auxpi.GithubResp{}
	var githubMsg = auxpi.GithubMsg{}
	err = githubResp.UnmarshalJSON(respBody)
	if err != nil {
		var maps = make(map[string]string)
		err1 := json.Unmarshal(respBody, &maps)
		if err1 != nil {
			return ImageReturn{}, err1
		}
		err = errors.New(maps["message"])
		return ImageReturn{}, err
	}
	//recursive
	if githubResp.Content.DownloadURL == "" {
		err = githubMsg.UnmarshalJSON(respBody)
		if githubMsg.DocumentationURL == "https://developer.github.com/v3/repos/contents/#update-a-file" {
			if githubResp.Content.DownloadURL != "" {
				return ImageReturn{
					Url:    githubResp.Content.DownloadURL,
					Delete: githubResp.Commit.Sha,
					ID:     18,
				}, nil
			}
			fmt.Println("xxxxxxxxxxxxxxxxxxxxxxx--->>>>>>")
			return g.Upload(image)
		}
	}

	if githubAccount.Proxy.Status {
		githubResp.Content.DownloadURL = githubAccount.Proxy.Node + githubResp.Content.DownloadURL
	}
	return ImageReturn{
		Url:    githubResp.Content.DownloadURL,
		Delete: githubResp.Commit.Sha,
		ID:     18,
	}, nil

}

//https://api.github.com/repos/aimerforreimu/blog/contents/tests/123.txt?
// access_token=37ffeac50a407259960bc38d526fde8f8db8fd41
func (g *Github) formatUrl(owner, repo, name, token string) string {
	return "https://api.github.com/repos/" + owner + "/" + repo + "/contents/" + name + "?access_token=" + token
}
