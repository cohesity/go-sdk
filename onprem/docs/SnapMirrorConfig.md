# SnapMirrorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewId** | Pointer to **NullableInt64** | Specifies the Id of the S3 view where data need to be written. | [optional] 
**IncrementalPrefix** | Pointer to **NullableString** | Specifies the incremental snapshot prefix value. | [optional] 

## Methods

### NewSnapMirrorConfig

`func NewSnapMirrorConfig() *SnapMirrorConfig`

NewSnapMirrorConfig instantiates a new SnapMirrorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapMirrorConfigWithDefaults

`func NewSnapMirrorConfigWithDefaults() *SnapMirrorConfig`

NewSnapMirrorConfigWithDefaults instantiates a new SnapMirrorConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewId

`func (o *SnapMirrorConfig) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *SnapMirrorConfig) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *SnapMirrorConfig) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *SnapMirrorConfig) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *SnapMirrorConfig) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *SnapMirrorConfig) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
### GetIncrementalPrefix

`func (o *SnapMirrorConfig) GetIncrementalPrefix() string`

GetIncrementalPrefix returns the IncrementalPrefix field if non-nil, zero value otherwise.

### GetIncrementalPrefixOk

`func (o *SnapMirrorConfig) GetIncrementalPrefixOk() (*string, bool)`

GetIncrementalPrefixOk returns a tuple with the IncrementalPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalPrefix

`func (o *SnapMirrorConfig) SetIncrementalPrefix(v string)`

SetIncrementalPrefix sets IncrementalPrefix field to given value.

### HasIncrementalPrefix

`func (o *SnapMirrorConfig) HasIncrementalPrefix() bool`

HasIncrementalPrefix returns a boolean if a field has been set.

### SetIncrementalPrefixNil

`func (o *SnapMirrorConfig) SetIncrementalPrefixNil(b bool)`

 SetIncrementalPrefixNil sets the value for IncrementalPrefix to be an explicit nil

### UnsetIncrementalPrefix
`func (o *SnapMirrorConfig) UnsetIncrementalPrefix()`

UnsetIncrementalPrefix ensures that no value is present for IncrementalPrefix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


