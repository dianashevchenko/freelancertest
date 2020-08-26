package services

import (
	"context"
	"freelancertest/repositories"
)

func CreateRepository(ctx context.Context, descr string) (*repositories.LdapRepository, error) {
	//tlsConfig := &tls.Config{InsecureSkipVerify: true, }
	//
	//l, err := ldap.DialTLS("tcp", "ldap.example.com:636", tlsConfig)
	//if err != nil {
	//	log.Println(err)
	//	return nil, err
	//}
	//
	//err = l.Bind("user@test.com", "password")
	//if err != nil {
	//	log.Println(err)
	//	return nil, err
	//}
	//return repositories.NewLdapRepository(l), nil

	return &repositories.LdapRepository{}, nil
}
