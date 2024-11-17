package mail

// import (
// 	"os"
// 	"testing"

// 	"github.com/joho/godotenv"
// 	_ "github.com/joho/godotenv/autoload"
// 	"github.com/nvr-ai/go-mail/templates"
// )

// func TestSend(t *testing.T) {
// 	godotenv.Load(".env")
// 	err := Send(os.Getenv("RESEND_API_KEY"), templates.RegistrationTemplate, Recipients{
// 		To: []string{"m@matthewdavis.io"},
// 	}, "Welcome to infraregistry.com! (test)", map[string]string{
// 		"link": "https://infraregistry.com",
// 		"code": "123456",
// 	})

// 	if err != nil {
// 		t.Error(err)
// 	}
// }
