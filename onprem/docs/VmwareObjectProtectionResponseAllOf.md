# VmwareObjectProtectionResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the objects to not be protected in this backup. This field only applies if provided object id is non leaf entity such as Tag or a folder. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 
**CdpInfo** | Pointer to [**VmwareCdpObject**](VmwareCdpObject.md) |  | [optional] 
**StandbyInfo** | Pointer to [**VmwareStandbyObject**](VmwareStandbyObject.md) |  | [optional] 

## Methods

### NewVmwareObjectProtectionResponseAllOf

`func NewVmwareObjectProtectionResponseAllOf() *VmwareObjectProtectionResponseAllOf`

NewVmwareObjectProtectionResponseAllOf instantiates a new VmwareObjectProtectionResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectProtectionResponseAllOfWithDefaults

`func NewVmwareObjectProtectionResponseAllOfWithDefaults() *VmwareObjectProtectionResponseAllOf`

NewVmwareObjectProtectionResponseAllOfWithDefaults instantiates a new VmwareObjectProtectionResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeObjectIds

`func (o *VmwareObjectProtectionResponseAllOf) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *VmwareObjectProtectionResponseAllOf) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *VmwareObjectProtectionResponseAllOf) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *VmwareObjectProtectionResponseAllOf) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### GetCdpInfo

`func (o *VmwareObjectProtectionResponseAllOf) GetCdpInfo() VmwareCdpObject`

GetCdpInfo returns the CdpInfo field if non-nil, zero value otherwise.

### GetCdpInfoOk

`func (o *VmwareObjectProtectionResponseAllOf) GetCdpInfoOk() (*VmwareCdpObject, bool)`

GetCdpInfoOk returns a tuple with the CdpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpInfo

`func (o *VmwareObjectProtectionResponseAllOf) SetCdpInfo(v VmwareCdpObject)`

SetCdpInfo sets CdpInfo field to given value.

### HasCdpInfo

`func (o *VmwareObjectProtectionResponseAllOf) HasCdpInfo() bool`

HasCdpInfo returns a boolean if a field has been set.

### GetStandbyInfo

`func (o *VmwareObjectProtectionResponseAllOf) GetStandbyInfo() VmwareStandbyObject`

GetStandbyInfo returns the StandbyInfo field if non-nil, zero value otherwise.

### GetStandbyInfoOk

`func (o *VmwareObjectProtectionResponseAllOf) GetStandbyInfoOk() (*VmwareStandbyObject, bool)`

GetStandbyInfoOk returns a tuple with the StandbyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyInfo

`func (o *VmwareObjectProtectionResponseAllOf) SetStandbyInfo(v VmwareStandbyObject)`

SetStandbyInfo sets StandbyInfo field to given value.

### HasStandbyInfo

`func (o *VmwareObjectProtectionResponseAllOf) HasStandbyInfo() bool`

HasStandbyInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


