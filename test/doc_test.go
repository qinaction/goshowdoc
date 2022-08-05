package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qinaction/goshowdoc"
	"net/http"
	"testing"
)

//请求绑定必须使用tool包中的BindJSON或者BindJSONNotWithValidate
//若不想使用rsp.go 请使用ReplyJson或者ReplyCodeJson
func TestGinResp(t *testing.T) {
	r := gin.New()
	UrlReq := "测试域名 ` http://xxx.cn`\n"
	goshowdoc.EnableDoc(2, goshowdoc.Doc{
		Url:      "https://XX=/api/item/updateByApi",
		ApiKey:   "XX",
		ApiToken: "XX",
		UrlReq: UrlReq,
	}, goshowdoc.Gin{})

	r.POST("/user/list", userList)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "ping")
	})
	err := r.Run("127.0.0.1:8199")
	if err != nil {
		fmt.Println(err)
	}
}

type docModel struct {
	ID   uint   `gorm:"primary_key;comment:'数据id'" json:"id" req:"-"`
	Name string `json:"name" gorm:"type:varchar(30);comment:姓名" req:"-"`
	Sex  string `json:"sex" gorm:"type:varchar(30);comment:性别" req:"-"`
	Age  int    `json:"age" gorm:"comment:年龄" req:"-"`
}

type u struct {
	Name string `json:"name" gorm:"type:varchar(30);comment:姓名" req:"-"`
	Sex  string `json:"sex" gorm:"type:varchar(30);comment:性别" req:"-"`
}


/*
使用公共响应体
*/
func userList(c *gin.Context) {
	var (
		rsp Rsp
		req u
	)
	// showDoc绑定参数
	if err := goshowdoc.BindJSON(c,req);err != nil{
		rsp.ReplyFailOperation(c, err.Error())
		return
	}
	//获取参数成功进行service、db操作
	//操作完成返回
	docModel := docModel{
		ID: 1,
		Name: "张三",
		Sex: "男",
		Age: 18,
	}
	goshowdoc.ReplyJson(c,docModel)
	//输出json
	c.JSON(http.StatusOK,docModel)
}


