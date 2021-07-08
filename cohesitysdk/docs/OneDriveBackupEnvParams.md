# OneDriveBackupEnvParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilteringPolicy** | Pointer to [**FilteringPolicyProto**](FilteringPolicyProto.md) |  | [optional] 
**ShouldBackupOnedrive** | Pointer to **NullableBool** | Specifies whether the OneDrive(s) for all the Office365 Users present in the protection job should be backed up. | [optional] 

## Methods

### NewOneDriveBackupEnvParams

`func NewOneDriveBackupEnvParams() *OneDriveBackupEnvParams`

NewOneDriveBackupEnvParams instantiates a new OneDriveBackupEnvParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneDriveBackupEnvParamsWithDefaults

`func NewOneDriveBackupEnvParamsWithDefaults() *OneDriveBackupEnvParams`

NewOneDriveBackupEnvParamsWithDefaults instantiates a new OneDriveBackupEnvParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilteringPolicy

`func (o *OneDriveBackupEnvParams) GetFilteringPolicy() FilteringPolicyProto`

GetFilteringPolicy returns the FilteringPolicy field if non-nil, zero value otherwise.

### GetFilteringPolicyOk

`func (o *OneDriveBackupEnvParams) GetFilteringPolicyOk() (*FilteringPolicyProto, bool)`

GetFilteringPolicyOk returns a tuple with the FilteringPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteringPolicy

`func (o *OneDriveBackupEnvParams) SetFilteringPolicy(v FilteringPolicyProto)`

SetFilteringPolicy sets FilteringPolicy field to given value.

### HasFilteringPolicy

`func (o *OneDriveBackupEnvParams) HasFilteringPolicy() bool`

HasFilteringPolicy returns a boolean if a field has been set.

### GetShouldBackupOnedrive

`func (o *OneDriveBackupEnvParams) GetShouldBackupOnedrive() bool`

GetShouldBackupOnedrive returns the ShouldBackupOnedrive field if non-nil, zero value otherwise.

### GetShouldBackupOnedriveOk

`func (o *OneDriveBackupEnvParams) GetShouldBackupOnedriveOk() (*bool, bool)`

GetShouldBackupOnedriveOk returns a tuple with the ShouldBackupOnedrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldBackupOnedrive

`func (o *OneDriveBackupEnvParams) SetShouldBackupOnedrive(v bool)`

SetShouldBackupOnedrive sets ShouldBackupOnedrive field to given value.

### HasShouldBackupOnedrive

`func (o *OneDriveBackupEnvParams) HasShouldBackupOnedrive() bool`

HasShouldBackupOnedrive returns a boolean if a field has been set.

### SetShouldBackupOnedriveNil

`func (o *OneDriveBackupEnvParams) SetShouldBackupOnedriveNil(b bool)`

 SetShouldBackupOnedriveNil sets the value for ShouldBackupOnedrive to be an explicit nil

### UnsetShouldBackupOnedrive
`func (o *OneDriveBackupEnvParams) UnsetShouldBackupOnedrive()`

UnsetShouldBackupOnedrive ensures that no value is present for ShouldBackupOnedrive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


