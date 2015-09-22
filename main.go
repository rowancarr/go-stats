package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"syscall"
)

//Stats ...
type Stats struct {
	Backups int    `json:"backups"`
	All     uint64 `json:"all"`
	Used    uint64 `json:"used"`
	Free    uint64 `json:"free"`
}

//SystemStats ...
func SystemStats(path string) (stat Stats) {
	backupdir := os.Getenv("EP_BACKUP_DIR")
	if backupdir == "" {
		backupdir = "/backups"
	}

	files, _ := ioutil.ReadDir(backupdir)
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	stat.All = fs.Blocks * uint64(fs.Bsize)
	stat.Free = fs.Bfree * uint64(fs.Bsize)
	stat.Used = stat.All - stat.Free
	stat.Backups = len(files)
	return
}

//const ..
const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", ShowStats)
	http.ListenAndServe(":"+port, nil)
}

//ShowStats ...
func ShowStats(w http.ResponseWriter, r *http.Request) {
	out := SystemStats("/")
	js, err := json.Marshal(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
