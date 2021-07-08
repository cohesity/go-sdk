# StopRemoteVaultSearchJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchJobUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the unique id of the Remote Vault search job in progress. | [optional] 

## Methods

### NewStopRemoteVaultSearchJobParameters

`func NewStopRemoteVaultSearchJobParameters() *StopRemoteVaultSearchJobParameters`

NewStopRemoteVaultSearchJobParameters instantiates a new StopRemoteVaultSearchJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopRemoteVaultSearchJobParametersWithDefaults

`func NewStopRemoteVaultSearchJobParametersWithDefaults() *StopRemoteVaultSearchJobParameters`

NewStopRemoteVaultSearchJobParametersWithDefaults instantiates a new StopRemoteVaultSearchJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchJobUid

`func (o *StopRemoteVaultSearchJobParameters) GetSearchJobUid() UniversalId`

GetSearchJobUid returns the SearchJobUid field if non-nil, zero value otherwise.

### GetSearchJobUidOk

`func (o *StopRemoteVaultSearchJobParameters) GetSearchJobUidOk() (*UniversalId, bool)`

GetSearchJobUidOk returns a tuple with the SearchJobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchJobUid

`func (o *StopRemoteVaultSearchJobParameters) SetSearchJobUid(v UniversalId)`

SetSearchJobUid sets SearchJobUid field to given value.

### HasSearchJobUid

`func (o *StopRemoteVaultSearchJobParameters) HasSearchJobUid() bool`

HasSearchJobUid returns a boolean if a field has been set.

### SetSearchJobUidNil

`func (o *StopRemoteVaultSearchJobParameters) SetSearchJobUidNil(b bool)`

 SetSearchJobUidNil sets the value for SearchJobUid to be an explicit nil

### UnsetSearchJobUid
`func (o *StopRemoteVaultSearchJobParameters) UnsetSearchJobUid()`

UnsetSearchJobUid ensures that no value is present for SearchJobUid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


