# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/runtimefabric/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrivateSpace**](DefaultApi.md#CreatePrivateSpace) | **Post** /organizations/{orgId}/privatespaces | 
[**DeletePrivateSpace**](DefaultApi.md#DeletePrivateSpace) | **Delete** /organizations/{orgId}/privatespaces/{privateSpaceId} | 
[**GetPrivateSpace**](DefaultApi.md#GetPrivateSpace) | **Get** /organizations/{orgId}/privatespaces/{privateSpaceId} | 
[**GetPrivateSpaces**](DefaultApi.md#GetPrivateSpaces) | **Get** /organizations/{orgId}/privatespaces | 
[**UpdatePrivateSpace**](DefaultApi.md#UpdatePrivateSpace) | **Patch** /organizations/{orgId}/privatespaces/{privateSpaceId} | 



## CreatePrivateSpace

> PrivateSpace CreatePrivateSpace(ctx, orgId).PrivateSpacePostBody(privateSpacePostBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    privateSpacePostBody := *openapiclient.NewPrivateSpacePostBody() // PrivateSpacePostBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreatePrivateSpace(context.Background(), orgId).PrivateSpacePostBody(privateSpacePostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreatePrivateSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrivateSpace`: PrivateSpace
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreatePrivateSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateSpacePostBody** | [**PrivateSpacePostBody**](PrivateSpacePostBody.md) |  | 

### Return type

[**PrivateSpace**](PrivateSpace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivateSpace

> DeletePrivateSpace(ctx, orgId, privateSpaceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeletePrivateSpace(context.Background(), orgId, privateSpaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeletePrivateSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateSpace

> PrivateSpace GetPrivateSpace(ctx, orgId, privateSpaceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPrivateSpace(context.Background(), orgId, privateSpaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPrivateSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateSpace`: PrivateSpace
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPrivateSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PrivateSpace**](PrivateSpace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateSpaces

> PrivateSpaceSummary GetPrivateSpaces(ctx, orgId).Offset(offset).Limit(limit).Sort(sort).Ascending(ascending).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    offset := int32(56) // int32 | The offset specifies the offset of the first row to return (optional)
    limit := int32(56) // int32 | Amount of objects retrieved in the response (optional)
    sort := "sort_example" // string | Property to sort by (optional)
    ascending := "ascending_example" // string | Order for sorting (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPrivateSpaces(context.Background(), orgId).Offset(offset).Limit(limit).Sort(sort).Ascending(ascending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPrivateSpaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateSpaces`: PrivateSpaceSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPrivateSpaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateSpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | The offset specifies the offset of the first row to return | 
 **limit** | **int32** | Amount of objects retrieved in the response | 
 **sort** | **string** | Property to sort by | 
 **ascending** | **string** | Order for sorting | 

### Return type

[**PrivateSpaceSummary**](PrivateSpaceSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivateSpace

> PrivateSpace UpdatePrivateSpace(ctx, orgId, privateSpaceId).PrivateSpacePatchBody(privateSpacePatchBody).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/private_space"
)

func main() {
    orgId := "orgId_example" // string | The ID of the organization in GUID format
    privateSpaceId := "privateSpaceId_example" // string | The ID of the private space in GUID format
    privateSpacePatchBody := *openapiclient.NewPrivateSpacePatchBody() // PrivateSpacePatchBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdatePrivateSpace(context.Background(), orgId, privateSpaceId).PrivateSpacePatchBody(privateSpacePatchBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdatePrivateSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePrivateSpace`: PrivateSpace
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdatePrivateSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The ID of the organization in GUID format | 
**privateSpaceId** | **string** | The ID of the private space in GUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrivateSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **privateSpacePatchBody** | [**PrivateSpacePatchBody**](PrivateSpacePatchBody.md) |  | 

### Return type

[**PrivateSpace**](PrivateSpace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

