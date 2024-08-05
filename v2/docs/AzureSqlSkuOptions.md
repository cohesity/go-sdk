# AzureSqlSkuOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **NullableInt32** | Specifies the capacity of the sku. For azure sql dbs, this is the number of cores. Default value is 4. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the sku name for azure sql databases and by default Hyperscale is selected. | [optional] 
**TierType** | Pointer to **NullableString** | Specifies the sku tier selection for azure sql databases and by default HS_Gen5 is selected. The tiers must match their sku name selected. | [optional] 

## Methods

### NewAzureSqlSkuOptions

`func NewAzureSqlSkuOptions() *AzureSqlSkuOptions`

NewAzureSqlSkuOptions instantiates a new AzureSqlSkuOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSqlSkuOptionsWithDefaults

`func NewAzureSqlSkuOptionsWithDefaults() *AzureSqlSkuOptions`

NewAzureSqlSkuOptionsWithDefaults instantiates a new AzureSqlSkuOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *AzureSqlSkuOptions) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *AzureSqlSkuOptions) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *AzureSqlSkuOptions) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *AzureSqlSkuOptions) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### SetCapacityNil

`func (o *AzureSqlSkuOptions) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *AzureSqlSkuOptions) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetName

`func (o *AzureSqlSkuOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureSqlSkuOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureSqlSkuOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureSqlSkuOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AzureSqlSkuOptions) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureSqlSkuOptions) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTierType

`func (o *AzureSqlSkuOptions) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *AzureSqlSkuOptions) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *AzureSqlSkuOptions) SetTierType(v string)`

SetTierType sets TierType field to given value.

### HasTierType

`func (o *AzureSqlSkuOptions) HasTierType() bool`

HasTierType returns a boolean if a field has been set.

### SetTierTypeNil

`func (o *AzureSqlSkuOptions) SetTierTypeNil(b bool)`

 SetTierTypeNil sets the value for TierType to be an explicit nil

### UnsetTierType
`func (o *AzureSqlSkuOptions) UnsetTierType()`

UnsetTierType ensures that no value is present for TierType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


