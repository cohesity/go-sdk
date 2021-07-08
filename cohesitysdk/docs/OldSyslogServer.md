# OldSyslogServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **NullableString** | Specifies the IP address or hostname of the syslog server. | 
**IsAlertAuditingEnabled** | Pointer to **NullableBool** | Specifies if cohesity alert should be sent to syslog server If &#39;true&#39;, alert audting message are sent to the server. If &#39;false&#39;, alert auditng message are not sent to the server.(default) | [optional] 
**IsClusterAuditingEnabled** | Pointer to **NullableBool** | Specifies if Cluster audit logs should be sent to this syslog server. If &#39;true&#39;, Cluster audit logs are sent to the syslog server. (default) If &#39;false&#39;, Cluster audit logs are not sent to the syslog server. Either cluster audit logs or filer audit logs should be enabled. | [optional] 
**IsDataProtectionEnabled** | Pointer to **NullableBool** | Specifies if dataprotection  logs should be sent to syslog server If &#39;true&#39;, dataprotection  logs are sent to the server. If &#39;false&#39;, dataprotection  logs are not sent to the server.(default) | [optional] 
**IsFilerAuditingEnabled** | Pointer to **NullableBool** | Specifies if filer audit logs should be sent to this syslog server. If &#39;true&#39;, filer audit logs are sent to the syslog server. (default) If &#39;false&#39;, filer audit logs are not sent to the syslog server. Either cluster audit logs or filer audit logs should be enabled. | [optional] 
**IsSshLogEnabled** | Pointer to **NullableBool** | Specifies if ssh login logs should be sent to syslog server If &#39;true&#39;, ssh login logs are sent to the server. If &#39;false&#39;, ssh login logs are not sent to the server.(default) | [optional] 
**Name** | Pointer to **NullableString** | Specifies a unique name for the syslog server on the Cluster. | [optional] 
**Port** | **NullableInt32** | Specifies the port where the syslog server listens. | 
**Protocol** | **NullableInt32** | Specifies the protocol used to send the logs. | 

## Methods

### NewOldSyslogServer

`func NewOldSyslogServer(address NullableString, port NullableInt32, protocol NullableInt32, ) *OldSyslogServer`

NewOldSyslogServer instantiates a new OldSyslogServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOldSyslogServerWithDefaults

`func NewOldSyslogServerWithDefaults() *OldSyslogServer`

NewOldSyslogServerWithDefaults instantiates a new OldSyslogServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *OldSyslogServer) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OldSyslogServer) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OldSyslogServer) SetAddress(v string)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *OldSyslogServer) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *OldSyslogServer) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetIsAlertAuditingEnabled

`func (o *OldSyslogServer) GetIsAlertAuditingEnabled() bool`

GetIsAlertAuditingEnabled returns the IsAlertAuditingEnabled field if non-nil, zero value otherwise.

### GetIsAlertAuditingEnabledOk

`func (o *OldSyslogServer) GetIsAlertAuditingEnabledOk() (*bool, bool)`

GetIsAlertAuditingEnabledOk returns a tuple with the IsAlertAuditingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAlertAuditingEnabled

`func (o *OldSyslogServer) SetIsAlertAuditingEnabled(v bool)`

SetIsAlertAuditingEnabled sets IsAlertAuditingEnabled field to given value.

### HasIsAlertAuditingEnabled

`func (o *OldSyslogServer) HasIsAlertAuditingEnabled() bool`

HasIsAlertAuditingEnabled returns a boolean if a field has been set.

### SetIsAlertAuditingEnabledNil

`func (o *OldSyslogServer) SetIsAlertAuditingEnabledNil(b bool)`

 SetIsAlertAuditingEnabledNil sets the value for IsAlertAuditingEnabled to be an explicit nil

### UnsetIsAlertAuditingEnabled
`func (o *OldSyslogServer) UnsetIsAlertAuditingEnabled()`

UnsetIsAlertAuditingEnabled ensures that no value is present for IsAlertAuditingEnabled, not even an explicit nil
### GetIsClusterAuditingEnabled

`func (o *OldSyslogServer) GetIsClusterAuditingEnabled() bool`

GetIsClusterAuditingEnabled returns the IsClusterAuditingEnabled field if non-nil, zero value otherwise.

### GetIsClusterAuditingEnabledOk

`func (o *OldSyslogServer) GetIsClusterAuditingEnabledOk() (*bool, bool)`

GetIsClusterAuditingEnabledOk returns a tuple with the IsClusterAuditingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClusterAuditingEnabled

`func (o *OldSyslogServer) SetIsClusterAuditingEnabled(v bool)`

SetIsClusterAuditingEnabled sets IsClusterAuditingEnabled field to given value.

### HasIsClusterAuditingEnabled

`func (o *OldSyslogServer) HasIsClusterAuditingEnabled() bool`

HasIsClusterAuditingEnabled returns a boolean if a field has been set.

### SetIsClusterAuditingEnabledNil

`func (o *OldSyslogServer) SetIsClusterAuditingEnabledNil(b bool)`

 SetIsClusterAuditingEnabledNil sets the value for IsClusterAuditingEnabled to be an explicit nil

### UnsetIsClusterAuditingEnabled
`func (o *OldSyslogServer) UnsetIsClusterAuditingEnabled()`

UnsetIsClusterAuditingEnabled ensures that no value is present for IsClusterAuditingEnabled, not even an explicit nil
### GetIsDataProtectionEnabled

`func (o *OldSyslogServer) GetIsDataProtectionEnabled() bool`

GetIsDataProtectionEnabled returns the IsDataProtectionEnabled field if non-nil, zero value otherwise.

### GetIsDataProtectionEnabledOk

`func (o *OldSyslogServer) GetIsDataProtectionEnabledOk() (*bool, bool)`

GetIsDataProtectionEnabledOk returns a tuple with the IsDataProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDataProtectionEnabled

`func (o *OldSyslogServer) SetIsDataProtectionEnabled(v bool)`

SetIsDataProtectionEnabled sets IsDataProtectionEnabled field to given value.

### HasIsDataProtectionEnabled

`func (o *OldSyslogServer) HasIsDataProtectionEnabled() bool`

HasIsDataProtectionEnabled returns a boolean if a field has been set.

### SetIsDataProtectionEnabledNil

`func (o *OldSyslogServer) SetIsDataProtectionEnabledNil(b bool)`

 SetIsDataProtectionEnabledNil sets the value for IsDataProtectionEnabled to be an explicit nil

### UnsetIsDataProtectionEnabled
`func (o *OldSyslogServer) UnsetIsDataProtectionEnabled()`

UnsetIsDataProtectionEnabled ensures that no value is present for IsDataProtectionEnabled, not even an explicit nil
### GetIsFilerAuditingEnabled

`func (o *OldSyslogServer) GetIsFilerAuditingEnabled() bool`

GetIsFilerAuditingEnabled returns the IsFilerAuditingEnabled field if non-nil, zero value otherwise.

### GetIsFilerAuditingEnabledOk

`func (o *OldSyslogServer) GetIsFilerAuditingEnabledOk() (*bool, bool)`

GetIsFilerAuditingEnabledOk returns a tuple with the IsFilerAuditingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFilerAuditingEnabled

`func (o *OldSyslogServer) SetIsFilerAuditingEnabled(v bool)`

SetIsFilerAuditingEnabled sets IsFilerAuditingEnabled field to given value.

### HasIsFilerAuditingEnabled

`func (o *OldSyslogServer) HasIsFilerAuditingEnabled() bool`

HasIsFilerAuditingEnabled returns a boolean if a field has been set.

### SetIsFilerAuditingEnabledNil

`func (o *OldSyslogServer) SetIsFilerAuditingEnabledNil(b bool)`

 SetIsFilerAuditingEnabledNil sets the value for IsFilerAuditingEnabled to be an explicit nil

### UnsetIsFilerAuditingEnabled
`func (o *OldSyslogServer) UnsetIsFilerAuditingEnabled()`

UnsetIsFilerAuditingEnabled ensures that no value is present for IsFilerAuditingEnabled, not even an explicit nil
### GetIsSshLogEnabled

`func (o *OldSyslogServer) GetIsSshLogEnabled() bool`

GetIsSshLogEnabled returns the IsSshLogEnabled field if non-nil, zero value otherwise.

### GetIsSshLogEnabledOk

`func (o *OldSyslogServer) GetIsSshLogEnabledOk() (*bool, bool)`

GetIsSshLogEnabledOk returns a tuple with the IsSshLogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSshLogEnabled

`func (o *OldSyslogServer) SetIsSshLogEnabled(v bool)`

SetIsSshLogEnabled sets IsSshLogEnabled field to given value.

### HasIsSshLogEnabled

`func (o *OldSyslogServer) HasIsSshLogEnabled() bool`

HasIsSshLogEnabled returns a boolean if a field has been set.

### SetIsSshLogEnabledNil

`func (o *OldSyslogServer) SetIsSshLogEnabledNil(b bool)`

 SetIsSshLogEnabledNil sets the value for IsSshLogEnabled to be an explicit nil

### UnsetIsSshLogEnabled
`func (o *OldSyslogServer) UnsetIsSshLogEnabled()`

UnsetIsSshLogEnabled ensures that no value is present for IsSshLogEnabled, not even an explicit nil
### GetName

`func (o *OldSyslogServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OldSyslogServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OldSyslogServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OldSyslogServer) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OldSyslogServer) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OldSyslogServer) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPort

`func (o *OldSyslogServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OldSyslogServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OldSyslogServer) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *OldSyslogServer) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *OldSyslogServer) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetProtocol

`func (o *OldSyslogServer) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *OldSyslogServer) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *OldSyslogServer) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.


### SetProtocolNil

`func (o *OldSyslogServer) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *OldSyslogServer) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


