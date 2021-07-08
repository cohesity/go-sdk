# OracleSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **NullableString** | Location is the path where Oracle is installed. | [optional] 
**SystemIdentifier** | Pointer to **NullableString** | SystemIdentifier is the unique Oracle System Identifier for the DB instance. | [optional] 

## Methods

### NewOracleSession

`func NewOracleSession() *OracleSession`

NewOracleSession instantiates a new OracleSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleSessionWithDefaults

`func NewOracleSessionWithDefaults() *OracleSession`

NewOracleSessionWithDefaults instantiates a new OracleSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *OracleSession) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OracleSession) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OracleSession) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OracleSession) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *OracleSession) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *OracleSession) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetSystemIdentifier

`func (o *OracleSession) GetSystemIdentifier() string`

GetSystemIdentifier returns the SystemIdentifier field if non-nil, zero value otherwise.

### GetSystemIdentifierOk

`func (o *OracleSession) GetSystemIdentifierOk() (*string, bool)`

GetSystemIdentifierOk returns a tuple with the SystemIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentifier

`func (o *OracleSession) SetSystemIdentifier(v string)`

SetSystemIdentifier sets SystemIdentifier field to given value.

### HasSystemIdentifier

`func (o *OracleSession) HasSystemIdentifier() bool`

HasSystemIdentifier returns a boolean if a field has been set.

### SetSystemIdentifierNil

`func (o *OracleSession) SetSystemIdentifierNil(b bool)`

 SetSystemIdentifierNil sets the value for SystemIdentifier to be an explicit nil

### UnsetSystemIdentifier
`func (o *OracleSession) UnsetSystemIdentifier()`

UnsetSystemIdentifier ensures that no value is present for SystemIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


