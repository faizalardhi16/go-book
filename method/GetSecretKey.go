package method

import "github.com/joho/godotenv"

func GetSecretKey() (string, error) {
	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
		return "", err
	}

	return envs["SECRET_KEY"], nil
}
