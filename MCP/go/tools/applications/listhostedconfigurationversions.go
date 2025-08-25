package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-appconfig/mcp-server/config"
	"github.com/amazon-appconfig/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListhostedconfigurationversionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["max_results"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("max_results=%v", val))
		}
		if val, ok := args["next_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("next_token=%v", val))
		}
		if val, ok := args["version_label"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("version_label=%v", val))
		}
		if val, ok := args["MaxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("MaxResults=%v", val))
		}
		if val, ok := args["NextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("NextToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/applications/%s/configurationprofiles/%s/hostedconfigurationversions%s", cfg.BaseURL, ApplicationId, ConfigurationProfileId, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.HostedConfigurationVersions
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

func CreateListhostedconfigurationversionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_applications_ApplicationId_configurationprofiles_ConfigurationProfileId_hostedconfigurationversions",
		mcp.WithDescription("Lists configurations stored in the AppConfig hosted configuration store by version."),
		mcp.WithString("X-Amz-Content-Sha256", mcp.Description("")),
		mcp.WithString("X-Amz-Date", mcp.Description("")),
		mcp.WithString("X-Amz-Algorithm", mcp.Description("")),
		mcp.WithString("X-Amz-Credential", mcp.Description("")),
		mcp.WithString("X-Amz-Security-Token", mcp.Description("")),
		mcp.WithString("X-Amz-Signature", mcp.Description("")),
		mcp.WithString("X-Amz-SignedHeaders", mcp.Description("")),
		mcp.WithString("ApplicationId", mcp.Required(), mcp.Description("The application ID.")),
		mcp.WithString("ConfigurationProfileId", mcp.Required(), mcp.Description("The configuration profile ID.")),
		mcp.WithNumber("max_results", mcp.Description("The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results.")),
		mcp.WithString("next_token", mcp.Description("A token to start the list. Use this token to get the next set of results. ")),
		mcp.WithString("version_label", mcp.Description("An optional filter that can be used to specify the version label of an AppConfig hosted configuration version. This parameter supports filtering by prefix using a wildcard, for example \"v2*\". If you don't specify an asterisk at the end of the value, only an exact match is returned.")),
		mcp.WithString("MaxResults", mcp.Description("Pagination limit")),
		mcp.WithString("NextToken", mcp.Description("Pagination token")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListhostedconfigurationversionsHandler(cfg),
	}
}
