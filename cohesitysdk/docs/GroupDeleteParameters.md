# GroupDeleteParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | Specifies the domain associated with the groups to delete. Only groups associated with the same domain can be deleted by a single request. If no domain is specified, the specified groups are deleted from the LOCAL domain on the Cohesity Cluster. If a non-LOCAL domain is specified, the specified groups are deleted on the Cohesity Cluster. However, the referenced group principals on the Active Directory are not deleted. | [optional] 
**Names** | Pointer to **[]string** | Array of Groups.  Specifies the list of groups to delete on the Cohesity Cluster. Only groups from the same domain can be deleted by a single request. | [optional] 

## Methods

### NewGroupDeleteParameters

`func NewGroupDeleteParameters() *GroupDeleteParameters`

NewGroupDeleteParameters instantiates a new GroupDeleteParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupDeleteParametersWithDefaults

`func NewGroupDeleteParametersWithDefaults() *GroupDeleteParameters`

NewGroupDeleteParametersWithDefaults instantiates a new GroupDeleteParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *GroupDeleteParameters) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GroupDeleteParameters) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GroupDeleteParameters) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GroupDeleteParameters) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *GroupDeleteParameters) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *GroupDeleteParameters) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetNames

`func (o *GroupDeleteParameters) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *GroupDeleteParameters) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *GroupDeleteParameters) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *GroupDeleteParameters) HasNames() bool`

HasNames returns a boolean if a field has been set.

### SetNamesNil

`func (o *GroupDeleteParameters) SetNamesNil(b bool)`

 SetNamesNil sets the value for Names to be an explicit nil

### UnsetNames
`func (o *GroupDeleteParameters) UnsetNames()`

UnsetNames ensures that no value is present for Names, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


