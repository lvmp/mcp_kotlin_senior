package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create MCP server
	s := server.NewMCPServer("mcp-kotlin-senior", "1.0.0",
		server.WithToolCapabilities(true),
	)

	// 1. analyze_architecture
	s.AddTool(mcp.NewTool("analyze_architecture",
		mcp.WithDescription("Analyze and suggest architectural improvements for a Kotlin project based on Clean Architecture, Hexagonal, or Monolith patterns."),
		mcp.WithString("projectType", mcp.Required(), mcp.Description("The type of the project (monolith, microservice, library).")),
		mcp.WithString("currentStructureDescription", mcp.Required(), mcp.Description("Brief description of the current folder/package structure.")),
		mcp.WithString("goal", mcp.Required(), mcp.Description("The architectural goal (clean_architecture, hexagonal, modular_monolith, microservices, refactor_legacy).")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		arguments := request.Params.Arguments.(map[string]interface{})
		projectType, _ := arguments["projectType"].(string)
		goal, _ := arguments["goal"].(string)

		var advice string
		var structureTemplate string

		switch goal {
		case "clean_architecture":
			advice = fmt.Sprintf("For a %s aiming for Clean Architecture in Kotlin, strict separation of concerns is key.\nDependency Rule: Source code dependencies can only point inwards.\nResult: Independent of Frameworks, Testable, Independent of UI, Independent of Database.", projectType)
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
    web/`
		case "hexagonal":
			advice = "Hexagonal Architecture (Ports and Adapters) focuses on isolating the domain logic from the outside world.\nPrimary Ports (Driving): Use Cases / Input Ports.\nSecondary Ports (Driven): Repository Interfaces / Output Ports."
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
    service/ (Implementation of Use Cases)`
		case "microservices":
			advice = "Microservices Architecture involves decomposing the domain into small, independent services.\nKey Principles:\n1. Independent Deployability.\n2. Shared Nothing (Database per Service).\n3. API First Design.\n\nKotlin fits well here with Spring Boot or Ktor/Quarkus. Focus on lightweight, fast startup containers."
			structureTemplate = `
k8s/ (Helm charts or Kustomize)
src/main/kotlin/com/example/serviceName/
  api/ (Controllers/gRPC)
  domain/ (Core Logic)
  data/ (Repositories/Entities)
  config/ (DI, Environment)
Dockerfile`
		default:
			advice = "Choose a goal like clean_architecture, hexagonal, or microservices for detailed advice."
			structureTemplate = "Standard Kotlin structure recommended."
		}

		result := fmt.Sprintf("### Architectural Analysis for %s\n\n**Goal**: %s\n\n%s\n\n### Suggested Kotlin Package Structure:\n```text%s\n```\n\n### Kotlin Specific Tips:\n- Use `data class` for Domain Entities.\n- Use `sealed class` for Domain Errors or Result types.\n- Use Coroutines `suspend` functions in your Ports/UseCases for I/O operations.", projectType, goal, advice, structureTemplate)
		return mcp.NewToolResultText(result), nil
	})

	// 2. generate_design_pattern
	s.AddTool(mcp.NewTool("generate_design_pattern",
		mcp.WithDescription("Generate a Kotlin implementation of a specific design pattern with best practices."),
		mcp.WithString("patternName", mcp.Required(), mcp.Description("The classic GoF design pattern to generate (singleton, factory_method, strategy, observer, etc.).")),
		mcp.WithString("context", mcp.Required(), mcp.Description("The context or use case for this pattern (e.g., 'PaymentProcessor', 'Logger').")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		arguments := request.Params.Arguments.(map[string]interface{})
		patternName, _ := arguments["patternName"].(string)
		contextStr, _ := arguments["context"].(string)
		className := strings.ReplaceAll(contextStr, " ", "")

		var code string
		var explanation string

		switch patternName {
		case "singleton":
			code = fmt.Sprintf(`object %sManager {
    init {
        println("%sManager initialized")
    }

    fun doSomething() {
        // Implementation
    }
}`, className, className)
			explanation = "In Kotlin, `object` is the idiomatic way to implement the Singleton pattern. It is thread-safe and lazy-loaded by default."
		case "strategy":
			code = fmt.Sprintf(`interface %sStrategy {
    fun execute(data: String): String
}

class Concrete%sStrategyA : %sStrategy {
    override fun execute(data: String) = "Strategy A: $data"
}

class Concrete%sStrategyB : %sStrategy {
    override fun execute(data: String) = "Strategy B: $data"
}

class %sContext(private var strategy: %sStrategy) {
    fun setStrategy(strategy: %sStrategy) {
        this.strategy = strategy
    }

    fun executeStrategy(data: String): String {
        return strategy.execute(data)
    }
}`, className, className, className, className, className, className, className, className)
			explanation = "The Strategy pattern defines a family of algorithms, encapsulates each one, and makes them interchangeable."
		case "observer":
			code = fmt.Sprintf(`interface %sObserver {
    fun update(event: String)
}

class %sSubject {
    private val observers = mutableListOf<%sObserver>()

    fun addObserver(observer: %sObserver) {
        observers.add(observer)
    }

    fun removeObserver(observer: %sObserver) {
        observers.remove(observer)
    }

    fun notifyObservers(event: String) {
        observers.forEach { it.update(event) }
    }
}`, className, className, className, className, className, className, className)
			explanation = "The Observer pattern defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically."
		default:
			code = fmt.Sprintf("// Implementation for %s Pattern in %s context coming soon.", patternName, contextStr)
			explanation = "This pattern is recognized but template is being expanded."
		}

		result := fmt.Sprintf("### %s Pattern for %s\n\n%s\n\n```kotlin\n%s\n```", patternName, contextStr, explanation, code)
		return mcp.NewToolResultText(result), nil
	})

	// 3. check_best_practices
	s.AddTool(mcp.NewTool("check_best_practices",
		mcp.WithDescription("Analyze Kotlin code snippets for common anti-patterns and suggest improvements based on SOLID and optimization principles."),
		mcp.WithString("codeSnippet", mcp.Required(), mcp.Description("The Kotlin code to analyze.")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		arguments := request.Params.Arguments.(map[string]interface{})
		codeSnippet, _ := arguments["codeSnippet"].(string)

		var suggestions []string
		if strings.Contains(codeSnippet, "!!") {
			suggestions = append(suggestions, "- Avoid using `!!` (not-null assertion). Use `?` safe calls or `?:` Elvis operator instead to prevent NullPointerExceptions.")
		}
		if strings.Contains(codeSnippet, "GlobalScope") {
			suggestions = append(suggestions, "- Avoid `GlobalScope`. Use structured concurrency with `viewModelScope`, `lifecycleScope`, or a custom CoroutineScope.")
		}
		if strings.Contains(codeSnippet, "var ") && !strings.Contains(codeSnippet, "val ") {
			suggestions = append(suggestions, "- Prefer `val` (immutable) over `var` (mutable) where possible to ensure thread safely and predictability.")
		}
		if strings.Contains(codeSnippet, "println") {
			suggestions = append(suggestions, "- Use a standard Logging framework (e.g., SLF4J or Timber) instead of `println`.")
		}

		if len(suggestions) == 0 {
			suggestions = append(suggestions, "Code looks clean regarding basic heuristics. Ensure you are following SOLID principles:",
				"- Single Responsibility: Each class should have one job.",
				"- Open/Closed: Open for extension, closed for modification.")
		}

		result := fmt.Sprintf("### Best Practices Analysis\n\n%s", strings.Join(suggestions, "\n"))
		return mcp.NewToolResultText(result), nil
	})

	// 4. generate_test_template
	s.AddTool(mcp.NewTool("generate_test_template",
		mcp.WithDescription("Generate a Kotlin test class template using JUnit 5 and MockK."),
		mcp.WithString("className", mcp.Required(), mcp.Description("The name of the class to be tested.")),
		mcp.WithString("testType", mcp.Required(), mcp.Description("The type of test to generate (unit, integration).")),
		// Dependencies as a list is harder with current mcp-go helper if not using jsonschema directly,
		// but let's try to handle it as a comma-separated string if needed, or if the SDK supports arrays.
		// Actually s.AddTool doesn't have WithArray easily, but we can use raw schema if needed.
		// For simplicity, let's use a string for dependencies.
		mcp.WithString("dependencies", mcp.Description("Comma-separated list of dependencies to mock (e.g. 'UserRepository,EmailService').")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		arguments := request.Params.Arguments.(map[string]interface{})
		className, _ := arguments["className"].(string)
		testType, _ := arguments["testType"].(string)
		depStr, _ := arguments["dependencies"].(string)

		var dependencies []string
		if depStr != "" {
			for _, d := range strings.Split(depStr, ",") {
				dependencies = append(dependencies, strings.TrimSpace(d))
			}
		}

		var content string
		if testType == "unit" {
			var mockDeclarations []string
			var constructionArgs []string
			for _, d := range dependencies {
				varName := strings.ToLower(d[0:1]) + d[1:]
				mockDeclarations = append(mockDeclarations, fmt.Sprintf("    @MockK private lateinit var %s: %s", varName, d))
				constructionArgs = append(constructionArgs, varName)
			}

			content = fmt.Sprintf(`
import io.mockk.impl.annotations.MockK
import io.mockk.junit5.MockKExtension
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.extension.ExtendWith
import io.mockk.verify
import io.mockk.every
import org.junit.jupiter.api.Assertions.*

@ExtendWith(MockKExtension::class)
class %sTest {

%s

    private lateinit var subject: %s

    @BeforeEach
    fun setUp() {
        // Initialize subject with mocks
        subject = %s(%s)
    }

    @Test
    fun `+"`"+`should do something expected`+"`"+`() {
        // Given
        // every { ... } returns ...

        // When
        // subject.doSomething()

        // Then
        // verify { ... }
        // assertTrue(...)
    }
}`, className, strings.Join(mockDeclarations, "\n"), className, className, strings.Join(constructionArgs, ", "))
		} else {
			content = fmt.Sprintf(`
import org.junit.jupiter.api.Test
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.test.context.ActiveProfiles

@SpringBootTest
@ActiveProfiles("test")
class %sIntegrationTest {

    @Test
    fun `+"`"+`context loads`+"`"+`() {
    }
}`, className)
		}

		result := fmt.Sprintf("### Generated %s Test Template for %s\n\n```kotlin%s\n```", testType, className, content)
		return mcp.NewToolResultText(result), nil
	})

	// 5. suggest_cloud_solution
	s.AddTool(mcp.NewTool("suggest_cloud_solution",
		mcp.WithDescription("Suggest Google Cloud Platform (GCP) services and DevOps strategies for a Kotlin application."),
		mcp.WithString("usageScenario", mcp.Required(), mcp.Description("What the application does (e.g. 'Event driven microservices', 'Simple CRUD API').")),
		mcp.WithString("requirements", mcp.Required(), mcp.Description("Comma-separated list of requirements (e.g. 'Serverless, SQL, Global Scale').")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		arguments := request.Params.Arguments.(map[string]interface{})
		usageScenario, _ := arguments["usageScenario"].(string)
		reqStr, _ := arguments["requirements"].(string)
		requirements := strings.Split(strings.ToLower(reqStr), ",")

		var suggestion strings.Builder

		hasServerless := false
		hasKubernetes := false
		hasSQL := false

		for _, r := range requirements {
			r = strings.TrimSpace(r)
			if strings.Contains(r, "serverless") {
				hasServerless = true
			}
			if strings.Contains(r, "kubernetes") || strings.Contains(r, "gke") {
				hasKubernetes = true
			}
			if strings.Contains(r, "sql") {
				hasSQL = true
			}
		}

		if hasServerless {
			suggestion.WriteString("- **Compute**: Cloud Run (fully managed container platform). Perfect for Kotlin (using Spring Boot or Ktor with GraalVM/JVM).\n")
		} else if hasKubernetes {
			suggestion.WriteString("- **Compute**: GKE (Google Kubernetes Engine). Standard for microservices orchestration.\n")
		} else {
			suggestion.WriteString("- **Compute**: Cloud Run is recommended as a default for modern stateless apps.\n")
		}

		if hasSQL {
			suggestion.WriteString("- **Database**: Cloud SQL (PostgreSQL recommended for Kotlin/JPA/Exposed).\n")
		} else {
			suggestion.WriteString("- **Database**: Firestore (NoSQL) for rapid development and mobile backends.\n")
		}

		suggestion.WriteString("- **CI/CD**: Cloud Build. Define `cloudbuild.yaml` to build and deploy your container.\n")
		suggestion.WriteString("- **Monitoring**: Cloud Operations Suite (formerly Stackdriver).")

		result := fmt.Sprintf("### GCP Cloud Solution for \"%s\"\n\n%s", usageScenario, suggestion.String())
		return mcp.NewToolResultText(result), nil
	})

	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
