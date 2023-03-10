

# 开篇

Kubernetes 资源扩展

![image-20230309204345606](assets/image-20230309204345606.png)



API类型的选择

| 考虑 API 聚合的情况                                          | 优先独立 API 的情况                                          |
| ------------------------------------------------------------ | :----------------------------------------------------------- |
| 你在开发新的 API                                             | 已经有一个提供API服务的程序并且工作良好                      |
| 你希望在 Kubernetes UI 中和其他内置类别一起查看你的新资源类别 | 不要求 kubectl 支持                                          |
| 你希望复用 Kubernetes API 支持特性                           | 不需要 Kubernetes UI 支持                                    |
| 你有意愿接受 Kubernets 对 REST 资源路径所做的格式限制，例如 API 组和命名空间 | 你不需要这些特性                                             |
| 你的 API 是声明式的                                          | 你的API 不符合声明式模型                                     |
| 你的资源可以自然地界定为集群作用域或集群中某个命名空间作用域 | 集群作用域的命名空间作用域这种二分法很不合适；需要对资源路径的细节进行控制 |



相关工具：

Client-go: https://github.com/kubernetes/client-go

sample-controller(controller示例): https://github.com/kubernetes/sample-controller

kubebuilder(快速创建controller的工具): https://github.com/kubernetes-sigs/kubebuilder  https://book.kubebuilder.io/

sample-apiserver(apiserver示例):https://github.com/kubernetes/sample-apiserver



# client-go













