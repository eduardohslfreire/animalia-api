package env

import "github.com/spf13/viper"

//Env uma representação de Env
type Env struct{}

//NewEnv cria uma novo Env
func NewEnv() *Env {
	v := new(Env)
	v.Init()
	return v
}

// Init habilita VIPER para ler variaveis de ambiente
func (e *Env) Init() {
	viper.AutomaticEnv()
}

//GetString ...
func (e *Env) GetString(key string) string {
	return viper.GetString(key)
}

//GetInt ...
func (e *Env) GetInt(key string) int {
	return viper.GetInt(key)
}

//GetBool ...
func (e *Env) GetBool(key string) bool {
	return viper.GetBool(key)
}
