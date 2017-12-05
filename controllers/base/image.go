package base

import (
	"github.com/astaxie/beego"
	"mime/multipart"
	"crypto/md5"
	"time"
	"encoding/hex"
	"errors"
	path2 "path"
)

type ImageController struct {
	beego.Controller
}

var GOODS_PIC_SAVE_PATH string = "static/img/goods/"

//picType 文件目录 static/img/PICTYPE/
func (this *ImageController) SavePic(picType string,picFile multipart.File, fileHeader *multipart.FileHeader)(string,error){
	md5Obj := md5.New()
	md5Obj.Write([]byte(fileHeader.Filename+time.Now().Format("2006-01-02 15:04:05")))
	newName := hex.EncodeToString(md5Obj.Sum(nil))
	path := GOODS_PIC_SAVE_PATH+newName+path2.Ext(fileHeader.Filename)
	picFile.Close()
	if fileHeader.Size > 2<<21 {
		return "",errors.New("图片过大，请重试")
	}else{
		if err := this.SaveToFile("pic_url",path);err==nil{
			return "/"+path, nil
		}else{
			return "",nil
		}
	}
}