# IsilonDataTieringParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]ProtectionObjectInput**](ProtectionObjectInput.md) | Specifies the objects to be included in the Protection Group. | 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the root of data tiering source. | [optional] 

## Methods

### NewIsilonDataTieringParams

`func NewIsilonDataTieringParams(objects []ProtectionObjectInput, ) *IsilonDataTieringParams`

NewIsilonDataTieringParams instantiates a new IsilonDataTieringParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsilonDataTieringParamsWithDefaults

`func NewIsilonDataTieringParamsWithDefaults() *IsilonDataTieringParams`

NewIsilonDataTieringParamsWithDefaults instantiates a new IsilonDataTieringParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *IsilonDataTieringParams) GetObjects() []ProtectionObjectInput`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *IsilonDataTieringParams) GetObjectsOk() (*[]ProtectionObjectInput, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *IsilonDataTieringParams) SetObjects(v []ProtectionObjectInput)`

SetObjects sets Objects field to given value.


### GetSourceId

`func (o *IsilonDataTieringParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *IsilonDataTieringParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *IsilonDataTieringParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *IsilonDataTieringParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *IsilonDataTieringParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *IsilonDataTieringParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


