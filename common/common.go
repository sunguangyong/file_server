package common

import (
	"file_server/config"
	"file_server/storage"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	ra = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GenRandomInt() int {
	return 100000 + ra.Intn(900000)
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CheckOrMakedir(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.MkdirAll(path, 0777)
	}
	return nil
}

func CurrentDirPath(app_name string, type_name string) string {
	return config.STATIC_DIR + "/" + app_name + "/" + type_name
}

func DoloadFilePath(app_name string, type_name string, file_name string) string {
	return "/download/" + app_name + "/" + type_name + "/" + file_name
}

func AbsoluteFilePath(app_name string, type_name string, file_name string) string {
	return config.STATIC_URL + DoloadFilePath(app_name, type_name, file_name)
}

func GenFileName(orig_file_name string) (file_name string) {
	items := strings.Split(orig_file_name, ".")
	items[len(items)-2] = fmt.Sprintf("%s_%d_%d", strings.Replace(items[len(items)-2], " ", "", -1), time.Now().Unix(), GenRandomInt())

	file_name = strings.Join(items, ".")
	file_name = strings.Replace(file_name, "(", "_", -1) // 替换所有（ 防止 bash: 未预期的符号 `(' 附近有语法错误

	file_name = strings.Replace(file_name, ")", "_", -1)
	return
}

func SaveFile(local_path string, fh *multipart.FileHeader) (err error) {
	localfd, err := os.OpenFile(local_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer localfd.Close()

	post_file, err := fh.Open()
	storage.CheckErr(err)
	defer post_file.Close()

	io.Copy(localfd, post_file)
	return nil
}

func ResponseDiskFile(w http.ResponseWriter, r *http.Request, disk_path string, data_type string, filename string) {
	if !FileExist(disk_path) {
		err_msg := "ERROR: Image Not Found"
		w.Write(FormatResponse("", -1, err_msg))
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if data_type == "file" {
		w.Header().Set("Content-Type", "octet-stream")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	} else if data_type == "mp3" {
		w.Header().Set("Content-Type", "audio/mp3") // 获取mp3
	} else if data_type == "wav" {
		w.Header().Set("Content-Type", "audio/wav")
	}
	http.ServeFile(w, r, disk_path)
}

func FormatResponse(data string, code int, msg string) []byte {
	return []byte(fmt.Sprintf(`{"data":%s, "code":"%d", "msg":"%s"}`, data, code, msg))
}
