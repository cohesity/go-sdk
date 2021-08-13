# CommonObjectsActionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectMap** | Pointer to [**[]ActionObjectMapping**](ActionObjectMapping.md) | Specifies the objectMap that will be used to perform bulk actions such as linking and unlinking. | [optional] 

## Methods

### NewCommonObjectsActionParams

`func NewCommonObjectsActionParams() *CommonObjectsActionParams`

NewCommonObjectsActionParams instantiates a new CommonObjectsActionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonObjectsActionParamsWithDefaults

`func NewCommonObjectsActionParamsWithDefaults() *CommonObjectsActionParams`

NewCommonObjectsActionParamsWithDefaults instantiates a new CommonObjectsActionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectMap

`func (o *CommonObjectsActionParams) GetObjectMap() []ActionObjectMapping`

GetObjectMap returns the ObjectMap field if non-nil, zero value otherwise.

### GetObjectMapOk

`func (o *CommonObjectsActionParams) GetObjectMapOk() (*[]ActionObjectMapping, bool)`

GetObjectMapOk returns a tuple with the ObjectMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectMap

`func (o *CommonObjectsActionParams) SetObjectMap(v []ActionObjectMapping)`

SetObjectMap sets ObjectMap field to given value.

### HasObjectMap

`func (o *CommonObjectsActionParams) HasObjectMap() bool`

HasObjectMap returns a boolean if a field has been set.

### SetObjectMapNil

`func (o *CommonObjectsActionParams) SetObjectMapNil(b bool)`

 SetObjectMapNil sets the value for ObjectMap to be an explicit nil

### UnsetObjectMap
`func (o *CommonObjectsActionParams) UnsetObjectMap()`

UnsetObjectMap ensures that no value is present for ObjectMap, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


