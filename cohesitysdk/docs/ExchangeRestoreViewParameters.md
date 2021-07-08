# ExchangeRestoreViewParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to **NullableBool** | Specifies whether we should white-list the restore view for all the IP addresses. If this parameter is set to false, then only the machine on which the view is mounted will be white-listed. | [optional] 
**MountPoint** | Pointer to **NullableString** | Specifies the path of the SMB share. | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the view to which the files of an Exchange database has to be cloned. | [optional] 

## Methods

### NewExchangeRestoreViewParameters

`func NewExchangeRestoreViewParameters() *ExchangeRestoreViewParameters`

NewExchangeRestoreViewParameters instantiates a new ExchangeRestoreViewParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRestoreViewParametersWithDefaults

`func NewExchangeRestoreViewParametersWithDefaults() *ExchangeRestoreViewParameters`

NewExchangeRestoreViewParametersWithDefaults instantiates a new ExchangeRestoreViewParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *ExchangeRestoreViewParameters) GetEndpoint() bool`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ExchangeRestoreViewParameters) GetEndpointOk() (*bool, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ExchangeRestoreViewParameters) SetEndpoint(v bool)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ExchangeRestoreViewParameters) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *ExchangeRestoreViewParameters) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *ExchangeRestoreViewParameters) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetMountPoint

`func (o *ExchangeRestoreViewParameters) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *ExchangeRestoreViewParameters) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *ExchangeRestoreViewParameters) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *ExchangeRestoreViewParameters) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### SetMountPointNil

`func (o *ExchangeRestoreViewParameters) SetMountPointNil(b bool)`

 SetMountPointNil sets the value for MountPoint to be an explicit nil

### UnsetMountPoint
`func (o *ExchangeRestoreViewParameters) UnsetMountPoint()`

UnsetMountPoint ensures that no value is present for MountPoint, not even an explicit nil
### GetViewName

`func (o *ExchangeRestoreViewParameters) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ExchangeRestoreViewParameters) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ExchangeRestoreViewParameters) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *ExchangeRestoreViewParameters) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *ExchangeRestoreViewParameters) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *ExchangeRestoreViewParameters) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


