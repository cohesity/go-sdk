# UpgradeClusterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Specifies a message describing the result of the request. | [optional] 
**StatusUrl** | Pointer to **NullableString** | Specifies the URL that can be queried to get the status of the operation once it has begun. | [optional] 

## Methods

### NewUpgradeClusterResult

`func NewUpgradeClusterResult() *UpgradeClusterResult`

NewUpgradeClusterResult instantiates a new UpgradeClusterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeClusterResultWithDefaults

`func NewUpgradeClusterResultWithDefaults() *UpgradeClusterResult`

NewUpgradeClusterResultWithDefaults instantiates a new UpgradeClusterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpgradeClusterResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpgradeClusterResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpgradeClusterResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpgradeClusterResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *UpgradeClusterResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *UpgradeClusterResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetStatusUrl

`func (o *UpgradeClusterResult) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *UpgradeClusterResult) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *UpgradeClusterResult) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.

### HasStatusUrl

`func (o *UpgradeClusterResult) HasStatusUrl() bool`

HasStatusUrl returns a boolean if a field has been set.

### SetStatusUrlNil

`func (o *UpgradeClusterResult) SetStatusUrlNil(b bool)`

 SetStatusUrlNil sets the value for StatusUrl to be an explicit nil

### UnsetStatusUrl
`func (o *UpgradeClusterResult) UnsetStatusUrl()`

UnsetStatusUrl ensures that no value is present for StatusUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


