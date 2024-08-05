# HeliosCommonSearchIndexedObjectsResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterErrors** | Pointer to [**[]HeliosSearchIndexedObjectsClusterError**](HeliosSearchIndexedObjectsClusterError.md) | A List of errors that occured on a subset of clusters. | [optional] 
**Count** | Pointer to **NullableInt32** | Specifies the total number of indexed objects that match the filter and search criteria. Use this value to determine how many additional requests are required to get the full result. | [optional] 
**ObjectType** | Pointer to **NullableString** | Specifies the object type. | [optional] 

## Methods

### NewHeliosCommonSearchIndexedObjectsResponseParams

`func NewHeliosCommonSearchIndexedObjectsResponseParams() *HeliosCommonSearchIndexedObjectsResponseParams`

NewHeliosCommonSearchIndexedObjectsResponseParams instantiates a new HeliosCommonSearchIndexedObjectsResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosCommonSearchIndexedObjectsResponseParamsWithDefaults

`func NewHeliosCommonSearchIndexedObjectsResponseParamsWithDefaults() *HeliosCommonSearchIndexedObjectsResponseParams`

NewHeliosCommonSearchIndexedObjectsResponseParamsWithDefaults instantiates a new HeliosCommonSearchIndexedObjectsResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterErrors

`func (o *HeliosCommonSearchIndexedObjectsResponseParams) GetClusterErrors() []HeliosSearchIndexedObjectsClusterError`

GetClusterErrors returns the ClusterErrors field if non-nil, zero value otherwise.

### GetClusterErrorsOk

`func (o *HeliosCommonSearchIndexedObjectsResponseParams) GetClusterErrorsOk() (*[]HeliosSearchIndexedObjectsClusterError, bool)`

GetClusterErrorsOk returns a tuple with the ClusterErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterErrors

`func (o *HeliosCommonSearchIndexedObjectsResponseParams) SetClusterErrors(v []HeliosSearchIndexedObjectsClusterError)`

SetClusterErrors sets ClusterErrors field to given value.

### HasClusterErrors

`func (o *HeliosCommonSearchIndexedObjectsResponseParams) HasClusterErrors() bool`

HasClusterErrors returns a boolean if a field has been set.

### GetCount

`func (o *HeliosCommonSearchIndexedObjectsResponseParams) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HeliosCommonSearchIndexedObjectsResponseParams) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HeliosCommonSearchIndexedObjectsResponseParams) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *HeliosCommonSearchIndexedObjectsResponseParams) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *HeliosCommonSearchIndexedObjectsResponseParams) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *HeliosCommonSearchIndexedObjectsResponseParams) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetObjectType

`func (o *HeliosCommonSearchIndexedObjectsResponseParams) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HeliosCommonSearchIndexedObjectsResponseParams) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HeliosCommonSearchIndexedObjectsResponseParams) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *HeliosCommonSearchIndexedObjectsResponseParams) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *HeliosCommonSearchIndexedObjectsResponseParams) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *HeliosCommonSearchIndexedObjectsResponseParams) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


