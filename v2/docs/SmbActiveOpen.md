# SmbActiveOpen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessInfoList** | Pointer to **[]string** | Specifies the File Access Type. Following documentation was taken from MSDN. https://msdn.microsoft.com/en-us/library/Cc246802.aspx  &#39;FileReadData&#39; indicates the right to read data from the file or named   pipe. &#39;FileWriteData&#39; indicates the right to write data into the file or named   pipe beyond the end of the file. &#39;FileAppendData&#39; indicates the right to append data into the file or named   pipe. &#39;FileReadEa&#39; indicates the right to read the extended attributes of the   file or named pipe. &#39;FileWriteEa&#39; indicates the right to write or change the extended   attributes to the file or named pipe. &#39;FileExecute&#39; indicates the right to delete entries within a directory. &#39;FileDeleteChild&#39; indicates the right to execute the file. &#39;FileReadAttributes&#39; indicates the right to read the attributes of the   file. &#39;FileWriteAttributes&#39; indicates the right to change the attributes of the   file. &#39;Delete&#39; indicates the right to delete the file. &#39;ReadControl&#39; indicates the right to read the security descriptor for the   file or named pipe. &#39;WriteDac&#39; indicates the right to change the discretionary access control   list (DACL) in the security descriptor for the file or named pipe. For   the DACL data structure, see ACL in [MS-DTYP]. &#39;WriteOwner&#39; indicates the right to change the owner in the security   descriptor for the file or named pipe. &#39;Synchronize&#39; is used only by SMB2 clients. &#39;AccessSystemSecurity&#39; indicates the right to read or change the system   access control list (SACL) in the security descriptor for the file or   named pipe. For the SACL data structure, see ACL in [MS-DTYP].&lt;42&gt; &#39;MaximumAllowed&#39; indicates that the client is requesting an open to the   file with the highest level of access the client has on this file.   If no access is granted for the client on this file, the server MUST   fail the open with STATUS_ACCESS_DENIED. &#39;GenericAll&#39; indicates a request for all the access flags that are   previously listed except MaximumAllowed and AccessSystemSecurity. &#39;GenericExecute&#39; indicates a request for the following combination of   access flags listed above:   FileReadAttributes| FileExecute| Synchronize| ReadControl. &#39;GenericWrite&#39; indicates a request for the following combination of   access flags listed above:   FileWriteData| FileAppendData| FileWriteAttributes| FileWriteEa|   Synchronize| ReadControl. &#39;GenericRead&#39; indicates a request for the following combination of   access flags listed above:   FileReadData| FileReadAttributes| FileReadEa| Synchronize|   ReadControl. | [optional] 
**AccessPrivilege** | Pointer to **[]string** | Specifies whether access privilege of others if they&#39;re allowed to read/write/delete. | [optional] 
**OpenId** | Pointer to **NullableInt64** | Specifies the id of the active open. | [optional] 

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
### GetAccessPrivilege

`func (o *SmbActiveOpen) GetAccessPrivilege() []string`

GetAccessPrivilege returns the AccessPrivilege field if non-nil, zero value otherwise.

### GetAccessPrivilegeOk

`func (o *SmbActiveOpen) GetAccessPrivilegeOk() (*[]string, bool)`

GetAccessPrivilegeOk returns a tuple with the AccessPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPrivilege

`func (o *SmbActiveOpen) SetAccessPrivilege(v []string)`

SetAccessPrivilege sets AccessPrivilege field to given value.

### HasAccessPrivilege

`func (o *SmbActiveOpen) HasAccessPrivilege() bool`

HasAccessPrivilege returns a boolean if a field has been set.

### SetAccessPrivilegeNil

`func (o *SmbActiveOpen) SetAccessPrivilegeNil(b bool)`

 SetAccessPrivilegeNil sets the value for AccessPrivilege to be an explicit nil

### UnsetAccessPrivilege
`func (o *SmbActiveOpen) UnsetAccessPrivilege()`

UnsetAccessPrivilege ensures that no value is present for AccessPrivilege, not even an explicit nil
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


