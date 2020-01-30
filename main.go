port := os.Getenv("PORT")
if port == "" {
    port = "8040"
    log.Println("[-] No PORT environment variable detected. Setting to ", port)
}
