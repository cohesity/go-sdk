# RecoverMongodbParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedConfigs** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies the advanced configuration for a recovery job. | [optional] 
**BandwidthMBPS** | Pointer to **NullableInt64** | Specifies the maximum network bandwidth that each concurrent IO Stream can use for exchanging data with the cluster. | [optional] 
**Concurrency** | Pointer to **NullableInt32** | Specifies the maximum number of concurrent IO Streams that will be created to exchange data with the cluster. | [optional] 
**Overwrite** | Pointer to **NullableBool** | Set to true to overwrite an existing object at the destination. If set to false, and the same object exists at the destination, then recovery will fail for that object. | [optional] 
**RecoverTo** | Pointer to **NullableInt64** | Specifies the &#39;Source Registration ID&#39; of the source where the objects are to be recovered. If this is not specified, the recovery job will recover to the original location. | [optional] 
**Warnings** | Pointer to **[]string** | This field will hold the warnings in cases where the job status is SucceededWithWarnings. | [optional] [readonly] 
**RecoverUserRoles** | Pointer to **NullableBool** | Specifies whether to recover User and roles at the time of recovery. | [optional] 
**RecoverZonesTags** | Pointer to **NullableBool** | Specifies whether to recover Zones/shard tags at the time of recovery. | [optional] 
**Snapshots** | [**[]RecoverMongodbSnapshotParams**](RecoverMongodbSnapshotParams.md) | Specifies the local snapshot ids of the Objects to be recovered. | 
**Suffix** | Pointer to **NullableString** | A suffix that is to be applied to all recovered objects. | [optional] 

## Methods

### NewRecoverMongodbParams

`func NewRecoverMongodbParams(snapshots []RecoverMongodbSnapshotParams, ) *RecoverMongodbParams`

NewRecoverMongodbParams instantiates a new RecoverMongodbParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverMongodbParamsWithDefaults

`func NewRecoverMongodbParamsWithDefaults() *RecoverMongodbParams`

NewRecoverMongodbParamsWithDefaults instantiates a new RecoverMongodbParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedConfigs

`func (o *RecoverMongodbParams) GetAdvancedConfigs() []KeyValuePair`

GetAdvancedConfigs returns the AdvancedConfigs field if non-nil, zero value otherwise.

### GetAdvancedConfigsOk

`func (o *RecoverMongodbParams) GetAdvancedConfigsOk() (*[]KeyValuePair, bool)`

GetAdvancedConfigsOk returns a tuple with the AdvancedConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfigs

`func (o *RecoverMongodbParams) SetAdvancedConfigs(v []KeyValuePair)`

SetAdvancedConfigs sets AdvancedConfigs field to given value.

### HasAdvancedConfigs

`func (o *RecoverMongodbParams) HasAdvancedConfigs() bool`

HasAdvancedConfigs returns a boolean if a field has been set.

### SetAdvancedConfigsNil

`func (o *RecoverMongodbParams) SetAdvancedConfigsNil(b bool)`

 SetAdvancedConfigsNil sets the value for AdvancedConfigs to be an explicit nil

### UnsetAdvancedConfigs
`func (o *RecoverMongodbParams) UnsetAdvancedConfigs()`

UnsetAdvancedConfigs ensures that no value is present for AdvancedConfigs, not even an explicit nil
### GetBandwidthMBPS

`func (o *RecoverMongodbParams) GetBandwidthMBPS() int64`

GetBandwidthMBPS returns the BandwidthMBPS field if non-nil, zero value otherwise.

### GetBandwidthMBPSOk

`func (o *RecoverMongodbParams) GetBandwidthMBPSOk() (*int64, bool)`

GetBandwidthMBPSOk returns a tuple with the BandwidthMBPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMBPS

`func (o *RecoverMongodbParams) SetBandwidthMBPS(v int64)`

SetBandwidthMBPS sets BandwidthMBPS field to given value.

### HasBandwidthMBPS

`func (o *RecoverMongodbParams) HasBandwidthMBPS() bool`

HasBandwidthMBPS returns a boolean if a field has been set.

### SetBandwidthMBPSNil

`func (o *RecoverMongodbParams) SetBandwidthMBPSNil(b bool)`

 SetBandwidthMBPSNil sets the value for BandwidthMBPS to be an explicit nil

### UnsetBandwidthMBPS
`func (o *RecoverMongodbParams) UnsetBandwidthMBPS()`

UnsetBandwidthMBPS ensures that no value is present for BandwidthMBPS, not even an explicit nil
### GetConcurrency

`func (o *RecoverMongodbParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *RecoverMongodbParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *RecoverMongodbParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *RecoverMongodbParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *RecoverMongodbParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *RecoverMongodbParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetOverwrite

`func (o *RecoverMongodbParams) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *RecoverMongodbParams) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *RecoverMongodbParams) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *RecoverMongodbParams) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### SetOverwriteNil

`func (o *RecoverMongodbParams) SetOverwriteNil(b bool)`

 SetOverwriteNil sets the value for Overwrite to be an explicit nil

### UnsetOverwrite
`func (o *RecoverMongodbParams) UnsetOverwrite()`

UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
### GetRecoverTo

`func (o *RecoverMongodbParams) GetRecoverTo() int64`

GetRecoverTo returns the RecoverTo field if non-nil, zero value otherwise.

### GetRecoverToOk

`func (o *RecoverMongodbParams) GetRecoverToOk() (*int64, bool)`

GetRecoverToOk returns a tuple with the RecoverTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverTo

`func (o *RecoverMongodbParams) SetRecoverTo(v int64)`

SetRecoverTo sets RecoverTo field to given value.

### HasRecoverTo

`func (o *RecoverMongodbParams) HasRecoverTo() bool`

HasRecoverTo returns a boolean if a field has been set.

### SetRecoverToNil

`func (o *RecoverMongodbParams) SetRecoverToNil(b bool)`

 SetRecoverToNil sets the value for RecoverTo to be an explicit nil

### UnsetRecoverTo
`func (o *RecoverMongodbParams) UnsetRecoverTo()`

UnsetRecoverTo ensures that no value is present for RecoverTo, not even an explicit nil
### GetWarnings

`func (o *RecoverMongodbParams) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *RecoverMongodbParams) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *RecoverMongodbParams) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *RecoverMongodbParams) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *RecoverMongodbParams) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *RecoverMongodbParams) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetRecoverUserRoles

`func (o *RecoverMongodbParams) GetRecoverUserRoles() bool`

GetRecoverUserRoles returns the RecoverUserRoles field if non-nil, zero value otherwise.

### GetRecoverUserRolesOk

`func (o *RecoverMongodbParams) GetRecoverUserRolesOk() (*bool, bool)`

GetRecoverUserRolesOk returns a tuple with the RecoverUserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverUserRoles

`func (o *RecoverMongodbParams) SetRecoverUserRoles(v bool)`

SetRecoverUserRoles sets RecoverUserRoles field to given value.

### HasRecoverUserRoles

`func (o *RecoverMongodbParams) HasRecoverUserRoles() bool`

HasRecoverUserRoles returns a boolean if a field has been set.

### SetRecoverUserRolesNil

`func (o *RecoverMongodbParams) SetRecoverUserRolesNil(b bool)`

 SetRecoverUserRolesNil sets the value for RecoverUserRoles to be an explicit nil

### UnsetRecoverUserRoles
`func (o *RecoverMongodbParams) UnsetRecoverUserRoles()`

UnsetRecoverUserRoles ensures that no value is present for RecoverUserRoles, not even an explicit nil
### GetRecoverZonesTags

`func (o *RecoverMongodbParams) GetRecoverZonesTags() bool`

GetRecoverZonesTags returns the RecoverZonesTags field if non-nil, zero value otherwise.

### GetRecoverZonesTagsOk

`func (o *RecoverMongodbParams) GetRecoverZonesTagsOk() (*bool, bool)`

GetRecoverZonesTagsOk returns a tuple with the RecoverZonesTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverZonesTags

`func (o *RecoverMongodbParams) SetRecoverZonesTags(v bool)`

SetRecoverZonesTags sets RecoverZonesTags field to given value.

### HasRecoverZonesTags

`func (o *RecoverMongodbParams) HasRecoverZonesTags() bool`

HasRecoverZonesTags returns a boolean if a field has been set.

### SetRecoverZonesTagsNil

`func (o *RecoverMongodbParams) SetRecoverZonesTagsNil(b bool)`

 SetRecoverZonesTagsNil sets the value for RecoverZonesTags to be an explicit nil

### UnsetRecoverZonesTags
`func (o *RecoverMongodbParams) UnsetRecoverZonesTags()`

UnsetRecoverZonesTags ensures that no value is present for RecoverZonesTags, not even an explicit nil
### GetSnapshots

`func (o *RecoverMongodbParams) GetSnapshots() []RecoverMongodbSnapshotParams`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *RecoverMongodbParams) GetSnapshotsOk() (*[]RecoverMongodbSnapshotParams, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *RecoverMongodbParams) SetSnapshots(v []RecoverMongodbSnapshotParams)`

SetSnapshots sets Snapshots field to given value.


### SetSnapshotsNil

`func (o *RecoverMongodbParams) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *RecoverMongodbParams) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil
### GetSuffix

`func (o *RecoverMongodbParams) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *RecoverMongodbParams) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *RecoverMongodbParams) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *RecoverMongodbParams) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *RecoverMongodbParams) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *RecoverMongodbParams) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


