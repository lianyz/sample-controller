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


## 代码生成示例

### 安装code-generator工具

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

### 创建示例项目

```
mkdir -p $GOPATH/src/github.com/lianyz/sample-controller
cd $GOPATH/src/github.com/lianyz/sample-controller
```

创建如下结构的目录结构和文件
```
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── hack
│         ├── boilerplate.go.txt
│         ├── tools.go
│         ├── update-codegen.sh
│         └── verify-codegen.sh
└── pkg
    └── apis
        └── samplecontroller
            └── v1
                ├── doc.go
                └── types.go
```

各文件的主要作用如下所示
```
├── go.mod                                      *  go mod init 出来的文件
├── go.sum                                      *  go mod init 出来的检测文件
├── hack                                   
│   ├── boilerplate.go.txt                      *  operator代码框架版权信息头内容
│   ├── tools.go                                *  用于自动下载code-generator项目
│   ├── update-codegen.sh                       *  官方sample-controller项目的模板文件，用于填写code-generator中的相关变量
│   └── verify-codegen.sh                       *  调用update-codegen.sh，并检测代码变化的脚本，CI过程中非常重要
├── pkg                                         *  标准目录
│   ├── apis                                    *  标准目录
│   │   └── samplecontroller                    *  group,API组
│   │       └── v1alpha1                        *  version,API版本
│   │           ├── doc.go                      *  代码生成器模板，package相关内容，全局标签填写文件
│   │           ├── types.go                    *  代码生成器模板，API类型相关内容，局部标签填写文件
```

### 生成代码

修改Makefile中group名称和version名称
```
hack/update-codegen.sh samplecontroller v1
```

在sample-controller项目根目录下执行make gen
```
➜  sample-controller git:(master) ✗ make gen
hack/update-codegen.sh
Generating deepcopy funcs
Generating clientset for samplecontroller:v1 at github.com/lianyz/sample-controller/pkg/client/clientset
Generating listers for samplecontroller:v1 at github.com/lianyz/sample-controller/pkg/client/listers
Generating informers for samplecontroller:v1 at github.com/lianyz/sample-controller/pkg/client/informers
```

自动生成以下代码文件
```
	new file:   pkg/apis/samplecontroller/v1/zz_generated.deepcopy.go
	new file:   pkg/client/clientset/versioned/clientset.go
	new file:   pkg/client/clientset/versioned/doc.go
	new file:   pkg/client/clientset/versioned/fake/clientset_generated.go
	new file:   pkg/client/clientset/versioned/fake/doc.go
	new file:   pkg/client/clientset/versioned/fake/register.go
	new file:   pkg/client/clientset/versioned/scheme/doc.go
	new file:   pkg/client/clientset/versioned/scheme/register.go
	new file:   pkg/client/clientset/versioned/typed/samplecontroller/v1/at.go
	new file:   pkg/client/clientset/versioned/typed/samplecontroller/v1/doc.go
	new file:   pkg/client/clientset/versioned/typed/samplecontroller/v1/fake/doc.go
	new file:   pkg/client/clientset/versioned/typed/samplecontroller/v1/fake/fake_at.go
	new file:   pkg/client/clientset/versioned/typed/samplecontroller/v1/fake/fake_samplecontroller_client.go
	new file:   pkg/client/clientset/versioned/typed/samplecontroller/v1/generated_expansion.go
	new file:   pkg/client/clientset/versioned/typed/samplecontroller/v1/samplecontroller_client.go
	new file:   pkg/client/informers/externalversions/factory.go
	new file:   pkg/client/informers/externalversions/generic.go
	new file:   pkg/client/informers/externalversions/internalinterfaces/factory_interfaces.go
	new file:   pkg/client/informers/externalversions/samplecontroller/interface.go
	new file:   pkg/client/informers/externalversions/samplecontroller/v1/at.go
	new file:   pkg/client/informers/externalversions/samplecontroller/v1/interface.go
	new file:   pkg/client/listers/samplecontroller/v1/at.go
	new file:   pkg/client/listers/samplecontroller/v1/expansion_generated.go
```

