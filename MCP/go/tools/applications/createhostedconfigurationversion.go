package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-appconfig/mcp-server/config"
	"github.com/amazon-appconfig/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreatehostedconfigurationversionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		ApplicationIdVal, ok := args["ApplicationId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: ApplicationId"), nil
		}
		ApplicationId, ok := ApplicationIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: ApplicationId"), nil
		}
		ConfigurationProfileIdVal, ok := args["ConfigurationProfileId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: ConfigurationProfileId"), nil
		}
		ConfigurationProfileId, ok := ConfigurationProfileIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: ConfigurationProfileId"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CreateHostedConfigurationVersionrequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/applications/%s/configurationprofiles/%s/hostedconfigurationversions#Content-Type", cfg.BaseURL, ApplicationId, ConfigurationProfileId)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Amz-Content-Sha256"]; ok {
			req.Header.Set("X-Amz-Content-Sha256", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Date"]; ok {
			req.Header.Set("X-Amz-Date", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Algorithm"]; ok {
			req.Header.Set("X-Amz-Algorithm", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Credential"]; ok {
			req.Header.Set("X-Amz-Credential", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Security-Token"]; ok {
			req.Header.Set("X-Amz-Security-Token", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Signature"]; ok {
			req.Header.Set("X-Amz-Signature", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-SignedHeaders"]; ok {
			req.Header.Set("X-Amz-SignedHeaders", fmt.Sprintf("%v", val))
		}
		if val, ok := args["Description"]; ok {
			req.Header.Set("Description", fmt.Sprintf("%v", val))
		}
		if val, ok := args["Content-Type"]; ok {
			req.Header.Set("Content-Type", fmt.Sprintf("%v", val))
		}
		if val, ok := args["Latest-Version-Number"]; ok {
			req.Header.Set("Latest-Version-Number", fmt.Sprintf("%v", val))
		}
		if val, ok := args["VersionLabel"]; ok {
			req.Header.Set("VersionLabel", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.HostedConfigurationVersion
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCreatehostedconfigurationversionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_applications_ApplicationId_configurationprofiles_ConfigurationProfileId_hostedconfigurationversions#Content-Type",
		mcp.WithDescription("Creates a new configuration in the AppConfig hosted configuration store."),
		mcp.WithString("X-Amz-Content-Sha256", mcp.Description("")),
		mcp.WithString("X-Amz-Date", mcp.Description("")),
		mcp.WithString("X-Amz-Algorithm", mcp.Description("")),
		mcp.WithString("X-Amz-Credential", mcp.Description("")),
		mcp.WithString("X-Amz-Security-Token", mcp.Description("")),
		mcp.WithString("X-Amz-Signature", mcp.Description("")),
		mcp.WithString("X-Amz-SignedHeaders", mcp.Description("")),
		mcp.WithString("ApplicationId", mcp.Required(), mcp.Description("The application ID.")),
		mcp.WithString("ConfigurationProfileId", mcp.Required(), mcp.Description("The configuration profile ID.")),
		mcp.WithString("Description", mcp.Description("A description of the configuration.")),
		mcp.WithString("Content-Type", mcp.Required(), mcp.Description("A standard MIME type describing the format of the configuration content. For more information, see <a href=\"https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17\">Content-Type</a>.")),
		mcp.WithNumber("Latest-Version-Number", mcp.Description("An optional locking token used to prevent race conditions from overwriting configuration updates when creating a new version. To ensure your data is not overwritten when creating multiple hosted configuration versions in rapid succession, specify the version number of the latest hosted configuration version.")),
		mcp.WithString("VersionLabel", mcp.Description("An optional, user-defined label for the AppConfig hosted configuration version. This value must contain at least one non-numeric character. For example, \"v2.2.0\".")),
		mcp.WithString("Content", mcp.Required(), mcp.Description("Input parameter: The content of the configuration or the configuration data.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatehostedconfigurationversionHandler(cfg),
	}
}
