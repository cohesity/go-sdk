# OracleConnectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **NullableString** | Specifies the unique identifier to locate the Oracle node or cluster. The host identifier can be IP address or FQDN. | 

## Methods

### NewOracleConnectionParams

`func NewOracleConnectionParams(hostname NullableString, ) *OracleConnectionParams`

NewOracleConnectionParams instantiates a new OracleConnectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleConnectionParamsWithDefaults

`func NewOracleConnectionParamsWithDefaults() *OracleConnectionParams`

NewOracleConnectionParamsWithDefaults instantiates a new OracleConnectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *OracleConnectionParams) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *OracleConnectionParams) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *OracleConnectionParams) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### SetHostnameNil

`func (o *OracleConnectionParams) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *OracleConnectionParams) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


