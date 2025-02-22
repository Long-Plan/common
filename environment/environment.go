package environment

var Env Mode

func InitEnv(env Mode) {
	Env = env
}

func GetEnv() Mode {
	return Env
}
