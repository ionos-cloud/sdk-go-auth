# \TokensApi

All URIs are relative to *https://api.ionos.com/auth/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**TokensDeleteByCriteria**](TokensApi.md#TokensDeleteByCriteria) | **Delete** /tokens | Delete tokens by criteria|
|[**TokensDeleteById**](TokensApi.md#TokensDeleteById) | **Delete** /tokens/{tokenId} | Delete tokens|
|[**TokensFindById**](TokensApi.md#TokensFindById) | **Get** /tokens/{tokenId} | Get tokens by Key ID|
|[**TokensGenerate**](TokensApi.md#TokensGenerate) | **Get** /tokens/generate | Create new tokens|
|[**TokensGet**](TokensApi.md#TokensGet) | **Get** /tokens | List all tokens|



## TokensDeleteByCriteria

```go
var result DeleteResponse = TokensDeleteByCriteria(ctx)
                      .Criteria(criteria)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete tokens by criteria



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-auth"
)

func main() {
    criteria := "criteria_example" // string | Delete tokens by criteria EXPIRED, ALL, or CURRENT. The tokens are deleted for the specified contract.
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which the token is generated. (optional)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.TokensApi.TokensDeleteByCriteria(context.Background()).Criteria(criteria).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TokensDeleteByCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TokensDeleteByCriteria`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TokensDeleteByCriteria`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiTokensDeleteByCriteriaRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **criteria** | **string** | Delete tokens by criteria EXPIRED, ALL, or CURRENT. The tokens are deleted for the specified contract. | |
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |

### Return type

[**DeleteResponse**](../models/DeleteResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"TokensApiService.TokensDeleteByCriteria"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "TokensApiService.TokensDeleteByCriteria": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "TokensApiService.TokensDeleteByCriteria": {
    "port": "8443",
},
})
```


## TokensDeleteById

```go
var result DeleteResponse = TokensDeleteById(ctx, tokenId)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete tokens



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-auth"
)

func main() {
    tokenId := "tokenId_example" // string | The Key ID of the token (can be retrieved from the header section of the token).
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which the token is generated. (optional)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.TokensApi.TokensDeleteById(context.Background(), tokenId).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TokensDeleteById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TokensDeleteById`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TokensDeleteById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**tokenId** | **string** | The Key ID of the token (can be retrieved from the header section of the token). | |

### Other Parameters

Other parameters are passed through a pointer to an apiTokensDeleteByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |

### Return type

[**DeleteResponse**](../models/DeleteResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"TokensApiService.TokensDeleteById"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "TokensApiService.TokensDeleteById": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "TokensApiService.TokensDeleteById": {
    "port": "8443",
},
})
```


## TokensFindById

```go
var result Token = TokensFindById(ctx, tokenId)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Get tokens by Key ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-auth"
)

func main() {
    tokenId := "tokenId_example" // string | The Key ID of the token (can be retrieved from the header section of the token).
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which the token is generated. (optional)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.TokensApi.TokensFindById(context.Background(), tokenId).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TokensFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TokensFindById`: Token
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TokensFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**tokenId** | **string** | The Key ID of the token (can be retrieved from the header section of the token). | |

### Other Parameters

Other parameters are passed through a pointer to an apiTokensFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |

### Return type

[**Token**](../models/Token.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"TokensApiService.TokensFindById"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "TokensApiService.TokensFindById": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "TokensApiService.TokensFindById": {
    "port": "8443",
},
})
```


## TokensGenerate

```go
var result Jwt = TokensGenerate(ctx)
                      .XContractNumber(xContractNumber)
                      .Ttl(ttl)
                      .Execute()
```

Create new tokens



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-auth"
)

func main() {
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which the token is generated. (optional)
    ttl := int32(56) // int32 | The maximum time that the access token will be valid for use within the application in seconds. (optional) (default to 31536000)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.TokensApi.TokensGenerate(context.Background()).XContractNumber(xContractNumber).Ttl(ttl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TokensGenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TokensGenerate`: Jwt
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TokensGenerate`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiTokensGenerateRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |
| **ttl** | **int32** | The maximum time that the access token will be valid for use within the application in seconds. | [default to 31536000]|

### Return type

[**Jwt**](../models/Jwt.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"TokensApiService.TokensGenerate"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "TokensApiService.TokensGenerate": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "TokensApiService.TokensGenerate": {
    "port": "8443",
},
})
```


## TokensGet

```go
var result Tokens = TokensGet(ctx)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List all tokens



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-auth"
)

func main() {
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which the token is generated. (optional)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.TokensApi.TokensGet(context.Background()).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TokensGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TokensGet`: Tokens
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TokensGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiTokensGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |

### Return type

[**Tokens**](../models/Tokens.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"TokensApiService.TokensGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "TokensApiService.TokensGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "TokensApiService.TokensGet": {
    "port": "8443",
},
})
```

