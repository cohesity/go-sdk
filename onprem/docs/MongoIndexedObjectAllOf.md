# MongoIndexedObjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | Specifies the Mongo Object Type. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the id of the indexed object. | [optional] 
**CdpInfo** | Pointer to [**CdpObjectInfo**](CdpObjectInfo.md) |  | [optional] 

## Methods

### NewMongoIndexedObjectAllOf

`func NewMongoIndexedObjectAllOf() *MongoIndexedObjectAllOf`

NewMongoIndexedObjectAllOf instantiates a new MongoIndexedObjectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoIndexedObjectAllOfWithDefaults

`func NewMongoIndexedObjectAllOfWithDefaults() *MongoIndexedObjectAllOf`

NewMongoIndexedObjectAllOfWithDefaults instantiates a new MongoIndexedObjectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MongoIndexedObjectAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MongoIndexedObjectAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MongoIndexedObjectAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MongoIndexedObjectAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MongoIndexedObjectAllOf) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MongoIndexedObjectAllOf) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *MongoIndexedObjectAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MongoIndexedObjectAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MongoIndexedObjectAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MongoIndexedObjectAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MongoIndexedObjectAllOf) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MongoIndexedObjectAllOf) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCdpInfo

`func (o *MongoIndexedObjectAllOf) GetCdpInfo() CdpObjectInfo`

GetCdpInfo returns the CdpInfo field if non-nil, zero value otherwise.

### GetCdpInfoOk

`func (o *MongoIndexedObjectAllOf) GetCdpInfoOk() (*CdpObjectInfo, bool)`

GetCdpInfoOk returns a tuple with the CdpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpInfo

`func (o *MongoIndexedObjectAllOf) SetCdpInfo(v CdpObjectInfo)`

SetCdpInfo sets CdpInfo field to given value.

### HasCdpInfo

`func (o *MongoIndexedObjectAllOf) HasCdpInfo() bool`

HasCdpInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


