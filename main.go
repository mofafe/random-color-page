package main

import (
	"crypto/rand"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		backgroundcolor, _ := makeRandomStr()
		color, _ := makeRandomStr()
		c.HTML(http.StatusOK, "random-color.html", gin.H{
			"backgroundcolor": backgroundcolor,
			"color":           color,
		})
	})

	r.Run()
}

func makeRandomStr() (string, error) {
	const letters = "abcdef0123456789"

	// 乱数を生成
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
