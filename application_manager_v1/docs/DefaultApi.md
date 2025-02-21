# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com/armui/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllDeployments**](DefaultApi.md#GetAllDeployments) | **Get** /applications | List deployments



## GetAllDeployments

> GetAllDeploymentsResponse GetAllDeployments(ctx).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).Execute()

List deployments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v1"
)

func main() {
    xAnypntEnvId := "xAnypntEnvId_example" // string | The environment id where the applications are deployed
    xAnypntOrgId := "xAnypntOrgId_example" // string | The org id where the applications are deployed

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAllDeployments(context.Background()).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllDeployments`: GetAllDeploymentsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnypntEnvId** | **string** | The environment id where the applications are deployed | 
 **xAnypntOrgId** | **string** | The org id where the applications are deployed | 

### Return type

[**GetAllDeploymentsResponse**](GetAllDeploymentsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

