# PostExchangeAppsReponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MasterOrganizationId** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**GrantTypes** | Pointer to **[]string** |  | [optional] 
**RedirectUri** | Pointer to **[]string** |  | [optional] 
**ClientProvider** | Pointer to [**PostExchangeAppsReponseClientProvider**](PostExchangeAppsReponseClientProvider.md) |  | [optional] 

## Methods

### NewPostExchangeAppsReponse

`func NewPostExchangeAppsReponse() *PostExchangeAppsReponse`

NewPostExchangeAppsReponse instantiates a new PostExchangeAppsReponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostExchangeAppsReponseWithDefaults

`func NewPostExchangeAppsReponseWithDefaults() *PostExchangeAppsReponse`

NewPostExchangeAppsReponseWithDefaults instantiates a new PostExchangeAppsReponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostExchangeAppsReponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostExchangeAppsReponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostExchangeAppsReponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PostExchangeAppsReponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PostExchangeAppsReponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostExchangeAppsReponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostExchangeAppsReponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostExchangeAppsReponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PostExchangeAppsReponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostExchangeAppsReponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostExchangeAppsReponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostExchangeAppsReponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMasterOrganizationId

`func (o *PostExchangeAppsReponse) GetMasterOrganizationId() string`

GetMasterOrganizationId returns the MasterOrganizationId field if non-nil, zero value otherwise.

### GetMasterOrganizationIdOk

`func (o *PostExchangeAppsReponse) GetMasterOrganizationIdOk() (*string, bool)`

GetMasterOrganizationIdOk returns a tuple with the MasterOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOrganizationId

`func (o *PostExchangeAppsReponse) SetMasterOrganizationId(v string)`

SetMasterOrganizationId sets MasterOrganizationId field to given value.

### HasMasterOrganizationId

`func (o *PostExchangeAppsReponse) HasMasterOrganizationId() bool`

HasMasterOrganizationId returns a boolean if a field has been set.

### GetClientId

`func (o *PostExchangeAppsReponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PostExchangeAppsReponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PostExchangeAppsReponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PostExchangeAppsReponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *PostExchangeAppsReponse) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PostExchangeAppsReponse) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PostExchangeAppsReponse) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PostExchangeAppsReponse) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetGrantTypes

`func (o *PostExchangeAppsReponse) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *PostExchangeAppsReponse) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *PostExchangeAppsReponse) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *PostExchangeAppsReponse) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetRedirectUri

`func (o *PostExchangeAppsReponse) GetRedirectUri() []string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *PostExchangeAppsReponse) GetRedirectUriOk() (*[]string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *PostExchangeAppsReponse) SetRedirectUri(v []string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *PostExchangeAppsReponse) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetClientProvider

`func (o *PostExchangeAppsReponse) GetClientProvider() PostExchangeAppsReponseClientProvider`

GetClientProvider returns the ClientProvider field if non-nil, zero value otherwise.

### GetClientProviderOk

`func (o *PostExchangeAppsReponse) GetClientProviderOk() (*PostExchangeAppsReponseClientProvider, bool)`

GetClientProviderOk returns a tuple with the ClientProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProvider

`func (o *PostExchangeAppsReponse) SetClientProvider(v PostExchangeAppsReponseClientProvider)`

SetClientProvider sets ClientProvider field to given value.

### HasClientProvider

`func (o *PostExchangeAppsReponse) HasClientProvider() bool`

HasClientProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


