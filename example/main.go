package main

import (
	"fmt"
	"github.com/bobacsmall/gocaptcha"
	"html/template"
	"log"
	"net/http"
)

// 自定义宽和高
const (
	dx = 150
	dy = 100
)

func main() {

	err := gocaptcha.ReadFonts("fonts", ".ttf")
	if err != nil {
		fmt.Println(err)
		return
	}
	http.HandleFunc("/", Index)
	http.HandleFunc("/get/", Get)
	fmt.Println("服务已启动...")
	// 可以自行修改端口
	err = http.ListenAndServe(":8800", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// 首页
func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tpl/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

// 生成图片流
func Get(w http.ResponseWriter, r *http.Request) {

	captchaImage, err := gocaptcha.NewCaptchaImage(dx, dy, gocaptcha.RandLightColor())

	captchaImage.DrawNoise(gocaptcha.CaptchaComplexLower)

	captchaImage.DrawTextNoise(gocaptcha.CaptchaComplexLower)

	captchaImage.DrawText(gocaptcha.RandText(4))
	//captchaImage.Drawline(3);
	captchaImage.DrawBorder(gocaptcha.ColorToRGB(0x17A7A7A))
	captchaImage.DrawSineLine()
	//captchaImage.DrawHollowLine()
	if err != nil {
		fmt.Println(err)
	}

	captchaImage.SaveImage(w, gocaptcha.ImageFormatJpeg)
}
