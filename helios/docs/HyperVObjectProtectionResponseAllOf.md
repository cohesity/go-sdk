# HyperVObjectProtectionResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the objects to not be protected by this Protection Group. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 

## Methods

### NewHyperVObjectProtectionResponseAllOf

`func NewHyperVObjectProtectionResponseAllOf() *HyperVObjectProtectionResponseAllOf`

NewHyperVObjectProtectionResponseAllOf instantiates a new HyperVObjectProtectionResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVObjectProtectionResponseAllOfWithDefaults

`func NewHyperVObjectProtectionResponseAllOfWithDefaults() *HyperVObjectProtectionResponseAllOf`

NewHyperVObjectProtectionResponseAllOfWithDefaults instantiates a new HyperVObjectProtectionResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeObjectIds

`func (o *HyperVObjectProtectionResponseAllOf) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *HyperVObjectProtectionResponseAllOf) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *HyperVObjectProtectionResponseAllOf) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *HyperVObjectProtectionResponseAllOf) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


