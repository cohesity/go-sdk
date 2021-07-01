# GpfsFileset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Specifies the id of the fileset. | [optional] 
**IsIndependentFileset** | Pointer to **NullableBool** | If the given fileset is an Independent fileset or not. | [optional] 
**Name** | Pointer to **NullableString** | Name of the filesystem associated with the fileset | [optional] 
**Path** | Pointer to **NullableString** | Specifies the absolute path of the fileset. | [optional] 
**Protocols** | Pointer to **[]string** | Specifies GPFS supported Protocol information enabled on GPFS File System &#39;kNfs&#39; indicates NFS exports in a GPFS fileset. &#39;kSmb&#39; indicates CIFS/SMB Shares in a GPFS fileset. | [optional] 

## Methods

### NewGpfsFileset

`func NewGpfsFileset() *GpfsFileset`

NewGpfsFileset instantiates a new GpfsFileset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpfsFilesetWithDefaults

`func NewGpfsFilesetWithDefaults() *GpfsFileset`

NewGpfsFilesetWithDefaults instantiates a new GpfsFileset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GpfsFileset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GpfsFileset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GpfsFileset) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GpfsFileset) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GpfsFileset) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GpfsFileset) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsIndependentFileset

`func (o *GpfsFileset) GetIsIndependentFileset() bool`

GetIsIndependentFileset returns the IsIndependentFileset field if non-nil, zero value otherwise.

### GetIsIndependentFilesetOk

`func (o *GpfsFileset) GetIsIndependentFilesetOk() (*bool, bool)`

GetIsIndependentFilesetOk returns a tuple with the IsIndependentFileset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIndependentFileset

`func (o *GpfsFileset) SetIsIndependentFileset(v bool)`

SetIsIndependentFileset sets IsIndependentFileset field to given value.

### HasIsIndependentFileset

`func (o *GpfsFileset) HasIsIndependentFileset() bool`

HasIsIndependentFileset returns a boolean if a field has been set.

### SetIsIndependentFilesetNil

`func (o *GpfsFileset) SetIsIndependentFilesetNil(b bool)`

 SetIsIndependentFilesetNil sets the value for IsIndependentFileset to be an explicit nil

### UnsetIsIndependentFileset
`func (o *GpfsFileset) UnsetIsIndependentFileset()`

UnsetIsIndependentFileset ensures that no value is present for IsIndependentFileset, not even an explicit nil
### GetName

`func (o *GpfsFileset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GpfsFileset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GpfsFileset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GpfsFileset) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GpfsFileset) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GpfsFileset) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *GpfsFileset) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GpfsFileset) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GpfsFileset) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GpfsFileset) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *GpfsFileset) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *GpfsFileset) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetProtocols

`func (o *GpfsFileset) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *GpfsFileset) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *GpfsFileset) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *GpfsFileset) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### SetProtocolsNil

`func (o *GpfsFileset) SetProtocolsNil(b bool)`

 SetProtocolsNil sets the value for Protocols to be an explicit nil

### UnsetProtocols
`func (o *GpfsFileset) UnsetProtocols()`

UnsetProtocols ensures that no value is present for Protocols, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


