package completion

import "github.com/posener/complete"

func Run() {
	gorew := complete.Command{
		Sub: complete.Commands{
			"install": complete.Command{
				Flags: complete.Flags{
					"-f":     complete.PredictDirs("*"),
					"--file": complete.PredictDirs("*"),
				},
			},
			"init": complete.Command{
				Flags: complete.Flags{
					"-b":        complete.PredictDirs("*"),
					"--binPath": complete.PredictDirs("*"),
					"-f":        complete.PredictDirs("*"),
					"--file":    complete.PredictDirs("*"),
					"-s":        complete.PredictDirs("*"),
					"--srcPath": complete.PredictDirs("*"),
				},
			},
		},
		Flags: complete.Flags{
			"help": complete.PredictNothing,
		},
		GlobalFlags: complete.Flags{
			"--help": complete.PredictNothing,
			"-h":     complete.PredictNothing,
		},
	}
	complete.New("gorew", gorew).Run()
}
