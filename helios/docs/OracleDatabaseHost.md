# OracleDatabaseHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **NullableString** | Specifies the id of the database host from which backup is allowed. | [optional] 
**ChannelCount** | Pointer to **int32** | Specifies the number of channels to be created for this host. Default value for the number of channels will be calculated as the minimum of number of nodes in Cohesity cluster and 2 * number of CPU on the host. | [optional] 
**Port** | Pointer to **int64** | Specifies the port where the Database is listening. | [optional] 
**SbtHostParams** | Pointer to [**OracleSbtHostParams**](OracleSbtHostParams.md) |  | [optional] 

## Methods

### NewOracleDatabaseHost

`func NewOracleDatabaseHost() *OracleDatabaseHost`

NewOracleDatabaseHost instantiates a new OracleDatabaseHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleDatabaseHostWithDefaults

`func NewOracleDatabaseHostWithDefaults() *OracleDatabaseHost`

NewOracleDatabaseHostWithDefaults instantiates a new OracleDatabaseHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *OracleDatabaseHost) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *OracleDatabaseHost) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *OracleDatabaseHost) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *OracleDatabaseHost) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### SetHostIdNil

`func (o *OracleDatabaseHost) SetHostIdNil(b bool)`

 SetHostIdNil sets the value for HostId to be an explicit nil

### UnsetHostId
`func (o *OracleDatabaseHost) UnsetHostId()`

UnsetHostId ensures that no value is present for HostId, not even an explicit nil
### GetChannelCount

`func (o *OracleDatabaseHost) GetChannelCount() int32`

GetChannelCount returns the ChannelCount field if non-nil, zero value otherwise.

### GetChannelCountOk

`func (o *OracleDatabaseHost) GetChannelCountOk() (*int32, bool)`

GetChannelCountOk returns a tuple with the ChannelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCount

`func (o *OracleDatabaseHost) SetChannelCount(v int32)`

SetChannelCount sets ChannelCount field to given value.

### HasChannelCount

`func (o *OracleDatabaseHost) HasChannelCount() bool`

HasChannelCount returns a boolean if a field has been set.

### GetPort

`func (o *OracleDatabaseHost) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OracleDatabaseHost) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OracleDatabaseHost) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *OracleDatabaseHost) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSbtHostParams

`func (o *OracleDatabaseHost) GetSbtHostParams() OracleSbtHostParams`

GetSbtHostParams returns the SbtHostParams field if non-nil, zero value otherwise.

### GetSbtHostParamsOk

`func (o *OracleDatabaseHost) GetSbtHostParamsOk() (*OracleSbtHostParams, bool)`

GetSbtHostParamsOk returns a tuple with the SbtHostParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbtHostParams

`func (o *OracleDatabaseHost) SetSbtHostParams(v OracleSbtHostParams)`

SetSbtHostParams sets SbtHostParams field to given value.

### HasSbtHostParams

`func (o *OracleDatabaseHost) HasSbtHostParams() bool`

HasSbtHostParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


