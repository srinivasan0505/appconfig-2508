package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// CreateApplicationrequest represents the CreateApplicationrequest schema from the OpenAPI specification
type CreateApplicationrequest struct {
	Description string `json:"Description,omitempty"` // A description of the application.
	Name string `json:"Name"` // A name for the application.
	Tags map[string]interface{} `json:"Tags,omitempty"` // Metadata to assign to the application. Tags help organize and categorize your AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
}

// CreateEnvironmentrequest represents the CreateEnvironmentrequest schema from the OpenAPI specification
type CreateEnvironmentrequest struct {
	Description string `json:"Description,omitempty"` // A description of the environment.
	Monitors []Monitor `json:"Monitors,omitempty"` // Amazon CloudWatch alarms to monitor during the deployment process.
	Name string `json:"Name"` // A name for the environment.
	Tags map[string]interface{} `json:"Tags,omitempty"` // Metadata to assign to the environment. Tags help organize and categorize your AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
}

// Parameter represents the Parameter schema from the OpenAPI specification
type Parameter struct {
	Description interface{} `json:"Description,omitempty"`
	Required interface{} `json:"Required,omitempty"`
}

// GetDeploymentRequest represents the GetDeploymentRequest schema from the OpenAPI specification
type GetDeploymentRequest struct {
}

// CreateExtensionAssociationRequest represents the CreateExtensionAssociationRequest schema from the OpenAPI specification
type CreateExtensionAssociationRequest struct {
	Extensionidentifier interface{} `json:"ExtensionIdentifier"`
	Extensionversionnumber interface{} `json:"ExtensionVersionNumber,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Resourceidentifier interface{} `json:"ResourceIdentifier"`
	Tags interface{} `json:"Tags,omitempty"`
}

// CreateApplicationRequest represents the CreateApplicationRequest schema from the OpenAPI specification
type CreateApplicationRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DeleteDeploymentStrategyRequest represents the DeleteDeploymentStrategyRequest schema from the OpenAPI specification
type DeleteDeploymentStrategyRequest struct {
}

// DeleteExtensionAssociationRequest represents the DeleteExtensionAssociationRequest schema from the OpenAPI specification
type DeleteExtensionAssociationRequest struct {
}

// Deployment represents the Deployment schema from the OpenAPI specification
type Deployment struct {
	Kmskeyidentifier interface{} `json:"KmsKeyIdentifier,omitempty"`
	Applicationid interface{} `json:"ApplicationId,omitempty"`
	Configurationlocationuri interface{} `json:"ConfigurationLocationUri,omitempty"`
	Configurationprofileid interface{} `json:"ConfigurationProfileId,omitempty"`
	Deploymentnumber interface{} `json:"DeploymentNumber,omitempty"`
	Finalbaketimeinminutes interface{} `json:"FinalBakeTimeInMinutes,omitempty"`
	Configurationname interface{} `json:"ConfigurationName,omitempty"`
	Deploymentdurationinminutes interface{} `json:"DeploymentDurationInMinutes,omitempty"`
	Environmentid interface{} `json:"EnvironmentId,omitempty"`
	Growthtype interface{} `json:"GrowthType,omitempty"`
	Kmskeyarn interface{} `json:"KmsKeyArn,omitempty"`
	Percentagecomplete interface{} `json:"PercentageComplete,omitempty"`
	Deploymentstrategyid interface{} `json:"DeploymentStrategyId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	State interface{} `json:"State,omitempty"`
	Appliedextensions interface{} `json:"AppliedExtensions,omitempty"`
	Startedat interface{} `json:"StartedAt,omitempty"`
	Eventlog interface{} `json:"EventLog,omitempty"`
	Growthfactor interface{} `json:"GrowthFactor,omitempty"`
	Completedat interface{} `json:"CompletedAt,omitempty"`
	Configurationversion interface{} `json:"ConfigurationVersion,omitempty"`
}

// GetHostedConfigurationVersionRequest represents the GetHostedConfigurationVersionRequest schema from the OpenAPI specification
type GetHostedConfigurationVersionRequest struct {
}

// ListDeploymentStrategiesRequest represents the ListDeploymentStrategiesRequest schema from the OpenAPI specification
type ListDeploymentStrategiesRequest struct {
}

// DeleteEnvironmentRequest represents the DeleteEnvironmentRequest schema from the OpenAPI specification
type DeleteEnvironmentRequest struct {
}

// ActionInvocation represents the ActionInvocation schema from the OpenAPI specification
type ActionInvocation struct {
	Uri interface{} `json:"Uri,omitempty"`
	Actionname interface{} `json:"ActionName,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Extensionidentifier interface{} `json:"ExtensionIdentifier,omitempty"`
	Invocationid interface{} `json:"InvocationId,omitempty"`
	Rolearn interface{} `json:"RoleArn,omitempty"`
}

// DeploymentEvent represents the DeploymentEvent schema from the OpenAPI specification
type DeploymentEvent struct {
	Eventtype interface{} `json:"EventType,omitempty"`
	Occurredat interface{} `json:"OccurredAt,omitempty"`
	Triggeredby interface{} `json:"TriggeredBy,omitempty"`
	Actioninvocations interface{} `json:"ActionInvocations,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// CreateHostedConfigurationVersionrequest represents the CreateHostedConfigurationVersionrequest schema from the OpenAPI specification
type CreateHostedConfigurationVersionrequest struct {
	Content string `json:"Content"` // The content of the configuration or the configuration data.
}

// Extensions represents the Extensions schema from the OpenAPI specification
type Extensions struct {
	Items interface{} `json:"Items,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteExtensionRequest represents the DeleteExtensionRequest schema from the OpenAPI specification
type DeleteExtensionRequest struct {
}

// CreateDeploymentStrategyrequest represents the CreateDeploymentStrategyrequest schema from the OpenAPI specification
type CreateDeploymentStrategyrequest struct {
	Name string `json:"Name"` // A name for the deployment strategy.
	Replicateto string `json:"ReplicateTo,omitempty"` // Save the deployment strategy to a Systems Manager (SSM) document.
	Tags map[string]interface{} `json:"Tags,omitempty"` // Metadata to assign to the deployment strategy. Tags help organize and categorize your AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
	Deploymentdurationinminutes int `json:"DeploymentDurationInMinutes"` // Total amount of time for a deployment to last.
	Description string `json:"Description,omitempty"` // A description of the deployment strategy.
	Finalbaketimeinminutes int `json:"FinalBakeTimeInMinutes,omitempty"` // Specifies the amount of time AppConfig monitors for Amazon CloudWatch alarms after the configuration has been deployed to 100% of its targets, before considering the deployment to be complete. If an alarm is triggered during this time, AppConfig rolls back the deployment. You must configure permissions for AppConfig to roll back based on CloudWatch alarms. For more information, see <a href="https://docs.aws.amazon.com/appconfig/latest/userguide/getting-started-with-appconfig-cloudwatch-alarms-permissions.html">Configuring permissions for rollback based on Amazon CloudWatch alarms</a> in the <i>AppConfig User Guide</i>.
	Growthfactor float32 `json:"GrowthFactor"` // The percentage of targets to receive a deployed configuration during each interval.
	Growthtype string `json:"GrowthType,omitempty"` // <p>The algorithm used to define how percentage grows over time. AppConfig supports the following growth types:</p> <p> <b>Linear</b>: For this type, AppConfig processes the deployment by dividing the total number of targets by the value specified for <code>Step percentage</code>. For example, a linear deployment that uses a <code>Step percentage</code> of 10 deploys the configuration to 10 percent of the hosts. After those deployments are complete, the system deploys the configuration to the next 10 percent. This continues until 100% of the targets have successfully received the configuration.</p> <p> <b>Exponential</b>: For this type, AppConfig processes the deployment exponentially using the following formula: <code>G*(2^N)</code>. In this formula, <code>G</code> is the growth factor specified by the user and <code>N</code> is the number of steps until the configuration is deployed to all targets. For example, if you specify a growth factor of 2, then the system rolls out the configuration as follows:</p> <p> <code>2*(2^0)</code> </p> <p> <code>2*(2^1)</code> </p> <p> <code>2*(2^2)</code> </p> <p>Expressed numerically, the deployment rolls out as follows: 2% of the targets, 4% of the targets, 8% of the targets, and continues until the configuration has been deployed to all targets.</p>
}

// CreateHostedConfigurationVersionRequest represents the CreateHostedConfigurationVersionRequest schema from the OpenAPI specification
type CreateHostedConfigurationVersionRequest struct {
	Content interface{} `json:"Content"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// HostedConfigurationVersions represents the HostedConfigurationVersions schema from the OpenAPI specification
type HostedConfigurationVersions struct {
	Items interface{} `json:"Items,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ConfigurationProfile represents the ConfigurationProfile schema from the OpenAPI specification
type ConfigurationProfile struct {
	TypeField interface{} `json:"Type,omitempty"`
	Validators interface{} `json:"Validators,omitempty"`
	Applicationid interface{} `json:"ApplicationId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Locationuri interface{} `json:"LocationUri,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Retrievalrolearn interface{} `json:"RetrievalRoleArn,omitempty"`
}

// GetConfigurationProfileRequest represents the GetConfigurationProfileRequest schema from the OpenAPI specification
type GetConfigurationProfileRequest struct {
}

// GetApplicationRequest represents the GetApplicationRequest schema from the OpenAPI specification
type GetApplicationRequest struct {
}

// ListHostedConfigurationVersionsRequest represents the ListHostedConfigurationVersionsRequest schema from the OpenAPI specification
type ListHostedConfigurationVersionsRequest struct {
}

// Monitor represents the Monitor schema from the OpenAPI specification
type Monitor struct {
	Alarmrolearn interface{} `json:"AlarmRoleArn,omitempty"`
	Alarmarn interface{} `json:"AlarmArn"`
}

// ExtensionAssociationSummary represents the ExtensionAssociationSummary schema from the OpenAPI specification
type ExtensionAssociationSummary struct {
	Extensionarn interface{} `json:"ExtensionArn,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Resourcearn interface{} `json:"ResourceArn,omitempty"`
}

// UpdateExtensionAssociationRequest represents the UpdateExtensionAssociationRequest schema from the OpenAPI specification
type UpdateExtensionAssociationRequest struct {
	Parameters interface{} `json:"Parameters,omitempty"`
}

// GetExtensionRequest represents the GetExtensionRequest schema from the OpenAPI specification
type GetExtensionRequest struct {
}

// CreateExtensionAssociationrequest represents the CreateExtensionAssociationrequest schema from the OpenAPI specification
type CreateExtensionAssociationrequest struct {
	Extensionversionnumber int `json:"ExtensionVersionNumber,omitempty"` // The version number of the extension. If not specified, AppConfig uses the maximum version of the extension.
	Parameters map[string]interface{} `json:"Parameters,omitempty"` // The parameter names and values defined in the extensions. Extension parameters marked <code>Required</code> must be entered for this field.
	Resourceidentifier string `json:"ResourceIdentifier"` // The ARN of an application, configuration profile, or environment.
	Tags map[string]interface{} `json:"Tags,omitempty"` // Adds one or more tags for the specified extension association. Tags are metadata that help you categorize resources in different ways, for example, by purpose, owner, or environment. Each tag consists of a key and an optional value, both of which you define.
	Extensionidentifier string `json:"ExtensionIdentifier"` // The name, the ID, or the Amazon Resource Name (ARN) of the extension.
}

// UpdateConfigurationProfileRequest represents the UpdateConfigurationProfileRequest schema from the OpenAPI specification
type UpdateConfigurationProfileRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Retrievalrolearn interface{} `json:"RetrievalRoleArn,omitempty"`
	Validators interface{} `json:"Validators,omitempty"`
}

// CreateDeploymentStrategyRequest represents the CreateDeploymentStrategyRequest schema from the OpenAPI specification
type CreateDeploymentStrategyRequest struct {
	Growthtype interface{} `json:"GrowthType,omitempty"`
	Name interface{} `json:"Name"`
	Replicateto interface{} `json:"ReplicateTo,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Deploymentdurationinminutes interface{} `json:"DeploymentDurationInMinutes"`
	Description interface{} `json:"Description,omitempty"`
	Finalbaketimeinminutes interface{} `json:"FinalBakeTimeInMinutes,omitempty"`
	Growthfactor interface{} `json:"GrowthFactor"`
}

// Extension represents the Extension schema from the OpenAPI specification
type Extension struct {
	Arn interface{} `json:"Arn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Versionnumber interface{} `json:"VersionNumber,omitempty"`
	Actions interface{} `json:"Actions,omitempty"`
}

// ListExtensionsRequest represents the ListExtensionsRequest schema from the OpenAPI specification
type ListExtensionsRequest struct {
}

// ListApplicationsRequest represents the ListApplicationsRequest schema from the OpenAPI specification
type ListApplicationsRequest struct {
}

// CreateConfigurationProfilerequest represents the CreateConfigurationProfilerequest schema from the OpenAPI specification
type CreateConfigurationProfilerequest struct {
	Name string `json:"Name"` // A name for the configuration profile.
	Retrievalrolearn string `json:"RetrievalRoleArn,omitempty"` // <p>The ARN of an IAM role with permission to access the configuration at the specified <code>LocationUri</code>.</p> <important> <p>A retrieval role ARN is not required for configurations stored in the AppConfig hosted configuration store. It is required for all other sources that store your configuration. </p> </important>
	Tags map[string]interface{} `json:"Tags,omitempty"` // Metadata to assign to the configuration profile. Tags help organize and categorize your AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
	TypeField string `json:"Type,omitempty"` // <p>The type of configurations contained in the profile. AppConfig supports <code>feature flags</code> and <code>freeform</code> configurations. We recommend you create feature flag configurations to enable or disable new features and freeform configurations to distribute configurations to an application. When calling this API, enter one of the following values for <code>Type</code>:</p> <p> <code>AWS.AppConfig.FeatureFlags</code> </p> <p> <code>AWS.Freeform</code> </p>
	Validators []Validator `json:"Validators,omitempty"` // A list of methods for validating the configuration.
	Description string `json:"Description,omitempty"` // A description of the configuration profile.
	Locationuri string `json:"LocationUri"` // <p>A URI to locate the configuration. You can specify the following:</p> <ul> <li> <p>For the AppConfig hosted configuration store and for feature flags, specify <code>hosted</code>.</p> </li> <li> <p>For an Amazon Web Services Systems Manager Parameter Store parameter, specify either the parameter name in the format <code>ssm-parameter://&lt;parameter name&gt;</code> or the ARN.</p> </li> <li> <p>For an Secrets Manager secret, specify the URI in the following format: <code>secrets-manager</code>://&lt;secret name&gt;.</p> </li> <li> <p>For an Amazon S3 object, specify the URI in the following format: <code>s3://&lt;bucket&gt;/&lt;objectKey&gt; </code>. Here is an example: <code>s3://my-bucket/my-app/us-east-1/my-config.json</code> </p> </li> <li> <p>For an SSM document, specify either the document name in the format <code>ssm-document://&lt;document name&gt;</code> or the Amazon Resource Name (ARN).</p> </li> </ul>
}

// UpdateDeploymentStrategyrequest represents the UpdateDeploymentStrategyrequest schema from the OpenAPI specification
type UpdateDeploymentStrategyrequest struct {
	Description string `json:"Description,omitempty"` // A description of the deployment strategy.
	Finalbaketimeinminutes int `json:"FinalBakeTimeInMinutes,omitempty"` // The amount of time that AppConfig monitors for alarms before considering the deployment to be complete and no longer eligible for automatic rollback.
	Growthfactor float32 `json:"GrowthFactor,omitempty"` // The percentage of targets to receive a deployed configuration during each interval.
	Growthtype string `json:"GrowthType,omitempty"` // <p>The algorithm used to define how percentage grows over time. AppConfig supports the following growth types:</p> <p> <b>Linear</b>: For this type, AppConfig processes the deployment by increments of the growth factor evenly distributed over the deployment time. For example, a linear deployment that uses a growth factor of 20 initially makes the configuration available to 20 percent of the targets. After 1/5th of the deployment time has passed, the system updates the percentage to 40 percent. This continues until 100% of the targets are set to receive the deployed configuration.</p> <p> <b>Exponential</b>: For this type, AppConfig processes the deployment exponentially using the following formula: <code>G*(2^N)</code>. In this formula, <code>G</code> is the growth factor specified by the user and <code>N</code> is the number of steps until the configuration is deployed to all targets. For example, if you specify a growth factor of 2, then the system rolls out the configuration as follows:</p> <p> <code>2*(2^0)</code> </p> <p> <code>2*(2^1)</code> </p> <p> <code>2*(2^2)</code> </p> <p>Expressed numerically, the deployment rolls out as follows: 2% of the targets, 4% of the targets, 8% of the targets, and continues until the configuration has been deployed to all targets.</p>
	Deploymentdurationinminutes int `json:"DeploymentDurationInMinutes,omitempty"` // Total amount of time for a deployment to last.
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// UpdateApplicationrequest represents the UpdateApplicationrequest schema from the OpenAPI specification
type UpdateApplicationrequest struct {
	Description string `json:"Description,omitempty"` // A description of the application.
	Name string `json:"Name,omitempty"` // The name of the application.
}

// UpdateApplicationRequest represents the UpdateApplicationRequest schema from the OpenAPI specification
type UpdateApplicationRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// ConfigurationProfileSummary represents the ConfigurationProfileSummary schema from the OpenAPI specification
type ConfigurationProfileSummary struct {
	Applicationid interface{} `json:"ApplicationId,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Locationuri interface{} `json:"LocationUri,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Validatortypes interface{} `json:"ValidatorTypes,omitempty"`
}

// Applications represents the Applications schema from the OpenAPI specification
type Applications struct {
	Items interface{} `json:"Items,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"Tags"`
}

// ExtensionSummary represents the ExtensionSummary schema from the OpenAPI specification
type ExtensionSummary struct {
	Arn interface{} `json:"Arn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Versionnumber interface{} `json:"VersionNumber,omitempty"`
}

// DeleteHostedConfigurationVersionRequest represents the DeleteHostedConfigurationVersionRequest schema from the OpenAPI specification
type DeleteHostedConfigurationVersionRequest struct {
}

// ResourceTags represents the ResourceTags schema from the OpenAPI specification
type ResourceTags struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// UpdateExtensionrequest represents the UpdateExtensionrequest schema from the OpenAPI specification
type UpdateExtensionrequest struct {
	Actions map[string]interface{} `json:"Actions,omitempty"` // The actions defined in the extension.
	Description string `json:"Description,omitempty"` // Information about the extension.
	Parameters map[string]interface{} `json:"Parameters,omitempty"` // One or more parameters for the actions called by the extension.
	Versionnumber int `json:"VersionNumber,omitempty"` // The extension version number.
}

// DeploymentStrategy represents the DeploymentStrategy schema from the OpenAPI specification
type DeploymentStrategy struct {
	Description interface{} `json:"Description,omitempty"`
	Finalbaketimeinminutes interface{} `json:"FinalBakeTimeInMinutes,omitempty"`
	Growthfactor interface{} `json:"GrowthFactor,omitempty"`
	Growthtype interface{} `json:"GrowthType,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Replicateto interface{} `json:"ReplicateTo,omitempty"`
	Deploymentdurationinminutes interface{} `json:"DeploymentDurationInMinutes,omitempty"`
}

// CreateExtensionRequest represents the CreateExtensionRequest schema from the OpenAPI specification
type CreateExtensionRequest struct {
	Name interface{} `json:"Name"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Actions interface{} `json:"Actions"`
	Description interface{} `json:"Description,omitempty"`
}

// Deployments represents the Deployments schema from the OpenAPI specification
type Deployments struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Items interface{} `json:"Items,omitempty"`
}

// ExtensionAssociation represents the ExtensionAssociation schema from the OpenAPI specification
type ExtensionAssociation struct {
	Id interface{} `json:"Id,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Resourcearn interface{} `json:"ResourceArn,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Extensionarn interface{} `json:"ExtensionArn,omitempty"`
	Extensionversionnumber interface{} `json:"ExtensionVersionNumber,omitempty"`
}

// TagResourcerequest represents the TagResourcerequest schema from the OpenAPI specification
type TagResourcerequest struct {
	Tags map[string]interface{} `json:"Tags"` // The key-value string map. The valid character set is [a-zA-Z+-=._:/]. The tag key can be up to 128 characters and must not start with <code>aws:</code>. The tag value can be up to 256 characters.
}

// HostedConfigurationVersionSummary represents the HostedConfigurationVersionSummary schema from the OpenAPI specification
type HostedConfigurationVersionSummary struct {
	Versionnumber interface{} `json:"VersionNumber,omitempty"`
	Applicationid interface{} `json:"ApplicationId,omitempty"`
	Configurationprofileid interface{} `json:"ConfigurationProfileId,omitempty"`
	Contenttype interface{} `json:"ContentType,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Versionlabel interface{} `json:"VersionLabel,omitempty"`
}

// Environments represents the Environments schema from the OpenAPI specification
type Environments struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Items interface{} `json:"Items,omitempty"`
}

// HostedConfigurationVersion represents the HostedConfigurationVersion schema from the OpenAPI specification
type HostedConfigurationVersion struct {
	Content interface{} `json:"Content,omitempty"`
}

// GetEnvironmentRequest represents the GetEnvironmentRequest schema from the OpenAPI specification
type GetEnvironmentRequest struct {
}

// UpdateExtensionAssociationrequest represents the UpdateExtensionAssociationrequest schema from the OpenAPI specification
type UpdateExtensionAssociationrequest struct {
	Parameters map[string]interface{} `json:"Parameters,omitempty"` // The parameter names and values defined in the extension.
}

// ListExtensionAssociationsRequest represents the ListExtensionAssociationsRequest schema from the OpenAPI specification
type ListExtensionAssociationsRequest struct {
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// ValidateConfigurationRequest represents the ValidateConfigurationRequest schema from the OpenAPI specification
type ValidateConfigurationRequest struct {
}

// UpdateConfigurationProfilerequest represents the UpdateConfigurationProfilerequest schema from the OpenAPI specification
type UpdateConfigurationProfilerequest struct {
	Description string `json:"Description,omitempty"` // A description of the configuration profile.
	Name string `json:"Name,omitempty"` // The name of the configuration profile.
	Retrievalrolearn string `json:"RetrievalRoleArn,omitempty"` // The ARN of an IAM role with permission to access the configuration at the specified <code>LocationUri</code>.
	Validators []Validator `json:"Validators,omitempty"` // A list of methods for validating the configuration.
}

// ConfigurationProfiles represents the ConfigurationProfiles schema from the OpenAPI specification
type ConfigurationProfiles struct {
	Items interface{} `json:"Items,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListDeploymentsRequest represents the ListDeploymentsRequest schema from the OpenAPI specification
type ListDeploymentsRequest struct {
}

// UpdateEnvironmentrequest represents the UpdateEnvironmentrequest schema from the OpenAPI specification
type UpdateEnvironmentrequest struct {
	Name string `json:"Name,omitempty"` // The name of the environment.
	Description string `json:"Description,omitempty"` // A description of the environment.
	Monitors []Monitor `json:"Monitors,omitempty"` // Amazon CloudWatch alarms to monitor during the deployment process.
}

// AppliedExtension represents the AppliedExtension schema from the OpenAPI specification
type AppliedExtension struct {
	Extensionassociationid interface{} `json:"ExtensionAssociationId,omitempty"`
	Extensionid interface{} `json:"ExtensionId,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Versionnumber interface{} `json:"VersionNumber,omitempty"`
}

// UpdateEnvironmentRequest represents the UpdateEnvironmentRequest schema from the OpenAPI specification
type UpdateEnvironmentRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Monitors interface{} `json:"Monitors,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// Validator represents the Validator schema from the OpenAPI specification
type Validator struct {
	TypeField interface{} `json:"Type"`
	Content interface{} `json:"Content"`
}

// DeploymentStrategies represents the DeploymentStrategies schema from the OpenAPI specification
type DeploymentStrategies struct {
	Items interface{} `json:"Items,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListEnvironmentsRequest represents the ListEnvironmentsRequest schema from the OpenAPI specification
type ListEnvironmentsRequest struct {
}

// ExtensionAssociations represents the ExtensionAssociations schema from the OpenAPI specification
type ExtensionAssociations struct {
	Items interface{} `json:"Items,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateDeploymentStrategyRequest represents the UpdateDeploymentStrategyRequest schema from the OpenAPI specification
type UpdateDeploymentStrategyRequest struct {
	Growthfactor interface{} `json:"GrowthFactor,omitempty"`
	Growthtype interface{} `json:"GrowthType,omitempty"`
	Deploymentdurationinminutes interface{} `json:"DeploymentDurationInMinutes,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Finalbaketimeinminutes interface{} `json:"FinalBakeTimeInMinutes,omitempty"`
}

// ActionsMap represents the ActionsMap schema from the OpenAPI specification
type ActionsMap struct {
}

// Configuration represents the Configuration schema from the OpenAPI specification
type Configuration struct {
	Content interface{} `json:"Content,omitempty"`
}

// Application represents the Application schema from the OpenAPI specification
type Application struct {
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// StartDeploymentRequest represents the StartDeploymentRequest schema from the OpenAPI specification
type StartDeploymentRequest struct {
	Kmskeyidentifier interface{} `json:"KmsKeyIdentifier,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Configurationprofileid interface{} `json:"ConfigurationProfileId"`
	Configurationversion interface{} `json:"ConfigurationVersion"`
	Deploymentstrategyid interface{} `json:"DeploymentStrategyId"`
	Description interface{} `json:"Description,omitempty"`
}

// DeleteApplicationRequest represents the DeleteApplicationRequest schema from the OpenAPI specification
type DeleteApplicationRequest struct {
}

// Environment represents the Environment schema from the OpenAPI specification
type Environment struct {
	Description interface{} `json:"Description,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Monitors interface{} `json:"Monitors,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	State interface{} `json:"State,omitempty"`
	Applicationid interface{} `json:"ApplicationId,omitempty"`
}

// GetConfigurationRequest represents the GetConfigurationRequest schema from the OpenAPI specification
type GetConfigurationRequest struct {
}

// CreateEnvironmentRequest represents the CreateEnvironmentRequest schema from the OpenAPI specification
type CreateEnvironmentRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Monitors interface{} `json:"Monitors,omitempty"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ParameterMap represents the ParameterMap schema from the OpenAPI specification
type ParameterMap struct {
}

// StopDeploymentRequest represents the StopDeploymentRequest schema from the OpenAPI specification
type StopDeploymentRequest struct {
}

// CreateConfigurationProfileRequest represents the CreateConfigurationProfileRequest schema from the OpenAPI specification
type CreateConfigurationProfileRequest struct {
	Name interface{} `json:"Name"`
	Retrievalrolearn interface{} `json:"RetrievalRoleArn,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Validators interface{} `json:"Validators,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Locationuri interface{} `json:"LocationUri"`
}

// UpdateExtensionRequest represents the UpdateExtensionRequest schema from the OpenAPI specification
type UpdateExtensionRequest struct {
	Actions interface{} `json:"Actions,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Versionnumber interface{} `json:"VersionNumber,omitempty"`
}

// Action represents the Action schema from the OpenAPI specification
type Action struct {
	Rolearn interface{} `json:"RoleArn,omitempty"`
	Uri interface{} `json:"Uri,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// GetDeploymentStrategyRequest represents the GetDeploymentStrategyRequest schema from the OpenAPI specification
type GetDeploymentStrategyRequest struct {
}

// ListConfigurationProfilesRequest represents the ListConfigurationProfilesRequest schema from the OpenAPI specification
type ListConfigurationProfilesRequest struct {
}

// CreateExtensionrequest represents the CreateExtensionrequest schema from the OpenAPI specification
type CreateExtensionrequest struct {
	Actions map[string]interface{} `json:"Actions"` // The actions defined in the extension.
	Description string `json:"Description,omitempty"` // Information about the extension.
	Name string `json:"Name"` // A name for the extension. Each extension name in your account must be unique. Extension versions use the same name.
	Parameters map[string]interface{} `json:"Parameters,omitempty"` // The parameters accepted by the extension. You specify parameter values when you associate the extension to an AppConfig resource by using the <code>CreateExtensionAssociation</code> API action. For Lambda extension actions, these parameters are included in the Lambda request object.
	Tags map[string]interface{} `json:"Tags,omitempty"` // Adds one or more tags for the specified extension. Tags are metadata that help you categorize resources in different ways, for example, by purpose, owner, or environment. Each tag consists of a key and an optional value, both of which you define.
}

// DeleteConfigurationProfileRequest represents the DeleteConfigurationProfileRequest schema from the OpenAPI specification
type DeleteConfigurationProfileRequest struct {
}

// ParameterValueMap represents the ParameterValueMap schema from the OpenAPI specification
type ParameterValueMap struct {
}

// StartDeploymentrequest represents the StartDeploymentrequest schema from the OpenAPI specification
type StartDeploymentrequest struct {
	Tags map[string]interface{} `json:"Tags,omitempty"` // Metadata to assign to the deployment. Tags help organize and categorize your AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
	Configurationprofileid string `json:"ConfigurationProfileId"` // The configuration profile ID.
	Configurationversion string `json:"ConfigurationVersion"` // The configuration version to deploy. If deploying an AppConfig hosted configuration version, you can specify either the version number or version label.
	Deploymentstrategyid string `json:"DeploymentStrategyId"` // The deployment strategy ID.
	Description string `json:"Description,omitempty"` // A description of the deployment.
	Kmskeyidentifier string `json:"KmsKeyIdentifier,omitempty"` // The KMS key identifier (key ID, key alias, or key ARN). AppConfig uses this ID to encrypt the configuration data using a customer managed key.
}

// DeploymentSummary represents the DeploymentSummary schema from the OpenAPI specification
type DeploymentSummary struct {
	Growthfactor interface{} `json:"GrowthFactor,omitempty"`
	Startedat interface{} `json:"StartedAt,omitempty"`
	Completedat interface{} `json:"CompletedAt,omitempty"`
	Growthtype interface{} `json:"GrowthType,omitempty"`
	Configurationname interface{} `json:"ConfigurationName,omitempty"`
	Deploymentdurationinminutes interface{} `json:"DeploymentDurationInMinutes,omitempty"`
	Deploymentnumber interface{} `json:"DeploymentNumber,omitempty"`
	Finalbaketimeinminutes interface{} `json:"FinalBakeTimeInMinutes,omitempty"`
	Percentagecomplete interface{} `json:"PercentageComplete,omitempty"`
	State interface{} `json:"State,omitempty"`
	Configurationversion interface{} `json:"ConfigurationVersion,omitempty"`
}

// GetExtensionAssociationRequest represents the GetExtensionAssociationRequest schema from the OpenAPI specification
type GetExtensionAssociationRequest struct {
}
