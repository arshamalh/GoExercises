package passgen

const (
	// Use this symbols in making your passwords.
	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

// Seed random should be called only once per program run
func SeedRandom() {
	// You can leave this function empty, but it's better to know how to seed math/rand package.
}

// generate a random pass with a given length
func GeneratePass(pass_len uint8) string {
	// Passwords should be more than 8 characters, and we correct users automatically.
	if pass_len < 8 {
		pass_len = 8
	}
	panic("Implement me")
}
