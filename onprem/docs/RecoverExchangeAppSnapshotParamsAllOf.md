# RecoverExchangeAppSnapshotParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppObjects** | Pointer to [**[]ExchangeRecoverDatabaseParams**](ExchangeRecoverDatabaseParams.md) | Specifies the list of params for all the databases which have to be recovered. | [optional] 

## Methods

### NewRecoverExchangeAppSnapshotParamsAllOf

`func NewRecoverExchangeAppSnapshotParamsAllOf() *RecoverExchangeAppSnapshotParamsAllOf`

NewRecoverExchangeAppSnapshotParamsAllOf instantiates a new RecoverExchangeAppSnapshotParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverExchangeAppSnapshotParamsAllOfWithDefaults

`func NewRecoverExchangeAppSnapshotParamsAllOfWithDefaults() *RecoverExchangeAppSnapshotParamsAllOf`

NewRecoverExchangeAppSnapshotParamsAllOfWithDefaults instantiates a new RecoverExchangeAppSnapshotParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppObjects

`func (o *RecoverExchangeAppSnapshotParamsAllOf) GetAppObjects() []ExchangeRecoverDatabaseParams`

GetAppObjects returns the AppObjects field if non-nil, zero value otherwise.

### GetAppObjectsOk

`func (o *RecoverExchangeAppSnapshotParamsAllOf) GetAppObjectsOk() (*[]ExchangeRecoverDatabaseParams, bool)`

GetAppObjectsOk returns a tuple with the AppObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppObjects

`func (o *RecoverExchangeAppSnapshotParamsAllOf) SetAppObjects(v []ExchangeRecoverDatabaseParams)`

SetAppObjects sets AppObjects field to given value.

### HasAppObjects

`func (o *RecoverExchangeAppSnapshotParamsAllOf) HasAppObjects() bool`

HasAppObjects returns a boolean if a field has been set.

### SetAppObjectsNil

`func (o *RecoverExchangeAppSnapshotParamsAllOf) SetAppObjectsNil(b bool)`

 SetAppObjectsNil sets the value for AppObjects to be an explicit nil

### UnsetAppObjects
`func (o *RecoverExchangeAppSnapshotParamsAllOf) UnsetAppObjects()`

UnsetAppObjects ensures that no value is present for AppObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


