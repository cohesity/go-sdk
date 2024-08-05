# HostSettingCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **NullableString** | Specifies the result of host checking performed by agent. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of host checking that was performed. | [optional] 

## Methods

### NewHostSettingCheck

`func NewHostSettingCheck() *HostSettingCheck`

NewHostSettingCheck instantiates a new HostSettingCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostSettingCheckWithDefaults

`func NewHostSettingCheckWithDefaults() *HostSettingCheck`

NewHostSettingCheckWithDefaults instantiates a new HostSettingCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *HostSettingCheck) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *HostSettingCheck) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *HostSettingCheck) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *HostSettingCheck) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *HostSettingCheck) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *HostSettingCheck) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetType

`func (o *HostSettingCheck) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HostSettingCheck) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HostSettingCheck) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HostSettingCheck) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HostSettingCheck) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HostSettingCheck) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


