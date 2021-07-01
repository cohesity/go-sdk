# SanStorageArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies a unique id of a SAN Storage Array. The id is unique across Cohesity Clusters. | [optional] 
**Ports** | Pointer to [**[]IscsiSanPort**](IscsiSanPort.md) | Specifies the SAN ports of the SAN Storage Array. | [optional] 
**Revision** | Pointer to **NullableString** | Specifies the revision of the SAN Storage Array. | [optional] 
**Version** | Pointer to **NullableString** | Specifies the version of the SAN Storage Array. | [optional] 

## Methods

### NewSanStorageArray

`func NewSanStorageArray() *SanStorageArray`

NewSanStorageArray instantiates a new SanStorageArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSanStorageArrayWithDefaults

`func NewSanStorageArrayWithDefaults() *SanStorageArray`

NewSanStorageArrayWithDefaults instantiates a new SanStorageArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SanStorageArray) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SanStorageArray) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SanStorageArray) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SanStorageArray) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SanStorageArray) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SanStorageArray) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPorts

`func (o *SanStorageArray) GetPorts() []IscsiSanPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *SanStorageArray) GetPortsOk() (*[]IscsiSanPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *SanStorageArray) SetPorts(v []IscsiSanPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *SanStorageArray) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *SanStorageArray) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *SanStorageArray) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetRevision

`func (o *SanStorageArray) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SanStorageArray) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SanStorageArray) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *SanStorageArray) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### SetRevisionNil

`func (o *SanStorageArray) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *SanStorageArray) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetVersion

`func (o *SanStorageArray) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SanStorageArray) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SanStorageArray) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SanStorageArray) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *SanStorageArray) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *SanStorageArray) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


