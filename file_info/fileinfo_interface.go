package file_info

import (
	"file_server/common"
	"file_server/config"
	"file_server/storage"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

//文件详情: /file/info/
func InfoFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		err_msg := "ERROR: method must be GET"
		w.Write(common.FormatResponse("", -1, err_msg))
		return
	}

	fmt.Println("start infoFIle")

	m, _ := url.ParseQuery(r.URL.RawQuery)

	fmt.Println(m)

	if len(m["file_id"]) == 0 {
		err_msg := "ERROR: parameters lack file_id"
		w.Write(common.FormatResponse("", -1, err_msg))
		return
	}

	file_id, err := strconv.Atoi(m["file_id"][0])
	if err != nil {
		err_msg := "ERROR: file_id not valid"
		w.Write(common.FormatResponse("", -1, err_msg))
		return
	}

	info_file := storage.GetFileInfo(file_id)

	fmt.Println(info_file)

	data_type := ""
	if len(m["data_type"]) > 0 {
		data_type = m["data_type"][0]
	}
	filename := "file"
	if len(m["filename"]) > 0 {
		filename = m["filename"][0]
	}

	fmt.Println("disk")
	if data_type == "file" {
		common.ResponseDiskFile(w, r, info_file["disk_path"], data_type, filename)
		return
	}
	//file_size := GetFileSize(info_file["disk_path"])

	data_fmt := `{"file_id":%d, "download_path": "%s", "static_url":"%s", "absolute_path":"%s", "hostname":"%s", "file_name": "%s", "file_size": %s}`
	data := fmt.Sprintf(data_fmt, file_id, info_file["download_path"], config.STATIC_URL, info_file["absolute_path"], info_file["hostname"], info_file["file_name"], info_file["file_size"])

	fmt.Println(data)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(common.FormatResponse(data, 0, "SUCCESS"))
}
