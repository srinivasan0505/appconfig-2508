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

func GetconfigurationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		ApplicationVal, ok := args["Application"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: Application"), nil
		}
		Application, ok := ApplicationVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: Application"), nil
		}
		EnvironmentVal, ok := args["Environment"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: Environment"), nil
		}
		Environment, ok := EnvironmentVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: Environment"), nil
		}
		ConfigurationVal, ok := args["Configuration"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: Configuration"), nil
		}
		Configuration, ok := ConfigurationVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: Configuration"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["client_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("client_id=%v", val))
		}
		if val, ok := args["client_configuration_version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("client_configuration_version=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/applications/%s/environments/%s/configurations/%s#client_id%s", cfg.BaseURL, Application, Environment, Configuration, queryString)
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
		var result models.Configuration
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

func CreateGetconfigurationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_applications_Application_environments_Environment_configurations_Configuration#client_id",
		mcp.WithDescription("<p>(Deprecated) Retrieves the latest deployed configuration.</p> <important> <p>Note the following important information.</p> <ul> <li> <p>This API action is deprecated. Calls to receive configuration data should use the <a href="https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_appconfigdata_StartConfigurationSession.html">StartConfigurationSession</a> and <a href="https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_appconfigdata_GetLatestConfiguration.html">GetLatestConfiguration</a> APIs instead. </p> </li> <li> <p> <code>GetConfiguration</code> is a priced call. For more information, see <a href="https://aws.amazon.com/systems-manager/pricing/">Pricing</a>.</p> </li> </ul> </important>"),
		mcp.WithString("X-Amz-Content-Sha256", mcp.Description("")),
		mcp.WithString("X-Amz-Date", mcp.Description("")),
		mcp.WithString("X-Amz-Algorithm", mcp.Description("")),
		mcp.WithString("X-Amz-Credential", mcp.Description("")),
		mcp.WithString("X-Amz-Security-Token", mcp.Description("")),
		mcp.WithString("X-Amz-Signature", mcp.Description("")),
		mcp.WithString("X-Amz-SignedHeaders", mcp.Description("")),
		mcp.WithString("Application", mcp.Required(), mcp.Description("The application to get. Specify either the application name or the application ID.")),
		mcp.WithString("Environment", mcp.Required(), mcp.Description("The environment to get. Specify either the environment name or the environment ID.")),
		mcp.WithString("Configuration", mcp.Required(), mcp.Description("The configuration to get. Specify either the configuration name or the configuration ID.")),
		mcp.WithString("client_id", mcp.Required(), mcp.Description("The clientId parameter in the following command is a unique, user-specified ID to identify the client for the configuration. This ID enables AppConfig to deploy the configuration in intervals, as defined in the deployment strategy. ")),
		mcp.WithString("client_configuration_version", mcp.Description("<p>The configuration version returned in the most recent <code>GetConfiguration</code> response.</p> <important> <p>AppConfig uses the value of the <code>ClientConfigurationVersion</code> parameter to identify the configuration version on your clients. If you don’t send <code>ClientConfigurationVersion</code> with each call to <code>GetConfiguration</code>, your clients receive the current configuration. You are charged each time your clients receive a configuration.</p> <p>To avoid excess charges, we recommend you use the <a href=\"https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/StartConfigurationSession.html\">StartConfigurationSession</a> and <a href=\"https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/GetLatestConfiguration.html\">GetLatestConfiguration</a> APIs, which track the client configuration version on your behalf. If you choose to continue using <code>GetConfiguration</code>, we recommend that you include the <code>ClientConfigurationVersion</code> value with every call to <code>GetConfiguration</code>. The value to use for <code>ClientConfigurationVersion</code> comes from the <code>ConfigurationVersion</code> attribute returned by <code>GetConfiguration</code> when there is new or updated data, and should be saved for subsequent calls to <code>GetConfiguration</code>.</p> </important> <p>For more information about working with configurations, see <a href=\"http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-retrieving-the-configuration.html\">Retrieving the Configuration</a> in the <i>AppConfig User Guide</i>.</p>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetconfigurationHandler(cfg),
	}
}
