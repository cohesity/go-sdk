# ElastifileProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**ElastifileCluster**](ElastifileCluster.md) |  | [optional] 
**Container** | Pointer to [**ElastifileContainer**](ElastifileContainer.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Specifies a unique name of the Protection Source. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the entity in an Elastifile file system like &#39;kCluster&#39;, &#39;kContainer&#39;. &#39;kCluster&#39; indicates an Elastifile Cluster. &#39;kContainer&#39; indicates a container on Elastifile cluster. | [optional] 

## Methods

### NewElastifileProtectionSource

`func NewElastifileProtectionSource() *ElastifileProtectionSource`

NewElastifileProtectionSource instantiates a new ElastifileProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElastifileProtectionSourceWithDefaults

`func NewElastifileProtectionSourceWithDefaults() *ElastifileProtectionSource`

NewElastifileProtectionSourceWithDefaults instantiates a new ElastifileProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ElastifileProtectionSource) GetCluster() ElastifileCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ElastifileProtectionSource) GetClusterOk() (*ElastifileCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ElastifileProtectionSource) SetCluster(v ElastifileCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ElastifileProtectionSource) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetContainer

`func (o *ElastifileProtectionSource) GetContainer() ElastifileContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ElastifileProtectionSource) GetContainerOk() (*ElastifileContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ElastifileProtectionSource) SetContainer(v ElastifileContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ElastifileProtectionSource) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetName

`func (o *ElastifileProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ElastifileProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ElastifileProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ElastifileProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ElastifileProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ElastifileProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *ElastifileProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ElastifileProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ElastifileProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ElastifileProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ElastifileProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ElastifileProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


