# UpdateAppInstanceStateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **NullableString** | Specifies the desired app instance state type. Specifies operational status of an app instance. kInitializing - The app instance has been launched or resumed, but is not fully running yet. kRunning - The app instance is running. Check health_status for the actual health. kPausing - The app instance is being paused. kPaused - The app instance has been paused. kTerminating - The app instance is being terminated. kTerminated -  The app instance has been terminated. kFailed - The app instance has failed due to an unrecoverable error. | [optional] 

## Methods

### NewUpdateAppInstanceStateParameters

`func NewUpdateAppInstanceStateParameters() *UpdateAppInstanceStateParameters`

NewUpdateAppInstanceStateParameters instantiates a new UpdateAppInstanceStateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAppInstanceStateParametersWithDefaults

`func NewUpdateAppInstanceStateParametersWithDefaults() *UpdateAppInstanceStateParameters`

NewUpdateAppInstanceStateParametersWithDefaults instantiates a new UpdateAppInstanceStateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *UpdateAppInstanceStateParameters) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateAppInstanceStateParameters) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateAppInstanceStateParameters) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateAppInstanceStateParameters) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *UpdateAppInstanceStateParameters) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *UpdateAppInstanceStateParameters) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


