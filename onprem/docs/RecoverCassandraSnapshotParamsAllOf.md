# RecoverCassandraSnapshotParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]RecoverCassandraNoSqlObjectParams**](RecoverCassandraNoSqlObjectParams.md) | Specifies details of objects to be recovered. | [optional] 

## Methods

### NewRecoverCassandraSnapshotParamsAllOf

`func NewRecoverCassandraSnapshotParamsAllOf() *RecoverCassandraSnapshotParamsAllOf`

NewRecoverCassandraSnapshotParamsAllOf instantiates a new RecoverCassandraSnapshotParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverCassandraSnapshotParamsAllOfWithDefaults

`func NewRecoverCassandraSnapshotParamsAllOfWithDefaults() *RecoverCassandraSnapshotParamsAllOf`

NewRecoverCassandraSnapshotParamsAllOfWithDefaults instantiates a new RecoverCassandraSnapshotParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *RecoverCassandraSnapshotParamsAllOf) GetObjects() []RecoverCassandraNoSqlObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverCassandraSnapshotParamsAllOf) GetObjectsOk() (*[]RecoverCassandraNoSqlObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverCassandraSnapshotParamsAllOf) SetObjects(v []RecoverCassandraNoSqlObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *RecoverCassandraSnapshotParamsAllOf) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *RecoverCassandraSnapshotParamsAllOf) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverCassandraSnapshotParamsAllOf) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


