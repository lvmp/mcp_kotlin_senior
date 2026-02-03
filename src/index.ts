import { Server } from "@modelcontextprotocol/sdk/server/index.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import { CallToolRequestSchema, ListToolsRequestSchema } from "@modelcontextprotocol/sdk/types.js";
import { z } from "zod";

import { AnalyzeArchitectureSchema, analyzeArchitecture } from "./tools/architectureTools.js";
import { GenerateDesignPatternSchema, generateDesignPattern } from "./tools/designPatternTools.js";
import { CheckBestPracticesSchema, checkBestPractices } from "./tools/codeAnalysisTools.js";
import { GenerateTestTemplateSchema, generateTestTemplate } from "./tools/testTools.js";
import { SuggestCloudSolutionSchema, suggestCloudSolution } from "./tools/cloudTools.js";

const server = new Server(
    {
        name: "mcp-kotlin-senior",
        version: "1.0.0",
    },
    {
        capabilities: {
            tools: {},
        },
    }
);

server.setRequestHandler(ListToolsRequestSchema, async () => {
    return {
        tools: [
            AnalyzeArchitectureSchema,
            GenerateDesignPatternSchema,
            CheckBestPracticesSchema,
            GenerateTestTemplateSchema,
            SuggestCloudSolutionSchema,
        ],
    };
});

server.setRequestHandler(CallToolRequestSchema, async (request) => {
    const { name, arguments: args } = request.params;

    if (!args) {
        throw new Error("No arguments provided");
    }

    // Type assertions for arguments based on schemas would go here for stricter runtime safety
    // For this implementation, we cast to any or the specific expected type in the helper functions.

    switch (name) {
        case "analyze_architecture":
            return analyzeArchitecture(args as any);
        case "generate_design_pattern":
            return generateDesignPattern(args as any);
        case "check_best_practices":
            return checkBestPractices(args as any);
        case "generate_test_template":
            return generateTestTemplate(args as any);
        case "suggest_cloud_solution":
            return suggestCloudSolution(args as any);
        default:
            throw new Error(`Unknown tool: ${name}`);
    }
});

async function run() {
    const transport = new StdioServerTransport();
    await server.connect(transport);
    console.error("Kotlin Senior MCP Server running on stdio");
}

run().catch((error) => {
    console.error("Server error:", error);
    process.exit(1);
});
