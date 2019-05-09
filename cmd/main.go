package main

import (
	"flag"
	"log"

	"github.com/louischan-oursky/fission-verifier"
)

func main() {
	secretDir := flag.String("secret-dir", "", "Path to shared secrets directory")
	configDir := flag.String("cfgmap-dir", "", "Path to shared configmap directory")
	dir := flag.String("output-dir", "", "Path to output directory")
	packageNamespace := flag.String("package-namespace", "", "The namespace of the package")
	packageName := flag.String("package-name", "", "The name of the package")

	flag.Parse()

	var err error

	err = verifier.Fetch(verifier.FetchOptions{
		Dir:              *dir,
		SecretDir:        *secretDir,
		ConfigDir:        *configDir,
		PackageNamespace: *packageNamespace,
		PackageName:      *packageName,
	})
	if err != nil {
		log.Fatalf("failed to fetch: %v\n", err)
	}
}
