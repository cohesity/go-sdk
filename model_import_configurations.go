/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"encoding/json"
)

// ImportConfigurations ImportConfigurations struct used for ImportConfig endpoint. This is the form of the request.
type ImportConfigurations struct {
	// Selective import of active directories.
	ActiveDirectories []string `json:"activeDirectories,omitempty"`
	// List of which entities to import all.
	All []string `json:"all,omitempty"`
	// Selective import certain cluster.
	Clusters []int64 `json:"clusters,omitempty"`
	// File is the config file.
	File NullableString `json:"file,omitempty"`
	// Selective import certain groups.
	Groups []string `json:"groups,omitempty"`
	// Selective import of Partition.
	Partitions []int64 `json:"partitions,omitempty"`
	// Selective import of principal sources.
	PrincipalSources []string `json:"principalSources,omitempty"`
	// Selective import of protection jobs.
	ProtectionJobs []int64 `json:"protectionJobs,omitempty"`
	// Selective import of protection policies.
	ProtectionPolicies []string `json:"protectionPolicies,omitempty"`
	// Selective import of protection sources.
	ProtectionSources []int64 `json:"protectionSources,omitempty"`
	// Selective import certain remote clusters.
	RemoteClusters []int64 `json:"remoteClusters,omitempty"`
	// Selective import certain roles (by username).
	Roles []string `json:"roles,omitempty"`
	// Selective import of sql.
	Sql []int64 `json:"sql,omitempty"`
	// Selective import certain users.
	Users []string `json:"users,omitempty"`
	// Selective import certain vaults.
	Vaults []int64 `json:"vaults,omitempty"`
	// Selective import certain Storage Domains (View Boxes).
	ViewBoxes []int64 `json:"viewBoxes,omitempty"`
	// Selective import of views.
	Views []int64 `json:"views,omitempty"`
}

// NewImportConfigurations instantiates a new ImportConfigurations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportConfigurations() *ImportConfigurations {
	this := ImportConfigurations{}
	return &this
}

// NewImportConfigurationsWithDefaults instantiates a new ImportConfigurations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportConfigurationsWithDefaults() *ImportConfigurations {
	this := ImportConfigurations{}
	return &this
}

// GetActiveDirectories returns the ActiveDirectories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetActiveDirectories() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.ActiveDirectories
}

// GetActiveDirectoriesOk returns a tuple with the ActiveDirectories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetActiveDirectoriesOk() (*[]string, bool) {
	if o == nil || o.ActiveDirectories == nil {
		return nil, false
	}
	return &o.ActiveDirectories, true
}

// HasActiveDirectories returns a boolean if a field has been set.
func (o *ImportConfigurations) HasActiveDirectories() bool {
	if o != nil && o.ActiveDirectories != nil {
		return true
	}

	return false
}

// SetActiveDirectories gets a reference to the given []string and assigns it to the ActiveDirectories field.
func (o *ImportConfigurations) SetActiveDirectories(v []string) {
	o.ActiveDirectories = v
}

// GetAll returns the All field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetAll() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetAllOk() (*[]string, bool) {
	if o == nil || o.All == nil {
		return nil, false
	}
	return &o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *ImportConfigurations) HasAll() bool {
	if o != nil && o.All != nil {
		return true
	}

	return false
}

// SetAll gets a reference to the given []string and assigns it to the All field.
func (o *ImportConfigurations) SetAll(v []string) {
	o.All = v
}

// GetClusters returns the Clusters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetClusters() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetClustersOk() (*[]int64, bool) {
	if o == nil || o.Clusters == nil {
		return nil, false
	}
	return &o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *ImportConfigurations) HasClusters() bool {
	if o != nil && o.Clusters != nil {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []int64 and assigns it to the Clusters field.
func (o *ImportConfigurations) SetClusters(v []int64) {
	o.Clusters = v
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetFile() string {
	if o == nil || o.File.Get() == nil {
		var ret string
		return ret
	}
	return *o.File.Get()
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetFileOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// HasFile returns a boolean if a field has been set.
func (o *ImportConfigurations) HasFile() bool {
	if o != nil && o.File.IsSet() {
		return true
	}

	return false
}

// SetFile gets a reference to the given NullableString and assigns it to the File field.
func (o *ImportConfigurations) SetFile(v string) {
	o.File.Set(&v)
}
// SetFileNil sets the value for File to be an explicit nil
func (o *ImportConfigurations) SetFileNil() {
	o.File.Set(nil)
}

// UnsetFile ensures that no value is present for File, not even an explicit nil
func (o *ImportConfigurations) UnsetFile() {
	o.File.Unset()
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetGroups() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetGroupsOk() (*[]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ImportConfigurations) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *ImportConfigurations) SetGroups(v []string) {
	o.Groups = v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetPartitions() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetPartitionsOk() (*[]int64, bool) {
	if o == nil || o.Partitions == nil {
		return nil, false
	}
	return &o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *ImportConfigurations) HasPartitions() bool {
	if o != nil && o.Partitions != nil {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []int64 and assigns it to the Partitions field.
func (o *ImportConfigurations) SetPartitions(v []int64) {
	o.Partitions = v
}

// GetPrincipalSources returns the PrincipalSources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetPrincipalSources() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.PrincipalSources
}

// GetPrincipalSourcesOk returns a tuple with the PrincipalSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetPrincipalSourcesOk() (*[]string, bool) {
	if o == nil || o.PrincipalSources == nil {
		return nil, false
	}
	return &o.PrincipalSources, true
}

// HasPrincipalSources returns a boolean if a field has been set.
func (o *ImportConfigurations) HasPrincipalSources() bool {
	if o != nil && o.PrincipalSources != nil {
		return true
	}

	return false
}

// SetPrincipalSources gets a reference to the given []string and assigns it to the PrincipalSources field.
func (o *ImportConfigurations) SetPrincipalSources(v []string) {
	o.PrincipalSources = v
}

// GetProtectionJobs returns the ProtectionJobs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetProtectionJobs() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.ProtectionJobs
}

// GetProtectionJobsOk returns a tuple with the ProtectionJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetProtectionJobsOk() (*[]int64, bool) {
	if o == nil || o.ProtectionJobs == nil {
		return nil, false
	}
	return &o.ProtectionJobs, true
}

// HasProtectionJobs returns a boolean if a field has been set.
func (o *ImportConfigurations) HasProtectionJobs() bool {
	if o != nil && o.ProtectionJobs != nil {
		return true
	}

	return false
}

// SetProtectionJobs gets a reference to the given []int64 and assigns it to the ProtectionJobs field.
func (o *ImportConfigurations) SetProtectionJobs(v []int64) {
	o.ProtectionJobs = v
}

// GetProtectionPolicies returns the ProtectionPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetProtectionPolicies() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.ProtectionPolicies
}

// GetProtectionPoliciesOk returns a tuple with the ProtectionPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetProtectionPoliciesOk() (*[]string, bool) {
	if o == nil || o.ProtectionPolicies == nil {
		return nil, false
	}
	return &o.ProtectionPolicies, true
}

// HasProtectionPolicies returns a boolean if a field has been set.
func (o *ImportConfigurations) HasProtectionPolicies() bool {
	if o != nil && o.ProtectionPolicies != nil {
		return true
	}

	return false
}

// SetProtectionPolicies gets a reference to the given []string and assigns it to the ProtectionPolicies field.
func (o *ImportConfigurations) SetProtectionPolicies(v []string) {
	o.ProtectionPolicies = v
}

// GetProtectionSources returns the ProtectionSources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetProtectionSources() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.ProtectionSources
}

// GetProtectionSourcesOk returns a tuple with the ProtectionSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetProtectionSourcesOk() (*[]int64, bool) {
	if o == nil || o.ProtectionSources == nil {
		return nil, false
	}
	return &o.ProtectionSources, true
}

// HasProtectionSources returns a boolean if a field has been set.
func (o *ImportConfigurations) HasProtectionSources() bool {
	if o != nil && o.ProtectionSources != nil {
		return true
	}

	return false
}

// SetProtectionSources gets a reference to the given []int64 and assigns it to the ProtectionSources field.
func (o *ImportConfigurations) SetProtectionSources(v []int64) {
	o.ProtectionSources = v
}

// GetRemoteClusters returns the RemoteClusters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetRemoteClusters() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.RemoteClusters
}

// GetRemoteClustersOk returns a tuple with the RemoteClusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetRemoteClustersOk() (*[]int64, bool) {
	if o == nil || o.RemoteClusters == nil {
		return nil, false
	}
	return &o.RemoteClusters, true
}

// HasRemoteClusters returns a boolean if a field has been set.
func (o *ImportConfigurations) HasRemoteClusters() bool {
	if o != nil && o.RemoteClusters != nil {
		return true
	}

	return false
}

// SetRemoteClusters gets a reference to the given []int64 and assigns it to the RemoteClusters field.
func (o *ImportConfigurations) SetRemoteClusters(v []int64) {
	o.RemoteClusters = v
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetRoles() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetRolesOk() (*[]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ImportConfigurations) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *ImportConfigurations) SetRoles(v []string) {
	o.Roles = v
}

// GetSql returns the Sql field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetSql() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetSqlOk() (*[]int64, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return &o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *ImportConfigurations) HasSql() bool {
	if o != nil && o.Sql != nil {
		return true
	}

	return false
}

// SetSql gets a reference to the given []int64 and assigns it to the Sql field.
func (o *ImportConfigurations) SetSql(v []int64) {
	o.Sql = v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetUsers() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetUsersOk() (*[]string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ImportConfigurations) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *ImportConfigurations) SetUsers(v []string) {
	o.Users = v
}

// GetVaults returns the Vaults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetVaults() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.Vaults
}

// GetVaultsOk returns a tuple with the Vaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetVaultsOk() (*[]int64, bool) {
	if o == nil || o.Vaults == nil {
		return nil, false
	}
	return &o.Vaults, true
}

// HasVaults returns a boolean if a field has been set.
func (o *ImportConfigurations) HasVaults() bool {
	if o != nil && o.Vaults != nil {
		return true
	}

	return false
}

// SetVaults gets a reference to the given []int64 and assigns it to the Vaults field.
func (o *ImportConfigurations) SetVaults(v []int64) {
	o.Vaults = v
}

// GetViewBoxes returns the ViewBoxes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetViewBoxes() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.ViewBoxes
}

// GetViewBoxesOk returns a tuple with the ViewBoxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetViewBoxesOk() (*[]int64, bool) {
	if o == nil || o.ViewBoxes == nil {
		return nil, false
	}
	return &o.ViewBoxes, true
}

// HasViewBoxes returns a boolean if a field has been set.
func (o *ImportConfigurations) HasViewBoxes() bool {
	if o != nil && o.ViewBoxes != nil {
		return true
	}

	return false
}

// SetViewBoxes gets a reference to the given []int64 and assigns it to the ViewBoxes field.
func (o *ImportConfigurations) SetViewBoxes(v []int64) {
	o.ViewBoxes = v
}

// GetViews returns the Views field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportConfigurations) GetViews() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportConfigurations) GetViewsOk() (*[]int64, bool) {
	if o == nil || o.Views == nil {
		return nil, false
	}
	return &o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *ImportConfigurations) HasViews() bool {
	if o != nil && o.Views != nil {
		return true
	}

	return false
}

// SetViews gets a reference to the given []int64 and assigns it to the Views field.
func (o *ImportConfigurations) SetViews(v []int64) {
	o.Views = v
}

func (o ImportConfigurations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveDirectories != nil {
		toSerialize["activeDirectories"] = o.ActiveDirectories
	}
	if o.All != nil {
		toSerialize["all"] = o.All
	}
	if o.Clusters != nil {
		toSerialize["clusters"] = o.Clusters
	}
	if o.File.IsSet() {
		toSerialize["file"] = o.File.Get()
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Partitions != nil {
		toSerialize["partitions"] = o.Partitions
	}
	if o.PrincipalSources != nil {
		toSerialize["principalSources"] = o.PrincipalSources
	}
	if o.ProtectionJobs != nil {
		toSerialize["protectionJobs"] = o.ProtectionJobs
	}
	if o.ProtectionPolicies != nil {
		toSerialize["protectionPolicies"] = o.ProtectionPolicies
	}
	if o.ProtectionSources != nil {
		toSerialize["protectionSources"] = o.ProtectionSources
	}
	if o.RemoteClusters != nil {
		toSerialize["remoteClusters"] = o.RemoteClusters
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Vaults != nil {
		toSerialize["vaults"] = o.Vaults
	}
	if o.ViewBoxes != nil {
		toSerialize["viewBoxes"] = o.ViewBoxes
	}
	if o.Views != nil {
		toSerialize["views"] = o.Views
	}
	return json.Marshal(toSerialize)
}

type NullableImportConfigurations struct {
	value *ImportConfigurations
	isSet bool
}

func (v NullableImportConfigurations) Get() *ImportConfigurations {
	return v.value
}

func (v *NullableImportConfigurations) Set(val *ImportConfigurations) {
	v.value = val
	v.isSet = true
}

func (v NullableImportConfigurations) IsSet() bool {
	return v.isSet
}

func (v *NullableImportConfigurations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportConfigurations(val *ImportConfigurations) *NullableImportConfigurations {
	return &NullableImportConfigurations{value: val, isSet: true}
}

func (v NullableImportConfigurations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportConfigurations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


