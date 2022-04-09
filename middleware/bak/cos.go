package bak

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Cos()  {
	u, _ := url.Parse("https://examplebucket-1250000000.cos.COS_REGION.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "SECRETID",
			SecretKey: "SECRETKEY",
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	name := "test/objectPut.go"
	// 1.通过字符串上传对象
	f := strings.NewReader("test")

	_, err := c.Object.Put(context.Background(), name, f, nil)
	if err != nil {
		panic(err)
	}
	// 2.通过本地文件上传对象
	_, err = c.Object.PutFromFile(context.Background(), name, "../test", nil)
	if err != nil {
		panic(err)
	}
	// 3.通过文件流上传对象
	fd, err := os.Open("./test")
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	_, err = c.Object.Put(context.Background(), name, fd, nil)
	if err != nil {
		panic(err)
	}
}
