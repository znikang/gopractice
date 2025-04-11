

# gopractice

副標
---

my self practice go lang

###  golang 運行時的arch 版本
> 當Mac M1 chip  使用的 docker desktop  裡面跑的是 amd64  所以 你得使用
>  OOS=linux GOARCH=amd64 go build -o myapp    在打包進  docker
> 
> 當 Mac M1 chip  的minikube 是跑  Architecture: aarch64
>  所以  docker 要用  去包file 
> docker build -t --platform linux/arm64 my-app:1.0 .
> 

### 放進 本地 minikube 
啟動 minikube 需要mount  外部地址 在使用PV PVC  去做內部 link 

minikube mount /Users/znikanghuang/k8s/mnt:/mnt

minikube start --mount --mount-string="/Users/znikanghuang/k8s/mnt:/mm/mnt"  --driver=qemu2    
> for mac m1 系列 


```aidl
 go lang
```
> --arch amd64 

> --arch arm64

minikube ssh 進入minikube 後可以查看 /mm 是否有link 到外面的資料夾

PersistentVolume 只在於 minikube ssh 進入容器內的路徑 還是需要從外部在mount 一次  所以開始最好 先 進入 mkdir

打包 docker 要用 minikube env 
eval $(minikube docker-env)

##產生 進minikube 的ar版本要很注意 
> docker info | grep Architecture
> Architecture: aarch64

> 所以編譯後  要打入的binary 要使用的 arm64 架構 不然能跑

> OOS=linux GOARCH=arm64 go build -o myapp
> 
> > OOS=linux GOARCH=amd64 go build -o myapp


docker build -t --platform linux/arm64 my-app:1.0 .

docker build -t my-webserver:1.0 . 

GOOS=linux GOARCH=amd64 go build -o webserver


### 需要這段
yaml   只讀取 local docker image
imagePullPolicy: Never or IfNotPresent





### minikube 指令

kubectl apply -f xxx.yaml 可以匯入 

kubectl delete -f xxx.yaml 刪除匯入的內容 

>minikube service gopractice-service 做外部的link service




### ingress

minikube addons list

>使用 minikube 需要
> minikube addons enable ingress



### protobuf 

protoc --go_out=. --go-grpc_out=. example.proto 
產生 protoc 相關檔案



