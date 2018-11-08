package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/textproto"
	"os"
	"os/exec"
	"strconv"
	"time"

	"gocv.io/x/gocv"
)

type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	// contains filtered or unexported fields
}

func init() {
	loadConfig()
}

func index(writer http.ResponseWriter, request *http.Request) {
	// threads := data.GetList()
	generateHTML(writer, nil, "layout", "public.navbar", "index")
}

// handle upload file
func upload(w http.ResponseWriter, r *http.Request) {
	// threads := data.GetList()
	// generateHTML(writer, nil, "layout", "public.navbar", "index")
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		// fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./sample/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		generateHTML(w, handler.Filename, "layout", "public.navbar", "showimage")
	}

}

func objectdetection(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	os.Remove("sample/out.jpg")
	name := r.FormValue("imagename")
	fmt.Printf("image name %s \n", name)
	// strcmd := fmt.Sprintf("../../handle_picture.py -i sample/%s -o sample/out.jpg", name)
	// fmt.Printf(strcmd)
	// cmd := exec.Command("python3", strcmd)
	strsource := fmt.Sprintf("sample/%s", name)
	cmd := exec.Command("python3", "../../handle_picture.py", "-i", strsource, "-o", "sample/out.jpg")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("cmd.Run() fail with status %s \n", err)
	}

	generateHTML(w, "out.jpg", "layout", "public.navbar", "index")
}

func imageprocess(w http.ResponseWriter, r *http.Request) {
	p("imageprocess")
	os.Remove("sample/out.jpg")
	name := r.FormValue("imagename")
	strsource := fmt.Sprintf("sample/%s", name)

	// using gocv
	src_img := gocv.IMRead(strsource, gocv.IMReadGrayScale)

	// using threshold
	threshold_img := gocv.NewMat()
	gocv.Threshold(src_img, &threshold_img, 100, 255, gocv.ThresholdType)
	gocv.IMWrite("sample/out.jpg", threshold_img)

	generateHTML(w, "out.jpg", "layout", "public.navbar", "index")
}
