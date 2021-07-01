# RestoreSourceSummaryByObjectTypeElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatastoreId** | Pointer to **NullableInt64** | Specifies the datastore where the object&#39;s files are recovered to. This field is populated when objects are recovered to a different resource pool or to a different parent source. This field is not populated when objects are recovered to their original datastore locations in the original parent source. | [optional] 
**FileRestoreInfo** | Pointer to [**[]FileRestoreInfo**](FileRestoreInfo.md) | Specifies a list of restore information of files. | [optional] 
**Name** | **NullableString** | Specifies the name of the Restore Task. This field must be set and must be a unique name. | 
**Objects** | Pointer to [**[]RestoreObjectDetails**](RestoreObjectDetails.md) | Array of Objects.  Specifies a list of Protection Source objects or Protection Job objects (with specified Protection Source objects). | [optional] 
**ProtectionSourceName** | Pointer to **NullableString** | The protection source name. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the start time of the Restore Task as a Unix epoch Timestamp (in microseconds). | [optional] 
**Type** | Pointer to **NullableString** | Specify the object type to filter with. Specifies the type of Restore Task.  &#39;kRecoverVMs&#39; specifies a Restore Task that recovers VMs. &#39;kCloneVMs&#39; specifies a Restore Task that clones VMs. &#39;kCloneView&#39; specifies a Restore Task that clones a View. &#39;kMountVolumes&#39; specifies a Restore Task that mounts volumes. &#39;kRestoreFiles&#39; specifies a Restore Task that recovers files and folders. &#39;kRecoverApp&#39; specifies a Restore Task that recovers app. &#39;kCloneApp&#39; specifies a Restore Task that clone app. &#39;kRecoverSanVolume&#39; specifies a Restore Task that recovers SAN volumes. &#39;kConvertAndDeployVMs&#39; specifies a Restore Task that converts and deploy VMs to a target environment. &#39;kMountFileVolume&#39; specifies a Restore Task that mounts a file volume. &#39;kSystem&#39; specifies a Restore Task that recovers a system. &#39;kRecoverVolumes&#39; specifies a Restore Task that recovers volumes via the physical agent. &#39;kDeployVolumes&#39; specifies a Restore Task that deploys volumes to a target environment. &#39;kDownloadFiles&#39; specifies a Restore Task that downloads the requested files and folders in zip format. &#39;kRecoverEmails&#39; specifies a Restore Task that recovers the mailbox/email items. &#39;kRecoverDisks&#39; specifies a Restore Task that recovers the virtual disks. &#39;kRecoverNamespaces&#39; specifies a Restore Task that recovers Kubernetes namespaces. &#39;kCloneVMsToView&#39; specifies a Restore Task that clones VMs into a View. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the Cohesity user who requested this Restore Task. | [optional] 

## Methods

### NewRestoreSourceSummaryByObjectTypeElement

`func NewRestoreSourceSummaryByObjectTypeElement(name NullableString, ) *RestoreSourceSummaryByObjectTypeElement`

NewRestoreSourceSummaryByObjectTypeElement instantiates a new RestoreSourceSummaryByObjectTypeElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSourceSummaryByObjectTypeElementWithDefaults

`func NewRestoreSourceSummaryByObjectTypeElementWithDefaults() *RestoreSourceSummaryByObjectTypeElement`

NewRestoreSourceSummaryByObjectTypeElementWithDefaults instantiates a new RestoreSourceSummaryByObjectTypeElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastoreId

`func (o *RestoreSourceSummaryByObjectTypeElement) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *RestoreSourceSummaryByObjectTypeElement) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *RestoreSourceSummaryByObjectTypeElement) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *RestoreSourceSummaryByObjectTypeElement) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### SetDatastoreIdNil

`func (o *RestoreSourceSummaryByObjectTypeElement) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *RestoreSourceSummaryByObjectTypeElement) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
### GetFileRestoreInfo

`func (o *RestoreSourceSummaryByObjectTypeElement) GetFileRestoreInfo() []FileRestoreInfo`

GetFileRestoreInfo returns the FileRestoreInfo field if non-nil, zero value otherwise.

### GetFileRestoreInfoOk

`func (o *RestoreSourceSummaryByObjectTypeElement) GetFileRestoreInfoOk() (*[]FileRestoreInfo, bool)`

GetFileRestoreInfoOk returns a tuple with the FileRestoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileRestoreInfo

`func (o *RestoreSourceSummaryByObjectTypeElement) SetFileRestoreInfo(v []FileRestoreInfo)`

SetFileRestoreInfo sets FileRestoreInfo field to given value.

### HasFileRestoreInfo

`func (o *RestoreSourceSummaryByObjectTypeElement) HasFileRestoreInfo() bool`

HasFileRestoreInfo returns a boolean if a field has been set.

### SetFileRestoreInfoNil

`func (o *RestoreSourceSummaryByObjectTypeElement) SetFileRestoreInfoNil(b bool)`

 SetFileRestoreInfoNil sets the value for FileRestoreInfo to be an explicit nil

### UnsetFileRestoreInfo
`func (o *RestoreSourceSummaryByObjectTypeElement) UnsetFileRestoreInfo()`

UnsetFileRestoreInfo ensures that no value is present for FileRestoreInfo, not even an explicit nil
### GetName

`func (o *RestoreSourceSummaryByObjectTypeElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestoreSourceSummaryByObjectTypeElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestoreSourceSummaryByObjectTypeElement) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *RestoreSourceSummaryByObjectTypeElement) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RestoreSourceSummaryByObjectTypeElement) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetObjects

`func (o *RestoreSourceSummaryByObjectTypeElement) GetObjects() []RestoreObjectDetails`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RestoreSourceSummaryByObjectTypeElement) GetObjectsOk() (*[]RestoreObjectDetails, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RestoreSourceSummaryByObjectTypeElement) SetObjects(v []RestoreObjectDetails)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *RestoreSourceSummaryByObjectTypeElement) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *RestoreSourceSummaryByObjectTypeElement) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RestoreSourceSummaryByObjectTypeElement) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetProtectionSourceName

`func (o *RestoreSourceSummaryByObjectTypeElement) GetProtectionSourceName() string`

GetProtectionSourceName returns the ProtectionSourceName field if non-nil, zero value otherwise.

### GetProtectionSourceNameOk

`func (o *RestoreSourceSummaryByObjectTypeElement) GetProtectionSourceNameOk() (*string, bool)`

GetProtectionSourceNameOk returns a tuple with the ProtectionSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceName

`func (o *RestoreSourceSummaryByObjectTypeElement) SetProtectionSourceName(v string)`

SetProtectionSourceName sets ProtectionSourceName field to given value.

### HasProtectionSourceName

`func (o *RestoreSourceSummaryByObjectTypeElement) HasProtectionSourceName() bool`

HasProtectionSourceName returns a boolean if a field has been set.

### SetProtectionSourceNameNil

`func (o *RestoreSourceSummaryByObjectTypeElement) SetProtectionSourceNameNil(b bool)`

 SetProtectionSourceNameNil sets the value for ProtectionSourceName to be an explicit nil

### UnsetProtectionSourceName
`func (o *RestoreSourceSummaryByObjectTypeElement) UnsetProtectionSourceName()`

UnsetProtectionSourceName ensures that no value is present for ProtectionSourceName, not even an explicit nil
### GetStartTimeUsecs

`func (o *RestoreSourceSummaryByObjectTypeElement) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *RestoreSourceSummaryByObjectTypeElement) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *RestoreSourceSummaryByObjectTypeElement) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *RestoreSourceSummaryByObjectTypeElement) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *RestoreSourceSummaryByObjectTypeElement) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *RestoreSourceSummaryByObjectTypeElement) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetType

`func (o *RestoreSourceSummaryByObjectTypeElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreSourceSummaryByObjectTypeElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreSourceSummaryByObjectTypeElement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RestoreSourceSummaryByObjectTypeElement) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RestoreSourceSummaryByObjectTypeElement) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RestoreSourceSummaryByObjectTypeElement) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUsername

`func (o *RestoreSourceSummaryByObjectTypeElement) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RestoreSourceSummaryByObjectTypeElement) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RestoreSourceSummaryByObjectTypeElement) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RestoreSourceSummaryByObjectTypeElement) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *RestoreSourceSummaryByObjectTypeElement) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *RestoreSourceSummaryByObjectTypeElement) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


