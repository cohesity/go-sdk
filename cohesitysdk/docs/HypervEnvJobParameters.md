# HypervEnvJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FallbackToCrashConsistent** | Pointer to **NullableBool** | If true, takes a crash-consistent snapshot when app-consistent snapshot fails. Otherwise, the snapshot attempt is marked failed. By default, this field is set to false. | [optional] 

## Methods

### NewHypervEnvJobParameters

`func NewHypervEnvJobParameters() *HypervEnvJobParameters`

NewHypervEnvJobParameters instantiates a new HypervEnvJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervEnvJobParametersWithDefaults

`func NewHypervEnvJobParametersWithDefaults() *HypervEnvJobParameters`

NewHypervEnvJobParametersWithDefaults instantiates a new HypervEnvJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFallbackToCrashConsistent

`func (o *HypervEnvJobParameters) GetFallbackToCrashConsistent() bool`

GetFallbackToCrashConsistent returns the FallbackToCrashConsistent field if non-nil, zero value otherwise.

### GetFallbackToCrashConsistentOk

`func (o *HypervEnvJobParameters) GetFallbackToCrashConsistentOk() (*bool, bool)`

GetFallbackToCrashConsistentOk returns a tuple with the FallbackToCrashConsistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToCrashConsistent

`func (o *HypervEnvJobParameters) SetFallbackToCrashConsistent(v bool)`

SetFallbackToCrashConsistent sets FallbackToCrashConsistent field to given value.

### HasFallbackToCrashConsistent

`func (o *HypervEnvJobParameters) HasFallbackToCrashConsistent() bool`

HasFallbackToCrashConsistent returns a boolean if a field has been set.

### SetFallbackToCrashConsistentNil

`func (o *HypervEnvJobParameters) SetFallbackToCrashConsistentNil(b bool)`

 SetFallbackToCrashConsistentNil sets the value for FallbackToCrashConsistent to be an explicit nil

### UnsetFallbackToCrashConsistent
`func (o *HypervEnvJobParameters) UnsetFallbackToCrashConsistent()`

UnsetFallbackToCrashConsistent ensures that no value is present for FallbackToCrashConsistent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


