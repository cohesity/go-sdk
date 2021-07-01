# VaultStatsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the Vault Id. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the Vault name. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the Vault type. | [optional] 
**UsageBytes** | Pointer to **NullableInt64** | Specifies the bytes used by the Vault. | [optional] 

## Methods

### NewVaultStatsInfo

`func NewVaultStatsInfo() *VaultStatsInfo`

NewVaultStatsInfo instantiates a new VaultStatsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultStatsInfoWithDefaults

`func NewVaultStatsInfoWithDefaults() *VaultStatsInfo`

NewVaultStatsInfoWithDefaults instantiates a new VaultStatsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VaultStatsInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VaultStatsInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VaultStatsInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VaultStatsInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VaultStatsInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VaultStatsInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *VaultStatsInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VaultStatsInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VaultStatsInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VaultStatsInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VaultStatsInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VaultStatsInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *VaultStatsInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VaultStatsInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VaultStatsInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VaultStatsInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *VaultStatsInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *VaultStatsInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUsageBytes

`func (o *VaultStatsInfo) GetUsageBytes() int64`

GetUsageBytes returns the UsageBytes field if non-nil, zero value otherwise.

### GetUsageBytesOk

`func (o *VaultStatsInfo) GetUsageBytesOk() (*int64, bool)`

GetUsageBytesOk returns a tuple with the UsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageBytes

`func (o *VaultStatsInfo) SetUsageBytes(v int64)`

SetUsageBytes sets UsageBytes field to given value.

### HasUsageBytes

`func (o *VaultStatsInfo) HasUsageBytes() bool`

HasUsageBytes returns a boolean if a field has been set.

### SetUsageBytesNil

`func (o *VaultStatsInfo) SetUsageBytesNil(b bool)`

 SetUsageBytesNil sets the value for UsageBytes to be an explicit nil

### UnsetUsageBytes
`func (o *VaultStatsInfo) UnsetUsageBytes()`

UnsetUsageBytes ensures that no value is present for UsageBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


