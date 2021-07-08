# AcropolisProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterUuid** | Pointer to **NullableString** | Specifies the UUID of the Acropolis cluster instance to which this entity belongs to. | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description about the Protection Source. | [optional] 
**MountPath** | Pointer to **NullableBool** | Specifies whether the VM is an agent VM. This is applicable to acropolis entity of type kVirtualMachine. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Acropolis Object. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of an Acropolis Protection Source Object such as &#39;kPrismCentral&#39;, &#39;kHost&#39;, &#39;kNetwork&#39;, etc. Specifies the type of an Acropolis source entity. &#39;kPrismCentral&#39; indicates a collection of multiple Nutanix clusters. &#39;kStandaloneCluster&#39; indicates a single Nutanix cluster. &#39;kCluster&#39; indicates a Nutanix cluster managed by a Prism Central. &#39;kHost&#39; indicates an Acropolis host. &#39;kVirtualMachine&#39; indicates a Virtual Machine. &#39;kNetwork&#39; indicates a Virtual Machine network object. &#39;kStorageContainer&#39; represents a storage container object. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID of the Acropolis Object. This is unique within the cluster instance. Together with clusterUuid, this entity is unique within the Acropolis environment. | [optional] 
**Version** | Pointer to **NullableString** | Specifies the version of an Acropolis cluster or standalone cluster. | [optional] 

## Methods

### NewAcropolisProtectionSource

`func NewAcropolisProtectionSource() *AcropolisProtectionSource`

NewAcropolisProtectionSource instantiates a new AcropolisProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcropolisProtectionSourceWithDefaults

`func NewAcropolisProtectionSourceWithDefaults() *AcropolisProtectionSource`

NewAcropolisProtectionSourceWithDefaults instantiates a new AcropolisProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterUuid

`func (o *AcropolisProtectionSource) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *AcropolisProtectionSource) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *AcropolisProtectionSource) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *AcropolisProtectionSource) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### SetClusterUuidNil

`func (o *AcropolisProtectionSource) SetClusterUuidNil(b bool)`

 SetClusterUuidNil sets the value for ClusterUuid to be an explicit nil

### UnsetClusterUuid
`func (o *AcropolisProtectionSource) UnsetClusterUuid()`

UnsetClusterUuid ensures that no value is present for ClusterUuid, not even an explicit nil
### GetDescription

`func (o *AcropolisProtectionSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AcropolisProtectionSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AcropolisProtectionSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AcropolisProtectionSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AcropolisProtectionSource) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AcropolisProtectionSource) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMountPath

`func (o *AcropolisProtectionSource) GetMountPath() bool`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *AcropolisProtectionSource) GetMountPathOk() (*bool, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *AcropolisProtectionSource) SetMountPath(v bool)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *AcropolisProtectionSource) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### SetMountPathNil

`func (o *AcropolisProtectionSource) SetMountPathNil(b bool)`

 SetMountPathNil sets the value for MountPath to be an explicit nil

### UnsetMountPath
`func (o *AcropolisProtectionSource) UnsetMountPath()`

UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
### GetName

`func (o *AcropolisProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcropolisProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcropolisProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AcropolisProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AcropolisProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AcropolisProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *AcropolisProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AcropolisProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AcropolisProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AcropolisProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AcropolisProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AcropolisProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUuid

`func (o *AcropolisProtectionSource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AcropolisProtectionSource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AcropolisProtectionSource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AcropolisProtectionSource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *AcropolisProtectionSource) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *AcropolisProtectionSource) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVersion

`func (o *AcropolisProtectionSource) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AcropolisProtectionSource) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AcropolisProtectionSource) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AcropolisProtectionSource) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *AcropolisProtectionSource) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *AcropolisProtectionSource) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


