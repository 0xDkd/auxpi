package models

import (
	"auxpi/auxpiAll"
	"auxpi/tools"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/astaxie/beego"
)

var resume chan int

type SyncImage struct {
	Model

	Link     string `json:"link"`
	Path     string `json:"path"`
	Delete   string `json:"delete"`
	External string `gorm:"-"json:"external"`

	ImageID uint  `gorm:"index"json:"image_id"`
	Image   Image `json:"image"`
}

func AddSyncImage(image SyncImage) {
	err := db.Create(&SyncImage{
		Link:    image.Link,
		Path:    image.Path,
		Delete:  image.Delete,
		ImageID: image.ImageID,
	}).Error
	if err != nil {
		beego.Alert("[Sync Image Create Record Error] :", err)
	}
}

func TestSyncImages(list []SyncImage, l int) {
	var msg = tools.Message{}
	var pMsg = tools.Message{}
	var mutex sync.Mutex
	var n = make(chan int)
	var success = make(chan bool)
	var i = 0
	for _, value := range list {
		go func(url string, id uint) {
			res, err := http.Get(url)
			if err == nil {
				lUrl, name, del := tools.LocalStoreInfo(res.Header.Get("Content-Type"), url)
				dst, _ := os.Create(name)
				io.Copy(dst, res.Body)
				value.ImageID = id
				value.Link = lUrl
				value.Delete = del
				mutex.Lock()
				AddSyncImage(value)
				n <- i
				if i == l {
					success <- true
				}
				i++
				mutex.Unlock()
			} else {
				beego.Alert("httperror")
			}

		}(value.External, value.ImageID)
	}

	for {
		select {
		case <-success:
			msg.Code = 200
			msg.Title = "Completed : )"
			msg.Msg = "Completion of Sync Images.Total : " + strconv.Itoa(l+1)
			msg.Data = l + 1
			msg.Status = "success"
			time.Sleep(5e8)
			tools.Send(&msg)
			beego.Alert("[Task Done!]")
			tools.UnLock()
			return
		case num := <-n:
			time.Sleep(1e5)
			pMsg.Data = num
			pMsg.Status = "running"
			tools.Send(&pMsg)
		case <-time.After(60 * time.Second):
			beego.Alert("[Task TimeOut]")
			msg.Status = "timeout"
			msg.Title = "Task timeout : ( "
			msg.Msg = "Task timeout.Please resubmit task or Contact the author"
			tools.Send(&msg)
			return
		}
	}

}

func GetSyncImages(pageNum int, pageSize int, maps interface{}) (images []SyncImage, count int) {
	db.Preload("Image").
		Model(&SyncImage{}).
		Where(maps).
		Count(&count).
		Offset(pageNum).Limit(pageSize).
		Find(&images)

	return
}

func DelSyncImage(ids []int) error {
	var images []SyncImage
	err := db.Model(&SyncImage{}).
		Select("path").
		Where(ids).
		Find(&images).Error

	modelsError(auxpi.ErrorToString(err))
	err = db.Where(ids).Delete(&SyncImage{}).Error
	modelsError(auxpi.ErrorToString(err))

	//删除图片
	for _, value := range images {

		err := os.Remove(value.Path)
		beego.Alert("remove :? )")
		if err != nil {
			AddLog("SYNC_IMAGE_DELETE", auxpi.ErrorToString(err), "SYSTEM", "ERROR")
		}
	}
	return nil
}

func MigrateSyncImage() error {
	if db.HasTable(&SyncImage{}) {
		err := db.DropTable(&SyncImage{}).Error
		err = db.CreateTable(&SyncImage{}).Error
		return err
	} else {
		err := db.CreateTable(&SyncImage{}).Error
		return err
	}
}
