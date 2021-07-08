# QStarServerCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | Specifies the IP address or DNS name of the server where QStar service is running. | [optional] 
**IntegralVolumeNames** | Pointer to **[]string** | Array of Integral Volume Names.  Specifies a list of existing Integral Volume names available on the QStar server for storing objects. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password used to access the QStar host. | [optional] 
**Port** | Pointer to **NullableInt32** | Specifies the listening port where QStar WEB API service is running. | [optional] 
**ShareType** | Pointer to **NullableString** | Specifies the sharing protocol type used by QStar to mount the integral volume. See the Cohesity online help for the recommended protocol for your environment. | [optional] 
**UseHttps** | Pointer to **NullableBool** | Specifies whether to use http or https to connect to the service. If true, a secure connection (https) is used. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the account name used to access the QStar host. | [optional] 

## Methods

### NewQStarServerCredentials

`func NewQStarServerCredentials() *QStarServerCredentials`

NewQStarServerCredentials instantiates a new QStarServerCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQStarServerCredentialsWithDefaults

`func NewQStarServerCredentialsWithDefaults() *QStarServerCredentials`

NewQStarServerCredentialsWithDefaults instantiates a new QStarServerCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *QStarServerCredentials) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *QStarServerCredentials) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *QStarServerCredentials) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *QStarServerCredentials) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *QStarServerCredentials) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *QStarServerCredentials) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetIntegralVolumeNames

`func (o *QStarServerCredentials) GetIntegralVolumeNames() []string`

GetIntegralVolumeNames returns the IntegralVolumeNames field if non-nil, zero value otherwise.

### GetIntegralVolumeNamesOk

`func (o *QStarServerCredentials) GetIntegralVolumeNamesOk() (*[]string, bool)`

GetIntegralVolumeNamesOk returns a tuple with the IntegralVolumeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegralVolumeNames

`func (o *QStarServerCredentials) SetIntegralVolumeNames(v []string)`

SetIntegralVolumeNames sets IntegralVolumeNames field to given value.

### HasIntegralVolumeNames

`func (o *QStarServerCredentials) HasIntegralVolumeNames() bool`

HasIntegralVolumeNames returns a boolean if a field has been set.

### SetIntegralVolumeNamesNil

`func (o *QStarServerCredentials) SetIntegralVolumeNamesNil(b bool)`

 SetIntegralVolumeNamesNil sets the value for IntegralVolumeNames to be an explicit nil

### UnsetIntegralVolumeNames
`func (o *QStarServerCredentials) UnsetIntegralVolumeNames()`

UnsetIntegralVolumeNames ensures that no value is present for IntegralVolumeNames, not even an explicit nil
### GetPassword

`func (o *QStarServerCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *QStarServerCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *QStarServerCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *QStarServerCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *QStarServerCredentials) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *QStarServerCredentials) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPort

`func (o *QStarServerCredentials) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *QStarServerCredentials) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *QStarServerCredentials) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *QStarServerCredentials) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *QStarServerCredentials) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *QStarServerCredentials) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetShareType

`func (o *QStarServerCredentials) GetShareType() string`

GetShareType returns the ShareType field if non-nil, zero value otherwise.

### GetShareTypeOk

`func (o *QStarServerCredentials) GetShareTypeOk() (*string, bool)`

GetShareTypeOk returns a tuple with the ShareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareType

`func (o *QStarServerCredentials) SetShareType(v string)`

SetShareType sets ShareType field to given value.

### HasShareType

`func (o *QStarServerCredentials) HasShareType() bool`

HasShareType returns a boolean if a field has been set.

### SetShareTypeNil

`func (o *QStarServerCredentials) SetShareTypeNil(b bool)`

 SetShareTypeNil sets the value for ShareType to be an explicit nil

### UnsetShareType
`func (o *QStarServerCredentials) UnsetShareType()`

UnsetShareType ensures that no value is present for ShareType, not even an explicit nil
### GetUseHttps

`func (o *QStarServerCredentials) GetUseHttps() bool`

GetUseHttps returns the UseHttps field if non-nil, zero value otherwise.

### GetUseHttpsOk

`func (o *QStarServerCredentials) GetUseHttpsOk() (*bool, bool)`

GetUseHttpsOk returns a tuple with the UseHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHttps

`func (o *QStarServerCredentials) SetUseHttps(v bool)`

SetUseHttps sets UseHttps field to given value.

### HasUseHttps

`func (o *QStarServerCredentials) HasUseHttps() bool`

HasUseHttps returns a boolean if a field has been set.

### SetUseHttpsNil

`func (o *QStarServerCredentials) SetUseHttpsNil(b bool)`

 SetUseHttpsNil sets the value for UseHttps to be an explicit nil

### UnsetUseHttps
`func (o *QStarServerCredentials) UnsetUseHttps()`

UnsetUseHttps ensures that no value is present for UseHttps, not even an explicit nil
### GetUsername

`func (o *QStarServerCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *QStarServerCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *QStarServerCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *QStarServerCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *QStarServerCredentials) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *QStarServerCredentials) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


