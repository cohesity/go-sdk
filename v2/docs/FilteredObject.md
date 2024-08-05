# FilteredObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies object id. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the source id to which this object belongs to. | [optional] 

## Methods

### NewFilteredObject

`func NewFilteredObject() *FilteredObject`

NewFilteredObject instantiates a new FilteredObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilteredObjectWithDefaults

`func NewFilteredObjectWithDefaults() *FilteredObject`

NewFilteredObjectWithDefaults instantiates a new FilteredObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FilteredObject) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilteredObject) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilteredObject) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FilteredObject) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *FilteredObject) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FilteredObject) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSourceId

`func (o *FilteredObject) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *FilteredObject) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *FilteredObject) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *FilteredObject) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *FilteredObject) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *FilteredObject) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


