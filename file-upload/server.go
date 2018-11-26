package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/upload", handleUploadCSV)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleUploadCSV(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// リクエストの情報を出力する
	requestDump, err := httputil.DumpRequest(r, true)
	handleError(err)
	log.Println(string(requestDump))

	// "file"というフィールド名に一致する最初のファイルが返却される
	// マルチパートフォームのデータはパースされていない場合ここでパースされる
	formFile, _, err := r.FormFile("file")
	handleError(err)
	defer formFile.Close()

	// データを保存するファイルを開く
	filename := fmt.Sprintf("uploaded_%d.txt", time.Now().UnixNano())
	saveFile, err := os.Create(filename)
	handleError(err)
	defer saveFile.Close()

	// ファイルにデータを書き込む
	_, err = io.Copy(saveFile, formFile)
	handleError(err)

	w.WriteHeader(http.StatusCreated)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
