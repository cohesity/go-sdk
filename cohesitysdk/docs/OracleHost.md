# OracleHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCount** | Pointer to **NullableInt32** | Specifies the count of CPU available on the host. | [optional] 
**IpAddresses** | Pointer to **[]string** | Specifies the IP address of the host. | [optional] 
**Ports** | Pointer to **[]int64** | Specifies ports available for this host. | [optional] 
**Sessions** | Pointer to [**[]OracleSession**](OracleSession.md) | Specifies multiple session configurations available for this host. | [optional] 

## Methods

### NewOracleHost

`func NewOracleHost() *OracleHost`

NewOracleHost instantiates a new OracleHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleHostWithDefaults

`func NewOracleHostWithDefaults() *OracleHost`

NewOracleHostWithDefaults instantiates a new OracleHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCount

`func (o *OracleHost) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *OracleHost) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *OracleHost) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *OracleHost) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### SetCpuCountNil

`func (o *OracleHost) SetCpuCountNil(b bool)`

 SetCpuCountNil sets the value for CpuCount to be an explicit nil

### UnsetCpuCount
`func (o *OracleHost) UnsetCpuCount()`

UnsetCpuCount ensures that no value is present for CpuCount, not even an explicit nil
### GetIpAddresses

`func (o *OracleHost) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *OracleHost) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *OracleHost) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *OracleHost) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *OracleHost) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *OracleHost) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil
### GetPorts

`func (o *OracleHost) GetPorts() []int64`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OracleHost) GetPortsOk() (*[]int64, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OracleHost) SetPorts(v []int64)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OracleHost) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *OracleHost) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *OracleHost) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetSessions

`func (o *OracleHost) GetSessions() []OracleSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *OracleHost) GetSessionsOk() (*[]OracleSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *OracleHost) SetSessions(v []OracleSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *OracleHost) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### SetSessionsNil

`func (o *OracleHost) SetSessionsNil(b bool)`

 SetSessionsNil sets the value for Sessions to be an explicit nil

### UnsetSessions
`func (o *OracleHost) UnsetSessions()`

UnsetSessions ensures that no value is present for Sessions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


