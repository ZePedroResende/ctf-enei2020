//https://pentesterlab.com/exercises/play_xxe/course
//https://github.com/RobotsAndPencils/go-saml/issues/14
    // fmt.Println("xmlsec1", "--sign", "--privkey-pem", privateKeyPath,
    //  "--id-attr:ID", id,
    //  "--output", samlXmlsecOutput.Name(), samlXmlsecInput.Name())
    output, err := exec.Command("xmlsec1", "--sign", "--privkey-pem", privateKeyPath,
        "--id-attr:ID", id,
        "--output", samlXmlsecOutput.Name(), samlXmlsecInput.Name()).CombinedOutput()
    if err != nil {
        return "", errors.New(err.Error() + " : " + string(output))

  defer deleteTempFile(samlXmlsecInput.Name())

    //fmt.Println("xmlsec1", "--verify", "--pubkey-cert-pem", publicCertPath, "--id-attr:ID", id, samlXmlsecInput.Name())
    _, err = exec.Command("xmlsec1", "--verify", "--pubkey-cert-pem", publicCertPath, "--id-attr:ID", id, samlXmlsecInput.Name()).CombinedOutput()
    if err != nil {
        return errors.New("error verifing signature: " + err.Error())
    }
    return nil
