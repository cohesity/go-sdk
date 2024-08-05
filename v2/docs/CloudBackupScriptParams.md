# CloudBackupScriptParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostBackupScript** | Pointer to [**CommonPrePostCloudScriptParams**](CommonPrePostCloudScriptParams.md) |  | [optional] 
**PostSnapshotScript** | Pointer to [**CommonPrePostCloudScriptParams**](CommonPrePostCloudScriptParams.md) |  | [optional] 
**PreBackupScript** | Pointer to [**CommonPrePostCloudScriptParams**](CommonPrePostCloudScriptParams.md) |  | [optional] 

## Methods

### NewCloudBackupScriptParams

`func NewCloudBackupScriptParams() *CloudBackupScriptParams`

NewCloudBackupScriptParams instantiates a new CloudBackupScriptParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBackupScriptParamsWithDefaults

`func NewCloudBackupScriptParamsWithDefaults() *CloudBackupScriptParams`

NewCloudBackupScriptParamsWithDefaults instantiates a new CloudBackupScriptParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostBackupScript

`func (o *CloudBackupScriptParams) GetPostBackupScript() CommonPrePostCloudScriptParams`

GetPostBackupScript returns the PostBackupScript field if non-nil, zero value otherwise.

### GetPostBackupScriptOk

`func (o *CloudBackupScriptParams) GetPostBackupScriptOk() (*CommonPrePostCloudScriptParams, bool)`

GetPostBackupScriptOk returns a tuple with the PostBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostBackupScript

`func (o *CloudBackupScriptParams) SetPostBackupScript(v CommonPrePostCloudScriptParams)`

SetPostBackupScript sets PostBackupScript field to given value.

### HasPostBackupScript

`func (o *CloudBackupScriptParams) HasPostBackupScript() bool`

HasPostBackupScript returns a boolean if a field has been set.

### GetPostSnapshotScript

`func (o *CloudBackupScriptParams) GetPostSnapshotScript() CommonPrePostCloudScriptParams`

GetPostSnapshotScript returns the PostSnapshotScript field if non-nil, zero value otherwise.

### GetPostSnapshotScriptOk

`func (o *CloudBackupScriptParams) GetPostSnapshotScriptOk() (*CommonPrePostCloudScriptParams, bool)`

GetPostSnapshotScriptOk returns a tuple with the PostSnapshotScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostSnapshotScript

`func (o *CloudBackupScriptParams) SetPostSnapshotScript(v CommonPrePostCloudScriptParams)`

SetPostSnapshotScript sets PostSnapshotScript field to given value.

### HasPostSnapshotScript

`func (o *CloudBackupScriptParams) HasPostSnapshotScript() bool`

HasPostSnapshotScript returns a boolean if a field has been set.

### GetPreBackupScript

`func (o *CloudBackupScriptParams) GetPreBackupScript() CommonPrePostCloudScriptParams`

GetPreBackupScript returns the PreBackupScript field if non-nil, zero value otherwise.

### GetPreBackupScriptOk

`func (o *CloudBackupScriptParams) GetPreBackupScriptOk() (*CommonPrePostCloudScriptParams, bool)`

GetPreBackupScriptOk returns a tuple with the PreBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreBackupScript

`func (o *CloudBackupScriptParams) SetPreBackupScript(v CommonPrePostCloudScriptParams)`

SetPreBackupScript sets PreBackupScript field to given value.

### HasPreBackupScript

`func (o *CloudBackupScriptParams) HasPreBackupScript() bool`

HasPreBackupScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


