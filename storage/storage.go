package storage


import (
    "database/sql"
    "file_server/config"
    "os"
    "fmt"
    "strconv"
    _ "github.com/go-sql-driver/mysql"
)


var  db * sql.DB 


func init() {
    db = ConnectMysql()
}


func CheckErr(err error) {
    if err != nil {
        fmt.Println("err: ", err)
        panic(err)
    }
}


func ConnectMysql() (* sql.DB) {
    config.Init()
    db, err := sql.Open("mysql", config.DB_ADDRESS)
    CheckErr(err)
    return db
}


func SaveFileInfo(file_name string, hostname, disk_path string, download_path string, absolute_path string, app_name string, file_type string, storage_ip string, file_size int64) (file_id int64) {
    stmt, err := db.Prepare("Insert into info_file_storage(file_name, hostname, disk_path, download_path, absolute_path, app_name, file_type, storage_ip, file_size) Values(?, ?, ?, ?, ?, ?, ?, ?, ?)") 
    CheckErr(err)

    res, err_stmt := stmt.Exec(file_name, hostname, disk_path, download_path, absolute_path, app_name, file_type, storage_ip, file_size)
    CheckErr(err_stmt)

    file_id, err = res.LastInsertId()
    CheckErr(err)
    return
} 


func GetFileInfo(file_id int) map[string]string {
    lang := fmt.Sprintf("SELECT file_name, disk_path, download_path, absolute_path, hostname, storage_ip, file_size FROM info_file_storage WHERE file_id=%d" , file_id)
    rows, err := db.Query(lang)
    CheckErr(err)

    var file_name, disk_path, download_path, absolute_path, hostname, storage_ip string
    var file_size int64

    rows.Next()
    err = rows.Scan(&file_name, &disk_path, &download_path, &absolute_path, &hostname, &storage_ip, &file_size)
    CheckErr(err)
    if file_size == 0 {
        file_size = GetFileSize(disk_path)
    }
    file_size /= 1024
    size := strconv.FormatInt(file_size, 10) 
    result := map[string]string{"file_name":file_name, "disk_path":disk_path, "download_path":download_path, "absolute_path":absolute_path, "hostname":hostname, "storage_ip": storage_ip, "file_size": size}
    return result
}

func GetFileSize(path string) int64 {
    file_info, err := os.Stat(path)
    if err != nil {
        return 0	
    }
    file_len := file_info.Size() 
    return file_len
}
