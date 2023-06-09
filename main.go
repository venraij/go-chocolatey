package chocolatey

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/windows"
)

func init() {
	if !amAdmin() {
		runMeElevated()
	}
	time.Sleep(3 * time.Second)
}

func Install() {
	cmd := exec.Command("powershell", "Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func InstallPackage(packageName string, args string) {
	log.Println("Running: choco install " + packageName + " " + args)

	cmd := exec.Command("choco", "install", packageName, args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func UninstallPackage(packageName string, args string) {
	log.Println("Running: choco uninstall " + packageName + " " + args)

	cmd := exec.Command("choco", "uninstall", packageName, args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func UpdatePackage(packageName string, args string) {
	log.Println("Running: choco upgrade " + packageName + " " + args)

	cmd := exec.Command("choco", "upgrade", packageName, args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func UpdateAllPackages(args string) {
	log.Println("Running: choco upgrade all " + args)

	cmd := exec.Command("choco", "upgrade", "all", args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func ListPackages(args string) {
	log.Println("Running: choco list " + args)

	cmd := exec.Command("choco", "list", args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func ListInstalledPackages(args string) {
	log.Println("Running: choco list --localonly" + args)

	cmd := exec.Command("choco", "list", "--localonly", args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func ListOutdatedPackages(args string) {
	log.Println("Running: choco list --outdated" + args)

	cmd := exec.Command("choco", "list", "--outdated", args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}

// Check if chocolatey is installed
func IsInstalled() bool {
	_, err := exec.LookPath("choco")
	if err != nil {
		return false
	}
	return true
}

func runMeElevated() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func amAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		fmt.Println("admin no")
		return false
	}
	fmt.Println("admin yes")
	return true
}
