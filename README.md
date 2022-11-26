# CRD

## code-generator

code-generator是Kubernetes提供的一个用于代码生成的项目，提供了以下工具为Kubernetes中的资源生成代码：
* deepcopy-gen: 生成深度拷贝方法，为每个类型生成 func (t* T) DeepCopy *T方法,API类型都需要实现深拷贝
* client-gen: 为资源生成标准的ClientSet
* informer-gen: 生成Informer，提供事件机制来响应资源的事件
* lister-gen: 生成Lister，为get和list请求提供只读缓存层（通过indexer获取）
* conversion-gen: 负责产生内外部类型的转换函数
* defaulter-gen: 负责处理字段默认值

Informer和Lister是构建控制器的基础，使用前四个代码生成器可以创建全功能的、和Kubernetes上游控制器工作机制相同的控制器
大部分生成器支持--input-dirs参数来读取一系列输入包，处理其中的每个类型，然后生成代码。
* 部分代码生成到输入包所在目录，如deepcopy-gen生成器, 也可以使用参数--output-file-base来定义输出文件名
* 其它代码生成到--output-package指定的目录，例如client-gen, informer-gen, lister-gen等生成器。


** 代码生成示例

*** 安装code-generator工具

将code-generator拉取到本地
```
mkdir -p $GOPATH/src/github.com/kubernetes
cd $GOPATH/src/github.com/kubernetes
git clone https://github.com/kubernetes/code-generator.git
cd code-generator
```

然后进入cmd目录下， 可以看到各个生成器
```
➜  code-generator git:(master) ls cmd
applyconfiguration-gen   defaulter-gen            lister-gen               set-gen
client-gen               go-to-protobuf           openapi-gen
conversion-gen           import-boss              prerelease-lifecycle-gen
deepcopy-gen             informer-gen             register-gen
```

安装各个生成器

```
go install ./cmd/*-gen
```

查看安装结果
```
➜  code-generator git:(master) cd $GOPATH/bin
➜  bin ls *-gen
applyconfiguration-gen   deepcopy-gen             lister-gen               register-gen
client-gen               defaulter-gen            openapi-gen              set-gen
conversion-gen           informer-gen             prerelease-lifecycle-gen
```