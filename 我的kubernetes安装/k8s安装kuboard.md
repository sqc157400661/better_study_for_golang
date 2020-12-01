安装 Kuboard。

> 如果您参考 https://kuboard.cn 网站上提供的 Kubernetes 安装文档，可在 master 节点上执行以下命令。

```
kubectl apply -f https://kuboard.cn/install-script/kuboard.yaml
kubectl apply -f https://addons.kuboard.cn/metrics-server/0.3.7/metrics-server.yaml

```

查看 Kuboard 运行状态：

```sh
kubectl get pods -l k8s.kuboard.cn/name=kuboard -n kube-system
```

## 获取Token

您可以获得管理员用户、只读用户的Token。

> - 默认情况下，您可以使用 ServiceAccount 的 Token 登录 Kuboard
> - 您还可以 [使用 GitLab/GitHub 账号登录 Kuboard/Kubectl](https://kuboard.cn/learning/k8s-advanced/sec/authenticate/install.html)
> - 您也可以 [为用户授权](https://kuboard.cn/learning/k8s-advanced/sec/kuboard.html)



```
# 如果您参考 www.kuboard.cn 提供的文档安装 Kuberenetes，可在第一个 Master 节点上执行此命令
echo $(kubectl -n kube-system get secret $(kubectl -n kube-system get secret | grep kuboard-user | awk '{print $1}') -o go-template='{{.data.token}}' | base64 -d)

```

