# CreateUser400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] [default to 400]
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateUser400Response

`func NewCreateUser400Response() *CreateUser400Response`

NewCreateUser400Response instantiates a new CreateUser400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUser400ResponseWithDefaults

`func NewCreateUser400ResponseWithDefaults() *CreateUser400Response`

NewCreateUser400ResponseWithDefaults instantiates a new CreateUser400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateUser400Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateUser400Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateUser400Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateUser400Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *CreateUser400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateUser400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateUser400Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateUser400Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


