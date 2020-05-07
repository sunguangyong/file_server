package main

import (
    "file_server/dowload"
    "file_server/file_info"
    "file_server/upload"
    "github.com/gorilla/mux"
    "file_server/config"
    "net/http"
    "fmt"
    _ "log"
)


func newRouter() http.Handler {
    r := mux.NewRouter()
    r.HandleFunc("/", upload.Index)
    r.HandleFunc("/upload/{app_name}/{type_name}", upload.Upload)   // 上传文件
    r.HandleFunc("/download/{app_name}/{type_name}/{file_name}", dowload.Download) // 下载文件
    r.HandleFunc("/file/info", file_info.InfoFile)   // 查看文件详情
    r.HandleFunc("/download_compress", dowload.DownloadZip)
    return r
}


func main() {
    fmt.Println("START LISTEN ON ", config.LISEN_HOST)
    err := http.ListenAndServe(config.LISEN_HOST, newRouter())
    fmt.Println("END RUN...", err)
}
