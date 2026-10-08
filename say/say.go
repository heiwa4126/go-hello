package say

import cowsay "github.com/Code-Hex/Neo-cowsay/v2"

// Say formats message as a cowsay speech bubble and returns the result.
func Say(message string) string {
	msg, err := cowsay.Say(
		message,
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		panic(err)
	}
	return msg
}
