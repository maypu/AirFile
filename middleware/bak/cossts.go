package bak

import (
	"fmt"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"os"
	"time"
)

func GetCOSTempToken() string {
	appid := "1251778150"
	bucket := "air-1251778150"
	region := "ap-chengdu"
	c := sts.NewClient(
		os.Getenv("AKIDlMhncBTTCfrmQnukKdfPs40e17vTjiQr"),
		os.Getenv("B5h38wShfMltlLegSxB7mAH7fxyiBOJx"),
		nil,
	)
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region: region,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
						"name/cos:PutObject",
						"name/cos:GetObject",
					},
					Effect: "allow",
					Resource: []string{
						//这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						"qcs::cos:" + region + ":uid/" + appid + ":" + bucket + "/*",
					},
				},
			},
		},
	}
	res, err := c.GetCredential(opt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)
	//fmt.Printf("%+v\n", res.Credentials)

	return "12313123"
}