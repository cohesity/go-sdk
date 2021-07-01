# CifsShareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acls** | Pointer to **[]string** | Array of Access Control Lists.  Specifies the ACLs for this share. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the CIFS share. This can be different from the volume name that this share belongs to. A single volume can export multiple CIFS shares, each with unique settings such as permissions. | [optional] 
**Path** | Pointer to **NullableString** | Specifies the path of this share under the Vserver&#39;s root. | [optional] 
**ServerName** | Pointer to **NullableString** | Specifies the CIFS server name (such as &#39;NETAPP-01&#39;) specified by the system administrator. This name is searchable within the active directory domain. | [optional] 

## Methods

### NewCifsShareInfo

`func NewCifsShareInfo() *CifsShareInfo`

NewCifsShareInfo instantiates a new CifsShareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCifsShareInfoWithDefaults

`func NewCifsShareInfoWithDefaults() *CifsShareInfo`

NewCifsShareInfoWithDefaults instantiates a new CifsShareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcls

`func (o *CifsShareInfo) GetAcls() []string`

GetAcls returns the Acls field if non-nil, zero value otherwise.

### GetAclsOk

`func (o *CifsShareInfo) GetAclsOk() (*[]string, bool)`

GetAclsOk returns a tuple with the Acls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcls

`func (o *CifsShareInfo) SetAcls(v []string)`

SetAcls sets Acls field to given value.

### HasAcls

`func (o *CifsShareInfo) HasAcls() bool`

HasAcls returns a boolean if a field has been set.

### SetAclsNil

`func (o *CifsShareInfo) SetAclsNil(b bool)`

 SetAclsNil sets the value for Acls to be an explicit nil

### UnsetAcls
`func (o *CifsShareInfo) UnsetAcls()`

UnsetAcls ensures that no value is present for Acls, not even an explicit nil
### GetName

`func (o *CifsShareInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CifsShareInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CifsShareInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CifsShareInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CifsShareInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CifsShareInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *CifsShareInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CifsShareInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CifsShareInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CifsShareInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *CifsShareInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *CifsShareInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetServerName

`func (o *CifsShareInfo) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *CifsShareInfo) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *CifsShareInfo) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *CifsShareInfo) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *CifsShareInfo) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *CifsShareInfo) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


