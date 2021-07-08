# ViewIdMappingProtoProtocolAccessInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IscsiAccess** | Pointer to **NullableInt32** | Access control for iSCSI protocol for this view. | [optional] 
**NfsAccess** | Pointer to **NullableInt32** | Access control for NFS protocol for this view. | [optional] 
**S3Access** | Pointer to **NullableInt32** | Access control for S3 protocol for this view. | [optional] 
**SmbAccess** | Pointer to **NullableInt32** | Access control for SMB protocol for this view. | [optional] 
**SwiftAccess** | Pointer to **NullableInt32** | Access control for Swift protocol for this view. | [optional] 

## Methods

### NewViewIdMappingProtoProtocolAccessInfo

`func NewViewIdMappingProtoProtocolAccessInfo() *ViewIdMappingProtoProtocolAccessInfo`

NewViewIdMappingProtoProtocolAccessInfo instantiates a new ViewIdMappingProtoProtocolAccessInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewIdMappingProtoProtocolAccessInfoWithDefaults

`func NewViewIdMappingProtoProtocolAccessInfoWithDefaults() *ViewIdMappingProtoProtocolAccessInfo`

NewViewIdMappingProtoProtocolAccessInfoWithDefaults instantiates a new ViewIdMappingProtoProtocolAccessInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIscsiAccess

`func (o *ViewIdMappingProtoProtocolAccessInfo) GetIscsiAccess() int32`

GetIscsiAccess returns the IscsiAccess field if non-nil, zero value otherwise.

### GetIscsiAccessOk

`func (o *ViewIdMappingProtoProtocolAccessInfo) GetIscsiAccessOk() (*int32, bool)`

GetIscsiAccessOk returns a tuple with the IscsiAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiAccess

`func (o *ViewIdMappingProtoProtocolAccessInfo) SetIscsiAccess(v int32)`

SetIscsiAccess sets IscsiAccess field to given value.

### HasIscsiAccess

`func (o *ViewIdMappingProtoProtocolAccessInfo) HasIscsiAccess() bool`

HasIscsiAccess returns a boolean if a field has been set.

### SetIscsiAccessNil

`func (o *ViewIdMappingProtoProtocolAccessInfo) SetIscsiAccessNil(b bool)`

 SetIscsiAccessNil sets the value for IscsiAccess to be an explicit nil

### UnsetIscsiAccess
`func (o *ViewIdMappingProtoProtocolAccessInfo) UnsetIscsiAccess()`

UnsetIscsiAccess ensures that no value is present for IscsiAccess, not even an explicit nil
### GetNfsAccess

`func (o *ViewIdMappingProtoProtocolAccessInfo) GetNfsAccess() int32`

GetNfsAccess returns the NfsAccess field if non-nil, zero value otherwise.

### GetNfsAccessOk

`func (o *ViewIdMappingProtoProtocolAccessInfo) GetNfsAccessOk() (*int32, bool)`

GetNfsAccessOk returns a tuple with the NfsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAccess

`func (o *ViewIdMappingProtoProtocolAccessInfo) SetNfsAccess(v int32)`

SetNfsAccess sets NfsAccess field to given value.

### HasNfsAccess

`func (o *ViewIdMappingProtoProtocolAccessInfo) HasNfsAccess() bool`

HasNfsAccess returns a boolean if a field has been set.

### SetNfsAccessNil

`func (o *ViewIdMappingProtoProtocolAccessInfo) SetNfsAccessNil(b bool)`

 SetNfsAccessNil sets the value for NfsAccess to be an explicit nil

### UnsetNfsAccess
`func (o *ViewIdMappingProtoProtocolAccessInfo) UnsetNfsAccess()`

UnsetNfsAccess ensures that no value is present for NfsAccess, not even an explicit nil
### GetS3Access

`func (o *ViewIdMappingProtoProtocolAccessInfo) GetS3Access() int32`

GetS3Access returns the S3Access field if non-nil, zero value otherwise.

### GetS3AccessOk

`func (o *ViewIdMappingProtoProtocolAccessInfo) GetS3AccessOk() (*int32, bool)`

GetS3AccessOk returns a tuple with the S3Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Access

`func (o *ViewIdMappingProtoProtocolAccessInfo) SetS3Access(v int32)`

SetS3Access sets S3Access field to given value.

### HasS3Access

`func (o *ViewIdMappingProtoProtocolAccessInfo) HasS3Access() bool`

HasS3Access returns a boolean if a field has been set.

### SetS3AccessNil

`func (o *ViewIdMappingProtoProtocolAccessInfo) SetS3AccessNil(b bool)`

 SetS3AccessNil sets the value for S3Access to be an explicit nil

### UnsetS3Access
`func (o *ViewIdMappingProtoProtocolAccessInfo) UnsetS3Access()`

UnsetS3Access ensures that no value is present for S3Access, not even an explicit nil
### GetSmbAccess

`func (o *ViewIdMappingProtoProtocolAccessInfo) GetSmbAccess() int32`

GetSmbAccess returns the SmbAccess field if non-nil, zero value otherwise.

### GetSmbAccessOk

`func (o *ViewIdMappingProtoProtocolAccessInfo) GetSmbAccessOk() (*int32, bool)`

GetSmbAccessOk returns a tuple with the SmbAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbAccess

`func (o *ViewIdMappingProtoProtocolAccessInfo) SetSmbAccess(v int32)`

SetSmbAccess sets SmbAccess field to given value.

### HasSmbAccess

`func (o *ViewIdMappingProtoProtocolAccessInfo) HasSmbAccess() bool`

HasSmbAccess returns a boolean if a field has been set.

### SetSmbAccessNil

`func (o *ViewIdMappingProtoProtocolAccessInfo) SetSmbAccessNil(b bool)`

 SetSmbAccessNil sets the value for SmbAccess to be an explicit nil

### UnsetSmbAccess
`func (o *ViewIdMappingProtoProtocolAccessInfo) UnsetSmbAccess()`

UnsetSmbAccess ensures that no value is present for SmbAccess, not even an explicit nil
### GetSwiftAccess

`func (o *ViewIdMappingProtoProtocolAccessInfo) GetSwiftAccess() int32`

GetSwiftAccess returns the SwiftAccess field if non-nil, zero value otherwise.

### GetSwiftAccessOk

`func (o *ViewIdMappingProtoProtocolAccessInfo) GetSwiftAccessOk() (*int32, bool)`

GetSwiftAccessOk returns a tuple with the SwiftAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftAccess

`func (o *ViewIdMappingProtoProtocolAccessInfo) SetSwiftAccess(v int32)`

SetSwiftAccess sets SwiftAccess field to given value.

### HasSwiftAccess

`func (o *ViewIdMappingProtoProtocolAccessInfo) HasSwiftAccess() bool`

HasSwiftAccess returns a boolean if a field has been set.

### SetSwiftAccessNil

`func (o *ViewIdMappingProtoProtocolAccessInfo) SetSwiftAccessNil(b bool)`

 SetSwiftAccessNil sets the value for SwiftAccess to be an explicit nil

### UnsetSwiftAccess
`func (o *ViewIdMappingProtoProtocolAccessInfo) UnsetSwiftAccess()`

UnsetSwiftAccess ensures that no value is present for SwiftAccess, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


