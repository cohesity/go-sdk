# NisNetgroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Description of the netgroup. | [optional] 
**Domain** | Pointer to **NullableString** | Specifies the domain of the netgroup. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the netgroup. | [optional] 
**NfsAccess** | Pointer to **NullableString** | Specifies whether clients from this netgroup can mount using NFS protocol. Protocol access level. &#39;kDisabled&#39; indicates Protocol access level &#39;Disabled&#39; &#39;kReadOnly&#39; indicates Protocol access level &#39;ReadOnly&#39; &#39;kReadWrite&#39; indicates Protocol access level &#39;ReadWrite&#39; | [optional] 
**NfsSquash** | Pointer to **NullableInt32** | Specifies the NFS squash type. | [optional] 

## Methods

### NewNisNetgroup

`func NewNisNetgroup() *NisNetgroup`

NewNisNetgroup instantiates a new NisNetgroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNisNetgroupWithDefaults

`func NewNisNetgroupWithDefaults() *NisNetgroup`

NewNisNetgroupWithDefaults instantiates a new NisNetgroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NisNetgroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NisNetgroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NisNetgroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NisNetgroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NisNetgroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NisNetgroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomain

`func (o *NisNetgroup) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *NisNetgroup) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *NisNetgroup) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *NisNetgroup) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *NisNetgroup) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *NisNetgroup) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetName

`func (o *NisNetgroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NisNetgroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NisNetgroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NisNetgroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NisNetgroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NisNetgroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNfsAccess

`func (o *NisNetgroup) GetNfsAccess() string`

GetNfsAccess returns the NfsAccess field if non-nil, zero value otherwise.

### GetNfsAccessOk

`func (o *NisNetgroup) GetNfsAccessOk() (*string, bool)`

GetNfsAccessOk returns a tuple with the NfsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAccess

`func (o *NisNetgroup) SetNfsAccess(v string)`

SetNfsAccess sets NfsAccess field to given value.

### HasNfsAccess

`func (o *NisNetgroup) HasNfsAccess() bool`

HasNfsAccess returns a boolean if a field has been set.

### SetNfsAccessNil

`func (o *NisNetgroup) SetNfsAccessNil(b bool)`

 SetNfsAccessNil sets the value for NfsAccess to be an explicit nil

### UnsetNfsAccess
`func (o *NisNetgroup) UnsetNfsAccess()`

UnsetNfsAccess ensures that no value is present for NfsAccess, not even an explicit nil
### GetNfsSquash

`func (o *NisNetgroup) GetNfsSquash() int32`

GetNfsSquash returns the NfsSquash field if non-nil, zero value otherwise.

### GetNfsSquashOk

`func (o *NisNetgroup) GetNfsSquashOk() (*int32, bool)`

GetNfsSquashOk returns a tuple with the NfsSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsSquash

`func (o *NisNetgroup) SetNfsSquash(v int32)`

SetNfsSquash sets NfsSquash field to given value.

### HasNfsSquash

`func (o *NisNetgroup) HasNfsSquash() bool`

HasNfsSquash returns a boolean if a field has been set.

### SetNfsSquashNil

`func (o *NisNetgroup) SetNfsSquashNil(b bool)`

 SetNfsSquashNil sets the value for NfsSquash to be an explicit nil

### UnsetNfsSquash
`func (o *NisNetgroup) UnsetNfsSquash()`

UnsetNfsSquash ensures that no value is present for NfsSquash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


