# Go API client for application_manager_v2

The Application Manager API exists to provide Mule Application management operations from Anypoint Control Planes to any Runtime Plane, with capabilities that include:
  - Deploying a Mule Application or API to a Mule Runtime
  - Scaling up or down a running application
  - Managing application settings (ie: properties)
  - Deleting a Mule Application Deployment
  - Changing artifact version or configurations of a deployment
  - Starting, Stopping or Restarting applications
  - etc.
This API currently supports deployments to Runtime Fabric and CloudHub 2.0 targets only.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import application_manager_v2 "github.com/mulesoft-anypoint/anypoint-client-go/application_manager_v2"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), application_manager_v2.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), application_manager_v2.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), application_manager_v2.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), application_manager_v2.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://anypoint.mulesoft.com/amc/application-manager/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**DeleteDeployment**](docs/DefaultApi.md#deletedeployment) | **Delete** /organizations/{orgId}/environments/{envId}/deployments/{deploymentId} | Delete deployment
*DefaultApi* | [**GetAllDeployments**](docs/DefaultApi.md#getalldeployments) | **Get** /organizations/{orgId}/environments/{envId}/deployments | List deployments
*DefaultApi* | [**GetDeploymentById**](docs/DefaultApi.md#getdeploymentbyid) | **Get** /organizations/{orgId}/environments/{envId}/deployments/{deploymentId} | Get a deploymnt
*DefaultApi* | [**PatchDeployment**](docs/DefaultApi.md#patchdeployment) | **Patch** /organizations/{orgId}/environments/{envId}/deployments/{deploymentId} | Patch a deployment
*DefaultApi* | [**PostDeployment**](docs/DefaultApi.md#postdeployment) | **Post** /organizations/{orgId}/environments/{envId}/deployments | Create deployment


## Documentation For Models

 - [AnypointMonitoring](docs/AnypointMonitoring.md)
 - [AnypointMonitoringResources](docs/AnypointMonitoringResources.md)
 - [AnypointMonitoringResourcesCpu](docs/AnypointMonitoringResourcesCpu.md)
 - [AppConfiguration](docs/AppConfiguration.md)
 - [Application](docs/Application.md)
 - [ApplicationIntegrations](docs/ApplicationIntegrations.md)
 - [Autoscaling](docs/Autoscaling.md)
 - [Deployment](docs/Deployment.md)
 - [DeploymentItem](docs/DeploymentItem.md)
 - [DeploymentItemApplication](docs/DeploymentItemApplication.md)
 - [DeploymentItemTarget](docs/DeploymentItemTarget.md)
 - [DeploymentRequestBody](docs/DeploymentRequestBody.md)
 - [DeploymentResponsePaged](docs/DeploymentResponsePaged.md)
 - [DeploymentSettings](docs/DeploymentSettings.md)
 - [Http](docs/Http.md)
 - [HttpInbound](docs/HttpInbound.md)
 - [Jvm](docs/Jvm.md)
 - [MuleAgentAppPropService](docs/MuleAgentAppPropService.md)
 - [MuleAgentLoggingService](docs/MuleAgentLoggingService.md)
 - [MuleAgentSchedulingService](docs/MuleAgentSchedulingService.md)
 - [ObjectStoreV2](docs/ObjectStoreV2.md)
 - [Ref](docs/Ref.md)
 - [Replicas](docs/Replicas.md)
 - [Resources](docs/Resources.md)
 - [ResourcesCpu](docs/ResourcesCpu.md)
 - [ResourcesMemory](docs/ResourcesMemory.md)
 - [ResourcesStorage](docs/ResourcesStorage.md)
 - [Runtime](docs/Runtime.md)
 - [Scheduler](docs/Scheduler.md)
 - [ScopeLoggingConfiguration](docs/ScopeLoggingConfiguration.md)
 - [Services](docs/Services.md)
 - [Sidecars](docs/Sidecars.md)
 - [State](docs/State.md)
 - [Target](docs/Target.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



