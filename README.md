# Gin-GoServer

這是一個基於 Go 語言, 與類別庫 Gin 實踐的 Web Api Server 並且包含 DockerFile 建構。

1. 如果想要使用 Docker 建構，請安裝 Docker, Docker Compose, 並執行
`docker-compose up --build`

2. 如果想使用本地環境建構，請安裝 Go 環境，並將資料夾加入 $GOPATH/src/ 底下

3. 結構如下，使用的依賴工具為 [dep](https://github.com/golang/dep)
.
├── Dockerfile
├── Gin-GoServer
├── Gopkg.lock
├── Gopkg.toml
├── README.md
├── conf
├── config
├── docker-compose.yaml
├── handler
├── log
├── main.go
├── routes
└── vendor