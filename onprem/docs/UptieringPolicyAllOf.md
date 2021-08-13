# UptieringPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileAge** | Pointer to [**UptieringFileAgePolicy**](UptieringFileAgePolicy.md) |  | [optional] 
**IncludeAllFiles** | Pointer to **NullableBool** | If set, all files in the view will be uptiered regardless of file_select_policy, num_file_access, hot_file_window, file_size constraints. | [optional] [default to false]
**Target** | Pointer to [**NullableUptieringTarget**](UptieringTarget.md) |  | [optional] 

## Methods

### NewUptieringPolicyAllOf

`func NewUptieringPolicyAllOf() *UptieringPolicyAllOf`

NewUptieringPolicyAllOf instantiates a new UptieringPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUptieringPolicyAllOfWithDefaults

`func NewUptieringPolicyAllOfWithDefaults() *UptieringPolicyAllOf`

NewUptieringPolicyAllOfWithDefaults instantiates a new UptieringPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileAge

`func (o *UptieringPolicyAllOf) GetFileAge() UptieringFileAgePolicy`

GetFileAge returns the FileAge field if non-nil, zero value otherwise.

### GetFileAgeOk

`func (o *UptieringPolicyAllOf) GetFileAgeOk() (*UptieringFileAgePolicy, bool)`

GetFileAgeOk returns a tuple with the FileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAge

`func (o *UptieringPolicyAllOf) SetFileAge(v UptieringFileAgePolicy)`

SetFileAge sets FileAge field to given value.

### HasFileAge

`func (o *UptieringPolicyAllOf) HasFileAge() bool`

HasFileAge returns a boolean if a field has been set.

### GetIncludeAllFiles

`func (o *UptieringPolicyAllOf) GetIncludeAllFiles() bool`

GetIncludeAllFiles returns the IncludeAllFiles field if non-nil, zero value otherwise.

### GetIncludeAllFilesOk

`func (o *UptieringPolicyAllOf) GetIncludeAllFilesOk() (*bool, bool)`

GetIncludeAllFilesOk returns a tuple with the IncludeAllFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllFiles

`func (o *UptieringPolicyAllOf) SetIncludeAllFiles(v bool)`

SetIncludeAllFiles sets IncludeAllFiles field to given value.

### HasIncludeAllFiles

`func (o *UptieringPolicyAllOf) HasIncludeAllFiles() bool`

HasIncludeAllFiles returns a boolean if a field has been set.

### SetIncludeAllFilesNil

`func (o *UptieringPolicyAllOf) SetIncludeAllFilesNil(b bool)`

 SetIncludeAllFilesNil sets the value for IncludeAllFiles to be an explicit nil

### UnsetIncludeAllFiles
`func (o *UptieringPolicyAllOf) UnsetIncludeAllFiles()`

UnsetIncludeAllFiles ensures that no value is present for IncludeAllFiles, not even an explicit nil
### GetTarget

`func (o *UptieringPolicyAllOf) GetTarget() UptieringTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *UptieringPolicyAllOf) GetTargetOk() (*UptieringTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *UptieringPolicyAllOf) SetTarget(v UptieringTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *UptieringPolicyAllOf) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *UptieringPolicyAllOf) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *UptieringPolicyAllOf) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


