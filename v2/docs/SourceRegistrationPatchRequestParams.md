# SourceRegistrationPatchRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CassandraParams** | Pointer to [**CassandraSourceRegistrationPatchParams**](CassandraSourceRegistrationPatchParams.md) |  | [optional] 
**Environment** | **NullableString** | Specifies the environment type of the Protection Source to be patched. Currently the only environment supported is kCassandra | 

## Methods

### NewSourceRegistrationPatchRequestParams

`func NewSourceRegistrationPatchRequestParams(environment NullableString, ) *SourceRegistrationPatchRequestParams`

NewSourceRegistrationPatchRequestParams instantiates a new SourceRegistrationPatchRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceRegistrationPatchRequestParamsWithDefaults

`func NewSourceRegistrationPatchRequestParamsWithDefaults() *SourceRegistrationPatchRequestParams`

NewSourceRegistrationPatchRequestParamsWithDefaults instantiates a new SourceRegistrationPatchRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraParams

`func (o *SourceRegistrationPatchRequestParams) GetCassandraParams() CassandraSourceRegistrationPatchParams`

GetCassandraParams returns the CassandraParams field if non-nil, zero value otherwise.

### GetCassandraParamsOk

`func (o *SourceRegistrationPatchRequestParams) GetCassandraParamsOk() (*CassandraSourceRegistrationPatchParams, bool)`

GetCassandraParamsOk returns a tuple with the CassandraParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraParams

`func (o *SourceRegistrationPatchRequestParams) SetCassandraParams(v CassandraSourceRegistrationPatchParams)`

SetCassandraParams sets CassandraParams field to given value.

### HasCassandraParams

`func (o *SourceRegistrationPatchRequestParams) HasCassandraParams() bool`

HasCassandraParams returns a boolean if a field has been set.

### GetEnvironment

`func (o *SourceRegistrationPatchRequestParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SourceRegistrationPatchRequestParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SourceRegistrationPatchRequestParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *SourceRegistrationPatchRequestParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *SourceRegistrationPatchRequestParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


