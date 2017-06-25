package main

import (
	"os/exec"
	"strconv"
	"log"
	"fmt"
	"strings"
	"flag"
	"path/filepath"
)

func main() {

	var (
		accountPath string
		mappingPath string
		versionCode int
		packageName string
		apiKey      string
		appId       string
	)
	flag.StringVar(&accountPath, "a", "", "FirebaseServiceAccountFilePath")
	flag.StringVar(&mappingPath, "m", "", "FirebaseCrashMappingFilePath")
	flag.IntVar(&versionCode, "c", 0, "FirebaseCrashVersionCode")
	flag.StringVar(&packageName, "p", "", "FirebaseCrashPackageName")
	flag.StringVar(&apiKey, "k", "", "PFirebaseCrashApiKey")
	flag.StringVar(&appId, "i", "", "PFirebaseCrashAppId")
	flag.Parse()

	absAccountPath, err := filepath.Abs(accountPath)
	if err != nil {
		log.Fatal(err)
	}

	absMappingPath, err := filepath.Abs(mappingPath)
	if err != nil {
		log.Fatal(err)
	}

	upload(absAccountPath, absMappingPath, versionCode, packageName, apiKey, appId)
}

func upload(accountPath string, mappingPath string, versionCode int, packageName string, apiKey string, appId string) {
	m := map[string]string{
		"-PFirebaseServiceAccountFilePath": accountPath,
		"-PFirebaseCrashMappingFilePath":   mappingPath,
		"-PFirebaseCrashVersionCode":       strconv.Itoa(versionCode),
		"-PFirebaseCrashPackageName":       packageName,
		"-PFirebaseCrashApiKey":            apiKey,
		"-PFirebaseCrashAppId":             appId}

	ks := []string{}
	for k, v := range m {
		if v != "" {
			ks = append(ks, strings.Join([]string{k, v}, "="))
		}
	}
	args := append(ks, "firebaseUploadArchivedProguardMapping")

	command := exec.Command("./gradlew", args...)
	fmt.Println(command.Args)
	out, err := command.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
