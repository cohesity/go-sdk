# MongoDBProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterInfo** | Pointer to [**MongoDBCluster**](MongoDBCluster.md) |  | [optional] 
**CollectionInfo** | Pointer to [**MongoDBCollection**](MongoDBCollection.md) |  | [optional] 
**DatabaseInfo** | Pointer to [**MongoDBDatabase**](MongoDBDatabase.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Specifies the instance name of the MongoDB entity. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the managed Object in MongoDB Protection Source. Specifies the type of an MongoDB source entity. &#39;kCluster&#39; indicates a mongodb cluster distributed over several physical nodes. &#39;kDatabase&#39; indicates a Database within the MongoDB environment. &#39;kCollection&#39; indicates a Collection in the MongoDB enironment. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID for the MongoDB entity. | [optional] 

## Methods

### NewMongoDBProtectionSource

`func NewMongoDBProtectionSource() *MongoDBProtectionSource`

NewMongoDBProtectionSource instantiates a new MongoDBProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBProtectionSourceWithDefaults

`func NewMongoDBProtectionSourceWithDefaults() *MongoDBProtectionSource`

NewMongoDBProtectionSourceWithDefaults instantiates a new MongoDBProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterInfo

`func (o *MongoDBProtectionSource) GetClusterInfo() MongoDBCluster`

GetClusterInfo returns the ClusterInfo field if non-nil, zero value otherwise.

### GetClusterInfoOk

`func (o *MongoDBProtectionSource) GetClusterInfoOk() (*MongoDBCluster, bool)`

GetClusterInfoOk returns a tuple with the ClusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterInfo

`func (o *MongoDBProtectionSource) SetClusterInfo(v MongoDBCluster)`

SetClusterInfo sets ClusterInfo field to given value.

### HasClusterInfo

`func (o *MongoDBProtectionSource) HasClusterInfo() bool`

HasClusterInfo returns a boolean if a field has been set.

### GetCollectionInfo

`func (o *MongoDBProtectionSource) GetCollectionInfo() MongoDBCollection`

GetCollectionInfo returns the CollectionInfo field if non-nil, zero value otherwise.

### GetCollectionInfoOk

`func (o *MongoDBProtectionSource) GetCollectionInfoOk() (*MongoDBCollection, bool)`

GetCollectionInfoOk returns a tuple with the CollectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionInfo

`func (o *MongoDBProtectionSource) SetCollectionInfo(v MongoDBCollection)`

SetCollectionInfo sets CollectionInfo field to given value.

### HasCollectionInfo

`func (o *MongoDBProtectionSource) HasCollectionInfo() bool`

HasCollectionInfo returns a boolean if a field has been set.

### GetDatabaseInfo

`func (o *MongoDBProtectionSource) GetDatabaseInfo() MongoDBDatabase`

GetDatabaseInfo returns the DatabaseInfo field if non-nil, zero value otherwise.

### GetDatabaseInfoOk

`func (o *MongoDBProtectionSource) GetDatabaseInfoOk() (*MongoDBDatabase, bool)`

GetDatabaseInfoOk returns a tuple with the DatabaseInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseInfo

`func (o *MongoDBProtectionSource) SetDatabaseInfo(v MongoDBDatabase)`

SetDatabaseInfo sets DatabaseInfo field to given value.

### HasDatabaseInfo

`func (o *MongoDBProtectionSource) HasDatabaseInfo() bool`

HasDatabaseInfo returns a boolean if a field has been set.

### GetName

`func (o *MongoDBProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MongoDBProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MongoDBProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MongoDBProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MongoDBProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MongoDBProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *MongoDBProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MongoDBProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MongoDBProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MongoDBProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MongoDBProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MongoDBProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUuid

`func (o *MongoDBProtectionSource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MongoDBProtectionSource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MongoDBProtectionSource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MongoDBProtectionSource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *MongoDBProtectionSource) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *MongoDBProtectionSource) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


