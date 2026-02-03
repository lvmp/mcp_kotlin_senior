import { z } from "zod";

export const GenerateTestTemplateSchema = {
    name: "generate_test_template",
    description: "Generate a Kotlin test class template using JUnit 5 and MockK.",
    inputSchema: z.object({
        className: z.string().describe("The name of the class to be tested."),
        testType: z.enum(["unit", "integration"]).describe("The type of test to generate."),
        dependencies: z.array(z.string()).optional().describe("List of dependencies to mock (e.g. ['UserRepository', 'EmailService'])."),
    }),
};

export function generateTestTemplate(args: { className: string; testType: string; dependencies?: string[] }) {
    const { className, testType, dependencies } = args;
    const mocks = dependencies || [];

    const mockDeclarations = mocks.map(d => `    @MockK private lateinit var ${d.charAt(0).toLowerCase() + d.slice(1)}: ${d}`).join("\n");
    const constructionArgs = mocks.map(d => `${d.charAt(0).toLowerCase() + d.slice(1)}`).join(", ");

    let content = "";

    if (testType === "unit") {
        content = `
import io.mockk.impl.annotations.MockK
import io.mockk.junit5.MockKExtension
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.extension.ExtendWith
import io.mockk.verify
import io.mockk.every
import org.junit.jupiter.api.Assertions.*

@ExtendWith(MockKExtension::class)
class ${className}Test {

${mockDeclarations}

    private lateinit var subject: ${className}

    @BeforeEach
    fun setUp() {
        // Initialize subject with mocks
        subject = ${className}(${constructionArgs})
    }

    @Test
    fun \`should do something expected\`() {
        // Given
        // every { ... } returns ...

        // When
        // subject.doSomething()

        // Then
        // verify { ... }
        // assertTrue(...)
    }
}
    `;
    } else {
        content = `
import org.junit.jupiter.api.Test
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.test.context.ActiveProfiles

@SpringBootTest
@ActiveProfiles("test")
class ${className}IntegrationTest {

    @Test
    fun \`context loads\`() {
    }
}
    `;
    }

    return {
        content: [
            {
                type: "text",
                text: `### Generated ${testType} Test Template for ${className}\n\n\`\`\`kotlin${content}\`\`\``
            }
        ]
    };
}
