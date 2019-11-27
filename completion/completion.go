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
