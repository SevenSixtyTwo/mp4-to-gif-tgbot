package env

import utilenv "mp4togif/utils/env"

var (
	TOKEN string
)

// import --> const --> var --> init()
func init() {
	utilenv.LoadFileEnv("./secrets/.smtp.env")

	utilenv.LoadStrVar(&TOKEN, "TOKEN")
}
