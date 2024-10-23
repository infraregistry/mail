package templates

import "testing"

func TestGetTemplate(t *testing.T) {
	template, err := GetTemplate("templates/output.html", map[string]string{"link": "https://localhost:5000/verify?token=test"})
	if err != nil {
		t.Errorf("GetTemplate() error = %v", err)
		return
	}

	if template == "" {
		t.Errorf("GetTemplate() template is empty")
		return
	}

	t.Logf("GetTemplate() template = %v", template)
}
