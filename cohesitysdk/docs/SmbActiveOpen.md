# SmbActiveOpen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessInfoList** | Pointer to **[]string** | Specifies the access information. &#39;kFileReadData&#39; indicates the right to read data from the file or named pipe. &#39;kFileWriteData&#39; indicates the right to write data into the file or named pipe beyond the end of the file. &#39;kFileAppendData&#39; indicates the right to append data into the file or named pipe. &#39;kFileReadEa&#39; indicates the right to read the extended attributes of the file or named pipe. &#39;kFileWriteEa&#39; indicates the right to write or change the extended attributes to the file or named pipe. &#39;kFileExecute&#39; indicates the right to delete entries within a directory. &#39;kFileDeleteChild&#39; indicates the right to execute the file. &#39;kFileReadAttributes&#39; indicates the right to read the attributes of the file. &#39;kFileWriteAttributes&#39; indicates the right to change the attributes of the file. &#39;kDelete&#39; indicates the right to delete the file. &#39;kReadControl&#39; indicates the right to read the security descriptor for the file or named pipe. &#39;kWriteDac&#39; indicates the right to change the discretionary access control list (DACL) in the security descriptor for the file or named pipe. For the DACL data structure, see ACL in [MS-DTYP]. &#39;kWriteOwner&#39; indicates the right to change the owner in the security descriptor for the file or named pipe. &#39;kSynchronize&#39; is used only by SMB2 clients. &#39;kAccessSystemSecurity&#39; indicates the right to read or change the system access control list (SACL) in the security descriptor for the file or named pipe. For the SACL data structure, see ACL in [MS-DTYP].&lt;42&gt; &#39;kMaximumAllowed&#39; indicates that the client is requesting an open to the file with the highest level of access the client has on this file. If no access is granted for the client on this file, the server MUST fail the open with STATUS_ACCESS_DENIED. &#39;kGenericAll&#39; indicates a request for all the access flags that are previously listed except kMaximumAllowed and kAccessSystemSecurity. &#39;kGenericExecute&#39; indicates a request for the following combination of access flags listed above: kFileReadAttributes| kFileExecute| kSynchronize| kReadControl. &#39;kGenericWrite&#39; indicates a request for the following combination of access flags listed above: kFileWriteData| kFileAppendData| kFileWriteAttributes| kFileWriteEa| kSynchronize| kReadControl. &#39;kGenericRead&#39; indicates a request for the following combination of access flags listed above: kFileReadData| kFileReadAttributes| kFileReadEa| kSynchronize| kReadControl. | [optional] 
**OpenId** | Pointer to **NullableInt64** | Specifies the id of the active open. | [optional] 
**OthersCanDelete** | Pointer to **NullableBool** | Specifies whether others are allowed to delete. | [optional] 
**OthersCanRead** | Pointer to **NullableBool** | Specifies whether others are allowed to read. | [optional] 
**OthersCanWrite** | Pointer to **NullableBool** | Specifies whether others are allowed to write. | [optional] 

## Methods

### NewSmbActiveOpen

`func NewSmbActiveOpen() *SmbActiveOpen`

NewSmbActiveOpen instantiates a new SmbActiveOpen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbActiveOpenWithDefaults

`func NewSmbActiveOpenWithDefaults() *SmbActiveOpen`

NewSmbActiveOpenWithDefaults instantiates a new SmbActiveOpen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessInfoList

`func (o *SmbActiveOpen) GetAccessInfoList() []string`

GetAccessInfoList returns the AccessInfoList field if non-nil, zero value otherwise.

### GetAccessInfoListOk

`func (o *SmbActiveOpen) GetAccessInfoListOk() (*[]string, bool)`

GetAccessInfoListOk returns a tuple with the AccessInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessInfoList

`func (o *SmbActiveOpen) SetAccessInfoList(v []string)`

SetAccessInfoList sets AccessInfoList field to given value.

### HasAccessInfoList

`func (o *SmbActiveOpen) HasAccessInfoList() bool`

HasAccessInfoList returns a boolean if a field has been set.

### SetAccessInfoListNil

`func (o *SmbActiveOpen) SetAccessInfoListNil(b bool)`

 SetAccessInfoListNil sets the value for AccessInfoList to be an explicit nil

### UnsetAccessInfoList
`func (o *SmbActiveOpen) UnsetAccessInfoList()`

UnsetAccessInfoList ensures that no value is present for AccessInfoList, not even an explicit nil
### GetOpenId

`func (o *SmbActiveOpen) GetOpenId() int64`

GetOpenId returns the OpenId field if non-nil, zero value otherwise.

### GetOpenIdOk

`func (o *SmbActiveOpen) GetOpenIdOk() (*int64, bool)`

GetOpenIdOk returns a tuple with the OpenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenId

`func (o *SmbActiveOpen) SetOpenId(v int64)`

SetOpenId sets OpenId field to given value.

### HasOpenId

`func (o *SmbActiveOpen) HasOpenId() bool`

HasOpenId returns a boolean if a field has been set.

### SetOpenIdNil

`func (o *SmbActiveOpen) SetOpenIdNil(b bool)`

 SetOpenIdNil sets the value for OpenId to be an explicit nil

### UnsetOpenId
`func (o *SmbActiveOpen) UnsetOpenId()`

UnsetOpenId ensures that no value is present for OpenId, not even an explicit nil
### GetOthersCanDelete

`func (o *SmbActiveOpen) GetOthersCanDelete() bool`

GetOthersCanDelete returns the OthersCanDelete field if non-nil, zero value otherwise.

### GetOthersCanDeleteOk

`func (o *SmbActiveOpen) GetOthersCanDeleteOk() (*bool, bool)`

GetOthersCanDeleteOk returns a tuple with the OthersCanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthersCanDelete

`func (o *SmbActiveOpen) SetOthersCanDelete(v bool)`

SetOthersCanDelete sets OthersCanDelete field to given value.

### HasOthersCanDelete

`func (o *SmbActiveOpen) HasOthersCanDelete() bool`

HasOthersCanDelete returns a boolean if a field has been set.

### SetOthersCanDeleteNil

`func (o *SmbActiveOpen) SetOthersCanDeleteNil(b bool)`

 SetOthersCanDeleteNil sets the value for OthersCanDelete to be an explicit nil

### UnsetOthersCanDelete
`func (o *SmbActiveOpen) UnsetOthersCanDelete()`

UnsetOthersCanDelete ensures that no value is present for OthersCanDelete, not even an explicit nil
### GetOthersCanRead

`func (o *SmbActiveOpen) GetOthersCanRead() bool`

GetOthersCanRead returns the OthersCanRead field if non-nil, zero value otherwise.

### GetOthersCanReadOk

`func (o *SmbActiveOpen) GetOthersCanReadOk() (*bool, bool)`

GetOthersCanReadOk returns a tuple with the OthersCanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthersCanRead

`func (o *SmbActiveOpen) SetOthersCanRead(v bool)`

SetOthersCanRead sets OthersCanRead field to given value.

### HasOthersCanRead

`func (o *SmbActiveOpen) HasOthersCanRead() bool`

HasOthersCanRead returns a boolean if a field has been set.

### SetOthersCanReadNil

`func (o *SmbActiveOpen) SetOthersCanReadNil(b bool)`

 SetOthersCanReadNil sets the value for OthersCanRead to be an explicit nil

### UnsetOthersCanRead
`func (o *SmbActiveOpen) UnsetOthersCanRead()`

UnsetOthersCanRead ensures that no value is present for OthersCanRead, not even an explicit nil
### GetOthersCanWrite

`func (o *SmbActiveOpen) GetOthersCanWrite() bool`

GetOthersCanWrite returns the OthersCanWrite field if non-nil, zero value otherwise.

### GetOthersCanWriteOk

`func (o *SmbActiveOpen) GetOthersCanWriteOk() (*bool, bool)`

GetOthersCanWriteOk returns a tuple with the OthersCanWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthersCanWrite

`func (o *SmbActiveOpen) SetOthersCanWrite(v bool)`

SetOthersCanWrite sets OthersCanWrite field to given value.

### HasOthersCanWrite

`func (o *SmbActiveOpen) HasOthersCanWrite() bool`

HasOthersCanWrite returns a boolean if a field has been set.

### SetOthersCanWriteNil

`func (o *SmbActiveOpen) SetOthersCanWriteNil(b bool)`

 SetOthersCanWriteNil sets the value for OthersCanWrite to be an explicit nil

### UnsetOthersCanWrite
`func (o *SmbActiveOpen) UnsetOthersCanWrite()`

UnsetOthersCanWrite ensures that no value is present for OthersCanWrite, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


