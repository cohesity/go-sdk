# OracleDataGuardInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | Specifies the role of the Oracle DataGuard database. | [optional] 
**StandbyType** | Pointer to **string** | Specifies the type of the standby oracle database. | [optional] 

## Methods

### NewOracleDataGuardInfo

`func NewOracleDataGuardInfo() *OracleDataGuardInfo`

NewOracleDataGuardInfo instantiates a new OracleDataGuardInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleDataGuardInfoWithDefaults

`func NewOracleDataGuardInfoWithDefaults() *OracleDataGuardInfo`

NewOracleDataGuardInfoWithDefaults instantiates a new OracleDataGuardInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *OracleDataGuardInfo) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OracleDataGuardInfo) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OracleDataGuardInfo) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OracleDataGuardInfo) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStandbyType

`func (o *OracleDataGuardInfo) GetStandbyType() string`

GetStandbyType returns the StandbyType field if non-nil, zero value otherwise.

### GetStandbyTypeOk

`func (o *OracleDataGuardInfo) GetStandbyTypeOk() (*string, bool)`

GetStandbyTypeOk returns a tuple with the StandbyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyType

`func (o *OracleDataGuardInfo) SetStandbyType(v string)`

SetStandbyType sets StandbyType field to given value.

### HasStandbyType

`func (o *OracleDataGuardInfo) HasStandbyType() bool`

HasStandbyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


