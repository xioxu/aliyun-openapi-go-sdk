## Inside the SDK Part 1
###Basic framework
###Profile
Profile object is used to store the user credential and other default parameter, such as region id, response format ...
####BasicClient

Aliyun provides many different products, so I encapulate each them as a client, and all of them is the sub class(maybe we should call it struct) of BaseClient.
And in the BaseClient, it firstl convert the requst object modal to a http requst, then do the concreate http request and finally convert response to a response modal.
####BasicRequest

Each api call needs some parameters, its correspond modal is named request object modal, its base class is BasicRequest.
BasicRequest contains the following methods:
1.  GetQuery() map[string]string
1.	GetBody() map[string]string
1.	GetHeaders() map[string]string
1.	GetDomainValues() map[string]string
1.	GetPath() string
1.	GetProtocol() string
1.	GetRequestMethod() string
1.	GetActionName() string

GetQuery(),GetBody(),GetHeaders() and GetDomainValues() is used to get the http request parameter from your reqeust object modal, since the properties of request modal maybe a query or a body parameter.

####BasicResponse
The interface of response modal, all of response modal   should  implement BuildPropertis() method, this method contains the logic about convert body to modal property.

###Aliyun Product Code
Aliyun product code is structed by product name and product version, like ecs, you should find  aliyun-sdk-go-ecs/20140526, it means in this folder it contains the ecs request and response modal and EcsClient code, which meta is 
published on 20140526.






