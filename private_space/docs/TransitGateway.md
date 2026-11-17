# TransitGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | AWS transit gateway id (e.g. tgw-017e20b9ce00c865c). | [optional] 
**Name** | Pointer to **string** | User-chosen display name for the attachment. | [optional] 
**AccountId** | Pointer to **string** | MuleSoft AWS account id (region-bound to the private space). | [optional] 
**Spec** | Pointer to [**TransitGatewaySpec**](TransitGatewaySpec.md) |  | [optional] 
**TgwResource** | Pointer to **string** | AWS console URL pointing at the transit gateway resource (organization list only). | [optional] 
**Attachments** | Pointer to [**[]TransitGateway**](TransitGateway.md) | Per-private-space attachments of this transit gateway (organization list only). | [optional] 
**Status** | Pointer to [**TransitGatewayStatus**](TransitGatewayStatus.md) |  | [optional] 

## Methods

### NewTransitGateway

`func NewTransitGateway() *TransitGateway`

NewTransitGateway instantiates a new TransitGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewayWithDefaults

`func NewTransitGatewayWithDefaults() *TransitGateway`

NewTransitGatewayWithDefaults instantiates a new TransitGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransitGateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransitGateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransitGateway) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransitGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TransitGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransitGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransitGateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransitGateway) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *TransitGateway) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransitGateway) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransitGateway) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TransitGateway) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetSpec

`func (o *TransitGateway) GetSpec() TransitGatewaySpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *TransitGateway) GetSpecOk() (*TransitGatewaySpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *TransitGateway) SetSpec(v TransitGatewaySpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *TransitGateway) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetTgwResource

`func (o *TransitGateway) GetTgwResource() string`

GetTgwResource returns the TgwResource field if non-nil, zero value otherwise.

### GetTgwResourceOk

`func (o *TransitGateway) GetTgwResourceOk() (*string, bool)`

GetTgwResourceOk returns a tuple with the TgwResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwResource

`func (o *TransitGateway) SetTgwResource(v string)`

SetTgwResource sets TgwResource field to given value.

### HasTgwResource

`func (o *TransitGateway) HasTgwResource() bool`

HasTgwResource returns a boolean if a field has been set.

### GetAttachments

`func (o *TransitGateway) GetAttachments() []TransitGateway`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TransitGateway) GetAttachmentsOk() (*[]TransitGateway, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TransitGateway) SetAttachments(v []TransitGateway)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TransitGateway) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetStatus

`func (o *TransitGateway) GetStatus() TransitGatewayStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransitGateway) GetStatusOk() (*TransitGatewayStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransitGateway) SetStatus(v TransitGatewayStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransitGateway) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


