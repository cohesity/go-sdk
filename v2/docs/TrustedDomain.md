# TrustedDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainControllersDenyList** | Pointer to **[]string** | Specifies a list of denied domain controllers of this domain. | [optional] 
**DomainName** | Pointer to **NullableString** | Specifies a domain name. | [optional] 
**PreferredDomainControllers** | Pointer to [**[]DomainController**](DomainController.md) | Specifies a list of preferred domain controllers for this domain. | [optional] 

## Methods

### NewTrustedDomain

`func NewTrustedDomain() *TrustedDomain`

NewTrustedDomain instantiates a new TrustedDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustedDomainWithDefaults

`func NewTrustedDomainWithDefaults() *TrustedDomain`

NewTrustedDomainWithDefaults instantiates a new TrustedDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainControllersDenyList

`func (o *TrustedDomain) GetDomainControllersDenyList() []*string`

GetDomainControllersDenyList returns the DomainControllersDenyList field if non-nil, zero value otherwise.

### GetDomainControllersDenyListOk

`func (o *TrustedDomain) GetDomainControllersDenyListOk() (*[]*string, bool)`

GetDomainControllersDenyListOk returns a tuple with the DomainControllersDenyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainControllersDenyList

`func (o *TrustedDomain) SetDomainControllersDenyList(v []*string)`

SetDomainControllersDenyList sets DomainControllersDenyList field to given value.

### HasDomainControllersDenyList

`func (o *TrustedDomain) HasDomainControllersDenyList() bool`

HasDomainControllersDenyList returns a boolean if a field has been set.

### GetDomainName

`func (o *TrustedDomain) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *TrustedDomain) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *TrustedDomain) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *TrustedDomain) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *TrustedDomain) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *TrustedDomain) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetPreferredDomainControllers

`func (o *TrustedDomain) GetPreferredDomainControllers() []DomainController`

GetPreferredDomainControllers returns the PreferredDomainControllers field if non-nil, zero value otherwise.

### GetPreferredDomainControllersOk

`func (o *TrustedDomain) GetPreferredDomainControllersOk() (*[]DomainController, bool)`

GetPreferredDomainControllersOk returns a tuple with the PreferredDomainControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDomainControllers

`func (o *TrustedDomain) SetPreferredDomainControllers(v []DomainController)`

SetPreferredDomainControllers sets PreferredDomainControllers field to given value.

### HasPreferredDomainControllers

`func (o *TrustedDomain) HasPreferredDomainControllers() bool`

HasPreferredDomainControllers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


