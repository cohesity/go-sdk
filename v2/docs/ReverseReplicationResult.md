# ReverseReplicationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorReason** | Pointer to **NullableString** | Specifies the reason of not enabling reverse replication. | [optional] 
**IsReverseReplicationEnabled** | Pointer to **NullableBool** | Specifies whether the reverse replication was enabled or not during group creation. It can be false, if source cluster is not reachable for reverse replication. | [optional] 

## Methods

### NewReverseReplicationResult

`func NewReverseReplicationResult() *ReverseReplicationResult`

NewReverseReplicationResult instantiates a new ReverseReplicationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseReplicationResultWithDefaults

`func NewReverseReplicationResultWithDefaults() *ReverseReplicationResult`

NewReverseReplicationResultWithDefaults instantiates a new ReverseReplicationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorReason

`func (o *ReverseReplicationResult) GetErrorReason() string`

GetErrorReason returns the ErrorReason field if non-nil, zero value otherwise.

### GetErrorReasonOk

`func (o *ReverseReplicationResult) GetErrorReasonOk() (*string, bool)`

GetErrorReasonOk returns a tuple with the ErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReason

`func (o *ReverseReplicationResult) SetErrorReason(v string)`

SetErrorReason sets ErrorReason field to given value.

### HasErrorReason

`func (o *ReverseReplicationResult) HasErrorReason() bool`

HasErrorReason returns a boolean if a field has been set.

### SetErrorReasonNil

`func (o *ReverseReplicationResult) SetErrorReasonNil(b bool)`

 SetErrorReasonNil sets the value for ErrorReason to be an explicit nil

### UnsetErrorReason
`func (o *ReverseReplicationResult) UnsetErrorReason()`

UnsetErrorReason ensures that no value is present for ErrorReason, not even an explicit nil
### GetIsReverseReplicationEnabled

`func (o *ReverseReplicationResult) GetIsReverseReplicationEnabled() bool`

GetIsReverseReplicationEnabled returns the IsReverseReplicationEnabled field if non-nil, zero value otherwise.

### GetIsReverseReplicationEnabledOk

`func (o *ReverseReplicationResult) GetIsReverseReplicationEnabledOk() (*bool, bool)`

GetIsReverseReplicationEnabledOk returns a tuple with the IsReverseReplicationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReverseReplicationEnabled

`func (o *ReverseReplicationResult) SetIsReverseReplicationEnabled(v bool)`

SetIsReverseReplicationEnabled sets IsReverseReplicationEnabled field to given value.

### HasIsReverseReplicationEnabled

`func (o *ReverseReplicationResult) HasIsReverseReplicationEnabled() bool`

HasIsReverseReplicationEnabled returns a boolean if a field has been set.

### SetIsReverseReplicationEnabledNil

`func (o *ReverseReplicationResult) SetIsReverseReplicationEnabledNil(b bool)`

 SetIsReverseReplicationEnabledNil sets the value for IsReverseReplicationEnabled to be an explicit nil

### UnsetIsReverseReplicationEnabled
`func (o *ReverseReplicationResult) UnsetIsReverseReplicationEnabled()`

UnsetIsReverseReplicationEnabled ensures that no value is present for IsReverseReplicationEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


