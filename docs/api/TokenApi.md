# TokenApi

All URIs are relative to *https://api.ionos.com/auth/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**CreateToken**](TokenApi.md#CreateToken) | **Get** /tokens/generate | Create new tokens|
|[**DeleteTokenByCriteria**](TokenApi.md#DeleteTokenByCriteria) | **Delete** /tokens | Delete tokens by criteria|
|[**DeleteTokenById**](TokenApi.md#DeleteTokenById) | **Delete** /tokens/{tokenId} | Delete tokens|
|[**GetAllTokens**](TokenApi.md#GetAllTokens) | **Get** /tokens | List all tokens|
|[**GetTokenById**](TokenApi.md#GetTokenById) | **Get** /tokens/{tokenId} | Get tokens by Key ID|



## CreateToken

```go
var result Jwt = CreateToken(ctx)
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
    openapiclient "./openapi"
)

func main() {
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which the token is generated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokenApi.CreateToken(context.Background()).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: Jwt
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |

### Return type

[**Jwt**](../models/Jwt.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DeleteTokenByCriteria

```go
var result DeleteResponse = DeleteTokenByCriteria(ctx)
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
    openapiclient "./openapi"
)

func main() {
    criteria := "criteria_example" // string | Delete tokens by criteria EXPIRED, ALL, or CURRENT. The tokens are deleted for the specified contract.
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which the token is generated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokenApi.DeleteTokenByCriteria(context.Background()).Criteria(criteria).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.DeleteTokenByCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTokenByCriteria`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.DeleteTokenByCriteria`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenByCriteriaRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **criteria** | **string** | Delete tokens by criteria EXPIRED, ALL, or CURRENT. The tokens are deleted for the specified contract. | |
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |

### Return type

[**DeleteResponse**](../models/DeleteResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DeleteTokenById

```go
var result DeleteResponse = DeleteTokenById(ctx, tokenId)
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
    openapiclient "./openapi"
)

func main() {
    tokenId := "tokenId_example" // string | The Key ID of the token (can be retrieved from the header section of the token).
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which the token is generated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokenApi.DeleteTokenById(context.Background(), tokenId).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.DeleteTokenById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTokenById`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.DeleteTokenById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**tokenId** | **string** | The Key ID of the token (can be retrieved from the header section of the token). | |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |

### Return type

[**DeleteResponse**](../models/DeleteResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GetAllTokens

```go
var result Tokens = GetAllTokens(ctx)
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
    openapiclient "./openapi"
)

func main() {
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which the token is generated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokenApi.GetAllTokens(context.Background()).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.GetAllTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTokens`: Tokens
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.GetAllTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTokensRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |

### Return type

[**Tokens**](../models/Tokens.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GetTokenById

```go
var result Token = GetTokenById(ctx, tokenId)
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
    openapiclient "./openapi"
)

func main() {
    tokenId := "tokenId_example" // string | The Key ID of the token (can be retrieved from the header section of the token).
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which the token is generated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokenApi.GetTokenById(context.Background(), tokenId).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.GetTokenById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenById`: Token
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.GetTokenById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**tokenId** | **string** | The Key ID of the token (can be retrieved from the header section of the token). | |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which the token is generated. | |

### Return type

[**Token**](../models/Token.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


