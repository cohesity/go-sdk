# ExchangeBackupJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCopyOnlyFull** | Pointer to **NullableBool** | Whether the backups should be copy-only. If this is set to true, then Exchange server will not truncate logs after backup. | [optional] 

## Methods

### NewExchangeBackupJobParams

`func NewExchangeBackupJobParams() *ExchangeBackupJobParams`

NewExchangeBackupJobParams instantiates a new ExchangeBackupJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeBackupJobParamsWithDefaults

`func NewExchangeBackupJobParamsWithDefaults() *ExchangeBackupJobParams`

NewExchangeBackupJobParamsWithDefaults instantiates a new ExchangeBackupJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCopyOnlyFull

`func (o *ExchangeBackupJobParams) GetIsCopyOnlyFull() bool`

GetIsCopyOnlyFull returns the IsCopyOnlyFull field if non-nil, zero value otherwise.

### GetIsCopyOnlyFullOk

`func (o *ExchangeBackupJobParams) GetIsCopyOnlyFullOk() (*bool, bool)`

GetIsCopyOnlyFullOk returns a tuple with the IsCopyOnlyFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCopyOnlyFull

`func (o *ExchangeBackupJobParams) SetIsCopyOnlyFull(v bool)`

SetIsCopyOnlyFull sets IsCopyOnlyFull field to given value.

### HasIsCopyOnlyFull

`func (o *ExchangeBackupJobParams) HasIsCopyOnlyFull() bool`

HasIsCopyOnlyFull returns a boolean if a field has been set.

### SetIsCopyOnlyFullNil

`func (o *ExchangeBackupJobParams) SetIsCopyOnlyFullNil(b bool)`

 SetIsCopyOnlyFullNil sets the value for IsCopyOnlyFull to be an explicit nil

### UnsetIsCopyOnlyFull
`func (o *ExchangeBackupJobParams) UnsetIsCopyOnlyFull()`

UnsetIsCopyOnlyFull ensures that no value is present for IsCopyOnlyFull, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


