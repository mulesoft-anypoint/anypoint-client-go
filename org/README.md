# Go API client for org

Description of the Organization API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.1
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
import org "github.com/mulesoft-anypoint/anypoint-client-go/org"
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
ctx := context.WithValue(context.Background(), org.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), org.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), org.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), org.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CreateBG**](docs/DefaultApi.md#createbg) | **Post** /organizations | Creates a new Business Group.
*DefaultApi* | [**DeleteBG**](docs/DefaultApi.md#deletebg) | **Delete** /organizations/{orgId} | Delete a Business Group by its id.
*DefaultApi* | [**GetBG**](docs/DefaultApi.md#getbg) | **Get** /organizations/{orgId} | Returns the Business Group instance with the given id.
*DefaultApi* | [**GetBGUsage**](docs/DefaultApi.md#getbgusage) | **Get** /cs/organizations/{orgId}/usage | Returns the usage for a given Business Group.
*DefaultApi* | [**GetClientCredentials**](docs/DefaultApi.md#getclientcredentials) | **Get** /organizations/{orgId}/clients/{clientId} | 
*DefaultApi* | [**UpdateBG**](docs/DefaultApi.md#updatebg) | **Put** /organizations/{orgId} | Update a Business Group by its id.


## Documentation For Models

 - [AddOnEntitlement](docs/AddOnEntitlement.md)
 - [AngGovernanceEntitlement](docs/AngGovernanceEntitlement.md)
 - [Api](docs/Api.md)
 - [ApiGovernanceDomainEntitlement](docs/ApiGovernanceDomainEntitlement.md)
 - [ApiGovernanceEntitlement](docs/ApiGovernanceEntitlement.md)
 - [ApiManagerEntitlement](docs/ApiManagerEntitlement.md)
 - [ApiMonitoringEntitlement](docs/ApiMonitoringEntitlement.md)
 - [ApiQueryEntitlement](docs/ApiQueryEntitlement.md)
 - [BGCore](docs/BGCore.md)
 - [BGPostReqBody](docs/BGPostReqBody.md)
 - [BGUpdateReqBody](docs/BGUpdateReqBody.md)
 - [Client](docs/Client.md)
 - [ComposerEntitlement](docs/ComposerEntitlement.md)
 - [CrowdEntitlement](docs/CrowdEntitlement.md)
 - [DesignCenterEntitlement](docs/DesignCenterEntitlement.md)
 - [EnablementResourceEntitlement](docs/EnablementResourceEntitlement.md)
 - [Entitlements](docs/Entitlements.md)
 - [EntitlementsCore](docs/EntitlementsCore.md)
 - [Environment](docs/Environment.md)
 - [Exchange2Entitlement](docs/Exchange2Entitlement.md)
 - [Exchange2EntitlementAssetUsageAndEngagement](docs/Exchange2EntitlementAssetUsageAndEngagement.md)
 - [GatewaysEntitlement](docs/GatewaysEntitlement.md)
 - [Governance](docs/Governance.md)
 - [HighAvailabilityEntitlement](docs/HighAvailabilityEntitlement.md)
 - [LoadBalancerEntitlement](docs/LoadBalancerEntitlement.md)
 - [ManagedGatewayLargeEntitlement](docs/ManagedGatewayLargeEntitlement.md)
 - [ManagedGatewaySmallEntitlement](docs/ManagedGatewaySmallEntitlement.md)
 - [MasterBGDetail](docs/MasterBGDetail.md)
 - [MasterBGSpecificDetails](docs/MasterBGSpecificDetails.md)
 - [MessagingEntitlement](docs/MessagingEntitlement.md)
 - [MonitoringCenterEntitlement](docs/MonitoringCenterEntitlement.md)
 - [MuleRuntimeIntegration](docs/MuleRuntimeIntegration.md)
 - [NetworkConnectionsEntitlement](docs/NetworkConnectionsEntitlement.md)
 - [PartnersProductionEntitlement](docs/PartnersProductionEntitlement.md)
 - [PartnersSandboxEntitlement](docs/PartnersSandboxEntitlement.md)
 - [Production](docs/Production.md)
 - [RpaEntitlement](docs/RpaEntitlement.md)
 - [Sandbox](docs/Sandbox.md)
 - [StaticIpsEntitlement](docs/StaticIpsEntitlement.md)
 - [Subscription](docs/Subscription.md)
 - [TradingPartnersProductionEntitlement](docs/TradingPartnersProductionEntitlement.md)
 - [TradingPartnersSandboxEntitlement](docs/TradingPartnersSandboxEntitlement.md)
 - [Usage](docs/Usage.md)
 - [UsageBasedPricing](docs/UsageBasedPricing.md)
 - [UsageBasedPricingEntitlement](docs/UsageBasedPricingEntitlement.md)
 - [UsageBasedPricingLimitsEntitlement](docs/UsageBasedPricingLimitsEntitlement.md)
 - [UsageBasedPricingLimitsEntitlementCpu](docs/UsageBasedPricingLimitsEntitlementCpu.md)
 - [User](docs/User.md)
 - [VCoresDesignEntitlement](docs/VCoresDesignEntitlement.md)
 - [VCoresProductionEntitlement](docs/VCoresProductionEntitlement.md)
 - [VCoresSandboxEntitlement](docs/VCoresSandboxEntitlement.md)
 - [VpcsEntitlement](docs/VpcsEntitlement.md)
 - [VpnsEntitlement](docs/VpnsEntitlement.md)
 - [WorkerCloudsEntitlement](docs/WorkerCloudsEntitlement.md)


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



