package goshowdoc

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitShowDoc() error {

	 fmt.Println(viper.GetString("e"))
	if viper.GetString("e") == "default" {
		//0不生成文档 1生成ShowDoc 2生成markDown
		UrlReq := "测试域名 ` http://backend.equity.common.turboradio.cn`\n" +
			"- R C 域名 ` http://equity.claim.backend.uniondrug.net`\n" +
			"- 生产域名 ` http://equity.claim.backend.uniondrug.cn`\n"
		EnableDoc(1, Doc{
			Url:      "https://exterior-doc.turboradio.cn/server/index.php?s=/api/item/updateByApi",
			ApiKey:   "5e8cbca2425b3c328621196aee0f5c8438371358",
			ApiToken: "d3b48d01bc4093c6e44f6e67a39a89bc413127168",
			UrlReq: UrlReq,
		}, Gin{})
	}
	return nil
}
