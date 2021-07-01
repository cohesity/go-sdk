# AdProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainController** | Pointer to [**AdDomainController**](AdDomainController.md) |  | [optional] 
**DomainName** | Pointer to **NullableString** | Specifies the domain name corresponding to the domain controller. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the domain name of the AD entity. | [optional] 
**OwnerId** | Pointer to **NullableInt64** | Specifies the entity id of the owner entity. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the managed object in AD Protection Source. Specifies the kind of AD protection source. &#39;kRootContainer&#39; indicates the entity is a root container to an AD domain controller. &#39;kDomainController&#39; indicates the domain controller hosted in this physical server. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID for the AD entity. | [optional] 

## Methods

### NewAdProtectionSource

`func NewAdProtectionSource() *AdProtectionSource`

NewAdProtectionSource instantiates a new AdProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdProtectionSourceWithDefaults

`func NewAdProtectionSourceWithDefaults() *AdProtectionSource`

NewAdProtectionSourceWithDefaults instantiates a new AdProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainController

`func (o *AdProtectionSource) GetDomainController() AdDomainController`

GetDomainController returns the DomainController field if non-nil, zero value otherwise.

### GetDomainControllerOk

`func (o *AdProtectionSource) GetDomainControllerOk() (*AdDomainController, bool)`

GetDomainControllerOk returns a tuple with the DomainController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainController

`func (o *AdProtectionSource) SetDomainController(v AdDomainController)`

SetDomainController sets DomainController field to given value.

### HasDomainController

`func (o *AdProtectionSource) HasDomainController() bool`

HasDomainController returns a boolean if a field has been set.

### GetDomainName

`func (o *AdProtectionSource) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *AdProtectionSource) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *AdProtectionSource) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *AdProtectionSource) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *AdProtectionSource) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *AdProtectionSource) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetName

`func (o *AdProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnerId

`func (o *AdProtectionSource) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AdProtectionSource) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AdProtectionSource) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AdProtectionSource) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *AdProtectionSource) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *AdProtectionSource) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetType

`func (o *AdProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AdProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AdProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUuid

`func (o *AdProtectionSource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AdProtectionSource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AdProtectionSource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AdProtectionSource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *AdProtectionSource) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *AdProtectionSource) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


