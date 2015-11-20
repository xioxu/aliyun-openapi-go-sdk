## Inside the SDK Part 3
###Endpoint
We provided endpoints.xml in default, but sometimes you may need to add you custom endpoint, it is also very simple:

`profile.AddEncpoint(...)`

### Package "parameterHelper"
Package parameterHelper contains some helper methods for buiding http request. Such as parsing iso8601/rfc2616 formated time, non-standed base64 encoding, and md5 method.

