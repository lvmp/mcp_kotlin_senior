import { z } from "zod";

export const SuggestCloudSolutionSchema = {
    name: "suggest_cloud_solution",
    description: "Suggest Google Cloud Platform (GCP) services and DevOps strategies for a Kotlin application.",
    inputSchema: z.object({
        usageScenario: z.string().describe("What the application does (e.g. 'Event driven microservices', 'Simple CRUD API')."),
        requirements: z.array(z.string()).describe("Specific requirements (e.g. 'Serverless', 'SQL', 'Global Scale')."),
    }),
};

export function suggestCloudSolution(args: { usageScenario: string; requirements: string[] }) {
    const { usageScenario, requirements } = args;

    let suggestion = "";

    if (requirements.some(r => r.toLowerCase().includes("serverless"))) {
        suggestion += "- **Compute**: Cloud Run (fully managed container platform). Perfect for Kotlin (using Spring Boot or Ktor with GraalVM/JVM).\n";
    } else if (requirements.some(r => r.toLowerCase().includes("kubernetes"))) {
        suggestion += "- **Compute**: GKE (Google Kubernetes Engine). Standard for microservices orchestration.\n";
    } else {
        suggestion += "- **Compute**: Cloud Run is recommended as a default for modern stateless apps.\n";
    }

    if (requirements.some(r => r.toLowerCase().includes("sql"))) {
        suggestion += "- **Database**: Cloud SQL (PostgreSQL recommended for Kotlin/JPA/Exposed).\n";
    } else {
        suggestion += "- **Database**: Firestore (NoSQL) for rapid development and mobile backends.\n";
    }

    suggestion += "- **CI/CD**: Cloud Build. Define `cloudbuild.yaml` to build and deploy your container.\n";
    suggestion += "- **Monitoring**: Cloud Operations Suite (formerly Stackdriver).";

    return {
        content: [
            {
                type: "text",
                text: `### GCP Cloud Solution for "${usageScenario}"\n\n${suggestion}`
            }
        ]
    };
}
