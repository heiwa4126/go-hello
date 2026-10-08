package say

import cowsay "github.com/Code-Hex/Neo-cowsay/v2"

// Say formats message as a cowsay speech bubble and returns the result.
func Say(message string) (string, error) {
	return cowsay.Say(
		message,
		// cowsay.Eyes("xx"),
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
}
