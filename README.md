# MyCloudGO

### 运行服务器
``go run main.go -p 5432``
其中6543只是一个端口例子，你可以改为其它可用的端口号

### 运行curl测试
测试hello：``curl -v http://localhost:5432/hello/testuser``
测试bye：``curl -v http://localhost:5432/bye/testuser``
混乱测试:``curl -v http://localhost:5432/Shit/You``、``curl -v http://localhost:5432/Shit``

### 运行ab测试
``ab -n 1000 -c 100 http://localhost:5432/hello/aivo``
