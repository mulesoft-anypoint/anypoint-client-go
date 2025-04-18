# Go API client for idp

Description of Identity Provider API

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
import idp "github.com/mulesoft-anypoint/anypoint-client-go/idp"
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
ctx := context.WithValue(context.Background(), idp.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), idp.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), idp.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), idp.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**OrganizationsOrgIdIdentityProvidersGet**](docs/DefaultApi.md#organizationsorgididentityprovidersget) | **Get** /organizations/{orgId}/identityProviders | 
*DefaultApi* | [**OrganizationsOrgIdIdentityProvidersIdpIdDelete**](docs/DefaultApi.md#organizationsorgididentityprovidersidpiddelete) | **Delete** /organizations/{orgId}/identityProviders/{idpId} | 
*DefaultApi* | [**OrganizationsOrgIdIdentityProvidersIdpIdGet**](docs/DefaultApi.md#organizationsorgididentityprovidersidpidget) | **Get** /organizations/{orgId}/identityProviders/{idpId} | 
*DefaultApi* | [**OrganizationsOrgIdIdentityProvidersIdpIdPatch**](docs/DefaultApi.md#organizationsorgididentityprovidersidpidpatch) | **Patch** /organizations/{orgId}/identityProviders/{idpId} | 
*DefaultApi* | [**OrganizationsOrgIdIdentityProvidersPost**](docs/DefaultApi.md#organizationsorgididentityproviderspost) | **Post** /organizations/{orgId}/identityProviders | 


## Documentation For Models

 - [Idp](docs/Idp.md)
 - [IdpPatchBody](docs/IdpPatchBody.md)
 - [IdpPostBody](docs/IdpPostBody.md)
 - [IdpSummary](docs/IdpSummary.md)
 - [IdpSummaryType](docs/IdpSummaryType.md)
 - [LdapProviderGet](docs/LdapProviderGet.md)
 - [LdapProviderPatch](docs/LdapProviderPatch.md)
 - [LdapProviderPatchConnection](docs/LdapProviderPatchConnection.md)
 - [LdapProviderPatchFilters](docs/LdapProviderPatchFilters.md)
 - [LdapProviderPatchGroupMapping](docs/LdapProviderPatchGroupMapping.md)
 - [LdapProviderPatchSearchBases](docs/LdapProviderPatchSearchBases.md)
 - [LdapProviderPatchUserMapping](docs/LdapProviderPatchUserMapping.md)
 - [LdapProviderPostBody](docs/LdapProviderPostBody.md)
 - [LdapProviderPostBodyConnection](docs/LdapProviderPostBodyConnection.md)
 - [LdapProviderPostBodyFilters](docs/LdapProviderPostBodyFilters.md)
 - [LdapProviderPostBodyGroupMapping](docs/LdapProviderPostBodyGroupMapping.md)
 - [LdapProviderPostBodySearchBases](docs/LdapProviderPostBodySearchBases.md)
 - [LdapProviderPostBodyType](docs/LdapProviderPostBodyType.md)
 - [LdapProviderPostBodyUserMapping](docs/LdapProviderPostBodyUserMapping.md)
 - [OpenIDProviderDynamicPostBody](docs/OpenIDProviderDynamicPostBody.md)
 - [OpenIDProviderDynamicPostBodyOidcProvider](docs/OpenIDProviderDynamicPostBodyOidcProvider.md)
 - [OpenIDProviderDynamicPostBodyOidcProviderClient](docs/OpenIDProviderDynamicPostBodyOidcProviderClient.md)
 - [OpenIDProviderDynamicPostBodyOidcProviderClientMetadata](docs/OpenIDProviderDynamicPostBodyOidcProviderClientMetadata.md)
 - [OpenIDProviderDynamicPostBodyOidcProviderClientRegistration](docs/OpenIDProviderDynamicPostBodyOidcProviderClientRegistration.md)
 - [OpenIDProviderDynamicPostBodyOidcProviderClientUrls](docs/OpenIDProviderDynamicPostBodyOidcProviderClientUrls.md)
 - [OpenIDProviderGet](docs/OpenIDProviderGet.md)
 - [OpenIDProviderGetOidcProvider](docs/OpenIDProviderGetOidcProvider.md)
 - [OpenIDProviderGetOidcProviderClient](docs/OpenIDProviderGetOidcProviderClient.md)
 - [OpenIDProviderGetOidcProviderClientCredentials](docs/OpenIDProviderGetOidcProviderClientCredentials.md)
 - [OpenIDProviderGetOidcProviderUrls](docs/OpenIDProviderGetOidcProviderUrls.md)
 - [OpenIDProviderGetServiceProvider](docs/OpenIDProviderGetServiceProvider.md)
 - [OpenIDProviderGetServiceProviderUrls](docs/OpenIDProviderGetServiceProviderUrls.md)
 - [OpenIDProviderManualPostBody](docs/OpenIDProviderManualPostBody.md)
 - [OpenIDProviderManualPostBodyOidcProvider](docs/OpenIDProviderManualPostBodyOidcProvider.md)
 - [OpenIDProviderManualPostBodyOidcProviderClaimsMapping](docs/OpenIDProviderManualPostBodyOidcProviderClaimsMapping.md)
 - [OpenIDProviderManualPostBodyOidcProviderClient](docs/OpenIDProviderManualPostBodyOidcProviderClient.md)
 - [OpenIDProviderManualPostBodyOidcProviderClientCredentials](docs/OpenIDProviderManualPostBodyOidcProviderClientCredentials.md)
 - [OpenIDProviderManualPostBodyOidcProviderUrls](docs/OpenIDProviderManualPostBodyOidcProviderUrls.md)
 - [OpenIDProviderManualPostBodyType](docs/OpenIDProviderManualPostBodyType.md)
 - [OpenIDProviderPatch](docs/OpenIDProviderPatch.md)
 - [OpenIDProviderPatchOidcProvider](docs/OpenIDProviderPatchOidcProvider.md)
 - [OpenIDProviderPatchOidcProviderClient](docs/OpenIDProviderPatchOidcProviderClient.md)
 - [OpenIDProviderPatchOidcProviderClientCredentials](docs/OpenIDProviderPatchOidcProviderClientCredentials.md)
 - [OpenIDProviderPatchOidcProviderClientRegistration](docs/OpenIDProviderPatchOidcProviderClientRegistration.md)
 - [OpenIDProviderPatchOidcProviderClientUrls](docs/OpenIDProviderPatchOidcProviderClientUrls.md)
 - [OpenIDProviderPatchOidcProviderUrls](docs/OpenIDProviderPatchOidcProviderUrls.md)
 - [OrganizationsOrgIdIdentityProvidersGet200Response](docs/OrganizationsOrgIdIdentityProvidersGet200Response.md)
 - [OrganizationsOrgIdIdentityProvidersIdpIdPatch400Response](docs/OrganizationsOrgIdIdentityProvidersIdpIdPatch400Response.md)
 - [OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner](docs/OrganizationsOrgIdIdentityProvidersIdpIdPatch400ResponseErrorsInner.md)
 - [SamlProviderGet](docs/SamlProviderGet.md)
 - [SamlProviderPatch](docs/SamlProviderPatch.md)
 - [SamlProviderPatchSaml](docs/SamlProviderPatchSaml.md)
 - [SamlProviderPatchServiceProvider](docs/SamlProviderPatchServiceProvider.md)
 - [SamlProviderPatchServiceProviderUrls](docs/SamlProviderPatchServiceProviderUrls.md)
 - [SamlProviderPatchType](docs/SamlProviderPatchType.md)
 - [SamlProviderPostBody](docs/SamlProviderPostBody.md)
 - [SamlProviderPostBodySaml](docs/SamlProviderPostBodySaml.md)
 - [SamlProviderPostBodySamlClaimsMapping](docs/SamlProviderPostBodySamlClaimsMapping.md)
 - [SamlProviderPostBodyServiceProvider](docs/SamlProviderPostBodyServiceProvider.md)
 - [SamlProviderPostBodyServiceProviderUrls](docs/SamlProviderPostBodyServiceProviderUrls.md)
 - [SamlProviderPostBodyType](docs/SamlProviderPostBodyType.md)


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



