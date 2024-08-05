# MigrateS3Views

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3MigrationAction** | **string** | Specifies the target S3 migration state for the Views specified in the viewIds parameter. Supported Migration States are: [Enable, Cancel, Pause, Resume]. | 
**ViewIds** | **[]int32** | Specifies the list Views IDs on which the migration action will be performed. | 

## Methods

### NewMigrateS3Views

`func NewMigrateS3Views(s3MigrationAction string, viewIds []int32, ) *MigrateS3Views`

NewMigrateS3Views instantiates a new MigrateS3Views object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateS3ViewsWithDefaults

`func NewMigrateS3ViewsWithDefaults() *MigrateS3Views`

NewMigrateS3ViewsWithDefaults instantiates a new MigrateS3Views object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3MigrationAction

`func (o *MigrateS3Views) GetS3MigrationAction() string`

GetS3MigrationAction returns the S3MigrationAction field if non-nil, zero value otherwise.

### GetS3MigrationActionOk

`func (o *MigrateS3Views) GetS3MigrationActionOk() (*string, bool)`

GetS3MigrationActionOk returns a tuple with the S3MigrationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3MigrationAction

`func (o *MigrateS3Views) SetS3MigrationAction(v string)`

SetS3MigrationAction sets S3MigrationAction field to given value.


### GetViewIds

`func (o *MigrateS3Views) GetViewIds() []int32`

GetViewIds returns the ViewIds field if non-nil, zero value otherwise.

### GetViewIdsOk

`func (o *MigrateS3Views) GetViewIdsOk() (*[]int32, bool)`

GetViewIdsOk returns a tuple with the ViewIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewIds

`func (o *MigrateS3Views) SetViewIds(v []int32)`

SetViewIds sets ViewIds field to given value.


### SetViewIdsNil

`func (o *MigrateS3Views) SetViewIdsNil(b bool)`

 SetViewIdsNil sets the value for ViewIds to be an explicit nil

### UnsetViewIds
`func (o *MigrateS3Views) UnsetViewIds()`

UnsetViewIds ensures that no value is present for ViewIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


