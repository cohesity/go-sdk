# PreferredDomainController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainControllers** | Pointer to **[]string** | List of Domain controllers DCs in FQDN format that are mapped to an Active Directory Domain name. | [optional] 
**DomainName** | Pointer to **NullableString** | Specifies the Domain name or the trusted domain of an Active Directory. | [optional] 

## Methods

### NewPreferredDomainController

`func NewPreferredDomainController() *PreferredDomainController`

NewPreferredDomainController instantiates a new PreferredDomainController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreferredDomainControllerWithDefaults

`func NewPreferredDomainControllerWithDefaults() *PreferredDomainController`

NewPreferredDomainControllerWithDefaults instantiates a new PreferredDomainController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainControllers

`func (o *PreferredDomainController) GetDomainControllers() []string`

GetDomainControllers returns the DomainControllers field if non-nil, zero value otherwise.

### GetDomainControllersOk

`func (o *PreferredDomainController) GetDomainControllersOk() (*[]string, bool)`

GetDomainControllersOk returns a tuple with the DomainControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainControllers

`func (o *PreferredDomainController) SetDomainControllers(v []string)`

SetDomainControllers sets DomainControllers field to given value.

### HasDomainControllers

`func (o *PreferredDomainController) HasDomainControllers() bool`

HasDomainControllers returns a boolean if a field has been set.

### SetDomainControllersNil

`func (o *PreferredDomainController) SetDomainControllersNil(b bool)`

 SetDomainControllersNil sets the value for DomainControllers to be an explicit nil

### UnsetDomainControllers
`func (o *PreferredDomainController) UnsetDomainControllers()`

UnsetDomainControllers ensures that no value is present for DomainControllers, not even an explicit nil
### GetDomainName

`func (o *PreferredDomainController) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *PreferredDomainController) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *PreferredDomainController) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *PreferredDomainController) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *PreferredDomainController) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *PreferredDomainController) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


