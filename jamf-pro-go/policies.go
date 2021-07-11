package jamf_pro_go

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"path"
)

const (
	APIVersionPolicies    = "classic"
	APIPathPolices        = "policies"
)

type Policy struct {
	XMLName               *xml.Name                   `yaml:"-" xml:"policy,omitempty"`
	General               *PolicyGeneral              `yaml:"general,omitempty" xml:"general,omitempty"`
	Scope                 *PolicyScope                `yaml:"scope,omitempty,omitempty ,omitempty" xml:"scope,omitempty"`
	SelfService           *PolicySelfService          `yaml:"self_service,omitempty" xml:"self_service,omitempty"`
	PackageConfiguration  *PolicyPackageConfiguration `yaml:"package_configuration,omitempty" xml:"package_configuration,omitempty"`
	Scripts               *PolicyScripts              `yaml:"scripts,omitempty" xml:"scripts,omitempty"`
	Printers              *PolicyPrinters             `yaml:"printers,omitempty" xml:"printers,omitempty"`
	DockItems             *PolicyDockItems            `yaml:"dock_items,omitempty" xml:"dock_items,omitempty"`
	AccountMaintenance    *PolicyAccountMaintenance   `yaml:"account_maintenance,omitempty" xml:"account_maintenance,omitempty"`
	RebootSettings        *PolicyRebootSettings       `yaml:"reboot_settings,omitempty" xml:"reboot,omitempty"`
	Maintenance           *PolicyMaintenance          `yaml:"maintenance,omitempty" xml:"maintenance,omitempty"`
	FilesProcesses        *PolicyFilesProcesses       `yaml:"files_processes,omitempty" xml:"files_processes,omitempty"`
	UserInteraction       *PolicyUserInteraction      `yaml:"user_interaction,omitempty" xml:"user_interaction,omitempty"`
	DiskEncryption        *PolicyDiskEncryption       `yaml:"disk_encryption,omitempty" xml:"disk_encryption,omitempty"`
}

type PolicyGeneral struct {
	ID                          uint32                      `yaml:"id,omitempty" xml:"id,omitempty"`
	Name                        string                      `yaml:"name,omitempty" xml:"name,omitempty"`
	Enabled                     bool                        `yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	Trigger                     string                      `yaml:"trigger,omitempty" xml:"trigger,omitempty,omitempty"`
	TriggerCheckin              bool                        `yaml:"trigger_checkin,omitempty" xml:"trigger_checkin,omitempty"`
	TriggerEnrollmentComplete   bool                        `yaml:"trigger_enrollment_complete,omitempty" xml:"trigger_enrollment_complete,omitempty"`
	TriggerLogin                bool                        `yaml:"trigger_login,omitempty" xml:"trigger_login,omitempty"`
	TriggerLogout               bool                        `yaml:"trigger_logout,omitempty" xml:"trigger_logout,omitempty"`
	TriggerNetworkStateChanged  bool                        `yaml:"trigger_network_state_changed,omitempty" xml:"trigger_network_state_changed,omitempty"`
	TriggerStartup              bool                        `yaml:"trigger_startup,omitempty" xml:"trigger_startup,omitempty"`
	TriggerOther                string                      `yaml:"trigger_other,omitempty" xml:"trigger_other,omitempty"`
	Frequency                   string                      `yaml:"frequency,omitempty" xml:"frequency,omitempty"`
	Offline                     bool                        `yaml:"offline,omitempty" xml:"offline,omitempty"`
	RetryEvent                  string                      `yaml:"retry_event,omitempty" xml:"retry_event,omitempty"`
	RetryAttempts               int32                       `yaml:"retry_attempts,omitempty" xml:"retry_attempts,omitempty"`
	NotifyOnEachFailedRetry     bool                        `yaml:"notify_on_each_failed_retry,omitempty" xml:"notify_on_each_failed_retry,omitempty"`
	LocationUserOnly            string                      `yaml:"location_user_only,omitempty" xml:"location_user_only,omitempty"`
	TargetDrive                 string                      `yaml:"target_drive,omitempty" xml:"target_drive,omitempty"`
	Category                    *PolicyCategory             `yaml:"category,omitempty" xml:"category,omitempty"`
	DateTimeLimitations         *PolicyDateTimeLimitations  `yaml:"date_time_limitations,omitempty" xml:"date_time_limitations,omitempty"`
	NetworkLimitations          *PolicyNetworkLimitations   `yaml:"network_limitations,omitempty" xml:"network_limitations,omitempty"`
	OverrideDefaultSettings     *PolicyOverrides            `yaml:"override_default_settings,omitempty" xml:"override_default_settings,omitempty"`
	NetworkRequirements         string                      `yaml:"network_requirements,omitempty" xml:"network_requirements,omitempty"`
	Site                        *PolicySite                 `yaml:"site,omitempty" xml:"site,omitempty"`
}

type PolicyCategory struct {
	ID    int32  `yaml:"id,omitempty" xml:"id,omitempty"`
	Name  string `yaml:"name,omitempty" xml:"name,omitempty"`
}

type PolicyDateTimeLimitations struct {
	ActivationDate      string `yaml:"activation_date,omitempty" xml:"activation_date,omitempty"`
	ActivationDateEpoch uint64 `yaml:"activation_date_epoch,omitempty" xml:"activation_date_epoch,omitempty"`
	ActivationDateUtc   string `yaml:"activation_date_utc,omitempty" xml:"activation_date_utc,omitempty"`
	ExpirationDate      string `yaml:"expiration_date,omitempty" xml:"expiration_date,omitempty"`
	ExpirationDateEpoch uint64 `yaml:"expiration_date_epoch,omitempty" xml:"expiration_date_epoch,omitempty"`
	ExpirationDateUtc   string `yaml:"expiration_date_utc,omitempty" xml:"expiration_date_utc,omitempty"`
	NoExecuteOn         string `yaml:"no_execute_on,omitempty" xml:"no_execute_on,omitempty"`
	NoExecuteStart      string `yaml:"no_execute_start,omitempty" xml:"no_execute_start,omitempty"`
	NoExecuteEnd        string `yaml:"no_execute_end,omitempty" xml:"no_execute_end,omitempty"`
}

type PolicyNetworkLimitations struct {
	MinimumNetworkConnection  string `yaml:"minimum_network_connection,omitempty" xml:"minimum_network_connection,omitempty"`
	AnyIPAddress              bool   `yaml:"any_ip_address,omitempty" xml:"any_ip_address,omitempty"`
}

type PolicyOverrides struct {
	TargetDrive         string `yaml:"target_drive,omitempty" xml:"target_drive,omitempty"`
	DistributionPoint   string `yaml:"distribution_point,omitempty" xml:"distribution_point,omitempty"`
	ForceAfpSmb         bool   `yaml:"force_afp_smb,omitempty" xml:"force_afp_smb,omitempty"`
	Sus                 string `yaml:"sus,omitempty" xml:"sus,omitempty"`
	NetbootServer       string `yaml:"netboot_server,omitempty" xml:"netboot_server,omitempty"`
}

type PolicySite struct {
	ID    int32  `yaml:"id,omitempty" xml:"id,omitempty"`
	Name  string `yaml:"name,omitempty" xml:"name,omitempty"`
}

type PolicyScope struct {
	AllComputers    bool                       `yaml:"all_computers,omitempty" xml:"all_computers,omitempty"`
	Computers       *PolicyScopeComputers      `yaml:"computers,omitempty" xml:"computers,omitempty"`
	ComputerGroups  *PolicyScopeComputerGroups `yaml:"computer_groups,omitempty" xml:"computer_groups,omitempty"`
	Buildings       *PolicyScopeBuildings      `yaml:"buildings,omitempty" xml:"buildings,omitempty"`
	Departments     *PolicyScopeDepartments    `yaml:"departments,omitempty" xml:"departments,omitempty"`
	LimitToUsers    *PolicyScopeLimitToUsers   `yaml:"limit_to_users,omitempty" xml:"limit_to_users,omitempty"`
	Limitations     *PolicyScopeLimitations    `yaml:"limitations,omitempty" xml:"limitations,omitempty"`
	Exclusions      *PolicyScopeExclusions     `yaml:"exclusions,omitempty" xml:"exclusions,omitempty"`
}

type PolicyScopeComputers struct {
	Computer  []*PolicyScopeComputer `yaml:"computer,omitempty" xml:"computer,omitempty"`
}

type PolicyScopeComputer struct {
	ID    uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name  string `yaml:"name,omitempty" xml:"name,omitempty"`
	UDID  string `yaml:"udid,omitempty" xml:"udid,omitempty"`
}

type PolicyScopeComputerGroups struct {
	ComputerGroup  []*PolicyScopeComputerGroup `yaml:"computer_group,omitempty" xml:"computer_group,omitempty"`
}

type PolicyScopeComputerGroup struct {
	ID    uint32  `yaml:"id,omitempty" xml:"id,omitempty"`
	Name  string  `yaml:"name,omitempty" xml:"name,omitempty"`
}

type PolicyScopeBuildings struct {
	Building  []*PolicyScopeBuilding `yaml:"building,omitempty" xml:"building,omitempty"`
}

type PolicyScopeBuilding struct {
	ID    uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name  string `yaml:"name,omitempty" xml:"name,omitempty"`
}

type PolicyScopeDepartments struct {
	Department  []*PolicyScopeDepartment `yaml:"department,omitempty" xml:"department,omitempty"`
}

type PolicyScopeDepartment struct {
	ID    uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name  string `yaml:"name,omitempty" xml:"name,omitempty"`
}

type PolicyScopeLimitToUsers struct {
	UserGroups  *PolicyScopeLimitUserGroups `yaml:"user_groups,omitempty" xml:"user_groups,omitempty"`
}

type PolicyScopeLimitUserGroups struct {
	UsrGroups  []*PolicyScopeLimitUserGroup `yaml:"usr_groups,omitempty" xml:"usr_groups,omitempty"`
}

type PolicyScopeLimitUserGroup struct {
	UserGroup  string `yaml:"user_group,omitempty" xml:"user_group,omitempty"`
}

type PolicyScopeLimitations struct {
	Users           *PolicyScopeUsers           `yaml:"users,omitempty" xml:"users,omitempty"`
	UserGroups      *PolicyScopeUserGroups      `yaml:"user_groups,omitempty" xml:"user_groups,omitempty"`
	NetworkSegments *PolicyScopeNetworkSegments `yaml:"network_segments,omitempty" xml:"network_segments,omitempty"`
	Ibeacons        *PolicyScopeIbeacons        `yaml:"ibeacons,omitempty" xml:"ibeacons,omitempty"`
}

type PolicyScopeUsers struct {
	User  []*PolicyScopeUsersUser `yaml:"user,omitempty" xml:"user,omitempty"`
}

type PolicyScopeUsersUser struct {
	ID    uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name  string `yaml:"name,omitempty" xml:"name,omitempty"`
}

type PolicyScopeUserGroups struct {
	UserGroup  []*PolicyScopeUserGroupsUserGroup `yaml:"user_group,omitempty" xml:"user_group,omitempty"`
}

type PolicyScopeUserGroupsUserGroup struct {
	ID    uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name  string `yaml:"name,omitempty" xml:"name,omitempty"`
}

type PolicyScopeNetworkSegments struct {
	NetworkSegment  []*PolicyScopeNetworkSegmentsNetworkSegment `yaml:"network_segment,omitempty" xml:"network_segment,omitempty"`
}

type PolicyScopeNetworkSegmentsNetworkSegment struct {
	ID    uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name  string `yaml:"name,omitempty" xml:"name,omitempty"`
}

type PolicyScopeIbeacons struct {
	Ibeacon  []*PolicyScopeIbeaconsIbeacon `yaml:"ibeacon,omitempty" xml:"ibeacon,omitempty"`
}

type PolicyScopeIbeaconsIbeacon struct {
	ID    uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name  string `yaml:"name,omitempty" xml:"name,omitempty"`
}

type PolicyScopeExclusions struct {
	Computers        *PolicyScopeComputers       `yaml:"computers,omitempty" xml:"computers,omitempty"`
	ComputerGroups   *PolicyScopeComputerGroups  `yaml:"computer_groups,omitempty" xml:"computer_groups,omitempty"`
	Buildings        *PolicyScopeBuildings       `yaml:"buildings,omitempty" xml:"buildings,omitempty"`
	Departments      *PolicyScopeDepartments     `yaml:"departments,omitempty" xml:"departments,omitempty"`
	Users            *PolicyScopeUsers           `yaml:"users,omitempty" xml:"users,omitempty"`
	UserGroups       *PolicyScopeUserGroups      `yaml:"user_groups,omitempty" xml:"user_groups,omitempty"`
	NetworkSegments  *PolicyScopeNetworkSegments `yaml:"network_segments,omitempty" xml:"network_segments,omitempty"`
	Ibeacons         *PolicyScopeIbeacons        `yaml:"ibeacons,omitempty" xml:"ibeacons,omitempty"`
}

type PolicySelfService struct {
	UseForSelfService           bool                         `yaml:"use_for_self_service,omitempty" xml:"use_for_self_service,omitempty"`
	SelfServiceDisplayName      string                       `yaml:"self_service_display_name,omitempty" xml:"self_service_display_name,omitempty"`
	InstallButtonText           string                       `yaml:"install_button_text,omitempty" xml:"install_button_text,omitempty"`
	ReinstallButtonText         string                       `yaml:"reinstall_button_text,omitempty" xml:"reinstall_button_text,omitempty"`
	SelfServiceDescription      string                       `yaml:"self_service_description,omitempty" xml:"self_service_description,omitempty"`
	ForceUsersToViewDescription bool                         `yaml:"force_users_to_view_description,omitempty" xml:"force_users_to_view_description,omitempty"`
	SelfServiceIcon             *PolicySelfServiceIcon       `yaml:"self_service_icon,omitempty" xml:"self_service_icon,omitempty"`
	FeatureOnMainPage           bool                         `yaml:"feature_on_main_page,omitempty" xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories       *PolicySelfServiceCategories `yaml:"self_service_categories,omitempty" xml:"self_service_categories,omitempty"`
}

type PolicySelfServiceIcon struct {
	ID    uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name  string `yaml:"name,omitempty" xml:"name,omitempty"`
	Url	  string `yaml:"url,omitempty" xml:"url,omitempty"`
}

type PolicySelfServiceCategories struct {
	Category  *PolicySelfServiceCategory `yaml:"category,omitempty" xml:"category,omitempty"`
}

type PolicySelfServiceCategory struct {
	ID         uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name       string `yaml:"name,omitempty" xml:"name,omitempty"`
	DisplayIn  bool   `yaml:"display_in,omitempty" xml:"display_in,omitempty"`
	FeatureIn  bool   `yaml:"feature_in,omitempty" xml:"feature_in,omitempty"`
}

type PolicyPackageConfiguration struct {
	Packages  *PolicyPackages `yaml:"packages,omitempty" xml:"packages,omitempty"`
}

type PolicyPackages struct {
	Size     uint32           `yaml:"size,omitempty" xml:"size,omitempty"`
	Package  []*PolicyPackage `yaml:"package,omitempty" xml:"package,omitempty"`
}

type PolicyPackage struct {
	ID             uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name           string `yaml:"name,omitempty" xml:"name,omitempty"`
	Action         string `yaml:"action,omitempty" xml:"action,omitempty"`
	Fut            bool   `yaml:"fut,omitempty" xml:"fut,omitempty"`
	Feu            bool   `yaml:"feu,omitempty" xml:"feu,omitempty"`
	UpdateAutorun  bool   `yaml:"update_autorun,omitempty" xml:"update_autorun,omitempty"`
}

type PolicyScripts struct {
	Size          uint32          `yaml:"size,omitempty" xml:"size,omitempty"`
	PolicyScript  []*PolicyScript `yaml:"script,omitempty" xml:"script,omitempty"`
}

type PolicyScript struct {
	ID           uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name         string `yaml:"name,omitempty" xml:"name,omitempty"`
	Priority     string `yaml:"priority,omitempty" xml:"priority,omitempty"`
	Parameter4   string `yaml:"parameter4,omitempty" xml:"parameter4,omitempty"`
	Parameter5   string `yaml:"parameter5,omitempty" xml:"parameter5,omitempty"`
	Parameter6   string `yaml:"parameter6,omitempty" xml:"parameter6,omitempty"`
	Parameter7   string `yaml:"parameter7,omitempty" xml:"parameter7,omitempty"`
	Parameter8   string `yaml:"parameter8,omitempty" xml:"parameter8,omitempty"`
	Parameter9   string `yaml:"parameter9,omitempty" xml:"parameter9,omitempty"`
	Parameter10  string `yaml:"parameter10,omitempty" xml:"parameter10,omitempty"`
	Parameter11  string `yaml:"parameter11,omitempty" xml:"parameter11,omitempty"`
}

type PolicyPrinters struct {
	Size                  uint32           `yaml:"size,omitempty" xml:"size,omitempty"`
	LeaveExistingDefault  string           `yaml:"leave_existing_default,omitempty" xml:"leave_existing_default,omitempty"`
	Printer               []*PolicyPrinter `yaml:"printer,omitempty" xml:"printer,omitempty"`
}

type PolicyPrinter struct {
	ID           uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name         string `yaml:"name,omitempty" xml:"name,omitempty"`
	Action       string `yaml:"action,omitempty" xml:"action,omitempty"`
	MakeDefault  string `yaml:"make_default,omitempty" xml:"make_default,omitempty"`
}

type PolicyDockItems struct {
	Size      uint32            `yaml:"size,omitempty" xml:"size,omitempty"`
	DockItem  []*PolicyDockItem `yaml:"dock_item,omitempty" xml:"dock_item,omitempty"`
}

type PolicyDockItem struct {
	ID      uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name    string `yaml:"name,omitempty" xml:"name,omitempty"`
	Action  string `yaml:"action,omitempty" xml:"action,omitempty"`
}

type PolicyAccountMaintenance struct {
	Accounts                 *PolicyAccounts                `yaml:"accounts,omitempty" xml:"accounts,omitempty"`
	DirectoryBindings        *PolicyDirectoryBindings       `yaml:"directory_bindings,omitempty" xml:"directory_bindings,omitempty"`
	ManagementAccount        *PolicyManagementAccount       `yaml:"management_account,omitempty" xml:"management_account,omitempty"`
	OpenFirmwareEfiPassword  *PolicyOpenFirmwareEfiPassword `yaml:"open_firmware_efi_password,omitempty" xml:"open_firmware_efi_password,omitempty"`
}

type PolicyAccounts struct {
	Size    uint32           `yaml:"size,omitempty" xml:"size,omitempty"`
	Account []*PolicyAccount `yaml:"account,omitempty" xml:"account,omitempty"`
}

type PolicyAccount struct {
	Action                  string `yaml:"action,omitempty" xml:"action,omitempty"`
	UserName                string `yaml:"user_name,omitempty" xml:"user_name,omitempty"`
	RealName                string `yaml:"real_name,omitempty" xml:"real_name,omitempty"`
	Password                string `yaml:"password,omitempty" xml:"password,omitempty"`
	ArchiveHomeDirectory    bool   `yaml:"archive_home_directory,omitempty" xml:"archive_home_directory,omitempty"`
	ArchiveHomeDirectoryTo  string `yaml:"archive_home_directory_to,omitempty" xml:"archive_home_directory_to,omitempty"`
	Home                    string `yaml:"home,omitempty" xml:"home,omitempty"`
	Picture                 string `yaml:"picture,omitempty" xml:"picture,omitempty"`
	Admin                   bool   `yaml:"admin,omitempty" xml:"admin,omitempty"`
	FileVaultEnabled        bool   `yaml:"filevault_enabled,omitempty" xml:"filevault_enabled,omitempty"`
}

type PolicyDirectoryBindings struct {
	Size     uint32                    `yaml:"size,omitempty" xml:"size,omitempty"`
	Binding  []*PolicyDirectoryBinding `yaml:"binding,omitempty" xml:"binding,omitempty"`
}

type PolicyDirectoryBinding struct {
	ID    uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name  string `yaml:"name,omitempty" xml:"name,omitempty"`
}

type PolicyManagementAccount struct {
	Action                 string `yaml:"action,omitempty" xml:"action,omitempty"`
	ManagedPassword        string `yaml:"managed_password,omitempty" xml:"managed_password,omitempty"`
	ManagedPasswordLength  uint32 `yaml:"managed_password_length,omitempty" xml:"managed_password_length,omitempty"`
}

type PolicyOpenFirmwareEfiPassword struct {
	OfMode      string `yaml:"of_mode,omitempty" xml:"of_mode,omitempty"`
	OfPassword  string `yaml:"of_password,omitempty" xml:"of_password,omitempty"`
}

type PolicyRebootSettings struct {
	Message                      string `yaml:"message,omitempty" xml:"message,omitempty"`
	StartupDisk                  string `yaml:"startup_disk,omitempty" xml:"startup_disk,omitempty"`
	SpecifyStartup               string `yaml:"specify_startup,omitempty" xml:"specify_startup,omitempty"`
	NoUserLoggedIn               string `yaml:"no_user_logged_in,omitempty" xml:"no_user_logged_in,omitempty"`
	UserLoggedIn                 string `yaml:"user_logged_in,omitempty" xml:"user_logged_in,omitempty"`
	MinutesUntilReboot           int32  `yaml:"minutes_until_reboot,omitempty" xml:"minutes_until_reboot,omitempty"`
	StartRebootTimerImmediately  bool  `yaml:"start_reboot_timer_immediately,omitempty" xml:"start_reboot_timer_immediately,omitempty"`
	FileVaultReboot              bool   `yaml:"file_value_2_reboot,omitempty" xml:"file_value_2_reboot,omitempty"`
}

type PolicyMaintenance struct {
	Recon                     bool `yaml:"recon,omitempty" xml:"recon,omitempty"`
	ResetName                 bool `yaml:"reset_name,omitempty" xml:"reset_name,omitempty"`
	InstallAllCachedPackages  bool `yaml:"install_all_cached_packages,omitempty" xml:"install_all_cached_packages,omitempty"`
	Heal                      bool `yaml:"heal,omitempty" xml:"heal,omitempty"`
	Prebindings               bool `yaml:"prebindings,omitempty" xml:"prebindings,omitempty"`
	Permissions               bool `yaml:"permissions,omitempty" xml:"permissions,omitempty"`
	Byhost                    bool `yaml:"byhost,omitempty" xml:"byhost,omitempty"`
	SystemCache               bool `yaml:"system_cache,omitempty" xml:"system_cache,omitempty"`
	UserCache                 bool `yaml:"user_cache,omitempty" xml:"user_cache,omitempty"`
	Verify                    bool `yaml:"verify,omitempty" xml:"verify,omitempty"`
}

type PolicyFilesProcesses struct {
	SearchByPath          string `yaml:"search_by_path,omitempty" xml:"search_by_path,omitempty"`
	DeleteFile            bool   `yaml:"delete_file,omitempty" xml:"delete_file,omitempty"`
	LocateFile            string `yaml:"locate_file,omitempty" xml:"locate_file,omitempty"`
	UpdateLocateDatabase  bool   `yaml:"update_locate_database,omitempty" xml:"update_locate_database,omitempty"`
	SpotlightSearch       string `yaml:"spotlight_search,omitempty" xml:"spotlight_search,omitempty"`
	SearchForProcess      string `yaml:"search_for_process,omitempty" xml:"search_for_process,omitempty"`
	KillProcess           bool   `yaml:"kill_process,omitempty" xml:"kill_process,omitempty"`
	RunCommand            string `yaml:"run_command,omitempty" xml:"run_command,omitempty"`
}

type PolicyUserInteraction struct {
	MessageStart           string `yaml:"message_start,omitempty" xml:"message_start,omitempty"`
	AllowUsersToDefer      bool   `yaml:"allow_users_to_defer,omitempty" xml:"allow_users_to_defer,omitempty"`
	AllowDeferralUntilUtc  string `yaml:"allow_deferral_until_utc,omitempty" xml:"allow_deferral_until_utc,omitempty"`
	AllowDeferralMinutes   uint32 `yaml:"allow_deferral_minutes,omitempty" xml:"allow_deferral_minutes,omitempty"`
	MessageFinish          string `yaml:"message_finish,omitempty" xml:"message_finish,omitempty"`
}

type PolicyDiskEncryption struct {
	Action                                  string `yaml:"action,omitempty" xml:"action,omitempty"`
	DiskEncryptionConfigurationID           uint32 `yaml:"disk_encryption_configuration_id,omitempty" xml:"disk_encryption_configuration_id,omitempty"`
	AuthRestart                             bool   `yaml:"auth_restart,omitempty" xml:"auth_restart,omitempty"`
	RemediateKeyType                        string `yaml:"remediate_key_type,omitempty" xml:"remediate_key_type,omitempty"`
	RemediateDiskEncryptionConfigurationID  uint32 `yaml:"remediate_disk_encryption_configuration_id,omitempty" xml:"remediate_disk_encryption_configuration_id,omitempty"`
}

type GetPoliciesResult struct {
	Size    uint32           `yaml:"size,omitempty" xml:"size,omitempty"`
	Policy  []PolicyOverview `yaml:"policy,omitempty" xml:"policy,omitempty"`
}

type PolicyOverview struct {
	ID    uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
	Name  string `yaml:"name,omitempty" xml:"name,omitempty"`
}


func (c *Client) GetPolicies() (*GetPoliciesResult, error) {
	var result GetPoliciesResult

	err := c.call(APIPathPolices, http.MethodGet,
		APIVersionPolicies, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPolicy(policyID uint32) (*Policy, error) {
	var result Policy

	err := c.call(path.Join(APIPathPolices, "id", fmt.Sprint(policyID)), http.MethodGet,
		APIVersionPolicies, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}


type PolicyResult struct {
	XMLName	xml.Name	`yaml:"-" xml:"policy,omitempty"`
	ID		uint32		`yaml:"id,omitempty" xml:"id,omitempty"`
}

func (c *Client) CreatePolicy (params *Policy) (*PolicyResult, error) {
	var result PolicyResult

	err := c.call(path.Join(APIPathPolices, "id", "0"), http.MethodPost,
		APIVersionPolicies, nil, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (c *Client) UpdatePolicy (policyID uint32, params *Policy) (*PolicyResult, error) {
	var result PolicyResult

	err := c.call(path.Join(APIPathPolices, "id", fmt.Sprint(policyID)), http.MethodPut,
		APIVersionPolicies, nil, params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeletePolicy (policyID uint32) error {
	err := c.call(path.Join(APIPathPolices, "id", fmt.Sprint(policyID)), http.MethodDelete,
		APIVersionPolicies, nil, nil, nil)
	if err != nil {
		return err
	}
	fmt.Println("[jamf-pro-go] Policy (ID: ", policyID, ") is deleted")

	return nil
}