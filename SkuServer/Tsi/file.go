package Tsi

import (
	"github.com/leesper/holmes"
	"fmt"
	"time"
	"os"
	"io"
	"sku/base"
	"io/ioutil"
	"sku/base/utils/encrypt"
)

func GetSavePath(saveType string) string {
	savePath := base.PATH_RUNTIME
	if saveType == TSI_RUN_TYPE_TEST_PRE {
		savePath = savePath + "/" + FILE_TSI
	} else {
		savePath = savePath + "/" + FILE_TSI_TEST
	}

	return savePath
}

func MoveFileToHistory() {
	fileTime := time.Now().Format("15-04-05")
	MoveFileToHistoryByType(TSI_RUN_TYPE_TEST_PRE,fileTime)
	MoveFileToHistoryByType(TSI_RUN_TYPE_TEST,fileTime)
}

func MoveFileToHistoryByType(fileType string, fileTime string) {
	//判断文件是否存在
	fileInfo, err := os.Stat(GetSavePath(fileType))
	if err == nil {
		fileName := fileInfo.Name()

		fileDay := time.Now().Format("20060102")

		newUri := base.PATH_HISTORY + "/" + fmt.Sprintf("%v/%v/", fileDay,fileTime) + fileName

		//创建目录
		newPath := string([]rune(newUri)[:len(newUri)-len(fileName)])
		os.MkdirAll(newPath, 0666)

		//迁移文件
		err = os.Rename(GetSavePath(fileType), newUri)
		if err != nil {
			holmes.Debugf("文件迁移失败：%v\n",err.Error())
		}
	}

}

type FileInfo struct {
	Name	string  `json:"name"`
	Size 	int64	`json:"size"`
	Md5 	string 	`json:"md5"`
}

func GetFileInfo(fileName string) (f FileInfo,err error) {
	filePath := GetSavePath(fileName)

	info, err := os.Stat(filePath)
	if err != nil {
		return f, err
	} else {
		f.Name = info.Name()
		f.Size = info.Size()

		md5,err := encrypt.Md5File(filePath)
		if err != nil {
			holmes.Errorf("获取文件信息Md5错误：%v\n",err.Error())
		}

		f.Md5 = md5
		return f,nil
	}
}

func GetFile(fileName string) (content []byte,err error) {
	filePath := GetSavePath(fileName)

	_, err = os.Stat(filePath)
	if err != nil {
		return content, err
	} else {
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			return content, err
		}

		return content,nil
	}
}

func SaveFile(saveType string,pm25 int) {
	//确定存储路径及文件名
	savePath := GetSavePath(saveType)

	var f *os.File
	defer f.Close()

	var err error

	//判断文件是否存在
	_, err = os.Stat(savePath)
	if err != nil {
		f, err = os.Create(savePath)  //创建文件
		if err != nil {
			holmes.Errorln(err.Error())
		}
	} else {
		f, err = os.OpenFile(savePath, os.O_APPEND, 0777)  //打开文件
		if err != nil {
			holmes.Errorln(err.Error())
		}
	}


	//拼装要存的数据
	saveStr := fmt.Sprintf("%d:%d\n",time.Now().Unix(),pm25)

	_, err = io.WriteString(f, saveStr) //写入文件(字符串)
	if err != nil {
		holmes.Errorln(err.Error())
	}
}