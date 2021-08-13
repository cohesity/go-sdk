# CouchbaseIndexedObjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | Specifies the Couchbase Object Type. For Couchbase this is alywas set to Bucket. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the id of the indexed object. | [optional] 

## Methods

### NewCouchbaseIndexedObjectAllOf

`func NewCouchbaseIndexedObjectAllOf() *CouchbaseIndexedObjectAllOf`

NewCouchbaseIndexedObjectAllOf instantiates a new CouchbaseIndexedObjectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouchbaseIndexedObjectAllOfWithDefaults

`func NewCouchbaseIndexedObjectAllOfWithDefaults() *CouchbaseIndexedObjectAllOf`

NewCouchbaseIndexedObjectAllOfWithDefaults instantiates a new CouchbaseIndexedObjectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouchbaseIndexedObjectAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouchbaseIndexedObjectAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouchbaseIndexedObjectAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CouchbaseIndexedObjectAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CouchbaseIndexedObjectAllOf) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CouchbaseIndexedObjectAllOf) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *CouchbaseIndexedObjectAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouchbaseIndexedObjectAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouchbaseIndexedObjectAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CouchbaseIndexedObjectAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CouchbaseIndexedObjectAllOf) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CouchbaseIndexedObjectAllOf) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


