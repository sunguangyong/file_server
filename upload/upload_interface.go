package upload

//package upload

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	_ "strings"
	"time"

	"github.com/file_server/common"
	"github.com/file_server/model"
	"github.com/gorilla/mux"
)

const uploadHTML = `
    <html>
      <head>
        <title>选择文件</title>
      </head>
      <body>
        <form enctype="multipart/form-data" action="/upload/jupiter/img" method="post">
          <input type="file" name="file"  />
          <input type="submit" value="提交" />
        </form>
      </body>
    </html>`

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(uploadHTML))
}

func save(app_name string, ftype_name string, orig_file_name string, fh *multipart.FileHeader) (fileInfo string,
	err error) {
	ctx := context.Background()
	size := fh.Size
	//var file_uploda storage.FileUpload
	file_name := common.GenFileName(orig_file_name)
	download_path := common.DoloadFilePath(app_name, ftype_name, url.QueryEscape(file_name)) // 将字符地址转为安全地址  防止特殊符号在服务器端无法获得正确的参数值
	absolute_path := common.AbsoluteFilePath(app_name, ftype_name, url.QueryEscape(file_name))

	disk_dir := common.CurrentDirPath(app_name, ftype_name)
	common.CheckOrMakedir(disk_dir)

	disk_path := fmt.Sprintf("%s/%s", disk_dir, file_name)
	err = common.SaveFile(disk_path, fh)

	if err != nil {
		log.Println(err)
	}

	files := &model.Files{}
	files.FileName = file_name
	files.DownloadPath = common.STATIC_URL + download_path
	files.AbsolutePath = absolute_path
	files.DiskPath = disk_path
	files.FileSize = size
	files.FileType = ftype_name
	files.AppName = app_name
	files.HostName, _ = os.Hostname()
	files.StorageIp = common.STORAGE_IP
	files.CreateTime = time.Now()
	files.UpdateTime = time.Now()

	result, err := model.FileConn.Insert(ctx, files)

	if err != nil {
		log.Println(err)
	}

	file_id, err := result.LastInsertId()

	if err != nil {
		log.Println(err)
	}

	fileInfo = fmt.Sprintf(`{"file_id":%d, "download_path": "%s", "static_url":"%s", "absolute_path":"%s", "disk_path":"%s"}`, file_id, files.DownloadPath, common.STATIC_URL, absolute_path, disk_path)

	return fileInfo, err
}

// 文件上传: /upload/<app_name>/<type_name>
func Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "multipart/form-data") //返回数据格式是json
	if r.Method != "POST" {
		err_msg := "ERROR: When Upload File, Method Must Be POST"
		w.Write(common.FormatResponse("", -1, err_msg))
		return
	}

	r.ParseForm()
	r.ParseMultipartForm(32 << 20) //max memory: 32M

	vars := mux.Vars(r)
	app_name, ftype_name := vars["app_name"], vars["type_name"]

	mp := r.MultipartForm
	if mp == nil {
		log.Println("not MultipartForm.")
		err_msg := "ERROR: When Upload File, Form Is Not MultipartForm"
		w.Write(common.FormatResponse("", -1, err_msg))
		return
	}

	fileHeaders, findFile := mp.File["file"]
	if !findFile || len(fileHeaders) == 0 {
		log.Println("file count == 0.")
		err_msg := "ERROR: Not Upload Any File"
		w.Write(common.FormatResponse("", -1, err_msg))
		return
	}

	records := make([]string, len(fileHeaders))
	for i, fh := range fileHeaders {
		orig_file_name := fh.Filename
		record, err := save(app_name, ftype_name, orig_file_name, fh)
		if err != nil {
			continue
		}
		records[i] = record
	}

	data := "[" + strings.Join(records, ",") + "]"
	if len(records) == 1 {
		data = data[1 : len(data)-1]
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(common.FormatResponse(data, 0, "SUCCESS"))
}
