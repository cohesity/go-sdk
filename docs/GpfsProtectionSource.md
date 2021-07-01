# GpfsProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to [**GpfsCluster**](GpfsCluster.md) |  | [optional] 
**Fileset** | Pointer to [**GpfsFileset**](GpfsFileset.md) |  | [optional] 
**Filesystem** | Pointer to [**GpfsFilesystem**](GpfsFilesystem.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Specifies a unique name of the Protection Source. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the entity in an GPFS file system like &#39;kCluster&#39;, &#39;kFilesystem&#39;, or, &#39;kFileset&#39;. &#39;kCluster&#39; indicates an GPFS Cluster. &#39;kFilesystem&#39; indicates a top level filesystem on GPFS cluster. &#39;kFileset&#39; indicates a fileset within a filesystem. | [optional] 

## Methods

### NewGpfsProtectionSource

`func NewGpfsProtectionSource() *GpfsProtectionSource`

NewGpfsProtectionSource instantiates a new GpfsProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpfsProtectionSourceWithDefaults

`func NewGpfsProtectionSourceWithDefaults() *GpfsProtectionSource`

NewGpfsProtectionSourceWithDefaults instantiates a new GpfsProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *GpfsProtectionSource) GetCluster() GpfsCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *GpfsProtectionSource) GetClusterOk() (*GpfsCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *GpfsProtectionSource) SetCluster(v GpfsCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *GpfsProtectionSource) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetFileset

`func (o *GpfsProtectionSource) GetFileset() GpfsFileset`

GetFileset returns the Fileset field if non-nil, zero value otherwise.

### GetFilesetOk

`func (o *GpfsProtectionSource) GetFilesetOk() (*GpfsFileset, bool)`

GetFilesetOk returns a tuple with the Fileset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileset

`func (o *GpfsProtectionSource) SetFileset(v GpfsFileset)`

SetFileset sets Fileset field to given value.

### HasFileset

`func (o *GpfsProtectionSource) HasFileset() bool`

HasFileset returns a boolean if a field has been set.

### GetFilesystem

`func (o *GpfsProtectionSource) GetFilesystem() GpfsFilesystem`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *GpfsProtectionSource) GetFilesystemOk() (*GpfsFilesystem, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *GpfsProtectionSource) SetFilesystem(v GpfsFilesystem)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *GpfsProtectionSource) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.

### GetName

`func (o *GpfsProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GpfsProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GpfsProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GpfsProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GpfsProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GpfsProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *GpfsProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GpfsProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GpfsProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GpfsProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GpfsProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GpfsProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


