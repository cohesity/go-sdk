# VmwareObjectProtectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeDisks** | Pointer to [**[]DiskInfo**](DiskInfo.md) | Specifies a list of disks to exclude from being protected. This is only applicable to VM objects. | [optional] 
**TruncateExchangeLogs** | Pointer to **NullableBool** | Specifies whether or not to truncate MS Exchange logs while taking an app consistent snapshot of this object. This is only applicable to objects which have a registered MS Exchange app. | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the objects to not be protected in this backup. This field only applies if provided object id is non leaf entity such as Tag or a folder. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 
**CdpInfo** | Pointer to [**VmwareCdpObject**](VmwareCdpObject.md) |  | [optional] 
**StandbyInfo** | Pointer to [**VmwareStandbyObject**](VmwareStandbyObject.md) |  | [optional] 

## Methods

### NewVmwareObjectProtectionResponse

`func NewVmwareObjectProtectionResponse() *VmwareObjectProtectionResponse`

NewVmwareObjectProtectionResponse instantiates a new VmwareObjectProtectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectProtectionResponseWithDefaults

`func NewVmwareObjectProtectionResponseWithDefaults() *VmwareObjectProtectionResponse`

NewVmwareObjectProtectionResponseWithDefaults instantiates a new VmwareObjectProtectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeDisks

`func (o *VmwareObjectProtectionResponse) GetExcludeDisks() []DiskInfo`

GetExcludeDisks returns the ExcludeDisks field if non-nil, zero value otherwise.

### GetExcludeDisksOk

`func (o *VmwareObjectProtectionResponse) GetExcludeDisksOk() (*[]DiskInfo, bool)`

GetExcludeDisksOk returns a tuple with the ExcludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDisks

`func (o *VmwareObjectProtectionResponse) SetExcludeDisks(v []DiskInfo)`

SetExcludeDisks sets ExcludeDisks field to given value.

### HasExcludeDisks

`func (o *VmwareObjectProtectionResponse) HasExcludeDisks() bool`

HasExcludeDisks returns a boolean if a field has been set.

### GetTruncateExchangeLogs

`func (o *VmwareObjectProtectionResponse) GetTruncateExchangeLogs() bool`

GetTruncateExchangeLogs returns the TruncateExchangeLogs field if non-nil, zero value otherwise.

### GetTruncateExchangeLogsOk

`func (o *VmwareObjectProtectionResponse) GetTruncateExchangeLogsOk() (*bool, bool)`

GetTruncateExchangeLogsOk returns a tuple with the TruncateExchangeLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateExchangeLogs

`func (o *VmwareObjectProtectionResponse) SetTruncateExchangeLogs(v bool)`

SetTruncateExchangeLogs sets TruncateExchangeLogs field to given value.

### HasTruncateExchangeLogs

`func (o *VmwareObjectProtectionResponse) HasTruncateExchangeLogs() bool`

HasTruncateExchangeLogs returns a boolean if a field has been set.

### SetTruncateExchangeLogsNil

`func (o *VmwareObjectProtectionResponse) SetTruncateExchangeLogsNil(b bool)`

 SetTruncateExchangeLogsNil sets the value for TruncateExchangeLogs to be an explicit nil

### UnsetTruncateExchangeLogs
`func (o *VmwareObjectProtectionResponse) UnsetTruncateExchangeLogs()`

UnsetTruncateExchangeLogs ensures that no value is present for TruncateExchangeLogs, not even an explicit nil
### GetExcludeObjectIds

`func (o *VmwareObjectProtectionResponse) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *VmwareObjectProtectionResponse) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *VmwareObjectProtectionResponse) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *VmwareObjectProtectionResponse) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### GetCdpInfo

`func (o *VmwareObjectProtectionResponse) GetCdpInfo() VmwareCdpObject`

GetCdpInfo returns the CdpInfo field if non-nil, zero value otherwise.

### GetCdpInfoOk

`func (o *VmwareObjectProtectionResponse) GetCdpInfoOk() (*VmwareCdpObject, bool)`

GetCdpInfoOk returns a tuple with the CdpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpInfo

`func (o *VmwareObjectProtectionResponse) SetCdpInfo(v VmwareCdpObject)`

SetCdpInfo sets CdpInfo field to given value.

### HasCdpInfo

`func (o *VmwareObjectProtectionResponse) HasCdpInfo() bool`

HasCdpInfo returns a boolean if a field has been set.

### GetStandbyInfo

`func (o *VmwareObjectProtectionResponse) GetStandbyInfo() VmwareStandbyObject`

GetStandbyInfo returns the StandbyInfo field if non-nil, zero value otherwise.

### GetStandbyInfoOk

`func (o *VmwareObjectProtectionResponse) GetStandbyInfoOk() (*VmwareStandbyObject, bool)`

GetStandbyInfoOk returns a tuple with the StandbyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyInfo

`func (o *VmwareObjectProtectionResponse) SetStandbyInfo(v VmwareStandbyObject)`

SetStandbyInfo sets StandbyInfo field to given value.

### HasStandbyInfo

`func (o *VmwareObjectProtectionResponse) HasStandbyInfo() bool`

HasStandbyInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


