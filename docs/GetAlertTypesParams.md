# GetAlertTypesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertCategoryList** | Pointer to **[]string** | Specifies a list of Alert Categories to filter alert types by. Specifies the category of an Alert. kDisk - Alerts that are related to Disk. kNode - Alerts that are related to Node. kCluster - Alerts that are related to Cluster. kNodeHealth - Alerts that are related to Node Health. kClusterHealth - Alerts that are related to Cluster Health. kBackupRestore - Alerts that are related to Backup/Restore. kEncryption - Alerts that are related to Encryption. kArchivalRestore - Alerts that are related to Archival/Restore. kRemoteReplication - Alerts that are related to Remote Replication. kQuota - Alerts that are related to Quota. kLicense - Alerts that are related to License. kHeliosProActiveWellness - Alerts that are related to Helios ProActive Wellness. kHeliosAnalyticsJobs - Alerts that are related to Helios Analytics Jobs. kHeliosSignatureJobs - Alerts that are related to Helios Signature Jobs. kSecurity - Alerts that are related to Security. kAppsInfra - Alerts that are related to applications infra. kAntivirus - Alerts that are related to antivirus. kArchivalCopy - Alerts that are related to archival copies. | [optional] 

## Methods

### NewGetAlertTypesParams

`func NewGetAlertTypesParams() *GetAlertTypesParams`

NewGetAlertTypesParams instantiates a new GetAlertTypesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlertTypesParamsWithDefaults

`func NewGetAlertTypesParamsWithDefaults() *GetAlertTypesParams`

NewGetAlertTypesParamsWithDefaults instantiates a new GetAlertTypesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertCategoryList

`func (o *GetAlertTypesParams) GetAlertCategoryList() []string`

GetAlertCategoryList returns the AlertCategoryList field if non-nil, zero value otherwise.

### GetAlertCategoryListOk

`func (o *GetAlertTypesParams) GetAlertCategoryListOk() (*[]string, bool)`

GetAlertCategoryListOk returns a tuple with the AlertCategoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCategoryList

`func (o *GetAlertTypesParams) SetAlertCategoryList(v []string)`

SetAlertCategoryList sets AlertCategoryList field to given value.

### HasAlertCategoryList

`func (o *GetAlertTypesParams) HasAlertCategoryList() bool`

HasAlertCategoryList returns a boolean if a field has been set.

### SetAlertCategoryListNil

`func (o *GetAlertTypesParams) SetAlertCategoryListNil(b bool)`

 SetAlertCategoryListNil sets the value for AlertCategoryList to be an explicit nil

### UnsetAlertCategoryList
`func (o *GetAlertTypesParams) UnsetAlertCategoryList()`

UnsetAlertCategoryList ensures that no value is present for AlertCategoryList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


