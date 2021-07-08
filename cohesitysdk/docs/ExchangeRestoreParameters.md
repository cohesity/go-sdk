# ExchangeRestoreParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | Specifies the Exchange restore type. Specifies the type of Exchange restore.  &#39;kNone&#39; specifies no special behaviour. &#39;kView&#39; specifies the option to create a view which cann be used by the external tools like Kroll to perform mailbox or mail-item recovery. &#39;kDatabase&#39; specifies the option to restore an Exchange database. | [optional] 
**ViewParameters** | Pointer to [**ExchangeRestoreViewParameters**](ExchangeRestoreViewParameters.md) |  | [optional] 

## Methods

### NewExchangeRestoreParameters

`func NewExchangeRestoreParameters() *ExchangeRestoreParameters`

NewExchangeRestoreParameters instantiates a new ExchangeRestoreParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRestoreParametersWithDefaults

`func NewExchangeRestoreParametersWithDefaults() *ExchangeRestoreParameters`

NewExchangeRestoreParametersWithDefaults instantiates a new ExchangeRestoreParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExchangeRestoreParameters) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExchangeRestoreParameters) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExchangeRestoreParameters) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExchangeRestoreParameters) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ExchangeRestoreParameters) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExchangeRestoreParameters) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetViewParameters

`func (o *ExchangeRestoreParameters) GetViewParameters() ExchangeRestoreViewParameters`

GetViewParameters returns the ViewParameters field if non-nil, zero value otherwise.

### GetViewParametersOk

`func (o *ExchangeRestoreParameters) GetViewParametersOk() (*ExchangeRestoreViewParameters, bool)`

GetViewParametersOk returns a tuple with the ViewParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewParameters

`func (o *ExchangeRestoreParameters) SetViewParameters(v ExchangeRestoreViewParameters)`

SetViewParameters sets ViewParameters field to given value.

### HasViewParameters

`func (o *ExchangeRestoreParameters) HasViewParameters() bool`

HasViewParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


