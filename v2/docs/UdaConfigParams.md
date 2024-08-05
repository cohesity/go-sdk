# UdaConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostOsSpecificConfigurations** | Pointer to [**[]UdaHostOsSpecificParams**](UdaHostOsSpecificParams.md) | Array of parameters for different host operating systems. | [optional] 
**Index** | Pointer to [**NullableUdaIndexParams**](UdaIndexParams.md) |  | [optional] 

## Methods

### NewUdaConfigParams

`func NewUdaConfigParams() *UdaConfigParams`

NewUdaConfigParams instantiates a new UdaConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdaConfigParamsWithDefaults

`func NewUdaConfigParamsWithDefaults() *UdaConfigParams`

NewUdaConfigParamsWithDefaults instantiates a new UdaConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostOsSpecificConfigurations

`func (o *UdaConfigParams) GetHostOsSpecificConfigurations() []UdaHostOsSpecificParams`

GetHostOsSpecificConfigurations returns the HostOsSpecificConfigurations field if non-nil, zero value otherwise.

### GetHostOsSpecificConfigurationsOk

`func (o *UdaConfigParams) GetHostOsSpecificConfigurationsOk() (*[]UdaHostOsSpecificParams, bool)`

GetHostOsSpecificConfigurationsOk returns a tuple with the HostOsSpecificConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOsSpecificConfigurations

`func (o *UdaConfigParams) SetHostOsSpecificConfigurations(v []UdaHostOsSpecificParams)`

SetHostOsSpecificConfigurations sets HostOsSpecificConfigurations field to given value.

### HasHostOsSpecificConfigurations

`func (o *UdaConfigParams) HasHostOsSpecificConfigurations() bool`

HasHostOsSpecificConfigurations returns a boolean if a field has been set.

### GetIndex

`func (o *UdaConfigParams) GetIndex() UdaIndexParams`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *UdaConfigParams) GetIndexOk() (*UdaIndexParams, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *UdaConfigParams) SetIndex(v UdaIndexParams)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *UdaConfigParams) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *UdaConfigParams) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *UdaConfigParams) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


