# PhysicalSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to **[]string** | Specifies the list of applications to be registered with Physical Source. | [optional] 
**Endpoint** | **string** | Specifies the endpoint IPaddress, URL or hostname of the physical host. | 
**ForceRegister** | Pointer to **NullableBool** | The agent running on a physical host will fail the registration if it is already registered as part of another cluster. By setting this option to true, agent can be forced to register with the current cluster. | [optional] 
**HostType** | Pointer to **NullableString** | Specifies the type of host. | [optional] 
**PhysicalType** | Pointer to **NullableString** | Specifies the type of physical server. | [optional] 

## Methods

### NewPhysicalSourceRegistrationParams

`func NewPhysicalSourceRegistrationParams(endpoint string, ) *PhysicalSourceRegistrationParams`

NewPhysicalSourceRegistrationParams instantiates a new PhysicalSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalSourceRegistrationParamsWithDefaults

`func NewPhysicalSourceRegistrationParamsWithDefaults() *PhysicalSourceRegistrationParams`

NewPhysicalSourceRegistrationParamsWithDefaults instantiates a new PhysicalSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *PhysicalSourceRegistrationParams) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *PhysicalSourceRegistrationParams) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *PhysicalSourceRegistrationParams) SetApplications(v []string)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *PhysicalSourceRegistrationParams) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### SetApplicationsNil

`func (o *PhysicalSourceRegistrationParams) SetApplicationsNil(b bool)`

 SetApplicationsNil sets the value for Applications to be an explicit nil

### UnsetApplications
`func (o *PhysicalSourceRegistrationParams) UnsetApplications()`

UnsetApplications ensures that no value is present for Applications, not even an explicit nil
### GetEndpoint

`func (o *PhysicalSourceRegistrationParams) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PhysicalSourceRegistrationParams) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PhysicalSourceRegistrationParams) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetForceRegister

`func (o *PhysicalSourceRegistrationParams) GetForceRegister() bool`

GetForceRegister returns the ForceRegister field if non-nil, zero value otherwise.

### GetForceRegisterOk

`func (o *PhysicalSourceRegistrationParams) GetForceRegisterOk() (*bool, bool)`

GetForceRegisterOk returns a tuple with the ForceRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRegister

`func (o *PhysicalSourceRegistrationParams) SetForceRegister(v bool)`

SetForceRegister sets ForceRegister field to given value.

### HasForceRegister

`func (o *PhysicalSourceRegistrationParams) HasForceRegister() bool`

HasForceRegister returns a boolean if a field has been set.

### SetForceRegisterNil

`func (o *PhysicalSourceRegistrationParams) SetForceRegisterNil(b bool)`

 SetForceRegisterNil sets the value for ForceRegister to be an explicit nil

### UnsetForceRegister
`func (o *PhysicalSourceRegistrationParams) UnsetForceRegister()`

UnsetForceRegister ensures that no value is present for ForceRegister, not even an explicit nil
### GetHostType

`func (o *PhysicalSourceRegistrationParams) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *PhysicalSourceRegistrationParams) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *PhysicalSourceRegistrationParams) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *PhysicalSourceRegistrationParams) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### SetHostTypeNil

`func (o *PhysicalSourceRegistrationParams) SetHostTypeNil(b bool)`

 SetHostTypeNil sets the value for HostType to be an explicit nil

### UnsetHostType
`func (o *PhysicalSourceRegistrationParams) UnsetHostType()`

UnsetHostType ensures that no value is present for HostType, not even an explicit nil
### GetPhysicalType

`func (o *PhysicalSourceRegistrationParams) GetPhysicalType() string`

GetPhysicalType returns the PhysicalType field if non-nil, zero value otherwise.

### GetPhysicalTypeOk

`func (o *PhysicalSourceRegistrationParams) GetPhysicalTypeOk() (*string, bool)`

GetPhysicalTypeOk returns a tuple with the PhysicalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalType

`func (o *PhysicalSourceRegistrationParams) SetPhysicalType(v string)`

SetPhysicalType sets PhysicalType field to given value.

### HasPhysicalType

`func (o *PhysicalSourceRegistrationParams) HasPhysicalType() bool`

HasPhysicalType returns a boolean if a field has been set.

### SetPhysicalTypeNil

`func (o *PhysicalSourceRegistrationParams) SetPhysicalTypeNil(b bool)`

 SetPhysicalTypeNil sets the value for PhysicalType to be an explicit nil

### UnsetPhysicalType
`func (o *PhysicalSourceRegistrationParams) UnsetPhysicalType()`

UnsetPhysicalType ensures that no value is present for PhysicalType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


