# IoPreferentialTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApolloIOPreferentialTier** | Pointer to **[]string** | Specifies the preferred storage tier used by Apollo as its working directory. | [optional] 
**ApolloWalIOPreferentialTier** | Pointer to **[]string** | Specifies the preferred storage tier used by Apollo as its actions WAL. | [optional] 
**AthenaIOPreferentialTier** | Pointer to **[]string** | Specifies the list of perferred storage tiers used by Athena. | [optional] 
**AthenaSlowerIOPreferentialTier** | Pointer to **[]string** | Specifies the list of perferred storage tiers used by Athena for slower storage. | [optional] 
**DownTierUsagePercentThresholds** | Pointer to **[]int32** | Specifies the usage percentage thresholds for the correponding storage tier. | [optional] 
**GrootIOPreferentialTier** | Pointer to **[]string** | Specifies the preferred storage tier used by Groot as its working directory. | [optional] 
**HydraDowntierIOPreferentialTier** | Pointer to **[]string** | Specifies the list of perferred storage tiers used by Hydra for offloading. | [optional] 
**HydraIOPreferentialTier** | Pointer to **[]string** | Specifies the list of perferred storage tiers used by Hydra. | [optional] 
**LibrarianIOPreferentialTier** | Pointer to **[]string** | Specifies the list of perferred storage tiers used by librarian. | [optional] 
**RandomIOPreferentialTier** | Pointer to **[]string** | Specifies the order of perferred storage tiers for random IO operations. | [optional] 
**ScribeIOPreferentialTier** | Pointer to **[]string** | Specifies the list of perferred storage tiers used by Scribe. | [optional] 
**SequentialIOPreferentialTier** | Pointer to **[]string** | Specifies the preferred storage tier for sequential IO operations. | [optional] 
**YodaIOPreferentialTier** | Pointer to **[]string** | Specifies the list of perferred storage tiers used by Yoda. | [optional] 

## Methods

### NewIoPreferentialTier

`func NewIoPreferentialTier() *IoPreferentialTier`

NewIoPreferentialTier instantiates a new IoPreferentialTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoPreferentialTierWithDefaults

`func NewIoPreferentialTierWithDefaults() *IoPreferentialTier`

NewIoPreferentialTierWithDefaults instantiates a new IoPreferentialTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApolloIOPreferentialTier

`func (o *IoPreferentialTier) GetApolloIOPreferentialTier() []string`

GetApolloIOPreferentialTier returns the ApolloIOPreferentialTier field if non-nil, zero value otherwise.

### GetApolloIOPreferentialTierOk

`func (o *IoPreferentialTier) GetApolloIOPreferentialTierOk() (*[]string, bool)`

GetApolloIOPreferentialTierOk returns a tuple with the ApolloIOPreferentialTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApolloIOPreferentialTier

`func (o *IoPreferentialTier) SetApolloIOPreferentialTier(v []string)`

SetApolloIOPreferentialTier sets ApolloIOPreferentialTier field to given value.

### HasApolloIOPreferentialTier

`func (o *IoPreferentialTier) HasApolloIOPreferentialTier() bool`

HasApolloIOPreferentialTier returns a boolean if a field has been set.

### GetApolloWalIOPreferentialTier

`func (o *IoPreferentialTier) GetApolloWalIOPreferentialTier() []string`

GetApolloWalIOPreferentialTier returns the ApolloWalIOPreferentialTier field if non-nil, zero value otherwise.

### GetApolloWalIOPreferentialTierOk

`func (o *IoPreferentialTier) GetApolloWalIOPreferentialTierOk() (*[]string, bool)`

GetApolloWalIOPreferentialTierOk returns a tuple with the ApolloWalIOPreferentialTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApolloWalIOPreferentialTier

`func (o *IoPreferentialTier) SetApolloWalIOPreferentialTier(v []string)`

SetApolloWalIOPreferentialTier sets ApolloWalIOPreferentialTier field to given value.

### HasApolloWalIOPreferentialTier

`func (o *IoPreferentialTier) HasApolloWalIOPreferentialTier() bool`

HasApolloWalIOPreferentialTier returns a boolean if a field has been set.

### GetAthenaIOPreferentialTier

`func (o *IoPreferentialTier) GetAthenaIOPreferentialTier() []string`

GetAthenaIOPreferentialTier returns the AthenaIOPreferentialTier field if non-nil, zero value otherwise.

### GetAthenaIOPreferentialTierOk

`func (o *IoPreferentialTier) GetAthenaIOPreferentialTierOk() (*[]string, bool)`

GetAthenaIOPreferentialTierOk returns a tuple with the AthenaIOPreferentialTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthenaIOPreferentialTier

`func (o *IoPreferentialTier) SetAthenaIOPreferentialTier(v []string)`

SetAthenaIOPreferentialTier sets AthenaIOPreferentialTier field to given value.

### HasAthenaIOPreferentialTier

`func (o *IoPreferentialTier) HasAthenaIOPreferentialTier() bool`

HasAthenaIOPreferentialTier returns a boolean if a field has been set.

### GetAthenaSlowerIOPreferentialTier

`func (o *IoPreferentialTier) GetAthenaSlowerIOPreferentialTier() []string`

GetAthenaSlowerIOPreferentialTier returns the AthenaSlowerIOPreferentialTier field if non-nil, zero value otherwise.

### GetAthenaSlowerIOPreferentialTierOk

`func (o *IoPreferentialTier) GetAthenaSlowerIOPreferentialTierOk() (*[]string, bool)`

GetAthenaSlowerIOPreferentialTierOk returns a tuple with the AthenaSlowerIOPreferentialTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthenaSlowerIOPreferentialTier

`func (o *IoPreferentialTier) SetAthenaSlowerIOPreferentialTier(v []string)`

SetAthenaSlowerIOPreferentialTier sets AthenaSlowerIOPreferentialTier field to given value.

### HasAthenaSlowerIOPreferentialTier

`func (o *IoPreferentialTier) HasAthenaSlowerIOPreferentialTier() bool`

HasAthenaSlowerIOPreferentialTier returns a boolean if a field has been set.

### GetDownTierUsagePercentThresholds

`func (o *IoPreferentialTier) GetDownTierUsagePercentThresholds() []int32`

GetDownTierUsagePercentThresholds returns the DownTierUsagePercentThresholds field if non-nil, zero value otherwise.

### GetDownTierUsagePercentThresholdsOk

`func (o *IoPreferentialTier) GetDownTierUsagePercentThresholdsOk() (*[]int32, bool)`

GetDownTierUsagePercentThresholdsOk returns a tuple with the DownTierUsagePercentThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownTierUsagePercentThresholds

`func (o *IoPreferentialTier) SetDownTierUsagePercentThresholds(v []int32)`

SetDownTierUsagePercentThresholds sets DownTierUsagePercentThresholds field to given value.

### HasDownTierUsagePercentThresholds

`func (o *IoPreferentialTier) HasDownTierUsagePercentThresholds() bool`

HasDownTierUsagePercentThresholds returns a boolean if a field has been set.

### GetGrootIOPreferentialTier

`func (o *IoPreferentialTier) GetGrootIOPreferentialTier() []string`

GetGrootIOPreferentialTier returns the GrootIOPreferentialTier field if non-nil, zero value otherwise.

### GetGrootIOPreferentialTierOk

`func (o *IoPreferentialTier) GetGrootIOPreferentialTierOk() (*[]string, bool)`

GetGrootIOPreferentialTierOk returns a tuple with the GrootIOPreferentialTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrootIOPreferentialTier

`func (o *IoPreferentialTier) SetGrootIOPreferentialTier(v []string)`

SetGrootIOPreferentialTier sets GrootIOPreferentialTier field to given value.

### HasGrootIOPreferentialTier

`func (o *IoPreferentialTier) HasGrootIOPreferentialTier() bool`

HasGrootIOPreferentialTier returns a boolean if a field has been set.

### GetHydraDowntierIOPreferentialTier

`func (o *IoPreferentialTier) GetHydraDowntierIOPreferentialTier() []string`

GetHydraDowntierIOPreferentialTier returns the HydraDowntierIOPreferentialTier field if non-nil, zero value otherwise.

### GetHydraDowntierIOPreferentialTierOk

`func (o *IoPreferentialTier) GetHydraDowntierIOPreferentialTierOk() (*[]string, bool)`

GetHydraDowntierIOPreferentialTierOk returns a tuple with the HydraDowntierIOPreferentialTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHydraDowntierIOPreferentialTier

`func (o *IoPreferentialTier) SetHydraDowntierIOPreferentialTier(v []string)`

SetHydraDowntierIOPreferentialTier sets HydraDowntierIOPreferentialTier field to given value.

### HasHydraDowntierIOPreferentialTier

`func (o *IoPreferentialTier) HasHydraDowntierIOPreferentialTier() bool`

HasHydraDowntierIOPreferentialTier returns a boolean if a field has been set.

### GetHydraIOPreferentialTier

`func (o *IoPreferentialTier) GetHydraIOPreferentialTier() []string`

GetHydraIOPreferentialTier returns the HydraIOPreferentialTier field if non-nil, zero value otherwise.

### GetHydraIOPreferentialTierOk

`func (o *IoPreferentialTier) GetHydraIOPreferentialTierOk() (*[]string, bool)`

GetHydraIOPreferentialTierOk returns a tuple with the HydraIOPreferentialTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHydraIOPreferentialTier

`func (o *IoPreferentialTier) SetHydraIOPreferentialTier(v []string)`

SetHydraIOPreferentialTier sets HydraIOPreferentialTier field to given value.

### HasHydraIOPreferentialTier

`func (o *IoPreferentialTier) HasHydraIOPreferentialTier() bool`

HasHydraIOPreferentialTier returns a boolean if a field has been set.

### GetLibrarianIOPreferentialTier

`func (o *IoPreferentialTier) GetLibrarianIOPreferentialTier() []string`

GetLibrarianIOPreferentialTier returns the LibrarianIOPreferentialTier field if non-nil, zero value otherwise.

### GetLibrarianIOPreferentialTierOk

`func (o *IoPreferentialTier) GetLibrarianIOPreferentialTierOk() (*[]string, bool)`

GetLibrarianIOPreferentialTierOk returns a tuple with the LibrarianIOPreferentialTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrarianIOPreferentialTier

`func (o *IoPreferentialTier) SetLibrarianIOPreferentialTier(v []string)`

SetLibrarianIOPreferentialTier sets LibrarianIOPreferentialTier field to given value.

### HasLibrarianIOPreferentialTier

`func (o *IoPreferentialTier) HasLibrarianIOPreferentialTier() bool`

HasLibrarianIOPreferentialTier returns a boolean if a field has been set.

### GetRandomIOPreferentialTier

`func (o *IoPreferentialTier) GetRandomIOPreferentialTier() []string`

GetRandomIOPreferentialTier returns the RandomIOPreferentialTier field if non-nil, zero value otherwise.

### GetRandomIOPreferentialTierOk

`func (o *IoPreferentialTier) GetRandomIOPreferentialTierOk() (*[]string, bool)`

GetRandomIOPreferentialTierOk returns a tuple with the RandomIOPreferentialTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomIOPreferentialTier

`func (o *IoPreferentialTier) SetRandomIOPreferentialTier(v []string)`

SetRandomIOPreferentialTier sets RandomIOPreferentialTier field to given value.

### HasRandomIOPreferentialTier

`func (o *IoPreferentialTier) HasRandomIOPreferentialTier() bool`

HasRandomIOPreferentialTier returns a boolean if a field has been set.

### GetScribeIOPreferentialTier

`func (o *IoPreferentialTier) GetScribeIOPreferentialTier() []string`

GetScribeIOPreferentialTier returns the ScribeIOPreferentialTier field if non-nil, zero value otherwise.

### GetScribeIOPreferentialTierOk

`func (o *IoPreferentialTier) GetScribeIOPreferentialTierOk() (*[]string, bool)`

GetScribeIOPreferentialTierOk returns a tuple with the ScribeIOPreferentialTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScribeIOPreferentialTier

`func (o *IoPreferentialTier) SetScribeIOPreferentialTier(v []string)`

SetScribeIOPreferentialTier sets ScribeIOPreferentialTier field to given value.

### HasScribeIOPreferentialTier

`func (o *IoPreferentialTier) HasScribeIOPreferentialTier() bool`

HasScribeIOPreferentialTier returns a boolean if a field has been set.

### GetSequentialIOPreferentialTier

`func (o *IoPreferentialTier) GetSequentialIOPreferentialTier() []string`

GetSequentialIOPreferentialTier returns the SequentialIOPreferentialTier field if non-nil, zero value otherwise.

### GetSequentialIOPreferentialTierOk

`func (o *IoPreferentialTier) GetSequentialIOPreferentialTierOk() (*[]string, bool)`

GetSequentialIOPreferentialTierOk returns a tuple with the SequentialIOPreferentialTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequentialIOPreferentialTier

`func (o *IoPreferentialTier) SetSequentialIOPreferentialTier(v []string)`

SetSequentialIOPreferentialTier sets SequentialIOPreferentialTier field to given value.

### HasSequentialIOPreferentialTier

`func (o *IoPreferentialTier) HasSequentialIOPreferentialTier() bool`

HasSequentialIOPreferentialTier returns a boolean if a field has been set.

### GetYodaIOPreferentialTier

`func (o *IoPreferentialTier) GetYodaIOPreferentialTier() []string`

GetYodaIOPreferentialTier returns the YodaIOPreferentialTier field if non-nil, zero value otherwise.

### GetYodaIOPreferentialTierOk

`func (o *IoPreferentialTier) GetYodaIOPreferentialTierOk() (*[]string, bool)`

GetYodaIOPreferentialTierOk returns a tuple with the YodaIOPreferentialTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYodaIOPreferentialTier

`func (o *IoPreferentialTier) SetYodaIOPreferentialTier(v []string)`

SetYodaIOPreferentialTier sets YodaIOPreferentialTier field to given value.

### HasYodaIOPreferentialTier

`func (o *IoPreferentialTier) HasYodaIOPreferentialTier() bool`

HasYodaIOPreferentialTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


