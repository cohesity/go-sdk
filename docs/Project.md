# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to [**Domain**](Domain.md) |  | [optional] 
**DomainId** | Pointer to **NullableString** | The ID of the domain to which the project is scoped. This field is used in the reponse of Keystone API. | [optional] 
**Id** | Pointer to **NullableString** | The ID of the project. | [optional] 
**Name** | Pointer to **NullableString** | The name of the project. | [optional] 

## Methods

### NewProject

`func NewProject() *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *Project) GetDomain() Domain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Project) GetDomainOk() (*Domain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Project) SetDomain(v Domain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Project) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDomainId

`func (o *Project) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *Project) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *Project) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *Project) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### SetDomainIdNil

`func (o *Project) SetDomainIdNil(b bool)`

 SetDomainIdNil sets the value for DomainId to be an explicit nil

### UnsetDomainId
`func (o *Project) UnsetDomainId()`

UnsetDomainId ensures that no value is present for DomainId, not even an explicit nil
### GetId

`func (o *Project) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Project) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Project) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Project) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Project) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Project) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Project) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


