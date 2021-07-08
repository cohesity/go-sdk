# MongoDBAdditionalParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecondaryNodeTag** | Pointer to **[]string** | The tag associated with the secondary nodes from which backups should be performed. | [optional] 
**UseSecondaryForBackup** | Pointer to **NullableBool** | Set to true if this cluster uses secondary nodes for backup. | [optional] 

## Methods

### NewMongoDBAdditionalParams

`func NewMongoDBAdditionalParams() *MongoDBAdditionalParams`

NewMongoDBAdditionalParams instantiates a new MongoDBAdditionalParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBAdditionalParamsWithDefaults

`func NewMongoDBAdditionalParamsWithDefaults() *MongoDBAdditionalParams`

NewMongoDBAdditionalParamsWithDefaults instantiates a new MongoDBAdditionalParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecondaryNodeTag

`func (o *MongoDBAdditionalParams) GetSecondaryNodeTag() []string`

GetSecondaryNodeTag returns the SecondaryNodeTag field if non-nil, zero value otherwise.

### GetSecondaryNodeTagOk

`func (o *MongoDBAdditionalParams) GetSecondaryNodeTagOk() (*[]string, bool)`

GetSecondaryNodeTagOk returns a tuple with the SecondaryNodeTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryNodeTag

`func (o *MongoDBAdditionalParams) SetSecondaryNodeTag(v []string)`

SetSecondaryNodeTag sets SecondaryNodeTag field to given value.

### HasSecondaryNodeTag

`func (o *MongoDBAdditionalParams) HasSecondaryNodeTag() bool`

HasSecondaryNodeTag returns a boolean if a field has been set.

### SetSecondaryNodeTagNil

`func (o *MongoDBAdditionalParams) SetSecondaryNodeTagNil(b bool)`

 SetSecondaryNodeTagNil sets the value for SecondaryNodeTag to be an explicit nil

### UnsetSecondaryNodeTag
`func (o *MongoDBAdditionalParams) UnsetSecondaryNodeTag()`

UnsetSecondaryNodeTag ensures that no value is present for SecondaryNodeTag, not even an explicit nil
### GetUseSecondaryForBackup

`func (o *MongoDBAdditionalParams) GetUseSecondaryForBackup() bool`

GetUseSecondaryForBackup returns the UseSecondaryForBackup field if non-nil, zero value otherwise.

### GetUseSecondaryForBackupOk

`func (o *MongoDBAdditionalParams) GetUseSecondaryForBackupOk() (*bool, bool)`

GetUseSecondaryForBackupOk returns a tuple with the UseSecondaryForBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSecondaryForBackup

`func (o *MongoDBAdditionalParams) SetUseSecondaryForBackup(v bool)`

SetUseSecondaryForBackup sets UseSecondaryForBackup field to given value.

### HasUseSecondaryForBackup

`func (o *MongoDBAdditionalParams) HasUseSecondaryForBackup() bool`

HasUseSecondaryForBackup returns a boolean if a field has been set.

### SetUseSecondaryForBackupNil

`func (o *MongoDBAdditionalParams) SetUseSecondaryForBackupNil(b bool)`

 SetUseSecondaryForBackupNil sets the value for UseSecondaryForBackup to be an explicit nil

### UnsetUseSecondaryForBackup
`func (o *MongoDBAdditionalParams) UnsetUseSecondaryForBackup()`

UnsetUseSecondaryForBackup ensures that no value is present for UseSecondaryForBackup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


