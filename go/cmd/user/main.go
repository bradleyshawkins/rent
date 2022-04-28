package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/bradleyshawkins/rent/berror"
	"github.com/bradleyshawkins/rent/cmd/user/rest"
	"github.com/bradleyshawkins/rent/config"
)

type cfg struct {
	ConnectionString string `envconfig:"DATABASE_URL" default:"postgresql://postgres:password@localhost:5432/rent-user?sslmode=disable"`
	Port             string `envconfig:"PORT" default:"8080"`
}

func main() {
	log.Println("Starting up rent-user")
	defer func() {
		log.Println("Shutting down rent-user")
	}()

	ctx := context.TODO()

	//secretManager, err := secretmanager.NewClient(ctx)

	//resp, err := secretManager.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
	//	Name: "projects/453363636036/secrets/jwt_key/versions/latest",
	//})
	//if err != nil {
	//	log.Println("Unable to get secret. Error:", err)
	//	os.Exit(1)
	//}

	conf, err := config.Load[cfg]()
	if err != nil {
		log.Println(berror.WrapInternal(err, "unable to parse config"))
		os.Exit(1)
	}

	fmt.Println(conf)

	server := rest.NewBasicServer()
	stop := server.Start("8080")

	if err := waitForShutdown(ctx, stop); err != nil {
		log.Println("Error shutting down. Error:", err)
		os.Exit(999)
	}
}

func waitForShutdown(ctx context.Context, stopFunc func(ctx context.Context) error) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	if err := stopFunc(ctx); err != nil {
		return err
	}
	return nil
}

//token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
//	Issuer:    "BH",
//	Subject:   "BH_JWT",
//	Audience:  nil,
//	ExpiresAt: nil,
//	NotBefore: nil,
//	IssuedAt:  nil,
//	ID:        "",
//})
//
//authenticator, err := bauth.NewJWTAuthenticator([]byte(conf.PublicKey))
//if err != nil {
//	log.Println("unable to create jwt authenticator")
//}
//
//pk, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(conf.PrivateKey))
//if err != nil {
//	log.Println("unable to parse rsa", err)
//	os.Exit(1)
//}
//
//str, err := token.SignedString(pk)
//if err != nil {
//	log.Println(err)
//	os.Exit(200)
//}
//
//fmt.Println(str)
//
//db, err := sql.Open("postgres", conf.ConnectionString)
//if err != nil {
//	log.Println(berror.WrapInternal(err, "unable to connect to database"))
//	os.Exit(2)
//}
//
//pDB, err := postgres.NewDatabase(db)
//if err != nil {
//	log.Println(err)
//	os.Exit(3)
//}
//
//sup := identity.NewSignUpManager(pDB)
//
//server := rest.NewServer(sup, authenticator)
