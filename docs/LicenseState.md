# LicenseState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedAttempts** | Pointer to **NullableInt64** | Specifies no of failed attempts at claiming the license server | [optional] 
**State** | Pointer to **NullableString** | Specifies the current state of licensing workflow. LicenseStateType specifies the licenseState type. &#39;kInProgressNewCluster&#39; indicates licensing server claim is in progress for &#39;New&#39; Cluster. &#39;kInProgressOldCluster&#39; indicates licensing server claim is in progress for &#39;Old&#39; Cluster. &#39;kClaimed&#39; indicates licensing server is claimed. &#39;kSkipped&#39; indicates licensing workflow has been skipped. &#39;kStarted&#39; indicates licensing UI workflow has started. | [optional] 

## Methods

### NewLicenseState

`func NewLicenseState() *LicenseState`

NewLicenseState instantiates a new LicenseState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseStateWithDefaults

`func NewLicenseStateWithDefaults() *LicenseState`

NewLicenseStateWithDefaults instantiates a new LicenseState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedAttempts

`func (o *LicenseState) GetFailedAttempts() int64`

GetFailedAttempts returns the FailedAttempts field if non-nil, zero value otherwise.

### GetFailedAttemptsOk

`func (o *LicenseState) GetFailedAttemptsOk() (*int64, bool)`

GetFailedAttemptsOk returns a tuple with the FailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttempts

`func (o *LicenseState) SetFailedAttempts(v int64)`

SetFailedAttempts sets FailedAttempts field to given value.

### HasFailedAttempts

`func (o *LicenseState) HasFailedAttempts() bool`

HasFailedAttempts returns a boolean if a field has been set.

### SetFailedAttemptsNil

`func (o *LicenseState) SetFailedAttemptsNil(b bool)`

 SetFailedAttemptsNil sets the value for FailedAttempts to be an explicit nil

### UnsetFailedAttempts
`func (o *LicenseState) UnsetFailedAttempts()`

UnsetFailedAttempts ensures that no value is present for FailedAttempts, not even an explicit nil
### GetState

`func (o *LicenseState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LicenseState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LicenseState) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LicenseState) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *LicenseState) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *LicenseState) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


