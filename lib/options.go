package lib

import (
        "os"
        "path/filepath"

        "github.com/rs/zerolog"
)

// Options contains all of the gowitness options
type Options struct {
	// Logging
	Logger         *zerolog.Logger
	Debug          bool
	DisableLogging bool

	// Screenshots
	ScreenshotPath string

	// Generic options
	Threads    int
	NoHTTPS    bool
	NoHTTP     bool
	ServerAddr string
	BasePath   string

	// Server command
	AllowInsecureURIs bool

	// File command
	File string

	// Scan command
	ScanCidr     []string
	ScanCidrFile string
	ScanRandom   bool
	ScanPorts    string
	PortsSmall   bool
	PortsMedium  bool
	PortsLarge   bool

	// Single
	ScreenshotFileName string

	// Nessus
	NessusPluginContains []string
	NessusServiceNames   []string
	NessusPluginOutput   []string
	NessusPorts          []int

	// Nmap
	NmapFile            string
	NmapService         []string
	NmapServiceContains string
	NmapPorts           []int
	NmapScanHostnames   bool
	NmapOpenPortsOnly   bool

	// Report List
	ReportJSON     bool
	ReportCSV      bool
	PerceptionSort bool

	// Merge
	MergeDBs        []string
	MergeSourcePath string
	MergeOutputDB   string
}

// NewOptions returns a new options struct
func NewOptions() *Options {
	return &Options{}
}

// PrepareScreenshotPath prepares the path to save screenshots in
func (opt *Options) PrepareScreenshotPath() error {

        if _, err := os.Stat(opt.ScreenshotPath); os.IsNotExist(err) {
                if err = os.Mkdir(opt.ScreenshotPath, 0750); err != nil {
                        return err
                }
        }

        testFile := filepath.Join(opt.ScreenshotPath, ".perm_check")
        if err := os.WriteFile(testFile, []byte{}, 0600); err != nil {
                return err
        }
        os.Remove(testFile)

        return nil
}
