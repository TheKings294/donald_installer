package main

import (
	"donald_installer/installer"
	"fmt"
	"os/exec"
)

func checkPackageManager(OsName string) {
	fmt.Printf("\nChecking package managers for %s...\n", OsName)

	switch OsName {
	case "Linux":
		if checkIfExists("apt") {

		}
		if checkIfExists("pacman") {

		}
		if checkIfExists("yum") {

		}
	case "macOS":
		if !checkIfExists("brew") {
			if installer.InstallBrew() {

			}
		}
	case "Windows":
		if !checkIfExists("choco") {
			if installer.InstallChoco() {

			}
		}
	default:
		fmt.Println("Unknown OS")
	}

	fmt.Println("Done checking.\n")
}

func checkIfExists(cmd string) bool {
	if _, err := exec.LookPath(cmd); err == nil {
		return true
	} else {
		return false
	}
}
