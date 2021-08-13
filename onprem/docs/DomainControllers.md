# DomainControllers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **NullableString** | Specifies the domain name. | [optional] 
**Controllers** | Pointer to [**[]DomainController**](DomainController.md) | Specifies a list of domain controllers of the domain. | [optional] 

## Methods

### NewDomainControllers

`func NewDomainControllers() *DomainControllers`

NewDomainControllers instantiates a new DomainControllers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainControllersWithDefaults

`func NewDomainControllersWithDefaults() *DomainControllers`

NewDomainControllersWithDefaults instantiates a new DomainControllers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *DomainControllers) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DomainControllers) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DomainControllers) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DomainControllers) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *DomainControllers) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *DomainControllers) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetControllers

`func (o *DomainControllers) GetControllers() []DomainController`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *DomainControllers) GetControllersOk() (*[]DomainController, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *DomainControllers) SetControllers(v []DomainController)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *DomainControllers) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### SetControllersNil

`func (o *DomainControllers) SetControllersNil(b bool)`

 SetControllersNil sets the value for Controllers to be an explicit nil

### UnsetControllers
`func (o *DomainControllers) UnsetControllers()`

UnsetControllers ensures that no value is present for Controllers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


