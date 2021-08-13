# ActiveDirectoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdMappingParams** | Pointer to [**NullableIdMappingParams**](IdMappingParams.md) | Specifies the params of the user id mapping info of an Active Directory. | [optional] 
**DomainName** | Pointer to **NullableString** | Specifies the domain name of the Active Directory. | [optional] 
**CentrifyZones** | Pointer to [**[]CentrifyZones**](CentrifyZones.md) | Specifies a list of centrify zones. | [optional] 
**DomainControllers** | Pointer to [**[]DomainControllers**](DomainControllers.md) | A list of domain names with a list of it&#39;s domain controllers. | [optional] 
**SecurityPrincipals** | Pointer to [**[]SecurityPrincipal**](SecurityPrincipal.md) | Specifies a list of security principals. | [optional] 
**Permissions** | Pointer to [**[]Tenant**](Tenant.md) | Specifies the list of tenants that have permissions for this Active Directory. | [optional] 
**TransitiveAdTrustLevelLimit** | Pointer to **NullableInt32** | Specifies level of transitive Active Directory trust domains to be used. | [optional] 

## Methods

### NewActiveDirectoryAllOf

`func NewActiveDirectoryAllOf() *ActiveDirectoryAllOf`

NewActiveDirectoryAllOf instantiates a new ActiveDirectoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDirectoryAllOfWithDefaults

`func NewActiveDirectoryAllOfWithDefaults() *ActiveDirectoryAllOf`

NewActiveDirectoryAllOfWithDefaults instantiates a new ActiveDirectoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdMappingParams

`func (o *ActiveDirectoryAllOf) GetIdMappingParams() IdMappingParams`

GetIdMappingParams returns the IdMappingParams field if non-nil, zero value otherwise.

### GetIdMappingParamsOk

`func (o *ActiveDirectoryAllOf) GetIdMappingParamsOk() (*IdMappingParams, bool)`

GetIdMappingParamsOk returns a tuple with the IdMappingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdMappingParams

`func (o *ActiveDirectoryAllOf) SetIdMappingParams(v IdMappingParams)`

SetIdMappingParams sets IdMappingParams field to given value.

### HasIdMappingParams

`func (o *ActiveDirectoryAllOf) HasIdMappingParams() bool`

HasIdMappingParams returns a boolean if a field has been set.

### SetIdMappingParamsNil

`func (o *ActiveDirectoryAllOf) SetIdMappingParamsNil(b bool)`

 SetIdMappingParamsNil sets the value for IdMappingParams to be an explicit nil

### UnsetIdMappingParams
`func (o *ActiveDirectoryAllOf) UnsetIdMappingParams()`

UnsetIdMappingParams ensures that no value is present for IdMappingParams, not even an explicit nil
### GetDomainName

`func (o *ActiveDirectoryAllOf) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ActiveDirectoryAllOf) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ActiveDirectoryAllOf) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ActiveDirectoryAllOf) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *ActiveDirectoryAllOf) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *ActiveDirectoryAllOf) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetCentrifyZones

`func (o *ActiveDirectoryAllOf) GetCentrifyZones() []CentrifyZones`

GetCentrifyZones returns the CentrifyZones field if non-nil, zero value otherwise.

### GetCentrifyZonesOk

`func (o *ActiveDirectoryAllOf) GetCentrifyZonesOk() (*[]CentrifyZones, bool)`

GetCentrifyZonesOk returns a tuple with the CentrifyZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentrifyZones

`func (o *ActiveDirectoryAllOf) SetCentrifyZones(v []CentrifyZones)`

SetCentrifyZones sets CentrifyZones field to given value.

### HasCentrifyZones

`func (o *ActiveDirectoryAllOf) HasCentrifyZones() bool`

HasCentrifyZones returns a boolean if a field has been set.

### SetCentrifyZonesNil

`func (o *ActiveDirectoryAllOf) SetCentrifyZonesNil(b bool)`

 SetCentrifyZonesNil sets the value for CentrifyZones to be an explicit nil

### UnsetCentrifyZones
`func (o *ActiveDirectoryAllOf) UnsetCentrifyZones()`

UnsetCentrifyZones ensures that no value is present for CentrifyZones, not even an explicit nil
### GetDomainControllers

`func (o *ActiveDirectoryAllOf) GetDomainControllers() []DomainControllers`

GetDomainControllers returns the DomainControllers field if non-nil, zero value otherwise.

### GetDomainControllersOk

`func (o *ActiveDirectoryAllOf) GetDomainControllersOk() (*[]DomainControllers, bool)`

GetDomainControllersOk returns a tuple with the DomainControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainControllers

`func (o *ActiveDirectoryAllOf) SetDomainControllers(v []DomainControllers)`

SetDomainControllers sets DomainControllers field to given value.

### HasDomainControllers

`func (o *ActiveDirectoryAllOf) HasDomainControllers() bool`

HasDomainControllers returns a boolean if a field has been set.

### SetDomainControllersNil

`func (o *ActiveDirectoryAllOf) SetDomainControllersNil(b bool)`

 SetDomainControllersNil sets the value for DomainControllers to be an explicit nil

### UnsetDomainControllers
`func (o *ActiveDirectoryAllOf) UnsetDomainControllers()`

UnsetDomainControllers ensures that no value is present for DomainControllers, not even an explicit nil
### GetSecurityPrincipals

`func (o *ActiveDirectoryAllOf) GetSecurityPrincipals() []SecurityPrincipal`

GetSecurityPrincipals returns the SecurityPrincipals field if non-nil, zero value otherwise.

### GetSecurityPrincipalsOk

`func (o *ActiveDirectoryAllOf) GetSecurityPrincipalsOk() (*[]SecurityPrincipal, bool)`

GetSecurityPrincipalsOk returns a tuple with the SecurityPrincipals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPrincipals

`func (o *ActiveDirectoryAllOf) SetSecurityPrincipals(v []SecurityPrincipal)`

SetSecurityPrincipals sets SecurityPrincipals field to given value.

### HasSecurityPrincipals

`func (o *ActiveDirectoryAllOf) HasSecurityPrincipals() bool`

HasSecurityPrincipals returns a boolean if a field has been set.

### SetSecurityPrincipalsNil

`func (o *ActiveDirectoryAllOf) SetSecurityPrincipalsNil(b bool)`

 SetSecurityPrincipalsNil sets the value for SecurityPrincipals to be an explicit nil

### UnsetSecurityPrincipals
`func (o *ActiveDirectoryAllOf) UnsetSecurityPrincipals()`

UnsetSecurityPrincipals ensures that no value is present for SecurityPrincipals, not even an explicit nil
### GetPermissions

`func (o *ActiveDirectoryAllOf) GetPermissions() []Tenant`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ActiveDirectoryAllOf) GetPermissionsOk() (*[]Tenant, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ActiveDirectoryAllOf) SetPermissions(v []Tenant)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ActiveDirectoryAllOf) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *ActiveDirectoryAllOf) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ActiveDirectoryAllOf) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetTransitiveAdTrustLevelLimit

`func (o *ActiveDirectoryAllOf) GetTransitiveAdTrustLevelLimit() int32`

GetTransitiveAdTrustLevelLimit returns the TransitiveAdTrustLevelLimit field if non-nil, zero value otherwise.

### GetTransitiveAdTrustLevelLimitOk

`func (o *ActiveDirectoryAllOf) GetTransitiveAdTrustLevelLimitOk() (*int32, bool)`

GetTransitiveAdTrustLevelLimitOk returns a tuple with the TransitiveAdTrustLevelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveAdTrustLevelLimit

`func (o *ActiveDirectoryAllOf) SetTransitiveAdTrustLevelLimit(v int32)`

SetTransitiveAdTrustLevelLimit sets TransitiveAdTrustLevelLimit field to given value.

### HasTransitiveAdTrustLevelLimit

`func (o *ActiveDirectoryAllOf) HasTransitiveAdTrustLevelLimit() bool`

HasTransitiveAdTrustLevelLimit returns a boolean if a field has been set.

### SetTransitiveAdTrustLevelLimitNil

`func (o *ActiveDirectoryAllOf) SetTransitiveAdTrustLevelLimitNil(b bool)`

 SetTransitiveAdTrustLevelLimitNil sets the value for TransitiveAdTrustLevelLimit to be an explicit nil

### UnsetTransitiveAdTrustLevelLimit
`func (o *ActiveDirectoryAllOf) UnsetTransitiveAdTrustLevelLimit()`

UnsetTransitiveAdTrustLevelLimit ensures that no value is present for TransitiveAdTrustLevelLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


