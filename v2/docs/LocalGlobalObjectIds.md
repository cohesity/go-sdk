# LocalGlobalObjectIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalId** | Pointer to **string** | Specifies the global ID of the object. | [optional] 
**LocalId** | Pointer to **int64** | Specifies the local ID of the object. | [optional] 

## Methods

### NewLocalGlobalObjectIds

`func NewLocalGlobalObjectIds() *LocalGlobalObjectIds`

NewLocalGlobalObjectIds instantiates a new LocalGlobalObjectIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalGlobalObjectIdsWithDefaults

`func NewLocalGlobalObjectIdsWithDefaults() *LocalGlobalObjectIds`

NewLocalGlobalObjectIdsWithDefaults instantiates a new LocalGlobalObjectIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalId

`func (o *LocalGlobalObjectIds) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *LocalGlobalObjectIds) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *LocalGlobalObjectIds) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *LocalGlobalObjectIds) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetLocalId

`func (o *LocalGlobalObjectIds) GetLocalId() int64`

GetLocalId returns the LocalId field if non-nil, zero value otherwise.

### GetLocalIdOk

`func (o *LocalGlobalObjectIds) GetLocalIdOk() (*int64, bool)`

GetLocalIdOk returns a tuple with the LocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalId

`func (o *LocalGlobalObjectIds) SetLocalId(v int64)`

SetLocalId sets LocalId field to given value.

### HasLocalId

`func (o *LocalGlobalObjectIds) HasLocalId() bool`

HasLocalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


