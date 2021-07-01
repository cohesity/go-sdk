# UpdateProtectionObjectParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PauseBackup** | Pointer to **NullableBool** | Specifies if the protection for the Protection Object is to be paused. | [optional] 
**ProtectedSourceUid** | [**UniversalId**](UniversalId.md) |  | 
**RpoPolicyId** | Pointer to **NullableString** | Specifies the unique id of the new RPO policy to assign to the object. | [optional] 
**SourceParameters** | Pointer to [**[]SourceSpecialParameter**](SourceSpecialParameter.md) | Specifies the additional special settings for a Protected Source. | [optional] 

## Methods

### NewUpdateProtectionObjectParameters

`func NewUpdateProtectionObjectParameters(protectedSourceUid UniversalId, ) *UpdateProtectionObjectParameters`

NewUpdateProtectionObjectParameters instantiates a new UpdateProtectionObjectParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProtectionObjectParametersWithDefaults

`func NewUpdateProtectionObjectParametersWithDefaults() *UpdateProtectionObjectParameters`

NewUpdateProtectionObjectParametersWithDefaults instantiates a new UpdateProtectionObjectParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPauseBackup

`func (o *UpdateProtectionObjectParameters) GetPauseBackup() bool`

GetPauseBackup returns the PauseBackup field if non-nil, zero value otherwise.

### GetPauseBackupOk

`func (o *UpdateProtectionObjectParameters) GetPauseBackupOk() (*bool, bool)`

GetPauseBackupOk returns a tuple with the PauseBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseBackup

`func (o *UpdateProtectionObjectParameters) SetPauseBackup(v bool)`

SetPauseBackup sets PauseBackup field to given value.

### HasPauseBackup

`func (o *UpdateProtectionObjectParameters) HasPauseBackup() bool`

HasPauseBackup returns a boolean if a field has been set.

### SetPauseBackupNil

`func (o *UpdateProtectionObjectParameters) SetPauseBackupNil(b bool)`

 SetPauseBackupNil sets the value for PauseBackup to be an explicit nil

### UnsetPauseBackup
`func (o *UpdateProtectionObjectParameters) UnsetPauseBackup()`

UnsetPauseBackup ensures that no value is present for PauseBackup, not even an explicit nil
### GetProtectedSourceUid

`func (o *UpdateProtectionObjectParameters) GetProtectedSourceUid() UniversalId`

GetProtectedSourceUid returns the ProtectedSourceUid field if non-nil, zero value otherwise.

### GetProtectedSourceUidOk

`func (o *UpdateProtectionObjectParameters) GetProtectedSourceUidOk() (*UniversalId, bool)`

GetProtectedSourceUidOk returns a tuple with the ProtectedSourceUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedSourceUid

`func (o *UpdateProtectionObjectParameters) SetProtectedSourceUid(v UniversalId)`

SetProtectedSourceUid sets ProtectedSourceUid field to given value.


### GetRpoPolicyId

`func (o *UpdateProtectionObjectParameters) GetRpoPolicyId() string`

GetRpoPolicyId returns the RpoPolicyId field if non-nil, zero value otherwise.

### GetRpoPolicyIdOk

`func (o *UpdateProtectionObjectParameters) GetRpoPolicyIdOk() (*string, bool)`

GetRpoPolicyIdOk returns a tuple with the RpoPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpoPolicyId

`func (o *UpdateProtectionObjectParameters) SetRpoPolicyId(v string)`

SetRpoPolicyId sets RpoPolicyId field to given value.

### HasRpoPolicyId

`func (o *UpdateProtectionObjectParameters) HasRpoPolicyId() bool`

HasRpoPolicyId returns a boolean if a field has been set.

### SetRpoPolicyIdNil

`func (o *UpdateProtectionObjectParameters) SetRpoPolicyIdNil(b bool)`

 SetRpoPolicyIdNil sets the value for RpoPolicyId to be an explicit nil

### UnsetRpoPolicyId
`func (o *UpdateProtectionObjectParameters) UnsetRpoPolicyId()`

UnsetRpoPolicyId ensures that no value is present for RpoPolicyId, not even an explicit nil
### GetSourceParameters

`func (o *UpdateProtectionObjectParameters) GetSourceParameters() []SourceSpecialParameter`

GetSourceParameters returns the SourceParameters field if non-nil, zero value otherwise.

### GetSourceParametersOk

`func (o *UpdateProtectionObjectParameters) GetSourceParametersOk() (*[]SourceSpecialParameter, bool)`

GetSourceParametersOk returns a tuple with the SourceParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceParameters

`func (o *UpdateProtectionObjectParameters) SetSourceParameters(v []SourceSpecialParameter)`

SetSourceParameters sets SourceParameters field to given value.

### HasSourceParameters

`func (o *UpdateProtectionObjectParameters) HasSourceParameters() bool`

HasSourceParameters returns a boolean if a field has been set.

### SetSourceParametersNil

`func (o *UpdateProtectionObjectParameters) SetSourceParametersNil(b bool)`

 SetSourceParametersNil sets the value for SourceParameters to be an explicit nil

### UnsetSourceParameters
`func (o *UpdateProtectionObjectParameters) UnsetSourceParameters()`

UnsetSourceParameters ensures that no value is present for SourceParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


