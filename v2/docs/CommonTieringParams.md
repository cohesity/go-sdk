# CommonTieringParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeAllFiles** | Pointer to **NullableBool** | If set, all files in the view will be uptiered regardless of file_select_policy, num_file_access, hot_file_window, file_size constraints. | [optional] [default to false]
**Target** | Pointer to [**NullableDataTieringTarget**](DataTieringTarget.md) |  | [optional] 

## Methods

### NewCommonTieringParams

`func NewCommonTieringParams() *CommonTieringParams`

NewCommonTieringParams instantiates a new CommonTieringParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonTieringParamsWithDefaults

`func NewCommonTieringParamsWithDefaults() *CommonTieringParams`

NewCommonTieringParamsWithDefaults instantiates a new CommonTieringParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeAllFiles

`func (o *CommonTieringParams) GetIncludeAllFiles() bool`

GetIncludeAllFiles returns the IncludeAllFiles field if non-nil, zero value otherwise.

### GetIncludeAllFilesOk

`func (o *CommonTieringParams) GetIncludeAllFilesOk() (*bool, bool)`

GetIncludeAllFilesOk returns a tuple with the IncludeAllFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllFiles

`func (o *CommonTieringParams) SetIncludeAllFiles(v bool)`

SetIncludeAllFiles sets IncludeAllFiles field to given value.

### HasIncludeAllFiles

`func (o *CommonTieringParams) HasIncludeAllFiles() bool`

HasIncludeAllFiles returns a boolean if a field has been set.

### SetIncludeAllFilesNil

`func (o *CommonTieringParams) SetIncludeAllFilesNil(b bool)`

 SetIncludeAllFilesNil sets the value for IncludeAllFiles to be an explicit nil

### UnsetIncludeAllFiles
`func (o *CommonTieringParams) UnsetIncludeAllFiles()`

UnsetIncludeAllFiles ensures that no value is present for IncludeAllFiles, not even an explicit nil
### GetTarget

`func (o *CommonTieringParams) GetTarget() DataTieringTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CommonTieringParams) GetTargetOk() (*DataTieringTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CommonTieringParams) SetTarget(v DataTieringTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CommonTieringParams) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *CommonTieringParams) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *CommonTieringParams) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


