# MongoDBCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCappedCollection** | Pointer to **NullableBool** | Set to true if this is a capped Collection. | [optional] 
**IsMongoView** | Pointer to **NullableBool** | Set to true if this Collection is a view. | [optional] 
**SizeBytes** | Pointer to **NullableInt64** | Size of this Collection. | [optional] 

## Methods

### NewMongoDBCollection

`func NewMongoDBCollection() *MongoDBCollection`

NewMongoDBCollection instantiates a new MongoDBCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBCollectionWithDefaults

`func NewMongoDBCollectionWithDefaults() *MongoDBCollection`

NewMongoDBCollectionWithDefaults instantiates a new MongoDBCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCappedCollection

`func (o *MongoDBCollection) GetIsCappedCollection() bool`

GetIsCappedCollection returns the IsCappedCollection field if non-nil, zero value otherwise.

### GetIsCappedCollectionOk

`func (o *MongoDBCollection) GetIsCappedCollectionOk() (*bool, bool)`

GetIsCappedCollectionOk returns a tuple with the IsCappedCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCappedCollection

`func (o *MongoDBCollection) SetIsCappedCollection(v bool)`

SetIsCappedCollection sets IsCappedCollection field to given value.

### HasIsCappedCollection

`func (o *MongoDBCollection) HasIsCappedCollection() bool`

HasIsCappedCollection returns a boolean if a field has been set.

### SetIsCappedCollectionNil

`func (o *MongoDBCollection) SetIsCappedCollectionNil(b bool)`

 SetIsCappedCollectionNil sets the value for IsCappedCollection to be an explicit nil

### UnsetIsCappedCollection
`func (o *MongoDBCollection) UnsetIsCappedCollection()`

UnsetIsCappedCollection ensures that no value is present for IsCappedCollection, not even an explicit nil
### GetIsMongoView

`func (o *MongoDBCollection) GetIsMongoView() bool`

GetIsMongoView returns the IsMongoView field if non-nil, zero value otherwise.

### GetIsMongoViewOk

`func (o *MongoDBCollection) GetIsMongoViewOk() (*bool, bool)`

GetIsMongoViewOk returns a tuple with the IsMongoView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMongoView

`func (o *MongoDBCollection) SetIsMongoView(v bool)`

SetIsMongoView sets IsMongoView field to given value.

### HasIsMongoView

`func (o *MongoDBCollection) HasIsMongoView() bool`

HasIsMongoView returns a boolean if a field has been set.

### SetIsMongoViewNil

`func (o *MongoDBCollection) SetIsMongoViewNil(b bool)`

 SetIsMongoViewNil sets the value for IsMongoView to be an explicit nil

### UnsetIsMongoView
`func (o *MongoDBCollection) UnsetIsMongoView()`

UnsetIsMongoView ensures that no value is present for IsMongoView, not even an explicit nil
### GetSizeBytes

`func (o *MongoDBCollection) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *MongoDBCollection) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *MongoDBCollection) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *MongoDBCollection) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### SetSizeBytesNil

`func (o *MongoDBCollection) SetSizeBytesNil(b bool)`

 SetSizeBytesNil sets the value for SizeBytes to be an explicit nil

### UnsetSizeBytes
`func (o *MongoDBCollection) UnsetSizeBytes()`

UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


