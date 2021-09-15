package main

import (
	"fmt"
	git "github.com/libgit2/git2go/v31"
	"log"
)

func main() {
	cloneOptions := &git.CloneOptions{}
	cloneOptions.FetchOptions = &git.FetchOptions{
		RemoteCallbacks: git.RemoteCallbacks{
			CredentialsCallback: func(url string, username string, allowedTypes git.CredentialType) (*git.Credential, error) {
				//credential, e := git.NewCredentialSSHKeyFromAgent(username)
				credential, e := git.NewCredentialSSHKey(username, "/Users/q/.ssh/id_rsa.pub", "/Users/q/.ssh/id_rsa", "")
				return credential, e
			},
			CertificateCheckCallback: func(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
				return 0
			},
		},
		Prune:           0,
		UpdateFetchhead: false,
		DownloadTags:    0,
		Headers:         nil,
		ProxyOptions:    git.ProxyOptions{},
	}
	repo, err := git.Clone("https://git.taiji.com.cn:443/TDP/tdp.git", "/tmp/code/tdp", cloneOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repo)
	//signature := &git.Signature{
	//	Name: "qiuzhanghua",
	//	Email: "qiuzhanghua@icloud.com",
	//	When: time.Now(),
	//}
}
