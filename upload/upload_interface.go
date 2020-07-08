package upload

//package upload

import (
	"file_server/common"
	"file_server/config"
	"file_server/storage"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	_ "strings"
	"time"
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

func save(app_name string, ftype_name string, orig_file_name string, fh *multipart.FileHeader) string {
	size := fh.Size
	var file_uploda storage.FileUpload
	file_name := common.GenFileName(orig_file_name)
	download_path := common.DoloadFilePath(app_name, ftype_name, url.QueryEscape(file_name)) // 将字符地址转为安全地址  防止特殊符号在服务器端无法获得正确的参数值
	absolute_path := common.AbsoluteFilePath(app_name, ftype_name, url.QueryEscape(file_name))

	disk_dir := common.CurrentDirPath(app_name, ftype_name)
	common.CheckOrMakedir(disk_dir)

	disk_path := fmt.Sprintf("%s/%s", disk_dir, file_name)
	err := common.SaveFile(disk_path, fh)
	storage.CheckErr(err)

	file_uploda.FileName = file_name
	file_uploda.DownloadPath = config.STATIC_URL + download_path
	file_uploda.AbsolutePath = absolute_path
	file_uploda.DiskPath = disk_path
	file_uploda.FileSize = size
	file_uploda.FileType = ftype_name
	file_uploda.AppName = app_name
	file_uploda.HostName, _ = os.Hostname()
	file_uploda.StorageIp = config.STORAGE_IP
	file_uploda.RTime = time.Now()
	file_id := storage.FileUploadInsert(file_uploda)
	return fmt.Sprintf(`{"file_id":%d, "download_path": "%s", "static_url":"%s", "absolute_path":"%s", "disk_path":"%s"}`, file_id, file_uploda.DownloadPath, config.STATIC_URL, absolute_path, disk_path)
}

//文件上传: /upload/<app_name>/<type_name>
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
		record := save(app_name, ftype_name, orig_file_name, fh)
		records[i] = record
	}

	data := "[" + strings.Join(records, ",") + "]"
	if len(records) == 1 {
		data = data[1 : len(data)-1]
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(common.FormatResponse(data, 0, "SUCCESS"))
}
