import { z } from "zod";

export const AnalyzeArchitectureSchema = {
    name: "analyze_architecture",
    description: "Analyze and suggest architectural improvements for a Kotlin project based on Clean Architecture, Hexagonal, or Monolith patterns.",
    inputSchema: z.object({
        projectType: z.enum(["monolith", "microservice", "library"]).describe("The type of the project."),
        currentStructureDescription: z.string().describe("Brief description of the current folder/package structure."),
        goal: z.enum(["clean_architecture", "hexagonal", "modular_monolith", "refactor_legacy"]).describe("The architectural goal."),
    }),
};

export function analyzeArchitecture(args: { projectType: string; currentStructureDescription: string; goal: string }) {
    const { projectType, currentStructureDescription, goal } = args;

    let advice = "";
    let structureTemplate = "";

    if (goal === "clean_architecture") {
        advice = `For a ${projectType} aiming for Clean Architecture in Kotlin, strict separation of concerns is key.
    Dependency Rule: Source code dependencies can only point inwards.
    Result: Independent of Frameworks, Testable, Independent of UI, Independent of Database.`;

        structureTemplate = `
    src/
      domain/          (Enterprise Business Rules - Entities)
      usecase/         (Application Business Rules)
      adapter/         (Interface Adapters)
        controller/
        presenter/
        gateway/       (Repo implementations)
      infrastructure/  (Frameworks & Drivers)
        db/
        web/
    `;
    } else if (goal === "hexagonal") {
        advice = `Hexagonal Architecture (Ports and Adapters) focuses on isolating the domain logic from the outside world.
    Primary Ports (Driving): Use Cases / Input Ports.
    Secondary Ports (Driven): Repository Interfaces / Output Ports.`;

        structureTemplate = `
    src/
      domain/
        model/
        port/
          in/  (Use Cases)
          out/ (Repository Interfaces)
      adapter/
        in/ 
          web/ (Controllers)
        out/
          persistence/ (Database Adapters)
      application/
        service/ (Implementation of Use Cases)
    `;
    }

    return {
        content: [
            {
                type: "text",
                text: `### Architectural Analysis for ${projectType}\n\n**Goal**: ${goal}\n\n${advice}\n\n### Suggested Kotlin Package Structure:\n\`\`\`text${structureTemplate}\`\`\`\n\n### Kotlin Specific Tips:\n- Use \`data class\` for Domain Entities.\n- Use \`sealed class\` for Domain Errors or Result types.\n- Use Coroutines \`suspend\` functions in your Ports/UseCases for I/O operations.`
            }
        ]
    };
}
