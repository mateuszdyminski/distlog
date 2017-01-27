# distlog

Example of distributed logging with gRPC and Zap based on net/context. It shows the possibility to pass context variables between the distributed services. 

## Run 

### Run http server
```
go run server.go
```

Sample logs from http server:

```
md@md distlog (master)$ go run server.go
{"level":"info","time":"2017-01-27T19:23:54Z","reqId":"7febbbcd-8bd1-4e9e-b6cb-6ff4aa55bf34","msg":"got arguments from http request","pid":23003,"exe":"server","arg1":1,"arg2":2,"operation":"ADD"}
{"level":"info","time":"2017-01-27T19:23:54Z","reqId":"7febbbcd-8bd1-4e9e-b6cb-6ff4aa55bf34","msg":"got result for arguments","pid":23003,"exe":"server","arg1":1,"arg2":2,"operation":"ADD","result":3
```

### Run gRPC server
```
go run calculator.go
```

Sample logs from gRPC server:
```
md@md distlog (master *#)$ go run calculator.go
{"level":"info","time":"2017-01-27T19:23:54Z","reqId":"7febbbcd-8bd1-4e9e-b6cb-6ff4aa55bf34","msg":"got arguments from rpc request","pid":22970,"exe":"calculator","arg1":1,"arg2":2,"operation":"ADD"}
```

