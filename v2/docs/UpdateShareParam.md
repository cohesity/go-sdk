# UpdateShareParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSubnetWhitelist** | Pointer to [**[]Subnet**](Subnet.md) | List of external client subnet IPs that are allowed to access the share. | [optional] 
**EnableFilerAuditLogging** | Pointer to **NullableBool** | This field is currently deprecated. Specifies if Filer Audit Logging is enabled for this Share. | [optional] 
**FileAuditLoggingState** | Pointer to **NullableString** | Specifies the state of File Audit logging for this Share. Inherited: Audit log setting is inherited from the  View. Enabled: Audit log is enabled for this Share. Disabled: Audit log is disabled for this Share. | [optional] 
**SmbConfig** | Pointer to [**UpdateShareParamSmbConfig**](UpdateShareParamSmbConfig.md) |  | [optional] 

## Methods

### NewUpdateShareParam

`func NewUpdateShareParam() *UpdateShareParam`

NewUpdateShareParam instantiates a new UpdateShareParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateShareParamWithDefaults

`func NewUpdateShareParamWithDefaults() *UpdateShareParam`

NewUpdateShareParamWithDefaults instantiates a new UpdateShareParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSubnetWhitelist

`func (o *UpdateShareParam) GetClientSubnetWhitelist() []Subnet`

GetClientSubnetWhitelist returns the ClientSubnetWhitelist field if non-nil, zero value otherwise.

### GetClientSubnetWhitelistOk

`func (o *UpdateShareParam) GetClientSubnetWhitelistOk() (*[]Subnet, bool)`

GetClientSubnetWhitelistOk returns a tuple with the ClientSubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSubnetWhitelist

`func (o *UpdateShareParam) SetClientSubnetWhitelist(v []Subnet)`

SetClientSubnetWhitelist sets ClientSubnetWhitelist field to given value.

### HasClientSubnetWhitelist

`func (o *UpdateShareParam) HasClientSubnetWhitelist() bool`

HasClientSubnetWhitelist returns a boolean if a field has been set.

### SetClientSubnetWhitelistNil

`func (o *UpdateShareParam) SetClientSubnetWhitelistNil(b bool)`

 SetClientSubnetWhitelistNil sets the value for ClientSubnetWhitelist to be an explicit nil

### UnsetClientSubnetWhitelist
`func (o *UpdateShareParam) UnsetClientSubnetWhitelist()`

UnsetClientSubnetWhitelist ensures that no value is present for ClientSubnetWhitelist, not even an explicit nil
### GetEnableFilerAuditLogging

`func (o *UpdateShareParam) GetEnableFilerAuditLogging() bool`

GetEnableFilerAuditLogging returns the EnableFilerAuditLogging field if non-nil, zero value otherwise.

### GetEnableFilerAuditLoggingOk

`func (o *UpdateShareParam) GetEnableFilerAuditLoggingOk() (*bool, bool)`

GetEnableFilerAuditLoggingOk returns a tuple with the EnableFilerAuditLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFilerAuditLogging

`func (o *UpdateShareParam) SetEnableFilerAuditLogging(v bool)`

SetEnableFilerAuditLogging sets EnableFilerAuditLogging field to given value.

### HasEnableFilerAuditLogging

`func (o *UpdateShareParam) HasEnableFilerAuditLogging() bool`

HasEnableFilerAuditLogging returns a boolean if a field has been set.

### SetEnableFilerAuditLoggingNil

`func (o *UpdateShareParam) SetEnableFilerAuditLoggingNil(b bool)`

 SetEnableFilerAuditLoggingNil sets the value for EnableFilerAuditLogging to be an explicit nil

### UnsetEnableFilerAuditLogging
`func (o *UpdateShareParam) UnsetEnableFilerAuditLogging()`

UnsetEnableFilerAuditLogging ensures that no value is present for EnableFilerAuditLogging, not even an explicit nil
### GetFileAuditLoggingState

`func (o *UpdateShareParam) GetFileAuditLoggingState() string`

GetFileAuditLoggingState returns the FileAuditLoggingState field if non-nil, zero value otherwise.

### GetFileAuditLoggingStateOk

`func (o *UpdateShareParam) GetFileAuditLoggingStateOk() (*string, bool)`

GetFileAuditLoggingStateOk returns a tuple with the FileAuditLoggingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAuditLoggingState

`func (o *UpdateShareParam) SetFileAuditLoggingState(v string)`

SetFileAuditLoggingState sets FileAuditLoggingState field to given value.

### HasFileAuditLoggingState

`func (o *UpdateShareParam) HasFileAuditLoggingState() bool`

HasFileAuditLoggingState returns a boolean if a field has been set.

### SetFileAuditLoggingStateNil

`func (o *UpdateShareParam) SetFileAuditLoggingStateNil(b bool)`

 SetFileAuditLoggingStateNil sets the value for FileAuditLoggingState to be an explicit nil

### UnsetFileAuditLoggingState
`func (o *UpdateShareParam) UnsetFileAuditLoggingState()`

UnsetFileAuditLoggingState ensures that no value is present for FileAuditLoggingState, not even an explicit nil
### GetSmbConfig

`func (o *UpdateShareParam) GetSmbConfig() UpdateShareParamSmbConfig`

GetSmbConfig returns the SmbConfig field if non-nil, zero value otherwise.

### GetSmbConfigOk

`func (o *UpdateShareParam) GetSmbConfigOk() (*UpdateShareParamSmbConfig, bool)`

GetSmbConfigOk returns a tuple with the SmbConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbConfig

`func (o *UpdateShareParam) SetSmbConfig(v UpdateShareParamSmbConfig)`

SetSmbConfig sets SmbConfig field to given value.

### HasSmbConfig

`func (o *UpdateShareParam) HasSmbConfig() bool`

HasSmbConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


