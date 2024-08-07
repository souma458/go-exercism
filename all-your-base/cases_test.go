package allyourbase

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

var testCases = []struct {
	description   string
	inputBase     int
	inputDigits   []int
	outputBase    int
	expected      []int
	expectedError string
}{
	// {
	// 	description:   "single bit one to decimal",
	// 	inputBase:     2,
	// 	inputDigits:   []int{1},
	// 	outputBase:    10,
	// 	expected:      []int{1},
	// 	expectedError: "",
	// },
	// {
	// 	description:   "binary to single decimal",
	// 	inputBase:     2,
	// 	inputDigits:   []int{1, 0, 1},
	// 	outputBase:    10,
	// 	expected:      []int{5},
	// 	expectedError: "",
	// },
	// {
	// 	description:   "single decimal to binary",
	// 	inputBase:     10,
	// 	inputDigits:   []int{5},
	// 	outputBase:    2,
	// 	expected:      []int{1, 0, 1},
	// 	expectedError: "",
	// },
	{
		description:   "binary to multiple decimal",
		inputBase:     2,
		inputDigits:   []int{1, 0, 1, 0, 1, 0},
		outputBase:    10,
		expected:      []int{4, 2},
		expectedError: "",
	},
	{
		description:   "decimal to binary",
		inputBase:     10,
		inputDigits:   []int{4, 2},
		outputBase:    2,
		expected:      []int{1, 0, 1, 0, 1, 0},
		expectedError: "",
	},
	{
		description:   "trinary to hexadecimal",
		inputBase:     3,
		inputDigits:   []int{1, 1, 2, 0},
		outputBase:    16,
		expected:      []int{2, 10},
		expectedError: "",
	},
	{
		description:   "hexadecimal to trinary",
		inputBase:     16,
		inputDigits:   []int{2, 10},
		outputBase:    3,
		expected:      []int{1, 1, 2, 0},
		expectedError: "",
	},
	{
		description:   "15-bit integer",
		inputBase:     97,
		inputDigits:   []int{3, 46, 60},
		outputBase:    73,
		expected:      []int{6, 10, 45},
		expectedError: "",
	},
	{
		description:   "empty list",
		inputBase:     2,
		inputDigits:   []int{},
		outputBase:    10,
		expected:      []int{0},
		expectedError: "",
	},
	{
		description:   "single zero",
		inputBase:     10,
		inputDigits:   []int{0},
		outputBase:    2,
		expected:      []int{0},
		expectedError: "",
	},
	{
		description:   "multiple zeros",
		inputBase:     10,
		inputDigits:   []int{0, 0, 0},
		outputBase:    2,
		expected:      []int{0},
		expectedError: "",
	},
	{
		description:   "leading zeros",
		inputBase:     7,
		inputDigits:   []int{0, 6, 0},
		outputBase:    10,
		expected:      []int{4, 2},
		expectedError: "",
	},
	{
		description:   "input base is one",
		inputBase:     1,
		inputDigits:   []int{0},
		outputBase:    10,
		expected:      []int(nil),
		expectedError: "input base must be >= 2",
	},
	{
		description:   "input base is zero",
		inputBase:     0,
		inputDigits:   []int{},
		outputBase:    10,
		expected:      []int(nil),
		expectedError: "input base must be >= 2",
	},
	{
		description:   "input base is negative",
		inputBase:     -2,
		inputDigits:   []int{1},
		outputBase:    10,
		expected:      []int(nil),
		expectedError: "input base must be >= 2",
	},
	{
		description:   "negative digit",
		inputBase:     2,
		inputDigits:   []int{1, -1, 1, 0, 1, 0},
		outputBase:    10,
		expected:      []int(nil),
		expectedError: "all digits must satisfy 0 <= d < input base",
	},
	{
		description:   "invalid positive digit",
		inputBase:     2,
		inputDigits:   []int{1, 2, 1, 0, 1, 0},
		outputBase:    10,
		expected:      []int(nil),
		expectedError: "all digits must satisfy 0 <= d < input base",
	},
	{
		description:   "output base is one",
		inputBase:     2,
		inputDigits:   []int{1, 0, 1, 0, 1, 0},
		outputBase:    1,
		expected:      []int(nil),
		expectedError: "output base must be >= 2",
	},
	{
		description:   "output base is zero",
		inputBase:     10,
		inputDigits:   []int{7},
		outputBase:    0,
		expected:      []int(nil),
		expectedError: "output base must be >= 2",
	},
	{
		description:   "output base is negative",
		inputBase:     2,
		inputDigits:   []int{1},
		outputBase:    -7,
		expected:      []int(nil),
		expectedError: "output base must be >= 2",
	},
	{
		description:   "both bases are negative",
		inputBase:     -2,
		inputDigits:   []int{1},
		outputBase:    -7,
		expected:      []int(nil),
		expectedError: "input base must be >= 2",
	},
}
