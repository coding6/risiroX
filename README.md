# RisiroX
## V1.0
实现简单的通信，支持单客户端和服务端的通信
## v2.0
支持多客户端与服务端的通信，建立客户端连接管理器，超过一定数量拒绝连接
## v3.0
读写分离，读单独开启goroutine。创建MessageHandler，读取到数据后，将数据提交给工作线程池，
具体的业务逻辑交给用户自己实现。

## v4.0
* 支持消息绑定Handler，比如消息类型是1，1可以执行用户自定义的操作
* 消息封装，封装为Message结构，现在是原始的[]byte
## v5.0
待做：
* 支持一个客户端发送消息广播到其他客户端（聊天服务器） done
* 支持connection连接启停前后的自定义操作 done
* 客户端的心跳检测机制
* 客户端下线时，连接池没有正确清空这个连接done
* 服务端下线，貌似客户端没有感知到，排查原因 done