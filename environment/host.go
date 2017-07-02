package environment

import "os"

func getHost() string {
	return os.Getenv("HOST")
}
