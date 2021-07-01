# SourceAppParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsVssCopyOnly** | Pointer to **NullableBool** | If the backup is a VSS full backup with the copy-only option specified. | [optional] 
**MsExchangeParams** | Pointer to [**MSExchangeParams**](MSExchangeParams.md) |  | [optional] 

## Methods

### NewSourceAppParams

`func NewSourceAppParams() *SourceAppParams`

NewSourceAppParams instantiates a new SourceAppParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAppParamsWithDefaults

`func NewSourceAppParamsWithDefaults() *SourceAppParams`

NewSourceAppParamsWithDefaults instantiates a new SourceAppParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsVssCopyOnly

`func (o *SourceAppParams) GetIsVssCopyOnly() bool`

GetIsVssCopyOnly returns the IsVssCopyOnly field if non-nil, zero value otherwise.

### GetIsVssCopyOnlyOk

`func (o *SourceAppParams) GetIsVssCopyOnlyOk() (*bool, bool)`

GetIsVssCopyOnlyOk returns a tuple with the IsVssCopyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVssCopyOnly

`func (o *SourceAppParams) SetIsVssCopyOnly(v bool)`

SetIsVssCopyOnly sets IsVssCopyOnly field to given value.

### HasIsVssCopyOnly

`func (o *SourceAppParams) HasIsVssCopyOnly() bool`

HasIsVssCopyOnly returns a boolean if a field has been set.

### SetIsVssCopyOnlyNil

`func (o *SourceAppParams) SetIsVssCopyOnlyNil(b bool)`

 SetIsVssCopyOnlyNil sets the value for IsVssCopyOnly to be an explicit nil

### UnsetIsVssCopyOnly
`func (o *SourceAppParams) UnsetIsVssCopyOnly()`

UnsetIsVssCopyOnly ensures that no value is present for IsVssCopyOnly, not even an explicit nil
### GetMsExchangeParams

`func (o *SourceAppParams) GetMsExchangeParams() MSExchangeParams`

GetMsExchangeParams returns the MsExchangeParams field if non-nil, zero value otherwise.

### GetMsExchangeParamsOk

`func (o *SourceAppParams) GetMsExchangeParamsOk() (*MSExchangeParams, bool)`

GetMsExchangeParamsOk returns a tuple with the MsExchangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsExchangeParams

`func (o *SourceAppParams) SetMsExchangeParams(v MSExchangeParams)`

SetMsExchangeParams sets MsExchangeParams field to given value.

### HasMsExchangeParams

`func (o *SourceAppParams) HasMsExchangeParams() bool`

HasMsExchangeParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


