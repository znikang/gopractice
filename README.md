# gopractice
my self practice go lang



### 放進 本地 minikube 
啟動 minikube 需要mount  外部地址 在使用PV PVC  去做內部 link 

minikube mount /Users/znikanghuang/k8s/mnt:/mnt

minikube start --mount --mount-string="/Users/znikanghuang/k8s/mnt:/mm/mnt"

minikube ssh 進入minikube 後可以查看 /mm 是否有link 到外面的資料夾

PersistentVolume 只在於 minikube ssh 進入容器內的路徑 還是需要從外部在mount 一次  所以開始最好 先 進入 mkdir

打包 docker 要用 minikube env 
eval $(minikube docker-env)

docker build -t my-app .
### 需要這段
yaml 
imagePullPolicy: Never
