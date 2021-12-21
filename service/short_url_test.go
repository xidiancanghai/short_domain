package service

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	s := NewShortUrlService()
	fmt.Println(s.AddLongUrl("https://jingyan.baidu.com/article/1974b289443178f4b1f77433.html"))
	fmt.Println(s.AddLongUrl("https://jingyan.baidu.com/article/1974b289443178f4b1f77433.html"))
	fmt.Println(s.GetByShortUrl("00000001"))
}
