package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash : Retorna a senha criptografada
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha : Compara a senha criptografada com a senha não criptografada
func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
