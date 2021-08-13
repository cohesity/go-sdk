# ExchangeProtectionGroupObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the id of the registered Exchange DAG(Database Availability Group) source or Exchange physical source. | 
**Name** | Pointer to **NullableString** | Specifies the name of the registered Exchange DAG(Database Availability Group) source or Exchange physical source. | [optional] [readonly] 
**AppParams** | Pointer to [**[]ExchangeAppParams**](ExchangeAppParams.md) | Specifies the specific parameters required for Exchange app configuration. | [optional] 

## Methods

### NewExchangeProtectionGroupObjectParams

`func NewExchangeProtectionGroupObjectParams(id NullableInt64, ) *ExchangeProtectionGroupObjectParams`

NewExchangeProtectionGroupObjectParams instantiates a new ExchangeProtectionGroupObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeProtectionGroupObjectParamsWithDefaults

`func NewExchangeProtectionGroupObjectParamsWithDefaults() *ExchangeProtectionGroupObjectParams`

NewExchangeProtectionGroupObjectParamsWithDefaults instantiates a new ExchangeProtectionGroupObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExchangeProtectionGroupObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExchangeProtectionGroupObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExchangeProtectionGroupObjectParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ExchangeProtectionGroupObjectParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExchangeProtectionGroupObjectParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ExchangeProtectionGroupObjectParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExchangeProtectionGroupObjectParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExchangeProtectionGroupObjectParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExchangeProtectionGroupObjectParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExchangeProtectionGroupObjectParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExchangeProtectionGroupObjectParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAppParams

`func (o *ExchangeProtectionGroupObjectParams) GetAppParams() []ExchangeAppParams`

GetAppParams returns the AppParams field if non-nil, zero value otherwise.

### GetAppParamsOk

`func (o *ExchangeProtectionGroupObjectParams) GetAppParamsOk() (*[]ExchangeAppParams, bool)`

GetAppParamsOk returns a tuple with the AppParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppParams

`func (o *ExchangeProtectionGroupObjectParams) SetAppParams(v []ExchangeAppParams)`

SetAppParams sets AppParams field to given value.

### HasAppParams

`func (o *ExchangeProtectionGroupObjectParams) HasAppParams() bool`

HasAppParams returns a boolean if a field has been set.

### SetAppParamsNil

`func (o *ExchangeProtectionGroupObjectParams) SetAppParamsNil(b bool)`

 SetAppParamsNil sets the value for AppParams to be an explicit nil

### UnsetAppParams
`func (o *ExchangeProtectionGroupObjectParams) UnsetAppParams()`

UnsetAppParams ensures that no value is present for AppParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


