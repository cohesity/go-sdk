# OraclePdbObjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbId** | **NullableString** | Specifies pluggable database id. | 
**DbName** | **NullableString** | Specifies name of the DB. | 

## Methods

### NewOraclePdbObjectInfo

`func NewOraclePdbObjectInfo(dbId NullableString, dbName NullableString, ) *OraclePdbObjectInfo`

NewOraclePdbObjectInfo instantiates a new OraclePdbObjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOraclePdbObjectInfoWithDefaults

`func NewOraclePdbObjectInfoWithDefaults() *OraclePdbObjectInfo`

NewOraclePdbObjectInfoWithDefaults instantiates a new OraclePdbObjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbId

`func (o *OraclePdbObjectInfo) GetDbId() string`

GetDbId returns the DbId field if non-nil, zero value otherwise.

### GetDbIdOk

`func (o *OraclePdbObjectInfo) GetDbIdOk() (*string, bool)`

GetDbIdOk returns a tuple with the DbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbId

`func (o *OraclePdbObjectInfo) SetDbId(v string)`

SetDbId sets DbId field to given value.


### SetDbIdNil

`func (o *OraclePdbObjectInfo) SetDbIdNil(b bool)`

 SetDbIdNil sets the value for DbId to be an explicit nil

### UnsetDbId
`func (o *OraclePdbObjectInfo) UnsetDbId()`

UnsetDbId ensures that no value is present for DbId, not even an explicit nil
### GetDbName

`func (o *OraclePdbObjectInfo) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *OraclePdbObjectInfo) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *OraclePdbObjectInfo) SetDbName(v string)`

SetDbName sets DbName field to given value.


### SetDbNameNil

`func (o *OraclePdbObjectInfo) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *OraclePdbObjectInfo) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


