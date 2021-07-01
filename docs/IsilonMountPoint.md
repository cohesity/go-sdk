# IsilonMountPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessZoneName** | Pointer to **NullableString** | Specifies the name of access zone. | [optional] 
**NfsMountPoint** | Pointer to [**IsilonNfsMountPoint**](IsilonNfsMountPoint.md) |  | [optional] 
**Path** | Pointer to **NullableString** | Specifies the path of the access zone in ifs. This should include the leading \&quot;/ifs/\&quot;. | [optional] 
**Protocols** | Pointer to **[]string** | List of Protocols on Isilon.  Specifies the list of protocols enabled on Isilon OneFs file system. &#39;kNfs&#39; indicates NFS exports in an Isilon Cluster. &#39;kSmb&#39; indicates CIFS/SMB Shares in an Isilon Cluster. | [optional] 
**SmbMountPoints** | Pointer to [**[]IsilonSmbMountPoint**](IsilonSmbMountPoint.md) | Specifies information about an SMB share. This field is set if the file system supports protocol type &#39;kSmb&#39;. | [optional] 

## Methods

### NewIsilonMountPoint

`func NewIsilonMountPoint() *IsilonMountPoint`

NewIsilonMountPoint instantiates a new IsilonMountPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsilonMountPointWithDefaults

`func NewIsilonMountPointWithDefaults() *IsilonMountPoint`

NewIsilonMountPointWithDefaults instantiates a new IsilonMountPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessZoneName

`func (o *IsilonMountPoint) GetAccessZoneName() string`

GetAccessZoneName returns the AccessZoneName field if non-nil, zero value otherwise.

### GetAccessZoneNameOk

`func (o *IsilonMountPoint) GetAccessZoneNameOk() (*string, bool)`

GetAccessZoneNameOk returns a tuple with the AccessZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessZoneName

`func (o *IsilonMountPoint) SetAccessZoneName(v string)`

SetAccessZoneName sets AccessZoneName field to given value.

### HasAccessZoneName

`func (o *IsilonMountPoint) HasAccessZoneName() bool`

HasAccessZoneName returns a boolean if a field has been set.

### SetAccessZoneNameNil

`func (o *IsilonMountPoint) SetAccessZoneNameNil(b bool)`

 SetAccessZoneNameNil sets the value for AccessZoneName to be an explicit nil

### UnsetAccessZoneName
`func (o *IsilonMountPoint) UnsetAccessZoneName()`

UnsetAccessZoneName ensures that no value is present for AccessZoneName, not even an explicit nil
### GetNfsMountPoint

`func (o *IsilonMountPoint) GetNfsMountPoint() IsilonNfsMountPoint`

GetNfsMountPoint returns the NfsMountPoint field if non-nil, zero value otherwise.

### GetNfsMountPointOk

`func (o *IsilonMountPoint) GetNfsMountPointOk() (*IsilonNfsMountPoint, bool)`

GetNfsMountPointOk returns a tuple with the NfsMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsMountPoint

`func (o *IsilonMountPoint) SetNfsMountPoint(v IsilonNfsMountPoint)`

SetNfsMountPoint sets NfsMountPoint field to given value.

### HasNfsMountPoint

`func (o *IsilonMountPoint) HasNfsMountPoint() bool`

HasNfsMountPoint returns a boolean if a field has been set.

### GetPath

`func (o *IsilonMountPoint) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IsilonMountPoint) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IsilonMountPoint) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IsilonMountPoint) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *IsilonMountPoint) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *IsilonMountPoint) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetProtocols

`func (o *IsilonMountPoint) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *IsilonMountPoint) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *IsilonMountPoint) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *IsilonMountPoint) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### SetProtocolsNil

`func (o *IsilonMountPoint) SetProtocolsNil(b bool)`

 SetProtocolsNil sets the value for Protocols to be an explicit nil

### UnsetProtocols
`func (o *IsilonMountPoint) UnsetProtocols()`

UnsetProtocols ensures that no value is present for Protocols, not even an explicit nil
### GetSmbMountPoints

`func (o *IsilonMountPoint) GetSmbMountPoints() []IsilonSmbMountPoint`

GetSmbMountPoints returns the SmbMountPoints field if non-nil, zero value otherwise.

### GetSmbMountPointsOk

`func (o *IsilonMountPoint) GetSmbMountPointsOk() (*[]IsilonSmbMountPoint, bool)`

GetSmbMountPointsOk returns a tuple with the SmbMountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbMountPoints

`func (o *IsilonMountPoint) SetSmbMountPoints(v []IsilonSmbMountPoint)`

SetSmbMountPoints sets SmbMountPoints field to given value.

### HasSmbMountPoints

`func (o *IsilonMountPoint) HasSmbMountPoints() bool`

HasSmbMountPoints returns a boolean if a field has been set.

### SetSmbMountPointsNil

`func (o *IsilonMountPoint) SetSmbMountPointsNil(b bool)`

 SetSmbMountPointsNil sets the value for SmbMountPoints to be an explicit nil

### UnsetSmbMountPoints
`func (o *IsilonMountPoint) UnsetSmbMountPoints()`

UnsetSmbMountPoints ensures that no value is present for SmbMountPoints, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


