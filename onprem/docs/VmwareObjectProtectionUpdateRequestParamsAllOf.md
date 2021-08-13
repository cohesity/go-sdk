# VmwareObjectProtectionUpdateRequestParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the objects to not be protected in this backup. This field only applies if provided object id is non leaf entity such as Tag or a folder. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 

## Methods

### NewVmwareObjectProtectionUpdateRequestParamsAllOf

`func NewVmwareObjectProtectionUpdateRequestParamsAllOf() *VmwareObjectProtectionUpdateRequestParamsAllOf`

NewVmwareObjectProtectionUpdateRequestParamsAllOf instantiates a new VmwareObjectProtectionUpdateRequestParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectProtectionUpdateRequestParamsAllOfWithDefaults

`func NewVmwareObjectProtectionUpdateRequestParamsAllOfWithDefaults() *VmwareObjectProtectionUpdateRequestParamsAllOf`

NewVmwareObjectProtectionUpdateRequestParamsAllOfWithDefaults instantiates a new VmwareObjectProtectionUpdateRequestParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeObjectIds

`func (o *VmwareObjectProtectionUpdateRequestParamsAllOf) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *VmwareObjectProtectionUpdateRequestParamsAllOf) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *VmwareObjectProtectionUpdateRequestParamsAllOf) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *VmwareObjectProtectionUpdateRequestParamsAllOf) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


