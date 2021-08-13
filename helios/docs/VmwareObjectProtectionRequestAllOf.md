# VmwareObjectProtectionRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the id of the object being protected. This can be a leaf level or non leaf level object. | 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the objects to not be protected in this backup. This field only applies if provided object id is non leaf entity such as Tag or a folder. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 

## Methods

### NewVmwareObjectProtectionRequestAllOf

`func NewVmwareObjectProtectionRequestAllOf(id NullableInt64, ) *VmwareObjectProtectionRequestAllOf`

NewVmwareObjectProtectionRequestAllOf instantiates a new VmwareObjectProtectionRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectProtectionRequestAllOfWithDefaults

`func NewVmwareObjectProtectionRequestAllOfWithDefaults() *VmwareObjectProtectionRequestAllOf`

NewVmwareObjectProtectionRequestAllOfWithDefaults instantiates a new VmwareObjectProtectionRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmwareObjectProtectionRequestAllOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmwareObjectProtectionRequestAllOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmwareObjectProtectionRequestAllOf) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *VmwareObjectProtectionRequestAllOf) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VmwareObjectProtectionRequestAllOf) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetExcludeObjectIds

`func (o *VmwareObjectProtectionRequestAllOf) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *VmwareObjectProtectionRequestAllOf) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *VmwareObjectProtectionRequestAllOf) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *VmwareObjectProtectionRequestAllOf) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


