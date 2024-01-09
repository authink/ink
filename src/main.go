package main

func main() {
	env := loadEnv()
	db := connectDB(env)

	ink := &Ink{
		env,
		db,
	}

	migrateSchema(ink)

	setupGracefulShutdown(
		ink,
		createServer(ink),
	)
}
