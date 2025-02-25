# \DefaultApi

All URIs are relative to *https://anypoint.mulesoft.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeployment**](DefaultApi.md#DeleteDeployment) | **Delete** /cloudhub/api/v2/applications/{domain} | Delete a single application
[**GetAllDeployments**](DefaultApi.md#GetAllDeployments) | **Get** /armui/api/v1/applications | List deployments
[**GetDeploymentDetails**](DefaultApi.md#GetDeploymentDetails) | **Get** /cloudhub/api/v2/applications/{domain} | Read deployment details
[**GetMuleVersions**](DefaultApi.md#GetMuleVersions) | **Get** /cloudhub/api/mule-versions | Read available mule versions
[**PostDeployment**](DefaultApi.md#PostDeployment) | **Post** /cloudhub/api/v2/applications | Create a new deployment
[**PutDeployment**](DefaultApi.md#PutDeployment) | **Put** /cloudhub/api/v2/applications/{domain} | Update a single application.
[**PutDeployments**](DefaultApi.md#PutDeployments) | **Put** /cloudhub/api/v2/applications | Updates deployment states in bulk



## DeleteDeployment

> DeleteDeployment(ctx, domain).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).Execute()

Delete a single application



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
    xAnypntEnvId := "xAnypntEnvId_example" // string | The environment id where the application is deployed
    xAnypntOrgId := "xAnypntOrgId_example" // string | The org id where the application is deployed
    domain := "domain_example" // string | The application domain in the path

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteDeployment(context.Background(), domain).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The application domain in the path | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnypntEnvId** | **string** | The environment id where the application is deployed | 
 **xAnypntOrgId** | **string** | The org id where the application is deployed | 


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


## GetDeploymentDetails

> DeploymentDetails GetDeploymentDetails(ctx, domain).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).Execute()

Read deployment details



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
    xAnypntEnvId := "xAnypntEnvId_example" // string | The environment id where the application is deployed
    xAnypntOrgId := "xAnypntOrgId_example" // string | The org id where the application is deployed
    domain := "domain_example" // string | The application domain in the path

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetDeploymentDetails(context.Background(), domain).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDeploymentDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentDetails`: DeploymentDetails
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDeploymentDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The application domain in the path | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnypntEnvId** | **string** | The environment id where the application is deployed | 
 **xAnypntOrgId** | **string** | The org id where the application is deployed | 


### Return type

[**DeploymentDetails**](DeploymentDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMuleVersions

> RuntimeCollection GetMuleVersions(ctx).Execute()

Read available mule versions



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetMuleVersions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMuleVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMuleVersions`: RuntimeCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMuleVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMuleVersionsRequest struct via the builder pattern


### Return type

[**RuntimeCollection**](RuntimeCollection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeployment

> DeploymentDetails PostDeployment(ctx).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).AppInfoJson(appInfoJson).File(file).AutoStart(autoStart).Execute()

Create a new deployment



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
    xAnypntEnvId := "xAnypntEnvId_example" // string | The environment id where the application is deployed
    xAnypntOrgId := "xAnypntOrgId_example" // string | The org id where the application is deployed
    appInfoJson := "appInfoJson_example" // string | A JSON string representing the deployment configuration. Look at \\\"DeploymentInfo\\\" for structure.  (optional)
    file := os.NewFile(1234, "some_file") // *os.File | The Mule application artifact to be deployed. (optional)
    autoStart := true // bool | Indicates whether the application should be automatically started after deployment. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PostDeployment(context.Background()).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).AppInfoJson(appInfoJson).File(file).AutoStart(autoStart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDeployment`: DeploymentDetails
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostDeployment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnypntEnvId** | **string** | The environment id where the application is deployed | 
 **xAnypntOrgId** | **string** | The org id where the application is deployed | 
 **appInfoJson** | **string** | A JSON string representing the deployment configuration. Look at \\\&quot;DeploymentInfo\\\&quot; for structure.  | 
 **file** | ***os.File** | The Mule application artifact to be deployed. | 
 **autoStart** | **bool** | Indicates whether the application should be automatically started after deployment. | 

### Return type

[**DeploymentDetails**](DeploymentDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDeployment

> DeploymentDetails PutDeployment(ctx, domain).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).AppInfoJson(appInfoJson).File(file).Execute()

Update a single application.



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
    xAnypntEnvId := "xAnypntEnvId_example" // string | The environment id where the application is deployed
    xAnypntOrgId := "xAnypntOrgId_example" // string | The org id where the application is deployed
    domain := "domain_example" // string | The application domain in the path
    appInfoJson := "appInfoJson_example" // string | A JSON string representing the deployment configuration. Look at \\\"DeploymentInfo\\\" for structure.  (optional)
    file := os.NewFile(1234, "some_file") // *os.File | The Mule application artifact to be deployed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PutDeployment(context.Background(), domain).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).AppInfoJson(appInfoJson).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutDeployment`: DeploymentDetails
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PutDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The application domain in the path | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnypntEnvId** | **string** | The environment id where the application is deployed | 
 **xAnypntOrgId** | **string** | The org id where the application is deployed | 

 **appInfoJson** | **string** | A JSON string representing the deployment configuration. Look at \\\&quot;DeploymentInfo\\\&quot; for structure.  | 
 **file** | ***os.File** | The Mule application artifact to be deployed. | 

### Return type

[**DeploymentDetails**](DeploymentDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDeployments

> PutDeployments(ctx).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).UpdateDeploymentBulkBody(updateDeploymentBulkBody).Execute()

Updates deployment states in bulk



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
    xAnypntEnvId := "xAnypntEnvId_example" // string | The environment id where the application is deployed
    xAnypntOrgId := "xAnypntOrgId_example" // string | The org id where the application is deployed
    updateDeploymentBulkBody := *openapiclient.NewUpdateDeploymentBulkBody() // UpdateDeploymentBulkBody | Application deployment body containing configuration and artifact file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.PutDeployments(context.Background()).XAnypntEnvId(xAnypntEnvId).XAnypntOrgId(xAnypntOrgId).UpdateDeploymentBulkBody(updateDeploymentBulkBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnypntEnvId** | **string** | The environment id where the application is deployed | 
 **xAnypntOrgId** | **string** | The org id where the application is deployed | 
 **updateDeploymentBulkBody** | [**UpdateDeploymentBulkBody**](UpdateDeploymentBulkBody.md) | Application deployment body containing configuration and artifact file. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

