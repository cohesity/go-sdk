# VmwareObjectProtectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the id of the object being protected. This can be a leaf level or non leaf level object. | 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the objects to not be protected in this backup. This field only applies if provided object id is non leaf entity such as Tag or a folder. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 
**ExcludeDisks** | Pointer to [**[]DiskInfo**](DiskInfo.md) | Specifies a list of disks to exclude from being protected. This is only applicable to VM objects. | [optional] 
**TruncateExchangeLogs** | Pointer to **NullableBool** | Specifies whether or not to truncate MS Exchange logs while taking an app consistent snapshot of this object. This is only applicable to objects which have a registered MS Exchange app. | [optional] 

## Methods

### NewVmwareObjectProtectionRequest

`func NewVmwareObjectProtectionRequest(id NullableInt64, ) *VmwareObjectProtectionRequest`

NewVmwareObjectProtectionRequest instantiates a new VmwareObjectProtectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectProtectionRequestWithDefaults

`func NewVmwareObjectProtectionRequestWithDefaults() *VmwareObjectProtectionRequest`

NewVmwareObjectProtectionRequestWithDefaults instantiates a new VmwareObjectProtectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmwareObjectProtectionRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmwareObjectProtectionRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmwareObjectProtectionRequest) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *VmwareObjectProtectionRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VmwareObjectProtectionRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetExcludeObjectIds

`func (o *VmwareObjectProtectionRequest) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *VmwareObjectProtectionRequest) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *VmwareObjectProtectionRequest) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *VmwareObjectProtectionRequest) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### GetExcludeDisks

`func (o *VmwareObjectProtectionRequest) GetExcludeDisks() []DiskInfo`

GetExcludeDisks returns the ExcludeDisks field if non-nil, zero value otherwise.

### GetExcludeDisksOk

`func (o *VmwareObjectProtectionRequest) GetExcludeDisksOk() (*[]DiskInfo, bool)`

GetExcludeDisksOk returns a tuple with the ExcludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDisks

`func (o *VmwareObjectProtectionRequest) SetExcludeDisks(v []DiskInfo)`

SetExcludeDisks sets ExcludeDisks field to given value.

### HasExcludeDisks

`func (o *VmwareObjectProtectionRequest) HasExcludeDisks() bool`

HasExcludeDisks returns a boolean if a field has been set.

### GetTruncateExchangeLogs

`func (o *VmwareObjectProtectionRequest) GetTruncateExchangeLogs() bool`

GetTruncateExchangeLogs returns the TruncateExchangeLogs field if non-nil, zero value otherwise.

### GetTruncateExchangeLogsOk

`func (o *VmwareObjectProtectionRequest) GetTruncateExchangeLogsOk() (*bool, bool)`

GetTruncateExchangeLogsOk returns a tuple with the TruncateExchangeLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateExchangeLogs

`func (o *VmwareObjectProtectionRequest) SetTruncateExchangeLogs(v bool)`

SetTruncateExchangeLogs sets TruncateExchangeLogs field to given value.

### HasTruncateExchangeLogs

`func (o *VmwareObjectProtectionRequest) HasTruncateExchangeLogs() bool`

HasTruncateExchangeLogs returns a boolean if a field has been set.

### SetTruncateExchangeLogsNil

`func (o *VmwareObjectProtectionRequest) SetTruncateExchangeLogsNil(b bool)`

 SetTruncateExchangeLogsNil sets the value for TruncateExchangeLogs to be an explicit nil

### UnsetTruncateExchangeLogs
`func (o *VmwareObjectProtectionRequest) UnsetTruncateExchangeLogs()`

UnsetTruncateExchangeLogs ensures that no value is present for TruncateExchangeLogs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


