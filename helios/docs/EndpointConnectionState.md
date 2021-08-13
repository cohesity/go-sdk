# EndpointConnectionState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the name of the endpoint | [optional] 
**Port** | Pointer to **NullableString** | Specifies the port of the endpoint | [optional] 
**Results** | Pointer to [**[]EndpointCheckResult**](EndpointCheckResult.md) | Specifies the results of the endpoints. | [optional] 

## Methods

### NewEndpointConnectionState

`func NewEndpointConnectionState() *EndpointConnectionState`

NewEndpointConnectionState instantiates a new EndpointConnectionState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointConnectionStateWithDefaults

`func NewEndpointConnectionStateWithDefaults() *EndpointConnectionState`

NewEndpointConnectionStateWithDefaults instantiates a new EndpointConnectionState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EndpointConnectionState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointConnectionState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointConnectionState) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EndpointConnectionState) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EndpointConnectionState) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EndpointConnectionState) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPort

`func (o *EndpointConnectionState) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EndpointConnectionState) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EndpointConnectionState) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *EndpointConnectionState) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *EndpointConnectionState) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *EndpointConnectionState) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetResults

`func (o *EndpointConnectionState) GetResults() []EndpointCheckResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *EndpointConnectionState) GetResultsOk() (*[]EndpointCheckResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *EndpointConnectionState) SetResults(v []EndpointCheckResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *EndpointConnectionState) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *EndpointConnectionState) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *EndpointConnectionState) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


