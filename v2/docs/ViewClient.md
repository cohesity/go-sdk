# ViewClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedTimeUsecs** | Pointer to **NullableInt64** | Specifies the time how long the client has connected to the server. | [optional] 
**Gid** | Pointer to **NullableInt32** | Specifies the GID of the client user. | [optional] 
**Ip** | Pointer to **NullableString** | Specifies the client ip. | [optional] 
**NodeIp** | Pointer to **NullableString** | Specifies the node ip which the client is connected to. | [optional] 
**Protocol** | Pointer to **NullableString** | Specifies the protocol the client uses. | [optional] 
**ServerIp** | Pointer to **NullableString** | Specifies the server ip which the client is connected to. | [optional] 
**SmbDialectVersion** | Pointer to **NullableInt32** | Specifies the dialect version for SMB client. | [optional] 
**Uid** | Pointer to **NullableInt32** | Specifies the UID of the client user. | [optional] 
**UserDomain** | Pointer to **NullableString** | Specifies the user domain of the client. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the username of the client. | [optional] 
**ViewId** | Pointer to **NullableInt64** | Specifies the id of the View which the client is connected to. | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the name of the View which the client is connected to. | [optional] 
**ViewPath** | Pointer to **NullableString** | Specifies the path of the View which the client is connected to. | [optional] 

## Methods

### NewViewClient

`func NewViewClient() *ViewClient`

NewViewClient instantiates a new ViewClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewClientWithDefaults

`func NewViewClientWithDefaults() *ViewClient`

NewViewClientWithDefaults instantiates a new ViewClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedTimeUsecs

`func (o *ViewClient) GetConnectedTimeUsecs() int64`

GetConnectedTimeUsecs returns the ConnectedTimeUsecs field if non-nil, zero value otherwise.

### GetConnectedTimeUsecsOk

`func (o *ViewClient) GetConnectedTimeUsecsOk() (*int64, bool)`

GetConnectedTimeUsecsOk returns a tuple with the ConnectedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedTimeUsecs

`func (o *ViewClient) SetConnectedTimeUsecs(v int64)`

SetConnectedTimeUsecs sets ConnectedTimeUsecs field to given value.

### HasConnectedTimeUsecs

`func (o *ViewClient) HasConnectedTimeUsecs() bool`

HasConnectedTimeUsecs returns a boolean if a field has been set.

### SetConnectedTimeUsecsNil

`func (o *ViewClient) SetConnectedTimeUsecsNil(b bool)`

 SetConnectedTimeUsecsNil sets the value for ConnectedTimeUsecs to be an explicit nil

### UnsetConnectedTimeUsecs
`func (o *ViewClient) UnsetConnectedTimeUsecs()`

UnsetConnectedTimeUsecs ensures that no value is present for ConnectedTimeUsecs, not even an explicit nil
### GetGid

`func (o *ViewClient) GetGid() int32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *ViewClient) GetGidOk() (*int32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *ViewClient) SetGid(v int32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *ViewClient) HasGid() bool`

HasGid returns a boolean if a field has been set.

### SetGidNil

`func (o *ViewClient) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *ViewClient) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil
### GetIp

`func (o *ViewClient) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ViewClient) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ViewClient) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ViewClient) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *ViewClient) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *ViewClient) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetNodeIp

`func (o *ViewClient) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *ViewClient) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *ViewClient) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *ViewClient) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### SetNodeIpNil

`func (o *ViewClient) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *ViewClient) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil
### GetProtocol

`func (o *ViewClient) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ViewClient) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ViewClient) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ViewClient) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *ViewClient) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *ViewClient) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetServerIp

`func (o *ViewClient) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *ViewClient) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *ViewClient) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *ViewClient) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### SetServerIpNil

`func (o *ViewClient) SetServerIpNil(b bool)`

 SetServerIpNil sets the value for ServerIp to be an explicit nil

### UnsetServerIp
`func (o *ViewClient) UnsetServerIp()`

UnsetServerIp ensures that no value is present for ServerIp, not even an explicit nil
### GetSmbDialectVersion

`func (o *ViewClient) GetSmbDialectVersion() int32`

GetSmbDialectVersion returns the SmbDialectVersion field if non-nil, zero value otherwise.

### GetSmbDialectVersionOk

`func (o *ViewClient) GetSmbDialectVersionOk() (*int32, bool)`

GetSmbDialectVersionOk returns a tuple with the SmbDialectVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbDialectVersion

`func (o *ViewClient) SetSmbDialectVersion(v int32)`

SetSmbDialectVersion sets SmbDialectVersion field to given value.

### HasSmbDialectVersion

`func (o *ViewClient) HasSmbDialectVersion() bool`

HasSmbDialectVersion returns a boolean if a field has been set.

### SetSmbDialectVersionNil

`func (o *ViewClient) SetSmbDialectVersionNil(b bool)`

 SetSmbDialectVersionNil sets the value for SmbDialectVersion to be an explicit nil

### UnsetSmbDialectVersion
`func (o *ViewClient) UnsetSmbDialectVersion()`

UnsetSmbDialectVersion ensures that no value is present for SmbDialectVersion, not even an explicit nil
### GetUid

`func (o *ViewClient) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ViewClient) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ViewClient) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ViewClient) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *ViewClient) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *ViewClient) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetUserDomain

`func (o *ViewClient) GetUserDomain() string`

GetUserDomain returns the UserDomain field if non-nil, zero value otherwise.

### GetUserDomainOk

`func (o *ViewClient) GetUserDomainOk() (*string, bool)`

GetUserDomainOk returns a tuple with the UserDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDomain

`func (o *ViewClient) SetUserDomain(v string)`

SetUserDomain sets UserDomain field to given value.

### HasUserDomain

`func (o *ViewClient) HasUserDomain() bool`

HasUserDomain returns a boolean if a field has been set.

### SetUserDomainNil

`func (o *ViewClient) SetUserDomainNil(b bool)`

 SetUserDomainNil sets the value for UserDomain to be an explicit nil

### UnsetUserDomain
`func (o *ViewClient) UnsetUserDomain()`

UnsetUserDomain ensures that no value is present for UserDomain, not even an explicit nil
### GetUsername

`func (o *ViewClient) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ViewClient) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ViewClient) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ViewClient) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ViewClient) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ViewClient) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetViewId

`func (o *ViewClient) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *ViewClient) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *ViewClient) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *ViewClient) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *ViewClient) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *ViewClient) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
### GetViewName

`func (o *ViewClient) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ViewClient) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ViewClient) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *ViewClient) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *ViewClient) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *ViewClient) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetViewPath

`func (o *ViewClient) GetViewPath() string`

GetViewPath returns the ViewPath field if non-nil, zero value otherwise.

### GetViewPathOk

`func (o *ViewClient) GetViewPathOk() (*string, bool)`

GetViewPathOk returns a tuple with the ViewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPath

`func (o *ViewClient) SetViewPath(v string)`

SetViewPath sets ViewPath field to given value.

### HasViewPath

`func (o *ViewClient) HasViewPath() bool`

HasViewPath returns a boolean if a field has been set.

### SetViewPathNil

`func (o *ViewClient) SetViewPathNil(b bool)`

 SetViewPathNil sets the value for ViewPath to be an explicit nil

### UnsetViewPath
`func (o *ViewClient) UnsetViewPath()`

UnsetViewPath ensures that no value is present for ViewPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


