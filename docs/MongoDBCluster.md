# MongoDBCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seeds** | Pointer to **[]string** | Seeds of this MongoDB Cluster. | [optional] 

## Methods

### NewMongoDBCluster

`func NewMongoDBCluster() *MongoDBCluster`

NewMongoDBCluster instantiates a new MongoDBCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBClusterWithDefaults

`func NewMongoDBClusterWithDefaults() *MongoDBCluster`

NewMongoDBClusterWithDefaults instantiates a new MongoDBCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeeds

`func (o *MongoDBCluster) GetSeeds() []string`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *MongoDBCluster) GetSeedsOk() (*[]string, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *MongoDBCluster) SetSeeds(v []string)`

SetSeeds sets Seeds field to given value.

### HasSeeds

`func (o *MongoDBCluster) HasSeeds() bool`

HasSeeds returns a boolean if a field has been set.

### SetSeedsNil

`func (o *MongoDBCluster) SetSeedsNil(b bool)`

 SetSeedsNil sets the value for Seeds to be an explicit nil

### UnsetSeeds
`func (o *MongoDBCluster) UnsetSeeds()`

UnsetSeeds ensures that no value is present for Seeds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


