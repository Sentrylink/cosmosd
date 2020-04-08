package lib

import "os"

// Run is the main loop, but returns an error
func Run(args []string) error {
	cfg, err := GetConfigFromEnv()
	if err != nil {
		return err
	}
	doUpgrade, err := LaunchProcess(cfg, args, os.Stdout, os.Stderr)

	// if RestartAfterUpgrade, we launch after a successful upgrade (only condition LaunchProcess returns nil)
	for cfg.RestartAfterUpgrade && err == nil && doUpgrade {
		doUpgrade, err = LaunchProcess(cfg, args, os.Stdout, os.Stderr)
	}
	return err
}
