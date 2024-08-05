# CassandraProtectionRunParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SetPrimaryForLog** | Pointer to **NullableBool** | Specifies the parameters to set this cluster as primary and trigger a new protection run for Log backup job. Default value is false. | [optional] [default to false]

## Methods

### NewCassandraProtectionRunParams

`func NewCassandraProtectionRunParams() *CassandraProtectionRunParams`

NewCassandraProtectionRunParams instantiates a new CassandraProtectionRunParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraProtectionRunParamsWithDefaults

`func NewCassandraProtectionRunParamsWithDefaults() *CassandraProtectionRunParams`

NewCassandraProtectionRunParamsWithDefaults instantiates a new CassandraProtectionRunParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSetPrimaryForLog

`func (o *CassandraProtectionRunParams) GetSetPrimaryForLog() bool`

GetSetPrimaryForLog returns the SetPrimaryForLog field if non-nil, zero value otherwise.

### GetSetPrimaryForLogOk

`func (o *CassandraProtectionRunParams) GetSetPrimaryForLogOk() (*bool, bool)`

GetSetPrimaryForLogOk returns a tuple with the SetPrimaryForLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetPrimaryForLog

`func (o *CassandraProtectionRunParams) SetSetPrimaryForLog(v bool)`

SetSetPrimaryForLog sets SetPrimaryForLog field to given value.

### HasSetPrimaryForLog

`func (o *CassandraProtectionRunParams) HasSetPrimaryForLog() bool`

HasSetPrimaryForLog returns a boolean if a field has been set.

### SetSetPrimaryForLogNil

`func (o *CassandraProtectionRunParams) SetSetPrimaryForLogNil(b bool)`

 SetSetPrimaryForLogNil sets the value for SetPrimaryForLog to be an explicit nil

### UnsetSetPrimaryForLog
`func (o *CassandraProtectionRunParams) UnsetSetPrimaryForLog()`

UnsetSetPrimaryForLog ensures that no value is present for SetPrimaryForLog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


