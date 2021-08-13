# HyperVObjectProtectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupsCopyOnly** | Pointer to **NullableBool** | Specifies whether the backups should be copy-only. | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the objects to not be protected by this Protection Group. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 

## Methods

### NewHyperVObjectProtectionResponse

`func NewHyperVObjectProtectionResponse() *HyperVObjectProtectionResponse`

NewHyperVObjectProtectionResponse instantiates a new HyperVObjectProtectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVObjectProtectionResponseWithDefaults

`func NewHyperVObjectProtectionResponseWithDefaults() *HyperVObjectProtectionResponse`

NewHyperVObjectProtectionResponseWithDefaults instantiates a new HyperVObjectProtectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupsCopyOnly

`func (o *HyperVObjectProtectionResponse) GetBackupsCopyOnly() bool`

GetBackupsCopyOnly returns the BackupsCopyOnly field if non-nil, zero value otherwise.

### GetBackupsCopyOnlyOk

`func (o *HyperVObjectProtectionResponse) GetBackupsCopyOnlyOk() (*bool, bool)`

GetBackupsCopyOnlyOk returns a tuple with the BackupsCopyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsCopyOnly

`func (o *HyperVObjectProtectionResponse) SetBackupsCopyOnly(v bool)`

SetBackupsCopyOnly sets BackupsCopyOnly field to given value.

### HasBackupsCopyOnly

`func (o *HyperVObjectProtectionResponse) HasBackupsCopyOnly() bool`

HasBackupsCopyOnly returns a boolean if a field has been set.

### SetBackupsCopyOnlyNil

`func (o *HyperVObjectProtectionResponse) SetBackupsCopyOnlyNil(b bool)`

 SetBackupsCopyOnlyNil sets the value for BackupsCopyOnly to be an explicit nil

### UnsetBackupsCopyOnly
`func (o *HyperVObjectProtectionResponse) UnsetBackupsCopyOnly()`

UnsetBackupsCopyOnly ensures that no value is present for BackupsCopyOnly, not even an explicit nil
### GetExcludeObjectIds

`func (o *HyperVObjectProtectionResponse) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *HyperVObjectProtectionResponse) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *HyperVObjectProtectionResponse) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *HyperVObjectProtectionResponse) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


