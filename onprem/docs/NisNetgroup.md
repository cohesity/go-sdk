# NisNetgroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the netgroup name. | 
**Domain** | **NullableString** | Specifies the domain name for the netgroup. | 
**NfsAccess** | Pointer to **NullableString** | Specifies NFS protocol acess level for clients from the netgroup. | [optional] 
**NfsSquash** | Pointer to **NullableString** | Specifies which nfsSquash Mounted. &#39;kNone&#39; mounts none. &#39;kRootSquash&#39; mounts nfsRootSquash. Whether clients from this subnet can mount as root on NFS. &#39;kAllSquash&#39; mounts nfsAllSquash. Whether all clients from this subnet can map view with view_all_squash_uid/view_all_squash_gid configured in the view. | [optional] 

## Methods

### NewNisNetgroup

`func NewNisNetgroup(name NullableString, domain NullableString, ) *NisNetgroup`

NewNisNetgroup instantiates a new NisNetgroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNisNetgroupWithDefaults

`func NewNisNetgroupWithDefaults() *NisNetgroup`

NewNisNetgroupWithDefaults instantiates a new NisNetgroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### SetNameNil

`func (o *NisNetgroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NisNetgroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
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


### SetDomainNil

`func (o *NisNetgroup) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *NisNetgroup) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
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

`func (o *NisNetgroup) GetNfsSquash() string`

GetNfsSquash returns the NfsSquash field if non-nil, zero value otherwise.

### GetNfsSquashOk

`func (o *NisNetgroup) GetNfsSquashOk() (*string, bool)`

GetNfsSquashOk returns a tuple with the NfsSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsSquash

`func (o *NisNetgroup) SetNfsSquash(v string)`

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


