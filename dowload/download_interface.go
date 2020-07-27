package dowload

import (
	"archive/zip"
	"file_server/common"
	"file_server/config"
	"file_server/storage"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/wonderivan/logger"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	_ "strings"
)

//文件下载: /download/<app_name>/<type_name>/<file_name>
func Download(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		err_msg := "ERROR: When Download File, Method Must Be GET"
		w.Write(common.FormatResponse("", 0, err_msg))
		return
	}
	logger.Info("load start --------------------------------")
	m, _ := url.ParseQuery(r.URL.RawQuery)
	data_type := ""
	if len(m["data_type"]) > 0 {
		data_type = m["data_type"][0]
	}
	vars := mux.Vars(r)
	app_name := vars["app_name"]
	ftype_name := vars["type_name"]
	file_name := vars["file_name"]
	disk_path := strings.Join([]string{config.STATIC_DIR, app_name, ftype_name, file_name}, "/")
	logger.Info("load    disk_path == ", disk_path)
	common.ResponseDiskFile(w, r, disk_path, data_type, file_name)
}

// 文件下载，传入 file_id, 返回压缩后的 zip 文件
func DownloadZip(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		err_msg := "ERROR: When Download File, Method Must Be GET"
		w.Write(common.FormatResponse("", 0, err_msg))
		return
	}

	m, _ := url.ParseQuery(r.URL.RawQuery)
	if len(m["file_id"]) == 0 {
		err_msg := "ERROR: lack of parameter: file_id"
		w.Write(common.FormatResponse("", 0, err_msg))
		return
	}
	if len(m["file_name"]) == 0 {
		err_msg := "ERROR: lack of parameter: file_name"
		w.Write(common.FormatResponse("", 0, err_msg))
		return
	}
	file_id := m["file_id"][0]
	file_name := m["file_name"][0]

	file_id_list := strings.Split(file_id, ",")

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file_name))

	GetZip(file_id_list, w)
}

func GetZip(file_id_list []string, w io.Writer) {
	zipw := zip.NewWriter(w)
	defer zipw.Close()

	for _, file_id := range file_id_list {
		file_id_int, _ := strconv.Atoi(file_id)
		info_file := storage.GetFileInfo(file_id_int)
		disk_path := info_file["disk_path"]
		f, _ := os.Stat(disk_path)
		zf, _ := ioutil.ReadFile(disk_path)
		zc, _ := zipw.Create(f.Name())
		_, err := zc.Write(zf)
		if err != nil {
			fmt.Println("err: ", err)
		}
	}
}
