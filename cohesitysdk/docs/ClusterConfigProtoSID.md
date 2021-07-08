# ClusterConfigProtoSID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentifierAuthority** | Pointer to **[]int32** | The authority under which the SID was created. This is always 6 bytes long. | [optional] 
**RevisionLevel** | Pointer to **NullableInt32** | The revision level of the SID. | [optional] 
**SubAuthority** | Pointer to **[]int32** | List of ids relative to the identifier_authority that uniquely identify a principal. The last entry in this list is the RID, which uniquely identifies the principal within a domain. | [optional] 

## Methods

### NewClusterConfigProtoSID

`func NewClusterConfigProtoSID() *ClusterConfigProtoSID`

NewClusterConfigProtoSID instantiates a new ClusterConfigProtoSID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigProtoSIDWithDefaults

`func NewClusterConfigProtoSIDWithDefaults() *ClusterConfigProtoSID`

NewClusterConfigProtoSIDWithDefaults instantiates a new ClusterConfigProtoSID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifierAuthority

`func (o *ClusterConfigProtoSID) GetIdentifierAuthority() []int32`

GetIdentifierAuthority returns the IdentifierAuthority field if non-nil, zero value otherwise.

### GetIdentifierAuthorityOk

`func (o *ClusterConfigProtoSID) GetIdentifierAuthorityOk() (*[]int32, bool)`

GetIdentifierAuthorityOk returns a tuple with the IdentifierAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierAuthority

`func (o *ClusterConfigProtoSID) SetIdentifierAuthority(v []int32)`

SetIdentifierAuthority sets IdentifierAuthority field to given value.

### HasIdentifierAuthority

`func (o *ClusterConfigProtoSID) HasIdentifierAuthority() bool`

HasIdentifierAuthority returns a boolean if a field has been set.

### SetIdentifierAuthorityNil

`func (o *ClusterConfigProtoSID) SetIdentifierAuthorityNil(b bool)`

 SetIdentifierAuthorityNil sets the value for IdentifierAuthority to be an explicit nil

### UnsetIdentifierAuthority
`func (o *ClusterConfigProtoSID) UnsetIdentifierAuthority()`

UnsetIdentifierAuthority ensures that no value is present for IdentifierAuthority, not even an explicit nil
### GetRevisionLevel

`func (o *ClusterConfigProtoSID) GetRevisionLevel() int32`

GetRevisionLevel returns the RevisionLevel field if non-nil, zero value otherwise.

### GetRevisionLevelOk

`func (o *ClusterConfigProtoSID) GetRevisionLevelOk() (*int32, bool)`

GetRevisionLevelOk returns a tuple with the RevisionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionLevel

`func (o *ClusterConfigProtoSID) SetRevisionLevel(v int32)`

SetRevisionLevel sets RevisionLevel field to given value.

### HasRevisionLevel

`func (o *ClusterConfigProtoSID) HasRevisionLevel() bool`

HasRevisionLevel returns a boolean if a field has been set.

### SetRevisionLevelNil

`func (o *ClusterConfigProtoSID) SetRevisionLevelNil(b bool)`

 SetRevisionLevelNil sets the value for RevisionLevel to be an explicit nil

### UnsetRevisionLevel
`func (o *ClusterConfigProtoSID) UnsetRevisionLevel()`

UnsetRevisionLevel ensures that no value is present for RevisionLevel, not even an explicit nil
### GetSubAuthority

`func (o *ClusterConfigProtoSID) GetSubAuthority() []int32`

GetSubAuthority returns the SubAuthority field if non-nil, zero value otherwise.

### GetSubAuthorityOk

`func (o *ClusterConfigProtoSID) GetSubAuthorityOk() (*[]int32, bool)`

GetSubAuthorityOk returns a tuple with the SubAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAuthority

`func (o *ClusterConfigProtoSID) SetSubAuthority(v []int32)`

SetSubAuthority sets SubAuthority field to given value.

### HasSubAuthority

`func (o *ClusterConfigProtoSID) HasSubAuthority() bool`

HasSubAuthority returns a boolean if a field has been set.

### SetSubAuthorityNil

`func (o *ClusterConfigProtoSID) SetSubAuthorityNil(b bool)`

 SetSubAuthorityNil sets the value for SubAuthority to be an explicit nil

### UnsetSubAuthority
`func (o *ClusterConfigProtoSID) UnsetSubAuthority()`

UnsetSubAuthority ensures that no value is present for SubAuthority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


