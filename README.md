# gopractice
my self practice go lang



### 放進 本地 minikube 
啟動 minikube 需要mount  外部地址 在使用PV PVC  去做內部 link 

minikube mount /Users/znikanghuang/k8s/mnt:/mnt

minikube start --mount --mount-string="/Users/znikanghuang/k8s/mnt:/mm/mnt" --network=192.168.130.0/24 
--driver=hyperkit

minikube ssh 進入minikube 後可以查看 /mm 是否有link 到外面的資料夾

PersistentVolume 只在於 minikube ssh 進入容器內的路徑 還是需要從外部在mount 一次  所以開始最好 先 進入 mkdir

打包 docker 要用 minikube env 
eval $(minikube docker-env)

docker build -t my-app .
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
產生 protuc 相關檔案



