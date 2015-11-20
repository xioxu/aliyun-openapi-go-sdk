## Inside the SDK Part 2
###RPC and ROA
Aliyun product have two different api styles, one is RPC and anoother is ROA.

####RPC
RPA likes a general http request, and the reqeust parameter is submitted with url query or request body.
####ROA
ROA uses http head to send the request parameter.

###Signature
Both RPC and ROA have same sign algorithm, it use the hmac-sh1 (the only supported method currentlly) to sign a string and send it in each http request. 
 





