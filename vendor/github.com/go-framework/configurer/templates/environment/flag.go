package environment

import (
	"github.com/spf13/pflag"
)

func AddEnvironmentFlag(flag *pflag.FlagSet, env *Environment) {
	flag.Int8Var((*int8)(env), "env", int8(*env), "Run environment")
}
