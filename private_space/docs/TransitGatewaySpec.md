# TransitGatewaySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceShare** | Pointer to [**TransitGatewaySpecResourceShare**](TransitGatewaySpecResourceShare.md) |  | [optional] 
**Region** | Pointer to **string** | AWS region (inherits the private space region). | [optional] 
**SpaceName** | Pointer to **string** | Private space display name. | [optional] 
**Ch2AwsAccountId** | Pointer to **string** | CloudHub 2.0 AWS account id of the transit gateway (organization list only). | [optional] 

## Methods

### NewTransitGatewaySpec

`func NewTransitGatewaySpec() *TransitGatewaySpec`

NewTransitGatewaySpec instantiates a new TransitGatewaySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewaySpecWithDefaults

`func NewTransitGatewaySpecWithDefaults() *TransitGatewaySpec`

NewTransitGatewaySpecWithDefaults instantiates a new TransitGatewaySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceShare

`func (o *TransitGatewaySpec) GetResourceShare() TransitGatewaySpecResourceShare`

GetResourceShare returns the ResourceShare field if non-nil, zero value otherwise.

### GetResourceShareOk

`func (o *TransitGatewaySpec) GetResourceShareOk() (*TransitGatewaySpecResourceShare, bool)`

GetResourceShareOk returns a tuple with the ResourceShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceShare

`func (o *TransitGatewaySpec) SetResourceShare(v TransitGatewaySpecResourceShare)`

SetResourceShare sets ResourceShare field to given value.

### HasResourceShare

`func (o *TransitGatewaySpec) HasResourceShare() bool`

HasResourceShare returns a boolean if a field has been set.

### GetRegion

`func (o *TransitGatewaySpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TransitGatewaySpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TransitGatewaySpec) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TransitGatewaySpec) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSpaceName

`func (o *TransitGatewaySpec) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *TransitGatewaySpec) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *TransitGatewaySpec) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *TransitGatewaySpec) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetCh2AwsAccountId

`func (o *TransitGatewaySpec) GetCh2AwsAccountId() string`

GetCh2AwsAccountId returns the Ch2AwsAccountId field if non-nil, zero value otherwise.

### GetCh2AwsAccountIdOk

`func (o *TransitGatewaySpec) GetCh2AwsAccountIdOk() (*string, bool)`

GetCh2AwsAccountIdOk returns a tuple with the Ch2AwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCh2AwsAccountId

`func (o *TransitGatewaySpec) SetCh2AwsAccountId(v string)`

SetCh2AwsAccountId sets Ch2AwsAccountId field to given value.

### HasCh2AwsAccountId

`func (o *TransitGatewaySpec) HasCh2AwsAccountId() bool`

HasCh2AwsAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


