package server

import (
	"auxpi/auxpiAll"
	"auxpi/bootstrap"
	"encoding/xml"
	"io"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"gopkg.in/masci/flickr.v2"
)

type Flickr struct {
}


func flickrGetOauth() *flickr.FlickrClient {
	api_key := bootstrap.SiteConfig.SiteUploadWay.FlickrAccount.Api_key
	api_secret := bootstrap.SiteConfig.SiteUploadWay.FlickrAccount.Api_secret
	client := flickr.NewFlickrClient(api_key, api_secret)
	client.OAuthToken = bootstrap.SiteConfig.SiteUploadWay.FlickrAccount.Oauth_token
	client.OAuthTokenSecret = bootstrap.SiteConfig.SiteUploadWay.FlickrAccount.Oauth_token_secret
	client.Id = bootstrap.SiteConfig.SiteUploadWay.FlickrAccount.Id
	client.OAuthSign()
	return client
}

func flickrGetPicInfo(id string) string {
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
	if v.Originalformat != "gif" && bootstrap.SiteConfig.SiteUploadWay.FlickrAccount.DefaultSize != "o" {
		picUrl := "https://" + "farm" + v.Farm + ".staticflickr.com/" + v.Server + "/" + v.Id + "_" + v.Secret + "_" + bootstrap.SiteConfig.SiteUploadWay.FlickrAccount.DefaultSize + ".jpg"
		return picUrl
	} else {
		picUrl := "https://" + "farm" + v.Farm + ".staticflickr.com/" + v.Server + "/" + v.Id + "_o-" + v.Secret + "_o." + v.Originalformat
		return picUrl
	}
}

func (this *Flickr) UploadToFlickr(file io.Reader, fileName string) string {
	client := flickrGetOauth()
	beego.Alert(file)
	beego.Alert(client)
	resp, err := flickr.UploadReader(client, file, fileName, nil)
	if err != nil {
		logs.Alert("Flickr ERROR :", err)
	}
	return flickrGetPicInfo(resp.ID)
}
