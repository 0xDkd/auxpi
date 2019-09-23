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
	"encoding/xml"
	"io"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/auxpi/auxpiAll"
	"github.com/auxpi/models"
	"gopkg.in/masci/flickr.v2"
)

type Flickr struct {
	FileLimit []string
	MaxSize   int
}

var flickrAccount = auxpi.FlickrAccount{}
var err = flickrAccount.UnmarshalJSON([]byte(models.GetOption("flickr", "conf")))

func (s *Flickr) Upload(image *ImageParam) (ImageReturn, error) {
	client := flickrGetOauth()
	re := bytes.NewReader(*image.Content)

	resp, err := flickr.UploadReader(client, re, image.Name, nil)
	if err != nil {
		logs.Alert("Flickr ERROR :", err)
		return ImageReturn{}, err
	}
	return ImageReturn{
			Url: flickrGetPicUrl(resp.ID),
			ID:  5,
		},
		nil
}

func flickrGetOauth() *flickr.FlickrClient {

	apiKey := flickrAccount.Api_key
	apiSecret := flickrAccount.Api_secret
	client := flickr.NewFlickrClient(apiKey, apiSecret)
	client.OAuthToken = flickrAccount.Oauth_token
	client.OAuthTokenSecret = flickrAccount.Oauth_token_secret
	client.Id = flickrAccount.Id
	client.OAuthSign()
	return client
}

func flickrGetPicUrl(id string) string {
	client := flickrGetOauth()
	client.Init()
	client.Args.Set("method", "flickr.photos.getInfo")
	client.Args.Set("photo_id", id)
	client.OAuthSign()
	response := &flickr.BasicResponse{}
	err := flickr.DoGet(client, response)

	if err != nil {
		logs.Alert("Flickr Error:", err)
		return ""
	}
	v := auxpi.FlickrGetPicResp{}
	xml.Unmarshal([]byte(response.Extra), &v)
	if v.Originalformat != "gif" && flickrAccount.DefaultSize != "o" {
		picUrl := "https://" + "farm" + v.Farm + ".staticflickr.com/" + v.Server + "/" + v.Id + "_" + v.Secret + "_" + flickrAccount.DefaultSize + ".jpg"
		return picUrl
	} else {
		picUrl := "https://" + "farm" + v.Farm + ".staticflickr.com/" + v.Server + "/" + v.Id + "_o-" + v.Originalsecret + "_o." + v.Originalformat
		return picUrl
	}
}

func (s *Flickr) UploadToFlickr(file io.Reader, fileName string) string {
	if !flickrAccount.Status {
		return ""
	}
	client := flickrGetOauth()
	beego.Alert(file)
	beego.Alert(client)
	resp, err := flickr.UploadReader(client, file, fileName, nil)
	if err != nil {
		logs.Alert("Flickr ERROR :", err)
	}
	return flickrGetPicUrl(resp.ID)
}
