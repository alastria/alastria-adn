package main


import (
    "crypto/tls"
	"fmt"
    "net/http"
    "io/ioutil"

)

func main() {
    var rsaCertPEM = [] byte(`-----BEGIN CERTIFICATE-----
MIIEKzCCAhOgAwIBAgIRAKlBaL+CEYXR1jXqJVDAtvIwDQYJKoZIhvcNAQELBQAw
MDEbMBkGA1UEChMSQ29uc29yY2lvIEFsYXN0cmlhMREwDwYDVQQDEwhhbGFzdHJp
YTAeFw0xODAxMjMxMTE2MjJaFw0xOTAxMjMxMTE2MjJaMEkxEDAOBgNVBAoTB21v
bml0b3IxNTAzBgNVBAMTLENsaWVudCBjZXJ0aWZpY2F0ZSBmb3IgVExTL1NTTCBh
Z2VudCBtb25pdG9yMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtoV2
clGh7lCSuR6VhTsbHbkhQqwoD0RmjQkft916Bv6hzQAQTJ4hve39V1mnMkMzxxmr
YGRX4LMirYw/laWuSGKNwpkbZlpgTm2x/AUIBPqXzEGbT25FuVIHbs7BJRNlTvQ7
RWOnldilO4mjr9MVX9G7kYofIbu6xXNReR31SHRIm/B314AWoxVoWcHtdvndcgkF
YtG50cB4aRauXGc+p+Tx6TPo1aexWTJIKJaJp8WeHuMr0KcPuNHyyCNW/PqvDYGZ
4eb28IKYidjozI4gESycsVenYMMuWyx5tzZ9q/IyoEHPcXBce/O6zSKtzgxyttY6
4d1G1LISrG6iVKUcHQIDAQABoycwJTAOBgNVHQ8BAf8EBAMCBaAwEwYDVR0lBAww
CgYIKwYBBQUHAwIwDQYJKoZIhvcNAQELBQADggIBAENgh/01kKvDSZg55Mve3o/p
E30F+P5Ys2WL5dea92h/x1Ccl96nGKtSny1NUJLnwHQRjt/A6CLEreI+9BxQenmH
G14l02YsylEtCQd3lG0bMSRLextW4pLKJDMgGzozw6iiytjU01y2xnl16PgTmeEm
5fAYMsQo5NguuzZB/Xi0b56DG9gd6qviqcQsLSaN5Tn9LhbhTngYJc7p02DRTNZh
1ugEc5Z3l4NMx0105z693PAOiVBX2ObKtcPDyyEmRMTTGLcZ3F+/osr6dcwRDIGO
M1UlkADr025Ae3EL2vdxB7MXgOSLK0S8I6pnSRBF063Xfa6i1Cp9Zdl5Mh1z77yi
iSzsLaCs+stYSWmV5As/pzmGttMJqyd5r7Eo3ZIDpkdk29e7rST6cAdv7mdBjYyw
fMgdzkA06MfZuW2DlKccp8SbeJiyhq75NxMtq4AXpufeZjy93O134ssjcY5r2BgU
4+820AH+O39Yg3HNa+iga4vcjqg25a0wF1+V+bkOn3IwW8CDuG6YmOlNSS5RM6cT
XaAtPt0La9V/9lgLPPz81tVggdg6/KH/WnsBlTWewCGdYbkpJQ3I9bpiUDo59yb2
15eGePlChmYyCyAr/vU21myni1EqZ4VwOZg2Y9pR+Nv8RT2ljq/NIgqpf7eV0WCQ
L/Bbhvz+bqpiiZnXk5Xj
-----END CERTIFICATE-----`)

    var rsaKeyPEM = [] byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAtoV2clGh7lCSuR6VhTsbHbkhQqwoD0RmjQkft916Bv6hzQAQ
TJ4hve39V1mnMkMzxxmrYGRX4LMirYw/laWuSGKNwpkbZlpgTm2x/AUIBPqXzEGb
T25FuVIHbs7BJRNlTvQ7RWOnldilO4mjr9MVX9G7kYofIbu6xXNReR31SHRIm/B3
14AWoxVoWcHtdvndcgkFYtG50cB4aRauXGc+p+Tx6TPo1aexWTJIKJaJp8WeHuMr
0KcPuNHyyCNW/PqvDYGZ4eb28IKYidjozI4gESycsVenYMMuWyx5tzZ9q/IyoEHP
cXBce/O6zSKtzgxyttY64d1G1LISrG6iVKUcHQIDAQABAoIBACjPwuJg+nJNNeGK
wygdRTzqLlO4JuTzCHM0vRDhxu1VdlxeTUa0fRr44hLsCwSkHinAxZ8yEKw/odto
ZrdRapzo3IXMsmG6h5pB0PBnN5nVZqfXa6DhKVn8y4itVmax5Y838SDc3ZYV8SxU
5HLmIftg7C3o0nk6ftKzrF4GotfmN5px8xNykH3RRBTpoM7cdCLYCVBifwcUQBc+
HkNKu/dGUW9hYU6DNxdlAziECsPh4NEOPRsDPUFe1e2Mm0CXCbuekGLYLXR/7WyO
7L40p24Jj0sexuQUOJQfPHDmIlxk2D68mgxsc0eh2QB8ieQS25rG8iYJMJYmvfvn
/4kBaoECgYEAzCAxhaXlqCa4DZUkzEJKmOJ4lcx9aYvoXOTqkStyRyrIG2+jWx/8
bN4IgMbm++DRGNIuOYjN6q3YCcFbPLK4qpSBqHyr45Ann+CLEA/3zR43HarrBpQD
5reGeL8gGFVyK286Ytv5HgQmYql8wR7ptaafxwYlFbS2wn+0xLVz3JECgYEA5OfA
P0fC0ouRRqzBXdoQqTusCP+I6E0N5L2J+DNrkEcpVWobdcczDGSXMqdDQTEN8bwc
FZUXyT8sg7sYB5yR2Rn+pm0VDnkaZ6en5bbeIW2UGOLABygp8BxDFLUZTwmT5bod
qrAdoUZX5wAEo2fhpvzzmPtH/bSx50aU1wAzvM0CgYB6ASf4TL7mcTYDEEitOVYi
6QHP7yhqZHAezcgRupkURlLSaziFJ6oVW+RTLA81LxtrLXzpcIY7JWsB8arZrazI
b/jLPrDyU+ALJAeaMyEWXV/uAJF3HhLy3HCoTPwe7ztNEK8iFX5hXXOf9tOVDif8
JbpMF7Uksx2lRPVDuXylsQKBgQDdXxMZH7lWDQvOIbxPm1iJkd/qQ9aPchWBpZwM
oQ3hVjCvHhK4cJD55z4iCoiMP3iQ068FYE7EnYfbTdELa2vRXcZcBRpcCNp0bDt/
eL6R7XtQJdo4EaudPnfRSuLTARw7Cyctr2y4T2MrT2Us8oXUDMU4qasRvPJgDG5d
DyD58QKBgFye0S1BDK72JZY2hmLeB7A5mO/ScUWdZHbUODvpEl57myAtf9qEX8sX
8YwLKVu4OkuIepz65ewyKdL7w8ijmwNcvcIN4wd1+MtXyA5Mbg8tyyrHkRagLjo8
g1UEr6GryQJZKfpyreDxSEOcLbzbB8odU7g2dernrIcbS/OJfV6t
-----END RSA PRIVATE KEY-----`)

    cert, err := tls.X509KeyPair([] byte(rsaCertPEM), [] byte(rsaKeyPEM))
	if err != nil {
		fmt.Printf(err.Error())
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true,
				Certificates: []tls.Certificate{cert},
			},
		},
	}
    resp, err := client.Get("https://52.56.69.220:8443/swagger/")
	if err != nil {
		fmt.Println(err)
	}

	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Printf("%v\n", resp.Status)
	fmt.Printf(string(htmlData))
}

