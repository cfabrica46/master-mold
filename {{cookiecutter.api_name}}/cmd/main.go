package main

func main() {
	cfg, err := config.GetAPIConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := openPostgresConn(cfg.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	runServer(cfg.Port, db)
}

func runServer(port string, db *sql.DB) {
	svc := service.GetService(db)

	getAllUsersHandler := httptransport.NewServer(
		endpoint.MakeGetAllUsersEndpoint(svc),
		transport.DecodeRequestWithoutBody(),
		transport.EncodeResponse,
	)

	getUserByIDHandler := httptransport.NewServer(
		endpoint.MakeGetUserByIDEndpoint(svc),
		transport.DecodeRequest(entity.IDRequest{}),
		transport.EncodeResponse,
	)

	getUserByUsernameAndPasswordHandler := httptransport.NewServer(
		endpoint.MakeGetUserByUsernameAndPasswordEndpoint(svc),
		transport.DecodeRequest(entity.UsernamePasswordRequest{}),
		transport.EncodeResponse,
	)

	getIDByUsernameHandler := httptransport.NewServer(
		endpoint.MakeGetIDByUsernameEndpoint(svc),
		transport.DecodeRequest(entity.UsernameRequest{}),
		transport.EncodeResponse,
	)

	insertUserHandler := httptransport.NewServer(
		endpoint.MakeInsertUserEndpoint(svc),
		transport.DecodeRequest(entity.UsernamePasswordEmailRequest{}),
		transport.EncodeResponse,
	)

	deleteUserHandler := httptransport.NewServer(
		endpoint.MakeDeleteUserEndpoint(svc),
		transport.DecodeRequest(entity.IDRequest{}),
		transport.EncodeResponse,
	)

	router := mux.NewRouter()
	router.Methods(http.MethodGet).Path("/users").Handler(getAllUsersHandler)
	router.Methods(http.MethodGet).Path("/user/id").Handler(getUserByIDHandler)
	router.Methods(http.MethodGet).Path("/user/username_password").
		Handler(getUserByUsernameAndPasswordHandler)
	router.Methods(http.MethodGet).Path("/id/username").Handler(getIDByUsernameHandler)
	router.Methods(http.MethodPost).Path("/user").Handler(insertUserHandler)
	router.Methods(http.MethodDelete).Path("/user").Handler(deleteUserHandler)

	log.Println("ListenAndServe on localhost:" + os.Getenv("PORT"))
	log.Println(http.ListenAndServe(":"+port, router))
}
