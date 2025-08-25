package main

import (
	"github.com/amazon-appconfig/mcp-server/config"
	"github.com/amazon-appconfig/mcp-server/models"
	tools_applications "github.com/amazon-appconfig/mcp-server/tools/applications"
	tools_extensions "github.com/amazon-appconfig/mcp-server/tools/extensions"
	tools_tags "github.com/amazon-appconfig/mcp-server/tools/tags"
	tools_extensionassociations "github.com/amazon-appconfig/mcp-server/tools/extensionassociations"
	tools_deploymentstrategies "github.com/amazon-appconfig/mcp-server/tools/deploymentstrategies"
	tools_deployementstrategies "github.com/amazon-appconfig/mcp-server/tools/deployementstrategies"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_applications.CreateListenvironmentsTool(cfg),
		tools_applications.CreateCreateenvironmentTool(cfg),
		tools_extensions.CreateUpdateextensionTool(cfg),
		tools_extensions.CreateDeleteextensionTool(cfg),
		tools_extensions.CreateGetextensionTool(cfg),
		tools_applications.CreateDeleteapplicationTool(cfg),
		tools_applications.CreateGetapplicationTool(cfg),
		tools_applications.CreateUpdateapplicationTool(cfg),
		tools_applications.CreateCreatehostedconfigurationversionTool(cfg),
		tools_applications.CreateValidateconfigurationTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_applications.CreateCreateapplicationTool(cfg),
		tools_applications.CreateListapplicationsTool(cfg),
		tools_applications.CreateListconfigurationprofilesTool(cfg),
		tools_applications.CreateCreateconfigurationprofileTool(cfg),
		tools_applications.CreateListdeploymentsTool(cfg),
		tools_applications.CreateStartdeploymentTool(cfg),
		tools_extensionassociations.CreateListextensionassociationsTool(cfg),
		tools_extensionassociations.CreateCreateextensionassociationTool(cfg),
		tools_applications.CreateUpdateconfigurationprofileTool(cfg),
		tools_applications.CreateDeleteconfigurationprofileTool(cfg),
		tools_applications.CreateGetconfigurationprofileTool(cfg),
		tools_applications.CreateStopdeploymentTool(cfg),
		tools_applications.CreateGetdeploymentTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_deploymentstrategies.CreateGetdeploymentstrategyTool(cfg),
		tools_deploymentstrategies.CreateUpdatedeploymentstrategyTool(cfg),
		tools_applications.CreateDeleteenvironmentTool(cfg),
		tools_applications.CreateGetenvironmentTool(cfg),
		tools_applications.CreateUpdateenvironmentTool(cfg),
		tools_applications.CreateDeletehostedconfigurationversionTool(cfg),
		tools_applications.CreateGethostedconfigurationversionTool(cfg),
		tools_deployementstrategies.CreateDeletedeploymentstrategyTool(cfg),
		tools_applications.CreateGetconfigurationTool(cfg),
		tools_deploymentstrategies.CreateListdeploymentstrategiesTool(cfg),
		tools_deploymentstrategies.CreateCreatedeploymentstrategyTool(cfg),
		tools_extensionassociations.CreateDeleteextensionassociationTool(cfg),
		tools_extensionassociations.CreateGetextensionassociationTool(cfg),
		tools_extensionassociations.CreateUpdateextensionassociationTool(cfg),
		tools_extensions.CreateListextensionsTool(cfg),
		tools_extensions.CreateCreateextensionTool(cfg),
		tools_applications.CreateListhostedconfigurationversionsTool(cfg),
	}
}
