package storage

//
//import (
//	"fmt"
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/postgres"
//	_ "github.com/lib/pq"
//	"github.com/wonderivan/logger"
//	"strconv"
//	"time"
//)
//
//var (
//	Db *gorm.DB
//)
//
//type FileUpload struct {
//	Id int64 `gorm:"primary_key"`
//	// Createtime              time.Time
//	FileName     string
//	HostName     string
//	DiskPath     string
//	DownloadPath string
//	AbsolutePath string
//	AppName      string
//	FileType     string
//	StorageIp    string
//	FileSize     int64
//	RTime        time.Time
//}
//
//func init() {
//	PgOrmOpen("10.0.3.122", "aicalldev", "Yplsec.com", "aicalldb", 5432)
//	//PgOrmOpen("49.235.38.211","aicalldev","Yplsec.com","aicalldb",5432)
//
//}
//
//func PgOrmOpen(host, user, password, dbname string, port int) {
//	var err error
//	pgInfo := fmt.Sprintf("host=%v port=%v user=%v  dbname=%v password=%v sslmode=disable", host, port, user, dbname, password)
//	Db, err = gorm.Open("postgres", pgInfo)
//	Db.SingularTable(true) // 全局禁用复数表名
//	Db.LogMode(true)       // 显示sql语句调试
//
//	if err != nil {
//		panic(err)
//	}
//}
//
//func FileUploadInsert(file_upload FileUpload) int64 {
//	er := Db.Create(&file_upload)
//	flag := Db.NewRecord(file_upload) // 主键是否为空
//	if er.Error != nil || flag {
//		logger.Error("insert file_upload err primary key is existing", er.Error)
//		return 0
//	}
//	return file_upload.Id
//}
//
//func FileUploadUpdate(file_upload FileUpload, data map[string]interface{}) bool {
//	Db.Model(&file_upload).Updates(data)
//	return true
//}
//
//func FileUploadFind(data map[string]interface{}) (file_upload_arry []FileUpload) {
//	Db.Where(data).Find(&file_upload_arry)
//	return file_upload_arry
//}
//
//func SavePgFileInfo(file_name string, hostname, disk_path string, download_path string, absolute_path string, app_name string, file_type string, storage_ip string, file_size int64) (file_id int64) {
//	stmt, err := db.Prepare("Insert into info_file_storage(file_name, hostname, disk_path, download_path, absolute_path, app_name, file_type, storage_ip, file_size) Values(?, ?, ?, ?, ?, ?, ?, ?, ?)")
//	CheckErr(err)
//
//	res, err_stmt := stmt.Exec(file_name, hostname, disk_path, download_path, absolute_path, app_name, file_type, storage_ip, file_size)
//	CheckErr(err_stmt)
//
//	file_id, err = res.LastInsertId()
//	CheckErr(err)
//	return
//}
//
//func GetPgFileInfo(file_id int) map[string]string {
//	lang := fmt.Sprintf("SELECT file_name, disk_path, download_path, absolute_path, hostname, storage_ip, file_size FROM info_file_storage WHERE file_id=%d", file_id)
//	rows, err := db.Query(lang)
//	CheckErr(err)
//
//	var file_name, disk_path, download_path, absolute_path, hostname, storage_ip string
//	var file_size int64
//
//	rows.Next()
//	err = rows.Scan(&file_name, &disk_path, &download_path, &absolute_path, &hostname, &storage_ip, &file_size)
//	CheckErr(err)
//	if file_size == 0 {
//		file_size = GetFileSize(disk_path)
//	}
//	file_size /= 1024
//	size := strconv.FormatInt(file_size, 10)
//	result := map[string]string{"file_name": file_name, "disk_path": disk_path, "download_path": download_path, "absolute_path": absolute_path, "hostname": hostname, "storage_ip": storage_ip, "file_size": size}
//	return result
//}
