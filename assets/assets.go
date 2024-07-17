package assets

import (
	"embed"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"path/filepath"
)

//go:embed browser
var browser embed.FS

func HttpHandler(req *http.Request, res http.ResponseWriter) {
	var err error
	vfs, err := fs.Sub(browser, "browser")
	if err != nil {
		panic(err)
	}

	path := filepath.ToSlash(filepath.Join(".", req.URL.Path))

	entry := []string{"index.html"}
	entry = append([]string{path}, entry...)

	var b []byte
	for _, s := range entry {
		b, err = fs.ReadFile(vfs, s)
		if err != nil {
			continue
		}
		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(s)))
		if _, err := res.Write(b); err != nil {
			log.Println(err)
		}
		return
	}
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
}
