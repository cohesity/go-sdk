# MarkBaseosUpgradeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Specifies optional message related to operation status. | [optional] 
**SetOperation** | **NullableBool** | Specifies whether the operation is set or not. | 

## Methods

### NewMarkBaseosUpgradeInfo

`func NewMarkBaseosUpgradeInfo(setOperation NullableBool, ) *MarkBaseosUpgradeInfo`

NewMarkBaseosUpgradeInfo instantiates a new MarkBaseosUpgradeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarkBaseosUpgradeInfoWithDefaults

`func NewMarkBaseosUpgradeInfoWithDefaults() *MarkBaseosUpgradeInfo`

NewMarkBaseosUpgradeInfoWithDefaults instantiates a new MarkBaseosUpgradeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MarkBaseosUpgradeInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MarkBaseosUpgradeInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MarkBaseosUpgradeInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MarkBaseosUpgradeInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MarkBaseosUpgradeInfo) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MarkBaseosUpgradeInfo) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSetOperation

`func (o *MarkBaseosUpgradeInfo) GetSetOperation() bool`

GetSetOperation returns the SetOperation field if non-nil, zero value otherwise.

### GetSetOperationOk

`func (o *MarkBaseosUpgradeInfo) GetSetOperationOk() (*bool, bool)`

GetSetOperationOk returns a tuple with the SetOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetOperation

`func (o *MarkBaseosUpgradeInfo) SetSetOperation(v bool)`

SetSetOperation sets SetOperation field to given value.


### SetSetOperationNil

`func (o *MarkBaseosUpgradeInfo) SetSetOperationNil(b bool)`

 SetSetOperationNil sets the value for SetOperation to be an explicit nil

### UnsetSetOperation
`func (o *MarkBaseosUpgradeInfo) UnsetSetOperation()`

UnsetSetOperation ensures that no value is present for SetOperation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


