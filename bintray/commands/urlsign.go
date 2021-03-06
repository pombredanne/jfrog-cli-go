package commands

import (
	"fmt"
	"github.com/jfrogdev/jfrog-cli-go/bintray/utils"
	"github.com/jfrogdev/jfrog-cli-go/utils/cliutils"
	"github.com/jfrogdev/jfrog-cli-go/utils/config"
	"github.com/jfrogdev/jfrog-cli-go/utils/ioutils"
)

func SignVersion(urlSigningDetails *utils.PathDetails, flags *UrlSigningFlags) {
	if flags.BintrayDetails.User == "" {
		flags.BintrayDetails.User = urlSigningDetails.Subject
	}
	path := urlSigningDetails.Subject + "/" + urlSigningDetails.Repo + "/" + urlSigningDetails.Path
	url := flags.BintrayDetails.ApiUrl + "signed_url/" + path
	data := builJson(flags)

	fmt.Println("Signing URL for: " + path)
	httpClientsDetails := utils.GetBintrayHttpClientDetails(flags.BintrayDetails)
	resp, body := ioutils.SendPost(url, []byte(data), httpClientsDetails)
	fmt.Println("Bintray response: " + resp.Status)
	fmt.Println(cliutils.IndentJson(body))
}

func builJson(flags *UrlSigningFlags) string {
	m := map[string]string{
		"expiry":          flags.Expiry,
		"valid_for_secs":  flags.ValidFor,
		"callback_id":     flags.CallbackId,
		"callback_email":  flags.CallbackEmail,
		"callback_url":    flags.CallbackUrl,
		"callback_method": flags.CallbackMethod,
	}
	return cliutils.MapToJson(m)
}

type UrlSigningFlags struct {
	BintrayDetails *config.BintrayDetails
	Expiry         string
	ValidFor       string
	CallbackId     string
	CallbackEmail  string
	CallbackUrl    string
	CallbackMethod string
}
