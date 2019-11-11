# MyCloudGO

### 运行服务器
``go run main.go -p 5432``  
其中5432只是一个端口例子，你可以改为其它可用的端口号，也可以默认使用8080端口号（如下）  
``go run main.go``  

### 运行curl测试
测试hello：``curl -v http://localhost:5432/hello/testuser``  
测试bye：``curl -v http://localhost:5432/bye/testuser``  
混乱测试:``curl -v http://localhost:5432/Shit/You``、``curl -v http://localhost:5432/Shit``  
  
### 运行ab测试  
``ab -n 1000 -c 100 http://localhost:5432/hello/aivo``

### 代码说明
1. main.go是入口文件，main函数如下。主要是通过pflag获取port的值，然后利用port的值创建server并运行
```go
const (
    PORT string = "8080"
)
func main() {
    port := os.Getenv("PORT")
    if len(port) == 0 {
        // 获取默认port值：8080
        port = PORT
    }
    // 获取显式指明的port值
    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }
    // 创建并运行server
    server := service.NewServer()
    server.Run(":" + port)
}
```
2. server.go的功能主要是初始化配置，监听访问。

```go
func NewServer() *negroni.Negroni {
    // 解析格式为Json格式
    formatter := render.New(render.Options{
        IndentJSON: true,
    })
    // 为handle提供接口
    n := negroni.Classic()
    // 多路运行（并发）
    mx := mux.NewRouter()
    // 初始化路由
    initRoutes(mx, formatter)
    // 使用路由
    n.UseHandler(mx)
    return n
}
```
初始化路由的代码如下：  
```go
func initRoutes(mx *mux.Router, formatter *render.Render) {
    // 处理类似localhost:8080/xxx的访问
	  mx.HandleFunc("/{op}", testHandler(formatter)).Methods("GET")
    // 处理类似localhost:8080/xxx/yyy的访问
    mx.HandleFunc("/{op}/{id}", testHandler(formatter)).Methods("GET")
}
```
处理访问的详细代码如下：
```go
func testHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        vars := mux.Vars(req)
        // 假设访问localhost:8080/hello/testuser
        // 获取op部分的值，比如hello
		    op := vars["op"]
        // 获取id部分的值，比如testuser
        id := vars["id"]
        
        switch {
          case op == "hello":
            // 如果是hello操作，那么返回客户端的json信息为Hello, + id
            formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello, " + id})
          case op == "bye":
            // 如果是bey操作，那么返回客户端的json信息为Bye, + id
            formatter.JSON(w, http.StatusOK, struct{ Test string }{"Bye, " + id})
          default:
            // 其它操作都是错误操作，直接返回Wrong url
            formatter.JSON(w, http.StatusOK, struct{ Test string }{"Wrong url"})
        }
    }
}
```
