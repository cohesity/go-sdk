# RecoverHiveParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshots** | [**[]RecoverHiveSnapshotParams**](RecoverHiveSnapshotParams.md) | Specifies the local snapshot ids of the Objects to be recovered. | 
**Suffix** | Pointer to **NullableString** | A suffix that is to be applied to all recovered objects. | [optional] 

## Methods

### NewRecoverHiveParamsAllOf

`func NewRecoverHiveParamsAllOf(snapshots []RecoverHiveSnapshotParams, ) *RecoverHiveParamsAllOf`

NewRecoverHiveParamsAllOf instantiates a new RecoverHiveParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHiveParamsAllOfWithDefaults

`func NewRecoverHiveParamsAllOfWithDefaults() *RecoverHiveParamsAllOf`

NewRecoverHiveParamsAllOfWithDefaults instantiates a new RecoverHiveParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshots

`func (o *RecoverHiveParamsAllOf) GetSnapshots() []RecoverHiveSnapshotParams`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *RecoverHiveParamsAllOf) GetSnapshotsOk() (*[]RecoverHiveSnapshotParams, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *RecoverHiveParamsAllOf) SetSnapshots(v []RecoverHiveSnapshotParams)`

SetSnapshots sets Snapshots field to given value.


### SetSnapshotsNil

`func (o *RecoverHiveParamsAllOf) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *RecoverHiveParamsAllOf) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil
### GetSuffix

`func (o *RecoverHiveParamsAllOf) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *RecoverHiveParamsAllOf) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *RecoverHiveParamsAllOf) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *RecoverHiveParamsAllOf) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *RecoverHiveParamsAllOf) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *RecoverHiveParamsAllOf) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


