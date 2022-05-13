# go_tls_server

实现了一个简单的go语言开发的基于pb通信的tls server

目录介绍：
cache:
  简单封装redis的使用，基于redigo
  
cmd：
  client && server
  
db：
  使用gorm
  
log：
  使用logrus
  
logic：
  业务逻辑实现
  
    协议参见protocol部分的pb文件
    
    共包括三个接口login get 和 set
    
    接口部分待后续优化
    
model：
  数据部分
  
protocol：
  协议
tools：
  生成测试数据的工具
  
util
