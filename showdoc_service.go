package goshowdoc

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitShowDoc() error {

	 fmt.Println(viper.GetString("e"))
	if viper.GetString("e") == "default" {
		//0不生成文档 1生成ShowDoc 2生成markDown
		UrlReq := "域名 ` http://m.cn`\n"
		EnableDoc(1, Doc{
			Url:      "",
			ApiKey:   "",
			ApiToken: "",
			UrlReq: UrlReq,
		}, Gin{})
	}
	return nil
}
