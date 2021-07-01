# SmbConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | Pointer to **NullableString** | Specifies the Client IP address of the connection. | [optional] 
**DomainName** | Pointer to **NullableString** | Domain name of the corresponding user. | [optional] 
**NodeIp** | Pointer to **NullableString** | Specifies a Node IP address where the connection request is received. | [optional] 
**Path** | Pointer to **NullableString** | Mount path. | [optional] 
**ServerIp** | Pointer to **NullableString** | Specifies the Server IP address of the connection. This could be a VIP, VLAN IP, or node IP on the Cluster. | [optional] 
**SessionId** | Pointer to **NullableInt64** | Session id. | [optional] 
**Sids** | Pointer to **[]string** | List of SIDs in the SMB session token. | [optional] 
**UserName** | Pointer to **NullableString** | User name used to login for this session. | [optional] 
**ViewId** | Pointer to **NullableInt64** | Specifies the id of the view. | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the name of the view. | [optional] 

## Methods

### NewSmbConnection

`func NewSmbConnection() *SmbConnection`

NewSmbConnection instantiates a new SmbConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbConnectionWithDefaults

`func NewSmbConnectionWithDefaults() *SmbConnection`

NewSmbConnectionWithDefaults instantiates a new SmbConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *SmbConnection) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *SmbConnection) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *SmbConnection) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *SmbConnection) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### SetClientIpNil

`func (o *SmbConnection) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *SmbConnection) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetDomainName

`func (o *SmbConnection) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *SmbConnection) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *SmbConnection) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *SmbConnection) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *SmbConnection) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *SmbConnection) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetNodeIp

`func (o *SmbConnection) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *SmbConnection) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *SmbConnection) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *SmbConnection) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### SetNodeIpNil

`func (o *SmbConnection) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *SmbConnection) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil
### GetPath

`func (o *SmbConnection) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SmbConnection) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SmbConnection) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SmbConnection) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *SmbConnection) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *SmbConnection) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetServerIp

`func (o *SmbConnection) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *SmbConnection) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *SmbConnection) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *SmbConnection) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### SetServerIpNil

`func (o *SmbConnection) SetServerIpNil(b bool)`

 SetServerIpNil sets the value for ServerIp to be an explicit nil

### UnsetServerIp
`func (o *SmbConnection) UnsetServerIp()`

UnsetServerIp ensures that no value is present for ServerIp, not even an explicit nil
### GetSessionId

`func (o *SmbConnection) GetSessionId() int64`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SmbConnection) GetSessionIdOk() (*int64, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SmbConnection) SetSessionId(v int64)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *SmbConnection) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *SmbConnection) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *SmbConnection) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetSids

`func (o *SmbConnection) GetSids() []string`

GetSids returns the Sids field if non-nil, zero value otherwise.

### GetSidsOk

`func (o *SmbConnection) GetSidsOk() (*[]string, bool)`

GetSidsOk returns a tuple with the Sids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSids

`func (o *SmbConnection) SetSids(v []string)`

SetSids sets Sids field to given value.

### HasSids

`func (o *SmbConnection) HasSids() bool`

HasSids returns a boolean if a field has been set.

### SetSidsNil

`func (o *SmbConnection) SetSidsNil(b bool)`

 SetSidsNil sets the value for Sids to be an explicit nil

### UnsetSids
`func (o *SmbConnection) UnsetSids()`

UnsetSids ensures that no value is present for Sids, not even an explicit nil
### GetUserName

`func (o *SmbConnection) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SmbConnection) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SmbConnection) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SmbConnection) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *SmbConnection) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *SmbConnection) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetViewId

`func (o *SmbConnection) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *SmbConnection) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *SmbConnection) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *SmbConnection) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *SmbConnection) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *SmbConnection) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
### GetViewName

`func (o *SmbConnection) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *SmbConnection) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *SmbConnection) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *SmbConnection) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *SmbConnection) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *SmbConnection) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


