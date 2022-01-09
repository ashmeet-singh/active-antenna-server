package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(os.Getenv("AA_DIR_WEBSITE"))))

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		file_path := path.Join(os.Getenv("AA_DIR_STORAGE"), r.Header.Get("X-FILE_PATH-X"))
		file, err1 := os.Open(file_path)
		if err1 != nil {
			panic("Unable to open file")
		}

		file_stat, err2 := os.Stat(file_path)
		if err2 != nil {
			panic("Unable to get stat")
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", strconv.FormatInt(file_stat.Size(), 10))
		w.Header().Set("Content-Disposition", "attachment; filename="+path.Base(file_path))

		io.Copy(w, file)
		fmt.Println("File sent")
	})

	http.ListenAndServe(":8080", nil)
}
