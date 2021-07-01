# OracleDataGuardInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **NullableString** | Specifies the role of the DataGuard database. Specifies the role of the DataGuard database.  A Data Guard configuration contains one production database, also referred to as the primary database, that functions in the primary role. The primary database can be either a single-instance Oracle database or an Oracle Real Application Clusters database.  A standby database is a transactionally consistent copy of the primary database. Similar to a primary database, a standby database can be either a single-instance Oracle database or an Oracle Real Application Clusters database. &#39;kPrimary&#39; indicates that the current database is primary database. &#39;kStandby&#39; indicates that the current database is standby database. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of standby database. Specifies the type of standby database. &#39;kPhysical&#39; indicates that the current database provides a physically identical copy of the primary database, with on disk structures identical to the primary database on a block-for-block basis. It is kept synchronized with the primary database, though Redo Apply, which recovers the redo data received from the primary database and applies the redo to the physical standby database. &#39;kLogical&#39; indicates that the current database provides the same logical information as the production database, although the physical structure can be different. It is kept synchronized with the primary database thorugh SQL Apply, which transforms the data in the redo received from the primary database into SQL statements and then executing the SQL statements on the standby database. &#39;kSnapshot&#39; indicates that the current database is a fully updateable standby created by converting a physical standby database into a snasphot standby database. It receives and archives but does not apply redo data from a primary database. | [optional] 

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

### SetRoleNil

`func (o *OracleDataGuardInfo) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *OracleDataGuardInfo) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetType

`func (o *OracleDataGuardInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OracleDataGuardInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OracleDataGuardInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OracleDataGuardInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *OracleDataGuardInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OracleDataGuardInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


