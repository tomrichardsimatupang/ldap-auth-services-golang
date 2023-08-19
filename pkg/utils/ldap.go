package utils

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

type LdapUser struct {
	Uid   string `json:"uid"`
	Name  string `json:"cn"`
	Mail  string `json:"mail"`
	Phone string `json:"phone"`
}

func GetUserLDAP(config LdapConfig, uid string, password string) (LdapUser, error) {

	var ldapUser LdapUser

	server := fmt.Sprintf("ldap://%s:%d", config.Host, config.Port)
	bindDN := fmt.Sprintf("uid=%s,%s", uid, config.Dn)
	bindPassword := password

	l, err := ldap.DialURL(server)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind(bindDN, bindPassword)
	if err != nil {
		return ldapUser, fmt.Errorf("Authentication failed: %w", err)
	}

	// Define the user's DN and the attributes you want to retrieve
	userDN := bindDN
	attributes := []string{"uid", "cn", "mail", "telephoneNumber"}

	searchRequest := ldap.NewSearchRequest(
		userDN,
		ldap.ScopeBaseObject, ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=*)", // You can modify the filter to match your user's entry
		attributes,
		nil,
	)

	searchResult, err := l.Search(searchRequest)
	if err != nil {
		return ldapUser, fmt.Errorf("Search error: %w", err)
	}

	if len(searchResult.Entries) > 0 {

		entry := searchResult.Entries[0]

		ldapUser.Uid = entry.GetAttributeValue("uid")
		ldapUser.Name = entry.GetAttributeValue("cn")
		ldapUser.Mail = entry.GetAttributeValue("mail")
		ldapUser.Phone = entry.GetAttributeValue("telephoneNumber")

		return ldapUser, nil

	}

	return ldapUser, errors.New("User not found")

}
