# ViewUserQuotaParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserQuotaPolicy** | Pointer to [**UserQuota**](UserQuota.md) |  | [optional] 
**ViewName** | Pointer to **NullableString** | View name of input view. | [optional] 

## Methods

### NewViewUserQuotaParameters

`func NewViewUserQuotaParameters() *ViewUserQuotaParameters`

NewViewUserQuotaParameters instantiates a new ViewUserQuotaParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewUserQuotaParametersWithDefaults

`func NewViewUserQuotaParametersWithDefaults() *ViewUserQuotaParameters`

NewViewUserQuotaParametersWithDefaults instantiates a new ViewUserQuotaParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserQuotaPolicy

`func (o *ViewUserQuotaParameters) GetUserQuotaPolicy() UserQuota`

GetUserQuotaPolicy returns the UserQuotaPolicy field if non-nil, zero value otherwise.

### GetUserQuotaPolicyOk

`func (o *ViewUserQuotaParameters) GetUserQuotaPolicyOk() (*UserQuota, bool)`

GetUserQuotaPolicyOk returns a tuple with the UserQuotaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuotaPolicy

`func (o *ViewUserQuotaParameters) SetUserQuotaPolicy(v UserQuota)`

SetUserQuotaPolicy sets UserQuotaPolicy field to given value.

### HasUserQuotaPolicy

`func (o *ViewUserQuotaParameters) HasUserQuotaPolicy() bool`

HasUserQuotaPolicy returns a boolean if a field has been set.

### GetViewName

`func (o *ViewUserQuotaParameters) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ViewUserQuotaParameters) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ViewUserQuotaParameters) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *ViewUserQuotaParameters) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *ViewUserQuotaParameters) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *ViewUserQuotaParameters) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


