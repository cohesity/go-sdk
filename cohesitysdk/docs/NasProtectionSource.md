# NasProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies a description about the Protection Source. | [optional] 
**MountPath** | Pointer to **NullableString** | Specifies the mount path of this NAS. For example, for a NFS mount point, this should be in the format of IP or hostname:/foo/bar. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the NetApp Object. | [optional] 
**Protocol** | Pointer to **NullableString** | Specifies the protocol used by the NAS server. Specifies the protocol used by a NAS server. &#39;kNfs3&#39; indicates NFS v3 protocol. &#39;kCifs1&#39; indicates CIFS v1.0 protocol. | [optional] 
**SkipValidation** | Pointer to **NullableBool** | Specifies whether to skip validation of the given mount point. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of a Protection Source Object in a generic NAS Source such as &#39;kGroup&#39;, or &#39;kHost&#39;. Specifies the kind of NAS mount. &#39;kGroup&#39; indicates top level node that holds individual NAS hosts. &#39;kHost&#39; indicates a single NAS path that can be mounted. &#39;kDfsGroup&#39; indicates a DFS group containing top level directories mapped to different servers. &#39;kDfsTopDir&#39; indicates a top level directory inside a DFS group, discovered when registering a DFS group. | [optional] 

## Methods

### NewNasProtectionSource

`func NewNasProtectionSource() *NasProtectionSource`

NewNasProtectionSource instantiates a new NasProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNasProtectionSourceWithDefaults

`func NewNasProtectionSourceWithDefaults() *NasProtectionSource`

NewNasProtectionSourceWithDefaults instantiates a new NasProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NasProtectionSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NasProtectionSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NasProtectionSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NasProtectionSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NasProtectionSource) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NasProtectionSource) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMountPath

`func (o *NasProtectionSource) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *NasProtectionSource) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *NasProtectionSource) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *NasProtectionSource) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### SetMountPathNil

`func (o *NasProtectionSource) SetMountPathNil(b bool)`

 SetMountPathNil sets the value for MountPath to be an explicit nil

### UnsetMountPath
`func (o *NasProtectionSource) UnsetMountPath()`

UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
### GetName

`func (o *NasProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NasProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NasProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NasProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NasProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NasProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProtocol

`func (o *NasProtectionSource) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NasProtectionSource) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NasProtectionSource) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NasProtectionSource) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *NasProtectionSource) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *NasProtectionSource) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetSkipValidation

`func (o *NasProtectionSource) GetSkipValidation() bool`

GetSkipValidation returns the SkipValidation field if non-nil, zero value otherwise.

### GetSkipValidationOk

`func (o *NasProtectionSource) GetSkipValidationOk() (*bool, bool)`

GetSkipValidationOk returns a tuple with the SkipValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipValidation

`func (o *NasProtectionSource) SetSkipValidation(v bool)`

SetSkipValidation sets SkipValidation field to given value.

### HasSkipValidation

`func (o *NasProtectionSource) HasSkipValidation() bool`

HasSkipValidation returns a boolean if a field has been set.

### SetSkipValidationNil

`func (o *NasProtectionSource) SetSkipValidationNil(b bool)`

 SetSkipValidationNil sets the value for SkipValidation to be an explicit nil

### UnsetSkipValidation
`func (o *NasProtectionSource) UnsetSkipValidation()`

UnsetSkipValidation ensures that no value is present for SkipValidation, not even an explicit nil
### GetType

`func (o *NasProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NasProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NasProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NasProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *NasProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NasProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


