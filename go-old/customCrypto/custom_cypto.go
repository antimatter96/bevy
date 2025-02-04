// Pacakge customcrypto stores all the constants used all over the service

package customcrypto

// CustomCrypto encapsulates all the functions performed by
// our crypto services
type CustomCrypto interface {
	Encrypt(string, string) (string, string, string, error)
	Decrypt(string, string, string, string) (string, error)
}

type PasswordChecker interface {
	IsSame([]byte, []byte) error
	GetHash([]byte) ([]byte, error)
	NoMatch(error) bool
}
