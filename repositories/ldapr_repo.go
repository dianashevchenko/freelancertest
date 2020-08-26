package repositories

import (
	"fmt"
	"freelancertest/models"
	"github.com/go-ldap/ldap"
	"math"
	"math/rand"
	"strconv"
)

type LdapRepository struct {
	conn *ldap.Conn
}

func NewLdapRepository(conn *ldap.Conn) *LdapRepository {

	return &LdapRepository{conn: conn}

}

func (repo *LdapRepository) Search(ueID string) (*models.AuthenticationSubscription, error) {

	//searchRequest := ldap.NewSearchRequest()
	//sr, err := repo.Conn.Search(searchRequest)
	//if err != nil {
	//	log.Fatal(err)
	//}

	randMock := strconv.Itoa(rand.Intn(math.MaxInt32))
	authMeth := models.AuthenticationSubscriptionAuthenticationMethodEAPAKAPRIME
	return &models.AuthenticationSubscription{
		AlgorithmID:                   randMock,
		AuthenticationManagementField: fmt.Sprintf("%s", randMock),
		AuthenticationMethod:          &authMeth,
	}, nil
}
