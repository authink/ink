package main

func main() {
	env := loadEnv()
	db := connectDB(env)

	ink := &Ink{
		env,
		db,
	}

	setupGracefulShutdown(
		ink,
		createServer(ink),
	)
}
