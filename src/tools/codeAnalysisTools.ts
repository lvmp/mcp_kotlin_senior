import { z } from "zod";

export const CheckBestPracticesSchema = {
    name: "check_best_practices",
    description: "Analyze Kotlin code snippets for common anti-patterns and suggest improvements based on SOLID and optimization principles.",
    inputSchema: z.object({
        codeSnippet: z.string().describe("The Kotlin code to analyze."),
    }),
};

export function checkBestPractices(args: { codeSnippet: string }) {
    const { codeSnippet } = args;

    // Very basic static analysis heuristics (for demonstration)
    // In a real world, this might call detekt or ktlint if available, or use a more complex parser.
    const suggestions = [];

    if (codeSnippet.includes("!!")) {
        suggestions.push("- Avoid using `!!` (not-null assertion). Use `?` safe calls or `?:` Elvis operator instead to prevent NullPointerExceptions.");
    }

    if (codeSnippet.includes("GlobalScope")) {
        suggestions.push("- Avoid `GlobalScope`. Use structured concurrency with `viewModelScope`, `lifecycleScope`, or a custom CoroutineScope.");
    }

    if (codeSnippet.includes("var ") && !codeSnippet.includes("val ")) {
        suggestions.push("- Prefer `val` (immutable) over `var` (mutable) where possible to ensure thread safely and predictability.");
    }

    if (codeSnippet.includes("println")) {
        suggestions.push("- Use a standard Logging framework (e.g., SLF4J or Timber) instead of `println`.");
    }

    if (suggestions.length === 0) {
        suggestions.push("Code looks clean regarding basic heuristics. Ensure you are following SOLID principles:");
        suggestions.push("- Single Responsibility: Each class should have one job.");
        suggestions.push("- Open/Closed: Open for extension, closed for modification.");
    }

    return {
        content: [
            {
                type: "text",
                text: `### Best Practices Analysis\n\n${suggestions.join("\n")}`
            }
        ]
    };
}
