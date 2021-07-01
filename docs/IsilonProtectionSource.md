# IsilonProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessZone** | Pointer to [**IsilonAccessZone**](IsilonAccessZone.md) |  | [optional] 
**Cluster** | Pointer to [**IsilonCluster**](IsilonCluster.md) |  | [optional] 
**MountPoint** | Pointer to [**IsilonMountPoint**](IsilonMountPoint.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Specifies a unique name of the Protection Source. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the entity in an Isilon OneFs file system like &#39;kCluster&#39;, &#39;kZone&#39;, or, &#39;kMountPoint&#39;. &#39;kCluster&#39; indicates an Isilon OneFs Cluster. &#39;kZone&#39; indicates an access zone in an Isilon OneFs Cluster. &#39;kMountPoint&#39; indicates a mount point exposed by an Isilon OneFs Cluster. | [optional] 

## Methods

### NewIsilonProtectionSource

`func NewIsilonProtectionSource() *IsilonProtectionSource`

NewIsilonProtectionSource instantiates a new IsilonProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsilonProtectionSourceWithDefaults

`func NewIsilonProtectionSourceWithDefaults() *IsilonProtectionSource`

NewIsilonProtectionSourceWithDefaults instantiates a new IsilonProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessZone

`func (o *IsilonProtectionSource) GetAccessZone() IsilonAccessZone`

GetAccessZone returns the AccessZone field if non-nil, zero value otherwise.

### GetAccessZoneOk

`func (o *IsilonProtectionSource) GetAccessZoneOk() (*IsilonAccessZone, bool)`

GetAccessZoneOk returns a tuple with the AccessZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessZone

`func (o *IsilonProtectionSource) SetAccessZone(v IsilonAccessZone)`

SetAccessZone sets AccessZone field to given value.

### HasAccessZone

`func (o *IsilonProtectionSource) HasAccessZone() bool`

HasAccessZone returns a boolean if a field has been set.

### GetCluster

`func (o *IsilonProtectionSource) GetCluster() IsilonCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *IsilonProtectionSource) GetClusterOk() (*IsilonCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *IsilonProtectionSource) SetCluster(v IsilonCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *IsilonProtectionSource) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetMountPoint

`func (o *IsilonProtectionSource) GetMountPoint() IsilonMountPoint`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *IsilonProtectionSource) GetMountPointOk() (*IsilonMountPoint, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *IsilonProtectionSource) SetMountPoint(v IsilonMountPoint)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *IsilonProtectionSource) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### GetName

`func (o *IsilonProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IsilonProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IsilonProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IsilonProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IsilonProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IsilonProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *IsilonProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IsilonProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IsilonProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IsilonProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *IsilonProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *IsilonProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


