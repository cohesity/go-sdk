# CouchbaseProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketInfo** | Pointer to [**CouchbaseBucket**](CouchbaseBucket.md) |  | [optional] 
**ClusterInfo** | Pointer to [**CouchbaseCluster**](CouchbaseCluster.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Specifies the instance name of the Couchbase entity. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the managed Object in Couchbase Protection Source. Specifies the type of an Couchbase source entity. &#39;kCluster&#39; indicates a Couchbase cluster distributed over several physical nodes. &#39;kBucket&#39; indicates a bucket within the Couchbase environment. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID for the Couchbase entity. | [optional] 

## Methods

### NewCouchbaseProtectionSource

`func NewCouchbaseProtectionSource() *CouchbaseProtectionSource`

NewCouchbaseProtectionSource instantiates a new CouchbaseProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouchbaseProtectionSourceWithDefaults

`func NewCouchbaseProtectionSourceWithDefaults() *CouchbaseProtectionSource`

NewCouchbaseProtectionSourceWithDefaults instantiates a new CouchbaseProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketInfo

`func (o *CouchbaseProtectionSource) GetBucketInfo() CouchbaseBucket`

GetBucketInfo returns the BucketInfo field if non-nil, zero value otherwise.

### GetBucketInfoOk

`func (o *CouchbaseProtectionSource) GetBucketInfoOk() (*CouchbaseBucket, bool)`

GetBucketInfoOk returns a tuple with the BucketInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketInfo

`func (o *CouchbaseProtectionSource) SetBucketInfo(v CouchbaseBucket)`

SetBucketInfo sets BucketInfo field to given value.

### HasBucketInfo

`func (o *CouchbaseProtectionSource) HasBucketInfo() bool`

HasBucketInfo returns a boolean if a field has been set.

### GetClusterInfo

`func (o *CouchbaseProtectionSource) GetClusterInfo() CouchbaseCluster`

GetClusterInfo returns the ClusterInfo field if non-nil, zero value otherwise.

### GetClusterInfoOk

`func (o *CouchbaseProtectionSource) GetClusterInfoOk() (*CouchbaseCluster, bool)`

GetClusterInfoOk returns a tuple with the ClusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterInfo

`func (o *CouchbaseProtectionSource) SetClusterInfo(v CouchbaseCluster)`

SetClusterInfo sets ClusterInfo field to given value.

### HasClusterInfo

`func (o *CouchbaseProtectionSource) HasClusterInfo() bool`

HasClusterInfo returns a boolean if a field has been set.

### GetName

`func (o *CouchbaseProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CouchbaseProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CouchbaseProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CouchbaseProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CouchbaseProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CouchbaseProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *CouchbaseProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouchbaseProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouchbaseProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CouchbaseProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CouchbaseProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CouchbaseProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUuid

`func (o *CouchbaseProtectionSource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CouchbaseProtectionSource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CouchbaseProtectionSource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CouchbaseProtectionSource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *CouchbaseProtectionSource) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *CouchbaseProtectionSource) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


