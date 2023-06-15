package mailer

import (
	"fmt"
	"io"
	"net/http"

	"github.com/GeorgeHN/email-backend/app/models"
	"gopkg.in/gomail.v2"
)

type MAILER struct {
	config Config
}

func NewMailer(c Config) *MAILER {
	return &MAILER{
		config: c,
	}
}

// Send Simple Email
func (s *MAILER) SendEmail(user string, campaing string, template string, to string, subject string, audience []*models.Audience) error {

	HTML, err := downloadFile(fmt.Sprintf("https://email.zkaia.com/clients/%s/campaings/%s/templates/%s.html", user, campaing, template))
	if err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetBody("text/html", string(HTML))
	for _, person := range audience {
		mailer.SetHeader("From", s.config.User)
		mailer.SetHeader("To", person.Email)
		mailer.SetHeader("Subject", fmt.Sprintf("%s %s - %s", person.First, person.Last, subject))

		dialer := gomail.NewDialer(s.config.Host, s.config.Port, s.config.User, s.config.Password)

		if err := dialer.DialAndSend(mailer); err != nil {
			return err
		}
		mailer.Reset()
	}

	return nil
}

// Send email with attachments
func (s *MAILER) SendEmailWithDocs(user string, campaing string, template string, to string, subject string, docs []string, audience []*models.Audience) error {
	HTML, err := downloadFile(fmt.Sprintf("https://email.zkaia.com/clients/%s/campaings/%s/templates/%s.html", user, campaing, template))
	if err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetBody("text/html", string(HTML))

	for _, people := range audience {
		mailer.SetHeader("From", s.config.User)
		mailer.SetHeader("To", people.Email)
		mailer.SetHeader("Subject", fmt.Sprintf("%s %s - %s", people.First, people.Last, subject))

		for _, doc := range docs {
			downloadFile, err := downloadFile(fmt.Sprintf("https://email.zkaia.com/clients/%s/campaings/%s/docs/%s", user, campaing, doc))
			if err != nil {
				return err
			}

			mailer.Attach(doc, gomail.SetCopyFunc(func(w io.Writer) error {
				_, err := w.Write(downloadFile)
				if err != nil {
					return err
				}
				return nil
			}))

		}

		dialer := gomail.NewDialer(s.config.Host, s.config.Port, s.config.User, s.config.Password)

		if err := dialer.DialAndSend(mailer); err != nil {
			return err
		}
		mailer.Reset()
	}

	return nil

}

// Welcome to new clients
func (s *MAILER) SendWelcomeEmail(to string, subject string) error {

	HTML, err := downloadFile("https://email.zkaia.com/templates/welcome.html")
	if err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", s.config.User)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)

	mailer.SetBody("text/html", string(HTML))

	dialer := gomail.NewDialer(s.config.Host, s.config.Port, s.config.User, s.config.Password)

	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}

// Weakly reports
func (s *MAILER) WeeklyRepports(clients []*models.Client, campaing string) error {
	HTML, err := downloadFile("https://email.zkaia.com/templates/statistics.html")
	if err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetBody("text/html", string(HTML))

	for _, client := range clients {
		mailer.SetHeader("From", s.config.User)
		mailer.SetHeader("To", client.Email)
		mailer.SetHeader("Subject", fmt.Sprintf("%s - Your weekly reports are here", client.Name))
		downloadFile, err := downloadFile(fmt.Sprintf("https://email.zkaia.com/clients/%s/campaings/%s/repport/repport.pdf", client.Name, campaing))
		if err != nil {
			return err
		}

		mailer.Attach(fmt.Sprintf("weekly_repport_%s.pdf", client.Name), gomail.SetCopyFunc(func(w io.Writer) error {
			_, err := w.Write(downloadFile)
			if err != nil {
				return err
			}
			return nil
		}))

		dialer := gomail.NewDialer(s.config.Host, s.config.Port, s.config.User, s.config.Password)

		if err := dialer.DialAndSend(mailer); err != nil {
			return err
		}
		mailer.Reset()
	}
	return nil

}

// Report Errors

func downloadFile(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
