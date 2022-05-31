package cms

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/MikelSot/metal-bat/model"
	"github.com/MikelSot/metal-bat/model/authorization"
)

func loadSignatures(conf model.Configuration, logger model.Logger) {
	private := conf.PrivateFileSign
	public := conf.PublicFileSign

	filePrivate, err := ioutil.ReadFile(private)
	checkErr(err, fmt.Sprintf("no se pudo leer el archivo de firma privado %s", private))

	filePublic, err := ioutil.ReadFile(public)
	checkErr(err, fmt.Sprintf("no se pudo leer el archivo de firma publica %s", public))

	authorization.LoadSignatures(filePrivate, filePublic, logger)
}

func checkErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
