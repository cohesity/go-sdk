# IdMappingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FallbackUserIdMappingInfo** | Pointer to [**UserIdMapping**](UserIdMapping.md) |  | [optional] 
**UnixRootSid** | Pointer to **NullableString** | Specifies the SID of the Active Directory domain user to be mapped to Unix root user. | [optional] 
**UserIdMappingInfo** | Pointer to [**UserIdMapping**](UserIdMapping.md) |  | [optional] 

## Methods

### NewIdMappingInfo

`func NewIdMappingInfo() *IdMappingInfo`

NewIdMappingInfo instantiates a new IdMappingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdMappingInfoWithDefaults

`func NewIdMappingInfoWithDefaults() *IdMappingInfo`

NewIdMappingInfoWithDefaults instantiates a new IdMappingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFallbackUserIdMappingInfo

`func (o *IdMappingInfo) GetFallbackUserIdMappingInfo() UserIdMapping`

GetFallbackUserIdMappingInfo returns the FallbackUserIdMappingInfo field if non-nil, zero value otherwise.

### GetFallbackUserIdMappingInfoOk

`func (o *IdMappingInfo) GetFallbackUserIdMappingInfoOk() (*UserIdMapping, bool)`

GetFallbackUserIdMappingInfoOk returns a tuple with the FallbackUserIdMappingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUserIdMappingInfo

`func (o *IdMappingInfo) SetFallbackUserIdMappingInfo(v UserIdMapping)`

SetFallbackUserIdMappingInfo sets FallbackUserIdMappingInfo field to given value.

### HasFallbackUserIdMappingInfo

`func (o *IdMappingInfo) HasFallbackUserIdMappingInfo() bool`

HasFallbackUserIdMappingInfo returns a boolean if a field has been set.

### GetUnixRootSid

`func (o *IdMappingInfo) GetUnixRootSid() string`

GetUnixRootSid returns the UnixRootSid field if non-nil, zero value otherwise.

### GetUnixRootSidOk

`func (o *IdMappingInfo) GetUnixRootSidOk() (*string, bool)`

GetUnixRootSidOk returns a tuple with the UnixRootSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixRootSid

`func (o *IdMappingInfo) SetUnixRootSid(v string)`

SetUnixRootSid sets UnixRootSid field to given value.

### HasUnixRootSid

`func (o *IdMappingInfo) HasUnixRootSid() bool`

HasUnixRootSid returns a boolean if a field has been set.

### SetUnixRootSidNil

`func (o *IdMappingInfo) SetUnixRootSidNil(b bool)`

 SetUnixRootSidNil sets the value for UnixRootSid to be an explicit nil

### UnsetUnixRootSid
`func (o *IdMappingInfo) UnsetUnixRootSid()`

UnsetUnixRootSid ensures that no value is present for UnixRootSid, not even an explicit nil
### GetUserIdMappingInfo

`func (o *IdMappingInfo) GetUserIdMappingInfo() UserIdMapping`

GetUserIdMappingInfo returns the UserIdMappingInfo field if non-nil, zero value otherwise.

### GetUserIdMappingInfoOk

`func (o *IdMappingInfo) GetUserIdMappingInfoOk() (*UserIdMapping, bool)`

GetUserIdMappingInfoOk returns a tuple with the UserIdMappingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdMappingInfo

`func (o *IdMappingInfo) SetUserIdMappingInfo(v UserIdMapping)`

SetUserIdMappingInfo sets UserIdMappingInfo field to given value.

### HasUserIdMappingInfo

`func (o *IdMappingInfo) HasUserIdMappingInfo() bool`

HasUserIdMappingInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


