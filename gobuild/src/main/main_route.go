package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"image"
	"io"
	"net/http"
	"net/textproto"
	"os"
	"os/exec"
	"strconv"
	"time"

	"gocv.io/x/gocv"
)

// =========================================
type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	// contains filtered or unexported fields
}

type PictureHandler struct {
	SrcFilename string
	DstFilename string
}

// =========================================
func init() {
	loadConfig()
}

func index(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "layout", "public.navbar", "index")
}

// handle upload file
func upload(w http.ResponseWriter, r *http.Request) {
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

		picthandler := PictureHandler{
			SrcFilename: handler.Filename,
			// DstFilename: nil,
		}

		generateHTML(w, picthandler, "layout", "public.navbar", "showimage")
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
	method := r.FormValue("method")
	strsource := fmt.Sprintf("sample/%s", name)

	p("name %v, method %v, strsource %v", name, method, strsource)

	// using gocv
	src_img := gocv.NewMat()
	defer src_img.Close()

	cvt_img := gocv.NewMat()
	defer cvt_img.Close()

	threshold_img := gocv.NewMat()
	defer threshold_img.Close()

	blur_img := gocv.NewMat()
	defer blur_img.Close()

	src_img = gocv.IMRead(strsource, gocv.IMReadColor)
	gocv.CvtColor(src_img, &cvt_img, gocv.ColorBGRToGray)
	gocv.IMWrite("sample/cvt.jpg", cvt_img)

	if method == "Threshold" {
		threshold_param, err := strconv.Atoi(r.FormValue("threshold_param"))
		if err != nil {
			p("err ", err)
		}
		gocv.Threshold(cvt_img, &threshold_img, float32(threshold_param), 255, gocv.ThresholdBinary)
		gocv.IMWrite("sample/out.jpg", threshold_img)
	} else if method == "Blur" {
		blur_param, err := strconv.Atoi(r.FormValue("blur_param"))
		if err != nil {
			p("err ", err)
		}
		gocv.GaussianBlur(cvt_img, &blur_img, image.Pt(blur_param, blur_param), 0, 0, gocv.BorderDefault)
		gocv.IMWrite("sample/out.jpg", blur_img)
	} else {
		p("Cannot find method")
	}

	picthandler := PictureHandler{
		SrcFilename: name,
		DstFilename: "out.jpg",
	}
	p("picthandler src %v, dst %v", picthandler.SrcFilename, picthandler.DstFilename)

	generateHTML(w, picthandler, "layout", "public.navbar", "showimage")
}
