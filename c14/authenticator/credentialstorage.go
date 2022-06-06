package authenticator


type CredentialStorage interface {
	GetSecretByAppid(appId string) string
}

type MySQLSto struct {

}

func NewMySQlSto() *MySQLSto {
	return &MySQLSto{}
}

func (m *MySQLSto) GetSecretByAppid(appId string) string {
	return "b"
}
