package main

import (
    "errors"
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
    "path/filepath"
    "regexp"
)

const version = "v0.1.0"

func main() {
    if len(os.Args) < 2 {
        printHelp()
        return
    }

    switch os.Args[1] {
    case "version":
        fmt.Println("Version:", version)
    case "new":
        if len(os.Args) < 3 {
            fmt.Println("Please provide a file name.")
            return
        }
        if err := createNewFile(os.Args[2]); err != nil {
            fmt.Println("Error:", err)
        }
    case "run":
        if len(os.Args) < 3 {
            fmt.Println("Please provide a file name to run.")
            return
        }
        if err := runFile(os.Args[2]); err != nil {
            fmt.Println("Error:", err)
        }
    case "help":
        printHelp()
    default:
        fmt.Println("Unknown command:", os.Args[1])
        printHelp()
    }
}

func createNewFile(filename string) error {
    sanitizedFilename, err := sanitizeFilename(filename)
    if err != nil {
        return err
    }
    if filepath.Ext(sanitizedFilename) == "" {
        sanitizedFilename += ".ze"
    }
    content := "// This is a new .ze file\n\n"
    if err := ioutil.WriteFile(sanitizedFilename, []byte(content), 0644); err != nil {
        return fmt.Errorf("creating file: %w", err)
    }
    fmt.Println("Created new file:", sanitizedFilename)
    return nil
}

func runFile(filename string) error {
    sanitizedFilename, err := sanitizeFilename(filename)
    if err != nil {
        return err
    }
    if _, err := os.Stat(sanitizedFilename); os.IsNotExist(err) {
        return fmt.Errorf("file does not exist: %s", sanitizedFilename)
    }
    cmd := exec.Command("go", "run", sanitizedFilename)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("running file: %w", err)
    }
    return nil
}

func sanitizeFilename(filename string) (string, error) {
    // Allow only alphanumeric characters, dashes, and underscores in filenames
    re := regexp.MustCompile(`^[a-zA-Z0-9_-]+(\.[a-zA-Z0-9]+)?$`)
    if !re.MatchString(filename) {
        return "", errors.New("invalid filename: filenames can only contain alphanumeric characters, dashes, and underscores")
    }
    return filename, nil
}

func printHelp() {
    fmt.Println("Package Manager - Help")
    fmt.Println("Commands:")
    fmt.Println("  version      - Display the current version")
    fmt.Println("  new <name>   - Create a new .ze file with the given name")
    fmt.Println("  run <name>   - Run the specified file")
    fmt.Println("  help         - Display help information")
}