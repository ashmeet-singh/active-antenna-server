package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(os.Getenv("AA_DIR_WEBSITE"))))

	http.HandleFunc("/storage/download/", func(w http.ResponseWriter, r *http.Request) {
		file_path, err := url.PathUnescape(r.URL.Path[len("/download/"):])
		if err != nil {
			fmt.Println("Unable to unescape path")
		}
		file_path = path.Join(os.Getenv("AA_DIR_STORAGE"), file_path)

		file, err1 := os.Open(file_path)
		if err1 != nil {
			fmt.Println("Unable to open file")
		}

		file_stat, err2 := os.Stat(file_path)
		if err2 != nil {
			fmt.Println("Unable to get stat")
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", strconv.FormatInt(file_stat.Size(), 10))
		w.Header().Set("Content-Disposition", "attachment; filename="+path.Base(file_path))

		io.Copy(w, file)
        fmt.Println("File sent: "+file_path)
	})

	http.ListenAndServe(":8080", nil)
}
