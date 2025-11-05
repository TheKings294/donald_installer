package installer

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func InstallChoco() bool {
	if runtime.GOOS != "windows" {
		fmt.Println("This is not Windows. Skipping Chocolatey installation.")
		return false
	}

	fmt.Println("Installing Chocolatey...")

	// PowerShell command to install Chocolatey
	cmd := exec.Command("powershell", "-NoProfile", "-InputFormat", "None",
		"-ExecutionPolicy", "Bypass", "-Command",
		`Set-ExecutionPolicy Bypass -Scope Process -Force; `+
			`[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; `+
			`iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))`)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error installing Chocolatey:", err)
		return false
	}

	path := exec.Command("powershell", "-Command",
		`[Environment]::SetEnvironmentVariable('PATH', $env:PATH + ';C:\ProgramData\chocolatey\bin', 'User')`)

	path.Stdout = os.Stdout
	path.Stderr = os.Stderr

	if err := path.Run(); err != nil {
		fmt.Println("Error adding Chocolatey to the path:", err)
		return false
	}

	fmt.Println("Chocolatey installed successfully!")

	return true
}

func InstallBrew() bool {
	if runtime.GOOS != "darwin" {
		fmt.Println("This is not macOS. Skipping Homebrew installation.")
		return false
	}

	fmt.Println("Installing Homebrew...")

	cmd := exec.Command("/bin/bash", "-c",
		`/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error installing Homebrew:", err)
		return false
	}

	path := exec.Command("/bin/bash", "-c",
		`echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> ~/.zprofile && eval "$(/opt/homebrew/bin/brew shellenv)"`)

	path.Stdout = os.Stdout
	path.Stderr = os.Stderr

	if err := path.Run(); err != nil {
		fmt.Println("Error adding Homebrew to PATH:", err)
		return false
	}

	fmt.Println("Homebrew installed !")
	
	return true
}
