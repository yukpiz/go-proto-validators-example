```bash
protoc -I. -I${GOPATH}/src --go_out=plugins=grpc:pb/ --govalidators_out=pb/ example.proto
```
