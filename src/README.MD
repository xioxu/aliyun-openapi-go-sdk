# Aliyun OpenAPi SDK for the Go programming language
Welcome to Aliyun SDK for Golang, this project provides a unified accessing interface for Aliyun OpenApi.

## Installing
NOTE: The code  has not been uploaded to github, so it not works in currenlly.

Install your specific product client through the `go tet` command likes the following:
```
get get github.com/aliyun/aliyun-sdk-go/ecs
```
You can also install the entire SDK by installing the root package:
```
go get github.com/aliyun/aliyun-sdk-go/...
```
## Endpoints
The product accessing url infomation is stored in file `Endpoints.xml`, you need to put it in your application root directory, here is the default `Endpoints.xml` [Download in Github](https://github.com/aliyun/aliyun-openapi-java-sdk/blob/b9d679ad3be9a73e14f1cbd83fd6ff1b0bdfd284/aliyun-java-sdk-core/src/main/resources/com/aliyuncs/endpoints/endpoints.xml)

## Usage

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

## License
licensed under the [Apache License 2.0](https://www.apache.org/licenses/

