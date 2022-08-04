# goshowdoc
1. 第一步获取：
   在showdoc文档中获取apikey、apitoken、showdocurl写入配置文件中

![img.png](img.png)

2. 在路由层添加前置拦截获取请求体
   ![img_1.png](img_1.png)
3. 在控制层post请使用goshowdoc.BindJSON()、goshowdoc.ReplyJson()
  ![img_2.png](img_2.png)
4. 登录showDoc查看文档