# OraclePluggableDatabaseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseId** | Pointer to **NullableString** | Specifies the ID of the Pluggable Database. | [optional] 
**Name** | Pointer to **NullableString** | Speicifes the name of the Pluggable Database. | [optional] 
**OpenMode** | Pointer to **NullableString** | Specifies the OPEN_MODE information. Specifies the OPEN_MODE type for the Oracle database. &#39;kMounted&#39; indicates that the database is open in Mounted mode. &#39;kReadWrite&#39; indicates that the database is open in R/W mode. &#39;kReadOnly&#39; indicates that the database is open in ReadOnly mode. &#39;kMigrate&#39; inidcates that the database is open in Migrate mode. | [optional] 
**SizeBytes** | Pointer to **NullableInt64** | Specifies the Size in Bytes of the Pluggable Database. | [optional] 

## Methods

### NewOraclePluggableDatabaseInfo

`func NewOraclePluggableDatabaseInfo() *OraclePluggableDatabaseInfo`

NewOraclePluggableDatabaseInfo instantiates a new OraclePluggableDatabaseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOraclePluggableDatabaseInfoWithDefaults

`func NewOraclePluggableDatabaseInfoWithDefaults() *OraclePluggableDatabaseInfo`

NewOraclePluggableDatabaseInfoWithDefaults instantiates a new OraclePluggableDatabaseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseId

`func (o *OraclePluggableDatabaseInfo) GetDatabaseId() string`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *OraclePluggableDatabaseInfo) GetDatabaseIdOk() (*string, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *OraclePluggableDatabaseInfo) SetDatabaseId(v string)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *OraclePluggableDatabaseInfo) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### SetDatabaseIdNil

`func (o *OraclePluggableDatabaseInfo) SetDatabaseIdNil(b bool)`

 SetDatabaseIdNil sets the value for DatabaseId to be an explicit nil

### UnsetDatabaseId
`func (o *OraclePluggableDatabaseInfo) UnsetDatabaseId()`

UnsetDatabaseId ensures that no value is present for DatabaseId, not even an explicit nil
### GetName

`func (o *OraclePluggableDatabaseInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OraclePluggableDatabaseInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OraclePluggableDatabaseInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OraclePluggableDatabaseInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OraclePluggableDatabaseInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OraclePluggableDatabaseInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOpenMode

`func (o *OraclePluggableDatabaseInfo) GetOpenMode() string`

GetOpenMode returns the OpenMode field if non-nil, zero value otherwise.

### GetOpenModeOk

`func (o *OraclePluggableDatabaseInfo) GetOpenModeOk() (*string, bool)`

GetOpenModeOk returns a tuple with the OpenMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenMode

`func (o *OraclePluggableDatabaseInfo) SetOpenMode(v string)`

SetOpenMode sets OpenMode field to given value.

### HasOpenMode

`func (o *OraclePluggableDatabaseInfo) HasOpenMode() bool`

HasOpenMode returns a boolean if a field has been set.

### SetOpenModeNil

`func (o *OraclePluggableDatabaseInfo) SetOpenModeNil(b bool)`

 SetOpenModeNil sets the value for OpenMode to be an explicit nil

### UnsetOpenMode
`func (o *OraclePluggableDatabaseInfo) UnsetOpenMode()`

UnsetOpenMode ensures that no value is present for OpenMode, not even an explicit nil
### GetSizeBytes

`func (o *OraclePluggableDatabaseInfo) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *OraclePluggableDatabaseInfo) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *OraclePluggableDatabaseInfo) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *OraclePluggableDatabaseInfo) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### SetSizeBytesNil

`func (o *OraclePluggableDatabaseInfo) SetSizeBytesNil(b bool)`

 SetSizeBytesNil sets the value for SizeBytes to be an explicit nil

### UnsetSizeBytes
`func (o *OraclePluggableDatabaseInfo) UnsetSizeBytes()`

UnsetSizeBytes ensures that no value is present for SizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


