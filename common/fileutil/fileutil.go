package fileutil

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/file_server/common/constant"
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

func CurrentDirPath(appName string, typeName string) string {
	return constant.STATIC_DIR + "/" + appName + "/" + typeName
}

func DoloadFilePath(appName string, typeName string, fileName string) string {
	return "/download/" + appName + "/" + typeName + "/" + fileName
}

func AbsoluteFilePath(appName string, typeName string, fileName string) string {
	return constant.STATIC_URL + DoloadFilePath(appName, typeName, fileName)
}

func GenFileName(origFileName string) (fileName string) {
	items := strings.Split(origFileName, ".")
	items[len(items)-2] = fmt.Sprintf("%s_%d_%d", strings.Replace(items[len(items)-2], " ", "", -1), time.Now().Unix(), GenRandomInt())

	fileName = strings.Join(items, ".")
	fileName = strings.Replace(fileName, "(", "_", -1) // 替换所有（ 防止 bash: 未预期的符号 `(' 附近有语法错误

	fileName = strings.Replace(fileName, ")", "_", -1)
	return
}

func SaveFile(localPath string, fh *multipart.FileHeader) (err error) {
	localfd, err := os.OpenFile(localPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer localfd.Close()

	postFile, err := fh.Open()

	if err != nil {
		return err
	}

	defer postFile.Close()

	io.Copy(localfd, postFile)
	return nil
}

func ResponseDiskFile(w http.ResponseWriter, r *http.Request, diskPath string, dataType string, filename string) {
	if !FileExist(diskPath) {
		err_msg := "ERROR: Image Not Found"
		w.Write(FormatResponse("", -1, err_msg))
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if dataType == "file" {
		w.Header().Set("Content-Type", "octet-stream")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	} else if dataType == "mp3" {
		w.Header().Set("Content-Type", "audio/mp3") // 获取mp3
	} else if dataType == "wav" {
		w.Header().Set("Content-Type", "audio/wav")
	}
	http.ServeFile(w, r, diskPath)
}

func FormatResponse(data string, code int, msg string) []byte {
	return []byte(fmt.Sprintf(`{"data":%s, "code":"%d", "msg":"%s"}`, data, code, msg))
}
