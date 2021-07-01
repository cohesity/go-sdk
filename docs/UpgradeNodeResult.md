# UpgradeNodeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Specifies a message describing the result of the request. | [optional] 
**StatusUrl** | Pointer to **NullableString** | Specifies a URL that can be queried to get the status of the operation once it has begun. | [optional] 

## Methods

### NewUpgradeNodeResult

`func NewUpgradeNodeResult() *UpgradeNodeResult`

NewUpgradeNodeResult instantiates a new UpgradeNodeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeNodeResultWithDefaults

`func NewUpgradeNodeResultWithDefaults() *UpgradeNodeResult`

NewUpgradeNodeResultWithDefaults instantiates a new UpgradeNodeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpgradeNodeResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpgradeNodeResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpgradeNodeResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpgradeNodeResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *UpgradeNodeResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *UpgradeNodeResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetStatusUrl

`func (o *UpgradeNodeResult) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *UpgradeNodeResult) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *UpgradeNodeResult) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.

### HasStatusUrl

`func (o *UpgradeNodeResult) HasStatusUrl() bool`

HasStatusUrl returns a boolean if a field has been set.

### SetStatusUrlNil

`func (o *UpgradeNodeResult) SetStatusUrlNil(b bool)`

 SetStatusUrlNil sets the value for StatusUrl to be an explicit nil

### UnsetStatusUrl
`func (o *UpgradeNodeResult) UnsetStatusUrl()`

UnsetStatusUrl ensures that no value is present for StatusUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


