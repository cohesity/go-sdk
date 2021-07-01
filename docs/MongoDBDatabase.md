# MongoDBDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeBytes** | Pointer to **NullableInt64** | Size of this Database. | [optional] 

## Methods

### NewMongoDBDatabase

`func NewMongoDBDatabase() *MongoDBDatabase`

NewMongoDBDatabase instantiates a new MongoDBDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBDatabaseWithDefaults

`func NewMongoDBDatabaseWithDefaults() *MongoDBDatabase`

NewMongoDBDatabaseWithDefaults instantiates a new MongoDBDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeBytes

`func (o *MongoDBDatabase) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *MongoDBDatabase) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *MongoDBDatabase) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *MongoDBDatabase) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### SetSizeBytesNil

`func (o *MongoDBDatabase) SetSizeBytesNil(b bool)`

 SetSizeBytesNil sets the value for SizeBytes to be an explicit nil

### UnsetSizeBytes
`func (o *MongoDBDatabase) UnsetSizeBytes()`

UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


