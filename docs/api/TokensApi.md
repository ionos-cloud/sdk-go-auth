# TokensApi

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

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.TokensDeleteByCriteria(context.Background()).Criteria(criteria).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TokensDeleteByCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokensDeleteByCriteria`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TokensDeleteByCriteria`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokensDeleteByCriteriaRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **criteria** | **string** | Delete tokens by criteria EXPIRED, ALL, or CURRENT. The tokens are deleted for the specified contract. | |
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |

### Return type

[**DeleteResponse**](../models/DeleteResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



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

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.TokensDeleteById(context.Background(), tokenId).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TokensDeleteById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokensDeleteById`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TokensDeleteById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**tokenId** | **string** | The Key ID of the token (can be retrieved from the header section of the token). | |

### Other Parameters

Other parameters are passed through a pointer to a apiTokensDeleteByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |

### Return type

[**DeleteResponse**](../models/DeleteResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



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

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.TokensFindById(context.Background(), tokenId).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TokensFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokensFindById`: Token
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TokensFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**tokenId** | **string** | The Key ID of the token (can be retrieved from the header section of the token). | |

### Other Parameters

Other parameters are passed through a pointer to a apiTokensFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |

### Return type

[**Token**](../models/Token.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## TokensGenerate

```go
var result Jwt = TokensGenerate(ctx)
                      .XContractNumber(xContractNumber)
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

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.TokensGenerate(context.Background()).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TokensGenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokensGenerate`: Jwt
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TokensGenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokensGenerateRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |

### Return type

[**Jwt**](../models/Jwt.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



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

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.TokensGet(context.Background()).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TokensGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokensGet`: Tokens
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TokensGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokensGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |

### Return type

[**Tokens**](../models/Tokens.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


