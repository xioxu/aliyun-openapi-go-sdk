## 项目结构说明
### codeGen
该目录包含有自动生成代码的node.js脚本及其模板，执行方法为node modeGen.js
### openapi-meta-master
该目录包含有阿里云官方提供的openapi meta文件，代码生成工具就依赖于此文件夹来生成go代码
### src
最终SDK所包含的go代码，其中真正写的代码只有aliyun-go-sdk-core文件夹下的内容，其他均为自动生成的代码。
### 设计思路
由于golang中目前不支持泛型功能，因此难以使用一个统一的DoAction方法，那么面临的选择就是要么像java SDK那样在Request对象中添加Do方法，要么就是对每一个Product封装一个Client，
最终我选择了第二种方法，这是因为每个Product都有一些共性的参数，如productName、version， 采用第一种方式的话，这些值在该product下的每个RequestObject中都需要重复出现，这是
一种显而易见的坏味道，而本来Product就是一个逻辑体，封装成Client也就顺理成章了。

### 快速开始for code reviewer:
1. 在本地创建一个文件夹，并checkout该项目
1. 将该文件夹添到gopath环境变量
1. 添加测试代码，同时将endpoints.xml放入与代码相同目录

### 示例代码
```
package main

import (
	"aliyun-go-sdk-core"
	"aliyun-go-sdk-ecs/20140526"
	"fmt"
)

func main() {
	accessProfile := core.GetDefaultProfile("Your Access Key", "Your Access Secret")
	client := ecs.New(accessProfile)

    //Call the operation 'DescribeInstances' of product ecs
	res, err := client.DescribeInstances(nil)

	if err != nil {
		fmt.Println("faild")
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}

```
##特别声明
本代码参考了openapi Jva SDK的部分实现，在此表示感谢

