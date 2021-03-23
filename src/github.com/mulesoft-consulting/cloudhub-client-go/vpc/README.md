# Go API client for vpc

Description of the VPC API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./vpc"
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
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://anypoint.mulesoft.com/cloudhub/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**OrganizationsOrgIdVpcsGet**](docs/DefaultApi.md#organizationsorgidvpcsget) | **Get** /organizations/{orgId}/vpcs | Returns a list of vpcs.
*DefaultApi* | [**OrganizationsOrgIdVpcsPost**](docs/DefaultApi.md#organizationsorgidvpcspost) | **Post** /organizations/{orgId}/vpcs | Creates a new VPC in the provided organization.
*DefaultApi* | [**OrganizationsOrgIdVpcsVpcIdDelete**](docs/DefaultApi.md#organizationsorgidvpcsvpciddelete) | **Delete** /organizations/{orgId}/vpcs/{vpcId} | Delete a VPC by its id.
*DefaultApi* | [**OrganizationsOrgIdVpcsVpcIdGet**](docs/DefaultApi.md#organizationsorgidvpcsvpcidget) | **Get** /organizations/{orgId}/vpcs/{vpcId} | Returns the vpc instance with the given id.
*DefaultApi* | [**OrganizationsOrgIdVpcsVpcIdPut**](docs/DefaultApi.md#organizationsorgidvpcsvpcidput) | **Put** /organizations/{orgId}/vpcs/{vpcId} | Update the VPC configuration


## Documentation For Models

 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse400](docs/InlineResponse400.md)
 - [TheFirstAnyOfSchema](docs/TheFirstAnyOfSchema.md)
 - [TheInternalDnsSchema](docs/TheInternalDnsSchema.md)
 - [Vpc](docs/Vpc.md)
 - [VpcAllOf](docs/VpcAllOf.md)
 - [VpcCore](docs/VpcCore.md)
 - [VpcCoreFirewallRules](docs/VpcCoreFirewallRules.md)


## Documentation For Authorization



### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARERTOKENSTRING")
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


