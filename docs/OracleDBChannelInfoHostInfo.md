# OracleDBChannelInfoHostInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | &#39;agent_id&#39; of the host from which we are allowed to take the backup/restore. | [optional] 
**NumChannels** | Pointer to **NullableInt32** | Number of channels we need to create for this host. Default value for num_channels will be calculated as minimum of number of nodes in cohesity cluster, 2 * number of cpu on Oracle host. | [optional] 
**Portnum** | Pointer to **NullableInt64** | port number where database is listening. | [optional] 
**SbtHostParams** | Pointer to [**OracleSbtHostParams**](OracleSbtHostParams.md) |  | [optional] 

## Methods

### NewOracleDBChannelInfoHostInfo

`func NewOracleDBChannelInfoHostInfo() *OracleDBChannelInfoHostInfo`

NewOracleDBChannelInfoHostInfo instantiates a new OracleDBChannelInfoHostInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleDBChannelInfoHostInfoWithDefaults

`func NewOracleDBChannelInfoHostInfoWithDefaults() *OracleDBChannelInfoHostInfo`

NewOracleDBChannelInfoHostInfoWithDefaults instantiates a new OracleDBChannelInfoHostInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *OracleDBChannelInfoHostInfo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *OracleDBChannelInfoHostInfo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *OracleDBChannelInfoHostInfo) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *OracleDBChannelInfoHostInfo) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *OracleDBChannelInfoHostInfo) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *OracleDBChannelInfoHostInfo) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetNumChannels

`func (o *OracleDBChannelInfoHostInfo) GetNumChannels() int32`

GetNumChannels returns the NumChannels field if non-nil, zero value otherwise.

### GetNumChannelsOk

`func (o *OracleDBChannelInfoHostInfo) GetNumChannelsOk() (*int32, bool)`

GetNumChannelsOk returns a tuple with the NumChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumChannels

`func (o *OracleDBChannelInfoHostInfo) SetNumChannels(v int32)`

SetNumChannels sets NumChannels field to given value.

### HasNumChannels

`func (o *OracleDBChannelInfoHostInfo) HasNumChannels() bool`

HasNumChannels returns a boolean if a field has been set.

### SetNumChannelsNil

`func (o *OracleDBChannelInfoHostInfo) SetNumChannelsNil(b bool)`

 SetNumChannelsNil sets the value for NumChannels to be an explicit nil

### UnsetNumChannels
`func (o *OracleDBChannelInfoHostInfo) UnsetNumChannels()`

UnsetNumChannels ensures that no value is present for NumChannels, not even an explicit nil
### GetPortnum

`func (o *OracleDBChannelInfoHostInfo) GetPortnum() int64`

GetPortnum returns the Portnum field if non-nil, zero value otherwise.

### GetPortnumOk

`func (o *OracleDBChannelInfoHostInfo) GetPortnumOk() (*int64, bool)`

GetPortnumOk returns a tuple with the Portnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortnum

`func (o *OracleDBChannelInfoHostInfo) SetPortnum(v int64)`

SetPortnum sets Portnum field to given value.

### HasPortnum

`func (o *OracleDBChannelInfoHostInfo) HasPortnum() bool`

HasPortnum returns a boolean if a field has been set.

### SetPortnumNil

`func (o *OracleDBChannelInfoHostInfo) SetPortnumNil(b bool)`

 SetPortnumNil sets the value for Portnum to be an explicit nil

### UnsetPortnum
`func (o *OracleDBChannelInfoHostInfo) UnsetPortnum()`

UnsetPortnum ensures that no value is present for Portnum, not even an explicit nil
### GetSbtHostParams

`func (o *OracleDBChannelInfoHostInfo) GetSbtHostParams() OracleSbtHostParams`

GetSbtHostParams returns the SbtHostParams field if non-nil, zero value otherwise.

### GetSbtHostParamsOk

`func (o *OracleDBChannelInfoHostInfo) GetSbtHostParamsOk() (*OracleSbtHostParams, bool)`

GetSbtHostParamsOk returns a tuple with the SbtHostParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbtHostParams

`func (o *OracleDBChannelInfoHostInfo) SetSbtHostParams(v OracleSbtHostParams)`

SetSbtHostParams sets SbtHostParams field to given value.

### HasSbtHostParams

`func (o *OracleDBChannelInfoHostInfo) HasSbtHostParams() bool`

HasSbtHostParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


