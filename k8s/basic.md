

## install for learn
### install kind
    https://kind.sigs.k8s.io/docs/user/quick-start/#installation

kind.yml
1. kind 创建的集群默认只监听127.0.0.1
2. kind 创建的集群默认从docker Hub拉取镜像
解决方法：通过一下配置创建集群
```
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  apiServerAddress: "{{your ip addr}}"
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"]
    endpoint = ["http://f1361db2.m.daocloud.io"]
```

```
kind create cluster --config kind.yml #创建一个集群--默认名称为kind
kind create cluster --name lin-yuan --config kind.yml #创建一个名称为lin-yuan的集群
kind get clusters #查看目前所有的集群

# 删除一个集群
# kind delete cluster --name [cluster name] 

cat ~/.kube/config # 集群的访问配置
```
### install kubectl
The Kubernetes command-line tool, kubectl, allows you to run commands against Kubernetes clusters. You can use kubectl to deploy applications, inspect and manage cluster resources, and view logs.
    https://kubernetes.io/docs/tasks/tools/install-kubectl-macos/

pay attention to this command when install kubectl:
    sudo chown root: /usr/local/bin/kubectl 
it mean set /usr/local/bin/kubectl to root user control group, which will give it the root power.

```
kubectl config view # 查看当前kubectl的配置
mkdir /Users/linyuan/kubectl
cd /Users/linyuan/kubectl
vim config
# 然后将 【cat ~/.kube/config】 的内容copy到这里来
# 将该配置文件的地址写入到环境变量KUBECONFIG ，如修改~/.zshrc等文件记得要source

# 配置成功则能看到对应配置
kubectl config view

# 切换集群
kubectl config use-context [cluster name]
```

## 配置kubctl访问clusters
    


## k8s 的常用几个类型

1. Pod
    pod代表一组共享namespace，cgroup和其他一些隔离的容器集合。（Pod通常不是直接创建的，而是根据诸如depolyment，job等工作负载进行创建的）
    一个最简单的Pod fileName: simple-pod.yaml
    ```
        apiVersion: v1
        kind: Pod
        metadata:
        name: nginx
        spec:
        containers:
        - name: nginx
            image: nginx:1.14.2
            ports:
            - containerPort: 80
    ```
    创建上面的Pod
    ```
        kubectl apply -f ./simple-pod.yaml
    ```
2. Depolyment


3. Service