## Using the SDK
###Installing with go

Please us `go get` to retrieve the latest version SDK, simply type the following command into a terminal window will instal the entire SDK:
```
go get github.com/aliyun/aliyun-sdk-go/...
```

## Importing
Once youo have installed the Aliyun SDK, you can import your expected product package into your application code, assuems you want to access `Aliyun Ecs` product,  then you need to import the specitied version Ecs SDK:
```
import (
	"aliyun-go-sdk-core"
	"aliyun-go-sdk-ecs/20140526"
)
```

## Accessing Profile
An accessing profile is a storage for store your api accessKey and accessSecret, which also provides a property to allows you change the default Signature method (You can omit it untill Aliyun supports a new signature method).

```
accessProfile := core.GetDefaultProfile("Your Access Key", "Your Access Secret")
```

## Product Accessing Client
Each product has its own cient and this client contains all the all apis of this product, so need to create its instance with a profile parameter which described in above:
```
client := ecs.New(accessProfile)
```

## Access the product api
With the above `client` instance, you will need create the required request object and set the required value, then the you will get the response or an error object:
```
req := &ecs.DescribeInstancesRequest{}
res, err := client.DescribeInstances(req)
```
