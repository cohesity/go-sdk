# HyperVObjectProtectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the id of the object being protected. This can be a leaf level or non leaf level object. | 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the objects to not be protected by this Protection Group. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 
**BackupsCopyOnly** | Pointer to **NullableBool** | Specifies whether the backups should be copy-only. | [optional] 

## Methods

### NewHyperVObjectProtectionRequest

`func NewHyperVObjectProtectionRequest(id NullableInt64, ) *HyperVObjectProtectionRequest`

NewHyperVObjectProtectionRequest instantiates a new HyperVObjectProtectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVObjectProtectionRequestWithDefaults

`func NewHyperVObjectProtectionRequestWithDefaults() *HyperVObjectProtectionRequest`

NewHyperVObjectProtectionRequestWithDefaults instantiates a new HyperVObjectProtectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HyperVObjectProtectionRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HyperVObjectProtectionRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HyperVObjectProtectionRequest) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *HyperVObjectProtectionRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HyperVObjectProtectionRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetExcludeObjectIds

`func (o *HyperVObjectProtectionRequest) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *HyperVObjectProtectionRequest) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *HyperVObjectProtectionRequest) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *HyperVObjectProtectionRequest) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### GetBackupsCopyOnly

`func (o *HyperVObjectProtectionRequest) GetBackupsCopyOnly() bool`

GetBackupsCopyOnly returns the BackupsCopyOnly field if non-nil, zero value otherwise.

### GetBackupsCopyOnlyOk

`func (o *HyperVObjectProtectionRequest) GetBackupsCopyOnlyOk() (*bool, bool)`

GetBackupsCopyOnlyOk returns a tuple with the BackupsCopyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsCopyOnly

`func (o *HyperVObjectProtectionRequest) SetBackupsCopyOnly(v bool)`

SetBackupsCopyOnly sets BackupsCopyOnly field to given value.

### HasBackupsCopyOnly

`func (o *HyperVObjectProtectionRequest) HasBackupsCopyOnly() bool`

HasBackupsCopyOnly returns a boolean if a field has been set.

### SetBackupsCopyOnlyNil

`func (o *HyperVObjectProtectionRequest) SetBackupsCopyOnlyNil(b bool)`

 SetBackupsCopyOnlyNil sets the value for BackupsCopyOnly to be an explicit nil

### UnsetBackupsCopyOnly
`func (o *HyperVObjectProtectionRequest) UnsetBackupsCopyOnly()`

UnsetBackupsCopyOnly ensures that no value is present for BackupsCopyOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


