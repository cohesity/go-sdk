# ArchivalQstarTapeExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **NullableString** | Specifies the host of the QStar Tape external target. | 
**WebServicesPort** | **NullableInt32** | Specifies the Web Services Port of the QStar Tape external target. | 
**Username** | **NullableString** | Specifies the Username of the QStar Tape external target. | 
**Password** | **NullableString** | Specifies the Password of the QStar Tape external target. | 
**ShareType** | Pointer to **NullableString** | Specifies the share type of QStar Tape external target. | [optional] 
**UseHttps** | Pointer to **NullableBool** | Specifies whether HTTPS is used or not. | [optional] 
**IntegralVolumeNames** | Pointer to **[]string** | Specifies the Integral Volume Names of the QStar Tape external target. | [optional] 
**IsIncrementalArchivalEnabled** | Pointer to **NullableBool** | Specifies if Incremental Archival setting is enabled or not. | [optional] 
**IsForeverIncrementalArchivalEnabled** | Pointer to **NullableBool** | Specifies if Forever Incremental Archival setting is enabled or not. | [optional] 

## Methods

### NewArchivalQstarTapeExternalTargetParams

`func NewArchivalQstarTapeExternalTargetParams(host NullableString, webServicesPort NullableInt32, username NullableString, password NullableString, ) *ArchivalQstarTapeExternalTargetParams`

NewArchivalQstarTapeExternalTargetParams instantiates a new ArchivalQstarTapeExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalQstarTapeExternalTargetParamsWithDefaults

`func NewArchivalQstarTapeExternalTargetParamsWithDefaults() *ArchivalQstarTapeExternalTargetParams`

NewArchivalQstarTapeExternalTargetParamsWithDefaults instantiates a new ArchivalQstarTapeExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ArchivalQstarTapeExternalTargetParams) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ArchivalQstarTapeExternalTargetParams) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ArchivalQstarTapeExternalTargetParams) SetHost(v string)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *ArchivalQstarTapeExternalTargetParams) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *ArchivalQstarTapeExternalTargetParams) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetWebServicesPort

`func (o *ArchivalQstarTapeExternalTargetParams) GetWebServicesPort() int32`

GetWebServicesPort returns the WebServicesPort field if non-nil, zero value otherwise.

### GetWebServicesPortOk

`func (o *ArchivalQstarTapeExternalTargetParams) GetWebServicesPortOk() (*int32, bool)`

GetWebServicesPortOk returns a tuple with the WebServicesPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebServicesPort

`func (o *ArchivalQstarTapeExternalTargetParams) SetWebServicesPort(v int32)`

SetWebServicesPort sets WebServicesPort field to given value.


### SetWebServicesPortNil

`func (o *ArchivalQstarTapeExternalTargetParams) SetWebServicesPortNil(b bool)`

 SetWebServicesPortNil sets the value for WebServicesPort to be an explicit nil

### UnsetWebServicesPort
`func (o *ArchivalQstarTapeExternalTargetParams) UnsetWebServicesPort()`

UnsetWebServicesPort ensures that no value is present for WebServicesPort, not even an explicit nil
### GetUsername

`func (o *ArchivalQstarTapeExternalTargetParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ArchivalQstarTapeExternalTargetParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ArchivalQstarTapeExternalTargetParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *ArchivalQstarTapeExternalTargetParams) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ArchivalQstarTapeExternalTargetParams) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ArchivalQstarTapeExternalTargetParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ArchivalQstarTapeExternalTargetParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ArchivalQstarTapeExternalTargetParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *ArchivalQstarTapeExternalTargetParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ArchivalQstarTapeExternalTargetParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetShareType

`func (o *ArchivalQstarTapeExternalTargetParams) GetShareType() string`

GetShareType returns the ShareType field if non-nil, zero value otherwise.

### GetShareTypeOk

`func (o *ArchivalQstarTapeExternalTargetParams) GetShareTypeOk() (*string, bool)`

GetShareTypeOk returns a tuple with the ShareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareType

`func (o *ArchivalQstarTapeExternalTargetParams) SetShareType(v string)`

SetShareType sets ShareType field to given value.

### HasShareType

`func (o *ArchivalQstarTapeExternalTargetParams) HasShareType() bool`

HasShareType returns a boolean if a field has been set.

### SetShareTypeNil

`func (o *ArchivalQstarTapeExternalTargetParams) SetShareTypeNil(b bool)`

 SetShareTypeNil sets the value for ShareType to be an explicit nil

### UnsetShareType
`func (o *ArchivalQstarTapeExternalTargetParams) UnsetShareType()`

UnsetShareType ensures that no value is present for ShareType, not even an explicit nil
### GetUseHttps

`func (o *ArchivalQstarTapeExternalTargetParams) GetUseHttps() bool`

GetUseHttps returns the UseHttps field if non-nil, zero value otherwise.

### GetUseHttpsOk

`func (o *ArchivalQstarTapeExternalTargetParams) GetUseHttpsOk() (*bool, bool)`

GetUseHttpsOk returns a tuple with the UseHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHttps

`func (o *ArchivalQstarTapeExternalTargetParams) SetUseHttps(v bool)`

SetUseHttps sets UseHttps field to given value.

### HasUseHttps

`func (o *ArchivalQstarTapeExternalTargetParams) HasUseHttps() bool`

HasUseHttps returns a boolean if a field has been set.

### SetUseHttpsNil

`func (o *ArchivalQstarTapeExternalTargetParams) SetUseHttpsNil(b bool)`

 SetUseHttpsNil sets the value for UseHttps to be an explicit nil

### UnsetUseHttps
`func (o *ArchivalQstarTapeExternalTargetParams) UnsetUseHttps()`

UnsetUseHttps ensures that no value is present for UseHttps, not even an explicit nil
### GetIntegralVolumeNames

`func (o *ArchivalQstarTapeExternalTargetParams) GetIntegralVolumeNames() []string`

GetIntegralVolumeNames returns the IntegralVolumeNames field if non-nil, zero value otherwise.

### GetIntegralVolumeNamesOk

`func (o *ArchivalQstarTapeExternalTargetParams) GetIntegralVolumeNamesOk() (*[]string, bool)`

GetIntegralVolumeNamesOk returns a tuple with the IntegralVolumeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegralVolumeNames

`func (o *ArchivalQstarTapeExternalTargetParams) SetIntegralVolumeNames(v []string)`

SetIntegralVolumeNames sets IntegralVolumeNames field to given value.

### HasIntegralVolumeNames

`func (o *ArchivalQstarTapeExternalTargetParams) HasIntegralVolumeNames() bool`

HasIntegralVolumeNames returns a boolean if a field has been set.

### SetIntegralVolumeNamesNil

`func (o *ArchivalQstarTapeExternalTargetParams) SetIntegralVolumeNamesNil(b bool)`

 SetIntegralVolumeNamesNil sets the value for IntegralVolumeNames to be an explicit nil

### UnsetIntegralVolumeNames
`func (o *ArchivalQstarTapeExternalTargetParams) UnsetIntegralVolumeNames()`

UnsetIntegralVolumeNames ensures that no value is present for IntegralVolumeNames, not even an explicit nil
### GetIsIncrementalArchivalEnabled

`func (o *ArchivalQstarTapeExternalTargetParams) GetIsIncrementalArchivalEnabled() bool`

GetIsIncrementalArchivalEnabled returns the IsIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsIncrementalArchivalEnabledOk

`func (o *ArchivalQstarTapeExternalTargetParams) GetIsIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsIncrementalArchivalEnabledOk returns a tuple with the IsIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncrementalArchivalEnabled

`func (o *ArchivalQstarTapeExternalTargetParams) SetIsIncrementalArchivalEnabled(v bool)`

SetIsIncrementalArchivalEnabled sets IsIncrementalArchivalEnabled field to given value.

### HasIsIncrementalArchivalEnabled

`func (o *ArchivalQstarTapeExternalTargetParams) HasIsIncrementalArchivalEnabled() bool`

HasIsIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsIncrementalArchivalEnabledNil

`func (o *ArchivalQstarTapeExternalTargetParams) SetIsIncrementalArchivalEnabledNil(b bool)`

 SetIsIncrementalArchivalEnabledNil sets the value for IsIncrementalArchivalEnabled to be an explicit nil

### UnsetIsIncrementalArchivalEnabled
`func (o *ArchivalQstarTapeExternalTargetParams) UnsetIsIncrementalArchivalEnabled()`

UnsetIsIncrementalArchivalEnabled ensures that no value is present for IsIncrementalArchivalEnabled, not even an explicit nil
### GetIsForeverIncrementalArchivalEnabled

`func (o *ArchivalQstarTapeExternalTargetParams) GetIsForeverIncrementalArchivalEnabled() bool`

GetIsForeverIncrementalArchivalEnabled returns the IsForeverIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsForeverIncrementalArchivalEnabledOk

`func (o *ArchivalQstarTapeExternalTargetParams) GetIsForeverIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsForeverIncrementalArchivalEnabledOk returns a tuple with the IsForeverIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForeverIncrementalArchivalEnabled

`func (o *ArchivalQstarTapeExternalTargetParams) SetIsForeverIncrementalArchivalEnabled(v bool)`

SetIsForeverIncrementalArchivalEnabled sets IsForeverIncrementalArchivalEnabled field to given value.

### HasIsForeverIncrementalArchivalEnabled

`func (o *ArchivalQstarTapeExternalTargetParams) HasIsForeverIncrementalArchivalEnabled() bool`

HasIsForeverIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsForeverIncrementalArchivalEnabledNil

`func (o *ArchivalQstarTapeExternalTargetParams) SetIsForeverIncrementalArchivalEnabledNil(b bool)`

 SetIsForeverIncrementalArchivalEnabledNil sets the value for IsForeverIncrementalArchivalEnabled to be an explicit nil

### UnsetIsForeverIncrementalArchivalEnabled
`func (o *ArchivalQstarTapeExternalTargetParams) UnsetIsForeverIncrementalArchivalEnabled()`

UnsetIsForeverIncrementalArchivalEnabled ensures that no value is present for IsForeverIncrementalArchivalEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


