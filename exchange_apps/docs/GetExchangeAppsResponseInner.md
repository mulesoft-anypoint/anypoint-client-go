# GetExchangeAppsResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**MasterOrganizationId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**ClientProvider** | Pointer to [**PostExchangeAppsReponseClientProvider**](PostExchangeAppsReponseClientProvider.md) |  | [optional] 
**GrantTypes** | Pointer to **[]string** |  | [optional] 
**RedirectUri** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetExchangeAppsResponseInner

`func NewGetExchangeAppsResponseInner() *GetExchangeAppsResponseInner`

NewGetExchangeAppsResponseInner instantiates a new GetExchangeAppsResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExchangeAppsResponseInnerWithDefaults

`func NewGetExchangeAppsResponseInnerWithDefaults() *GetExchangeAppsResponseInner`

NewGetExchangeAppsResponseInnerWithDefaults instantiates a new GetExchangeAppsResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetExchangeAppsResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetExchangeAppsResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetExchangeAppsResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetExchangeAppsResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetExchangeAppsResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetExchangeAppsResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetExchangeAppsResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetExchangeAppsResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *GetExchangeAppsResponseInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetExchangeAppsResponseInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetExchangeAppsResponseInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetExchangeAppsResponseInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *GetExchangeAppsResponseInner) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *GetExchangeAppsResponseInner) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetMasterOrganizationId

`func (o *GetExchangeAppsResponseInner) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *GetExchangeAppsResponseInner) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *GetExchangeAppsResponseInner) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *GetExchangeAppsResponseInner) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetId

`func (o *GetExchangeAppsResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetExchangeAppsResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetExchangeAppsResponseInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetExchangeAppsResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientId

`func (o *GetExchangeAppsResponseInner) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetExchangeAppsResponseInner) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetExchangeAppsResponseInner) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetExchangeAppsResponseInner) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *GetExchangeAppsResponseInner) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *GetExchangeAppsResponseInner) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *GetExchangeAppsResponseInner) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *GetExchangeAppsResponseInner) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientProvider

`func (o *GetExchangeAppsResponseInner) GetClientProvider() PostExchangeAppsReponseClientProvider`

GetClientProvider returns the ClientProvider field if non-nil, zero value otherwise.

### GetClientProviderOk

`func (o *GetExchangeAppsResponseInner) GetClientProviderOk() (*PostExchangeAppsReponseClientProvider, bool)`

GetClientProviderOk returns a tuple with the ClientProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProvider

`func (o *GetExchangeAppsResponseInner) SetClientProvider(v PostExchangeAppsReponseClientProvider)`

SetClientProvider sets ClientProvider field to given value.

### HasClientProvider

`func (o *GetExchangeAppsResponseInner) HasClientProvider() bool`

HasClientProvider returns a boolean if a field has been set.

### GetGrantTypes

`func (o *GetExchangeAppsResponseInner) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *GetExchangeAppsResponseInner) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *GetExchangeAppsResponseInner) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *GetExchangeAppsResponseInner) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetRedirectUri

`func (o *GetExchangeAppsResponseInner) GetRedirectUri() []string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *GetExchangeAppsResponseInner) GetRedirectUriOk() (*[]string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *GetExchangeAppsResponseInner) SetRedirectUri(v []string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *GetExchangeAppsResponseInner) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


