# UdaCreateRunResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternallyTriggeredRunId** | Pointer to **NullableString** | Specifies a unique run id for an externally triggered run. This id can be used by the caller to identify the corresponding backup process on the source host. | [optional] 

## Methods

### NewUdaCreateRunResponseParams

`func NewUdaCreateRunResponseParams() *UdaCreateRunResponseParams`

NewUdaCreateRunResponseParams instantiates a new UdaCreateRunResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdaCreateRunResponseParamsWithDefaults

`func NewUdaCreateRunResponseParamsWithDefaults() *UdaCreateRunResponseParams`

NewUdaCreateRunResponseParamsWithDefaults instantiates a new UdaCreateRunResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternallyTriggeredRunId

`func (o *UdaCreateRunResponseParams) GetExternallyTriggeredRunId() string`

GetExternallyTriggeredRunId returns the ExternallyTriggeredRunId field if non-nil, zero value otherwise.

### GetExternallyTriggeredRunIdOk

`func (o *UdaCreateRunResponseParams) GetExternallyTriggeredRunIdOk() (*string, bool)`

GetExternallyTriggeredRunIdOk returns a tuple with the ExternallyTriggeredRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyTriggeredRunId

`func (o *UdaCreateRunResponseParams) SetExternallyTriggeredRunId(v string)`

SetExternallyTriggeredRunId sets ExternallyTriggeredRunId field to given value.

### HasExternallyTriggeredRunId

`func (o *UdaCreateRunResponseParams) HasExternallyTriggeredRunId() bool`

HasExternallyTriggeredRunId returns a boolean if a field has been set.

### SetExternallyTriggeredRunIdNil

`func (o *UdaCreateRunResponseParams) SetExternallyTriggeredRunIdNil(b bool)`

 SetExternallyTriggeredRunIdNil sets the value for ExternallyTriggeredRunId to be an explicit nil

### UnsetExternallyTriggeredRunId
`func (o *UdaCreateRunResponseParams) UnsetExternallyTriggeredRunId()`

UnsetExternallyTriggeredRunId ensures that no value is present for ExternallyTriggeredRunId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


