# SyslogServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCertificate** | Pointer to **NullableString** | Syslog server CA certificate. | [optional] 
**Enabled** | Pointer to **NullableBool** | Specifies whether to enable the syslog server on the Cluster. | [optional] 
**FacilityList** | Pointer to **[]string** | Send enabled syslog facilities related logs to logging server. | [optional] 
**Id** | Pointer to **NullableInt32** | The id of the syslog server. | [optional] 
**Ip** | Pointer to **NullableString** | Specifies the IP address or hostname of the syslog server. | [optional] 
**IsTlsEnabled** | Pointer to **NullableBool** | Specify whether to enable tls support. | [optional] 
**MsgPatternList** | Pointer to **[]string** | Send logs including the msg patterns to logging server. | [optional] 
**Name** | Pointer to **NullableString** | Specifies a unique name for the syslog server on the Cluster. | [optional] 
**Port** | Pointer to **NullableInt32** | Specifies the port where the syslog server listens. | [optional] 
**ProgramNameList** | Pointer to **[]string** | Send programes related logs to logging server. | [optional] 
**Protocol** | Pointer to **NullableString** | Specifies the protocol used to send the logs. | [optional] 
**RawMsgPatternList** | Pointer to **[]string** | Send logs including the msg patterns to logging server. | [optional] 
**TokenId** | Pointer to **NullableString** | TokenId used for filtering messages on a relay or collector | [optional] 

## Methods

### NewSyslogServer

`func NewSyslogServer() *SyslogServer`

NewSyslogServer instantiates a new SyslogServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogServerWithDefaults

`func NewSyslogServerWithDefaults() *SyslogServer`

NewSyslogServerWithDefaults instantiates a new SyslogServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCertificate

`func (o *SyslogServer) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *SyslogServer) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *SyslogServer) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *SyslogServer) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### SetCaCertificateNil

`func (o *SyslogServer) SetCaCertificateNil(b bool)`

 SetCaCertificateNil sets the value for CaCertificate to be an explicit nil

### UnsetCaCertificate
`func (o *SyslogServer) UnsetCaCertificate()`

UnsetCaCertificate ensures that no value is present for CaCertificate, not even an explicit nil
### GetEnabled

`func (o *SyslogServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SyslogServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SyslogServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SyslogServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *SyslogServer) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *SyslogServer) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetFacilityList

`func (o *SyslogServer) GetFacilityList() []string`

GetFacilityList returns the FacilityList field if non-nil, zero value otherwise.

### GetFacilityListOk

`func (o *SyslogServer) GetFacilityListOk() (*[]string, bool)`

GetFacilityListOk returns a tuple with the FacilityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityList

`func (o *SyslogServer) SetFacilityList(v []string)`

SetFacilityList sets FacilityList field to given value.

### HasFacilityList

`func (o *SyslogServer) HasFacilityList() bool`

HasFacilityList returns a boolean if a field has been set.

### GetId

`func (o *SyslogServer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyslogServer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyslogServer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SyslogServer) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SyslogServer) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SyslogServer) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIp

`func (o *SyslogServer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SyslogServer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SyslogServer) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SyslogServer) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *SyslogServer) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *SyslogServer) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetIsTlsEnabled

`func (o *SyslogServer) GetIsTlsEnabled() bool`

GetIsTlsEnabled returns the IsTlsEnabled field if non-nil, zero value otherwise.

### GetIsTlsEnabledOk

`func (o *SyslogServer) GetIsTlsEnabledOk() (*bool, bool)`

GetIsTlsEnabledOk returns a tuple with the IsTlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTlsEnabled

`func (o *SyslogServer) SetIsTlsEnabled(v bool)`

SetIsTlsEnabled sets IsTlsEnabled field to given value.

### HasIsTlsEnabled

`func (o *SyslogServer) HasIsTlsEnabled() bool`

HasIsTlsEnabled returns a boolean if a field has been set.

### SetIsTlsEnabledNil

`func (o *SyslogServer) SetIsTlsEnabledNil(b bool)`

 SetIsTlsEnabledNil sets the value for IsTlsEnabled to be an explicit nil

### UnsetIsTlsEnabled
`func (o *SyslogServer) UnsetIsTlsEnabled()`

UnsetIsTlsEnabled ensures that no value is present for IsTlsEnabled, not even an explicit nil
### GetMsgPatternList

`func (o *SyslogServer) GetMsgPatternList() []string`

GetMsgPatternList returns the MsgPatternList field if non-nil, zero value otherwise.

### GetMsgPatternListOk

`func (o *SyslogServer) GetMsgPatternListOk() (*[]string, bool)`

GetMsgPatternListOk returns a tuple with the MsgPatternList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgPatternList

`func (o *SyslogServer) SetMsgPatternList(v []string)`

SetMsgPatternList sets MsgPatternList field to given value.

### HasMsgPatternList

`func (o *SyslogServer) HasMsgPatternList() bool`

HasMsgPatternList returns a boolean if a field has been set.

### GetName

`func (o *SyslogServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyslogServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyslogServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyslogServer) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SyslogServer) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SyslogServer) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPort

`func (o *SyslogServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SyslogServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SyslogServer) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SyslogServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *SyslogServer) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *SyslogServer) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetProgramNameList

`func (o *SyslogServer) GetProgramNameList() []string`

GetProgramNameList returns the ProgramNameList field if non-nil, zero value otherwise.

### GetProgramNameListOk

`func (o *SyslogServer) GetProgramNameListOk() (*[]string, bool)`

GetProgramNameListOk returns a tuple with the ProgramNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramNameList

`func (o *SyslogServer) SetProgramNameList(v []string)`

SetProgramNameList sets ProgramNameList field to given value.

### HasProgramNameList

`func (o *SyslogServer) HasProgramNameList() bool`

HasProgramNameList returns a boolean if a field has been set.

### GetProtocol

`func (o *SyslogServer) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SyslogServer) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SyslogServer) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SyslogServer) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *SyslogServer) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *SyslogServer) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetRawMsgPatternList

`func (o *SyslogServer) GetRawMsgPatternList() []string`

GetRawMsgPatternList returns the RawMsgPatternList field if non-nil, zero value otherwise.

### GetRawMsgPatternListOk

`func (o *SyslogServer) GetRawMsgPatternListOk() (*[]string, bool)`

GetRawMsgPatternListOk returns a tuple with the RawMsgPatternList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawMsgPatternList

`func (o *SyslogServer) SetRawMsgPatternList(v []string)`

SetRawMsgPatternList sets RawMsgPatternList field to given value.

### HasRawMsgPatternList

`func (o *SyslogServer) HasRawMsgPatternList() bool`

HasRawMsgPatternList returns a boolean if a field has been set.

### GetTokenId

`func (o *SyslogServer) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *SyslogServer) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *SyslogServer) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *SyslogServer) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### SetTokenIdNil

`func (o *SyslogServer) SetTokenIdNil(b bool)`

 SetTokenIdNil sets the value for TokenId to be an explicit nil

### UnsetTokenId
`func (o *SyslogServer) UnsetTokenId()`

UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


