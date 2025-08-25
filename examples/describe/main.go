package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== gt Describe() and Describef() Demo ===")
	fmt.Println()

	fmt.Println("The gt library now supports Describe() and Describef() methods for all test types!")
	fmt.Println()

	fmt.Println("Usage Examples:")
	fmt.Println()

	fmt.Println("1. Basic Describe() usage:")
	fmt.Println("   gt.Value(t, actual).Describe(\"User ID should match\").Equal(expected)")
	fmt.Println("   Error format:")
	fmt.Println("   Description: User ID should match")
	fmt.Println("   Error: values are not matched")
	fmt.Println()

	fmt.Println("2. Describef() usage with formatting:")
	fmt.Printf("   gt.Value(t, userID).Describef(\"User ID should be %%d for user %%s\", 456, \"Alice\").Equal(456)\n")
	fmt.Println("   Error format:")
	fmt.Println("   Description: User ID should be 456 for user Alice")
	fmt.Println("   Error: values are not matched")
	fmt.Println()

	fmt.Println("3. Array test with long description (80+ characters):")
	longDesc := "Array should contain all expected fruits including orange, grape, and pineapple for the complete collection"
	fmt.Printf("   gt.Array(t, items).Describe(\"%s\").Length(5)\n", longDesc)
	fmt.Println("   Error format:")
	fmt.Printf("   Description: %s\n", longDesc)
	fmt.Println("   Error: array length is expected to be 5, but actual is 2")
	fmt.Println()

	fmt.Println("4. Without description (traditional behavior):")
	fmt.Println("   gt.Array(t, items).Length(5)")
	fmt.Println("   Error format:")
	fmt.Println("   array length is expected to be 5, but actual is 2")
	fmt.Println()

	fmt.Println("5. Method chaining:")
	fmt.Println("   gt.Value(t, result).Describe(\"Validation step\").Required().Equal(expected)")
	fmt.Println()

	fmt.Println("=== Key Features ===")
	fmt.Println("• Describe(string) - Sets a plain description")
	fmt.Println("• Describef(format, args...) - Sets a formatted description using fmt.Sprintf")
	fmt.Println("• Long text friendly layout with 'Description:' and 'Error:' labels")
	fmt.Println("• Works with all test types (Value, Array, Map, Number, String, Bool, Error, File, Cast)")
	fmt.Println("• Integrates with Required() method for fail-fast behavior")
	fmt.Println("• Maintains full backward compatibility - no existing code needs to change")
	fmt.Println()

	fmt.Println("=== To see actual error messages ===")
	fmt.Println("Run: go test -v -run TestDemo_ValueWithDescription examples/describe/")
	fmt.Println("(Uncomment the t.Skip() line in the test to see the formatted error output)")
}