####summen10
#####请求地址
```text
测试域名 ` http://xxx.cn`

```
#####请求方式
```text
POST
```
#####请求参数

u

|参数名|必选|类型|说明|
|:----:|:---:|:-----:|:-----:|
|name|否|string|姓名|
|sex|否|string|性别|

#####请求示例
```text
{
	"name": "",
	"sex": ""
}
```
#####返回示例
```text
{
	"id": 1,
	"name": "张三",
	"sex": "男",
	"age": 18
}
```
#####返回参数说明

docModel

|参数名|类型|说明|
|:----:|:-----:|:-----:|
|id|uint|数据id|
|name|string|姓名|
|sex|string|性别|
|age|int|年龄|

#####备注
