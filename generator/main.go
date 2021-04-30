package generator

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func GenerateShFile(CMDs []string) (fileName string) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	fileName = strconv.Itoa(r.Int()) + ".sh"

	content := "#!/bin/bash\n\n"

	for _, cmd := range CMDs {
		content += cmd + "\n"
	}
	err := ioutil.WriteFile(fileName, []byte(content), 0755)
	if err != nil {
		log.Fatalln(err)
	}

	return
}
