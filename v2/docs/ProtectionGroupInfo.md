# ProtectionGroupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **NullableInt64** | This field is deprecated. &#39;protectionGroupId&#39; should be used instead. Specifies the id of the Protection Group. | [optional] 
**GroupName** | Pointer to **NullableString** | Specifies the name of the Protection Group. | [optional] 
**IsPaused** | Pointer to **NullableBool** | Specifies if the Protection Group&#39;s run is paused. | [optional] 
**LastRun** | Pointer to [**ProtectionGroupRun**](ProtectionGroupRun.md) |  | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the protection group id. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the Protection Group such as View or Puppeteer. &#39;Puppeteer&#39; refers to a Remote Adapter Group. Supported environment types such as &#39;View&#39;, &#39;SQL&#39;, &#39;VMware&#39;, etc. NOTE: &#39;Puppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;VMware&#39; indicates the VMware Protection Source environment. &#39;HyperV&#39; indicates the HyperV Protection Source environment. &#39;SQL&#39; indicates the SQL Protection Source environment. &#39;View&#39; indicates the View Protection Source environment. &#39;Puppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;Physical&#39; indicates the physical Protection Source environment. &#39;Pure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;Azure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;Netapp&#39; indicates the Netapp Protection Source environment. &#39;Agent&#39; indicates the Agent Protection Source environment. &#39;GenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;Acropolis&#39; indicates the Acropolis Protection Source environment. &#39;PhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;Isilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;GPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;KVM&#39; indicates the KVM Protection Source environment. &#39;AWS&#39; indicates the AWS Protection Source environment. &#39;Exchange&#39; indicates the Exchange Protection Source environment. &#39;HyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;Oracle&#39; indicates the Oracle Protection Source environment. &#39;GCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;FlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;AWSNative&#39; indicates the AWS Native Protection Source environment. &#39;O365&#39; indicates the Office 365 Protection Source environment. &#39;O365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;HyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;GCPNative&#39; indicates the GCP Native Protection Source environment. &#39;AzureNative&#39; indicates the Azure Native Protection Source environment. &#39;Kubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;Elastifile&#39; indicates Elastifile Protection Source environment. &#39;AD&#39; indicates Active Directory Protection Source environment. | [optional] 

## Methods

### NewProtectionGroupInfo

`func NewProtectionGroupInfo() *ProtectionGroupInfo`

NewProtectionGroupInfo instantiates a new ProtectionGroupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionGroupInfoWithDefaults

`func NewProtectionGroupInfoWithDefaults() *ProtectionGroupInfo`

NewProtectionGroupInfoWithDefaults instantiates a new ProtectionGroupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ProtectionGroupInfo) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ProtectionGroupInfo) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ProtectionGroupInfo) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ProtectionGroupInfo) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *ProtectionGroupInfo) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *ProtectionGroupInfo) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *ProtectionGroupInfo) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ProtectionGroupInfo) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ProtectionGroupInfo) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ProtectionGroupInfo) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *ProtectionGroupInfo) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *ProtectionGroupInfo) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetIsPaused

`func (o *ProtectionGroupInfo) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *ProtectionGroupInfo) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *ProtectionGroupInfo) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *ProtectionGroupInfo) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### SetIsPausedNil

`func (o *ProtectionGroupInfo) SetIsPausedNil(b bool)`

 SetIsPausedNil sets the value for IsPaused to be an explicit nil

### UnsetIsPaused
`func (o *ProtectionGroupInfo) UnsetIsPaused()`

UnsetIsPaused ensures that no value is present for IsPaused, not even an explicit nil
### GetLastRun

`func (o *ProtectionGroupInfo) GetLastRun() ProtectionGroupRun`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *ProtectionGroupInfo) GetLastRunOk() (*ProtectionGroupRun, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *ProtectionGroupInfo) SetLastRun(v ProtectionGroupRun)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *ProtectionGroupInfo) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetProtectionGroupId

`func (o *ProtectionGroupInfo) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *ProtectionGroupInfo) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *ProtectionGroupInfo) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *ProtectionGroupInfo) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *ProtectionGroupInfo) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *ProtectionGroupInfo) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetType

`func (o *ProtectionGroupInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProtectionGroupInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProtectionGroupInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProtectionGroupInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ProtectionGroupInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ProtectionGroupInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


