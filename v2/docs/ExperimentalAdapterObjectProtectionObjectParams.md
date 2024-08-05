# ExperimentalAdapterObjectProtectionObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedObjectIds** | Pointer to **[]int32** | Specifies the ids of the objects to be excluded in the Object Protection. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 
**Id** | **int64** | Specifies the id of the object being protected. This can be a leaf level or non leaf level object. | 

## Methods

### NewExperimentalAdapterObjectProtectionObjectParams

`func NewExperimentalAdapterObjectProtectionObjectParams(id int64, ) *ExperimentalAdapterObjectProtectionObjectParams`

NewExperimentalAdapterObjectProtectionObjectParams instantiates a new ExperimentalAdapterObjectProtectionObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentalAdapterObjectProtectionObjectParamsWithDefaults

`func NewExperimentalAdapterObjectProtectionObjectParamsWithDefaults() *ExperimentalAdapterObjectProtectionObjectParams`

NewExperimentalAdapterObjectProtectionObjectParamsWithDefaults instantiates a new ExperimentalAdapterObjectProtectionObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedObjectIds

`func (o *ExperimentalAdapterObjectProtectionObjectParams) GetExcludedObjectIds() []int32`

GetExcludedObjectIds returns the ExcludedObjectIds field if non-nil, zero value otherwise.

### GetExcludedObjectIdsOk

`func (o *ExperimentalAdapterObjectProtectionObjectParams) GetExcludedObjectIdsOk() (*[]int32, bool)`

GetExcludedObjectIdsOk returns a tuple with the ExcludedObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedObjectIds

`func (o *ExperimentalAdapterObjectProtectionObjectParams) SetExcludedObjectIds(v []int32)`

SetExcludedObjectIds sets ExcludedObjectIds field to given value.

### HasExcludedObjectIds

`func (o *ExperimentalAdapterObjectProtectionObjectParams) HasExcludedObjectIds() bool`

HasExcludedObjectIds returns a boolean if a field has been set.

### SetExcludedObjectIdsNil

`func (o *ExperimentalAdapterObjectProtectionObjectParams) SetExcludedObjectIdsNil(b bool)`

 SetExcludedObjectIdsNil sets the value for ExcludedObjectIds to be an explicit nil

### UnsetExcludedObjectIds
`func (o *ExperimentalAdapterObjectProtectionObjectParams) UnsetExcludedObjectIds()`

UnsetExcludedObjectIds ensures that no value is present for ExcludedObjectIds, not even an explicit nil
### GetId

`func (o *ExperimentalAdapterObjectProtectionObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExperimentalAdapterObjectProtectionObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExperimentalAdapterObjectProtectionObjectParams) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


