package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/GeorgeHN/email-backend/app/db"
	"github.com/GeorgeHN/email-backend/app/ftp"
	"github.com/GeorgeHN/email-backend/app/models"
	"github.com/GeorgeHN/email-backend/app/utils"
)

const (
	DATABASE_URI  = "mongodb+srv://j:rootroot@cluster0.rj0tg.mongodb.net/"
	DATABASE_NAME = "email_test"
	FTP_ADDR      = "ftp.zkaia.com:21"
	FTP_USER      = "zkaiacom"
	FTP_PASS      = "log.Fatal(1$)"
	FTP_PATH      = "email.zkaia.com/clients"
	ServerError   = http.StatusInternalServerError
	BadRequest    = http.StatusBadRequest
	Created       = http.StatusCreated
	OK            = http.StatusOK
	NotFound      = http.StatusNotFound
	NotAcceptable = http.StatusNotAcceptable
)

func NewAdmin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var ad models.Admin

	err := utils.ReadJSON(w, r, &ad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	databaseConfig := db.Config{
		URI:      DATABASE_URI,
		Database: DATABASE_NAME,
	}

	err = db.NewDBConn(databaseConfig).InsertAdmin(ad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	response.Error = false
	response.Message = "Admin added successfuly"

	utils.WriteJSON(w, r, response, http.StatusCreated)

}

// Create new client
func NewClientEP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var client models.Client

	err := utils.ReadJSON(w, r, &client)
	if err != nil {
		http.Error(w, http.StatusText(ServerError), ServerError)
		return
	}

	config := db.Config{
		URI:      DATABASE_URI,
		Database: DATABASE_NAME,
	}

	err = db.NewDBConn(config).InsertClient(client)
	if err != nil {
		http.Error(w, err.Error(), NotAcceptable)
		return
	}

	ftpC := ftp.Config{
		Addr:     FTP_ADDR,
		User:     FTP_USER,
		Password: FTP_PASS,
	}

	err = ftp.NewFTPServer(ftpC).CreateDir(fmt.Sprintf("%s/%s", FTP_PATH, client.Company))
	if err != nil {
		fmt.Println("Here")
		http.Error(w, err.Error(), ServerError)
		return
	}

	err = ftp.NewFTPServer(ftpC).CreateDir(fmt.Sprintf("%s/%s/%s", FTP_PATH, client.Company, "campaings"))
	if err != nil {
		fmt.Println("over here")

		http.Error(w, err.Error(), ServerError)
		return
	}

	var response struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	response.Error = false
	response.Message = "Client added successfuly"

	utils.WriteJSON(w, r, response, http.StatusCreated)

}

func DeleteClientEP(w http.ResponseWriter, r *http.Request) {
	//  We will going to delete everything
	err := ftp.NewFTPServer(ftp.Config{Addr: FTP_ADDR, User: FTP_USER, Password: FTP_PASS}).DelDir(fmt.Sprintf("%s/%s", FTP_PATH, r.URL.Query().Get("cpn")))
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	err = db.NewDBConn(db.Config{URI: DATABASE_URI, Database: DATABASE_NAME}).DeleteClient(r.URL.Query().Get("i"))
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	var response struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	response.Error = false
	response.Message = "Clients founded"

	utils.WriteJSON(w, r, response, OK)
}

func GetClients(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	user, err := db.NewDBConn(db.Config{URI: DATABASE_URI, Database: DATABASE_NAME}).GetClients()
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	var response struct {
		Error   bool             `json:"error"`
		Message string           `json:"message"`
		Clients []*models.Client `json:"clients"`
	}
	response.Error = false
	response.Message = "Clients founded"
	response.Clients = user

	utils.WriteJSON(w, r, response, OK)
}

func GetClient(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	client, err := db.NewDBConn(db.Config{URI: DATABASE_URI, Database: DATABASE_NAME}).GetClient(r.URL.Query().Get("zkmi"))
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	var response struct {
		Error   bool           `json:"error"`
		Message string         `json:"message"`
		Clients *models.Client `json:"client"`
	}
	response.Error = false
	response.Message = "Clients founded"
	response.Clients = client

	utils.WriteJSON(w, r, response, OK)
}

// Create new campaing
func InsertCampaingEP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var campaing models.Campaing

	err := utils.ReadJSON(w, r, &campaing)
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	config := db.Config{
		URI:      DATABASE_URI,
		Database: DATABASE_NAME,
	}

	err = db.NewDBConn(config).InsertCampaing(campaing)
	if err != nil {
		http.Error(w, err.Error(), NotAcceptable)
		return
	}

	err = ftp.NewFTPServer(ftp.Config{Addr: FTP_ADDR, User: FTP_USER, Password: FTP_PASS}).CreateDir(fmt.Sprintf("%s/%s/campaings/%s", FTP_PATH, campaing.CompanyName, campaing.CampaingName))
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	var response struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	response.Error = false
	response.Message = "Campaing created successfuly"

	utils.WriteJSON(w, r, response, http.StatusCreated)

}

func GetCampaings(w http.ResponseWriter, r *http.Request) {

	campaings, err := db.NewDBConn(db.Config{URI: DATABASE_URI, Database: DATABASE_NAME}).GetCampaings(r.URL.Query().Get("ains"))
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	var response struct {
		Error     bool               `json:"error"`
		Message   string             `json:"message"`
		Campaings []*models.Campaing `json:"campaings"`
	}
	response.Error = false
	response.Message = "Clients founded"
	response.Campaings = campaings

	utils.WriteJSON(w, r, response, OK)

}

func DeleteCampaingEP(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	cpn := r.URL.Query().Get("cpn")
	name := r.URL.Query().Get("name")

	err := db.NewDBConn(db.Config{URI: DATABASE_URI, Database: DATABASE_NAME}).DeleteCampaing(id)
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	// Delete the ftp folders
	err = ftp.NewFTPServer(ftp.Config{Addr: FTP_ADDR, User: FTP_USER, Password: FTP_PASS}).DelDir(fmt.Sprintf("%s/%s/campaings/%s", FTP_PATH, cpn, name))
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	var response struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	response.Error = false
	response.Message = "Campaing created successfuly"

	utils.WriteJSON(w, r, response, http.StatusCreated)

}

// Create new schedule
func InsertScheduleEP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var schedule models.Schedule

	err := utils.ReadJSON(w, r, &schedule)
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	err = db.NewDBConn(db.Config{URI: DATABASE_URI, Database: DATABASE_NAME}).InsertSchedules(schedule)
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	err = ftp.NewFTPServer(ftp.Config{Addr: FTP_ADDR, User: FTP_USER, Password: FTP_PASS}).CreateDir(fmt.Sprintf("%s/%s/campaings/%s/%s", FTP_PATH, schedule.CompanyName, schedule.CampaingName, schedule.Name))
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	err = ftp.NewFTPServer(ftp.Config{Addr: FTP_ADDR, User: FTP_USER, Password: FTP_PASS}).CreateDir(fmt.Sprintf("%s/%s/campaings/%s/%s/template", FTP_PATH, schedule.CompanyName, schedule.CampaingName, schedule.Name))
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	err = ftp.NewFTPServer(ftp.Config{Addr: FTP_ADDR, User: FTP_USER, Password: FTP_PASS}).CreateDir(fmt.Sprintf("%s/%s/campaings/%s/%s/audience", FTP_PATH, schedule.CompanyName, schedule.CampaingName, schedule.Name))
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	var response struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	response.Error = false
	response.Message = "schedule created successfuly"

	utils.WriteJSON(w, r, response, http.StatusCreated)

}

func StoreTemplate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "multipart/form-data")

	r.ParseMultipartForm(128)
	_, file, err := r.FormFile("template")
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	cpn := r.FormValue("cpn")
	cn := r.FormValue("cn")
	name := r.FormValue("name")

	msg, err := ftp.NewFTPServer(ftp.Config{Addr: FTP_ADDR, User: FTP_USER, Password: FTP_PASS}).AddFile(fmt.Sprintf("%s/%s/campaings/%s/%s/template/", FTP_PATH, cpn, cn, name), file)
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	var response struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	response.Error = false
	response.Message = msg

	utils.WriteJSON(w, r, response, http.StatusCreated)
}

func StoreAudience(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "multipart/form-data")

	r.ParseMultipartForm(128)
	_, file, err := r.FormFile("audience")
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	cpn := r.FormValue("cpn")
	cn := r.FormValue("cn")
	name := r.FormValue("name")

	msg, err := ftp.NewFTPServer(ftp.Config{Addr: FTP_ADDR, User: FTP_USER, Password: FTP_PASS}).AddFile(fmt.Sprintf("%s/%s/campaings/%s/%s/audience", FTP_PATH, cpn, cn, name), file)
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	var response struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	response.Error = false
	response.Message = msg

	utils.WriteJSON(w, r, response, http.StatusCreated)
}

func GetSchedules(w http.ResponseWriter, r *http.Request) {
	schedules, err := db.NewDBConn(db.Config{URI: DATABASE_URI, Database: DATABASE_NAME}).GetSchedules(r.URL.Query().Get("ins"))
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	var response struct {
		Error     bool               `json:"error"`
		Message   string             `json:"message"`
		Schedules []*models.Schedule `json:"schedules"`
	}
	response.Error = false
	response.Message = "Clients founded"
	response.Schedules = schedules

	utils.WriteJSON(w, r, response, OK)
}

func DeleteScheduleEP(w http.ResponseWriter, r *http.Request) {

	err := ftp.NewFTPServer(ftp.Config{Addr: FTP_ADDR, User: FTP_USER, Password: FTP_PASS}).DelDir(fmt.Sprintf("%s/%s/campaings/%s/%s", FTP_PATH, r.URL.Query().Get("cpn"), r.URL.Query().Get("cn"), r.URL.Query().Get("name")))
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	err = db.NewDBConn(db.Config{URI: DATABASE_URI, Database: DATABASE_NAME}).DeleteSchedules(r.URL.Query().Get("i"))
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}

	var response struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	response.Error = false
	response.Message = "Clients founded"

	utils.WriteJSON(w, r, response, OK)
}

// StoreFile
func AddFileEP(w http.ResponseWriter, r *http.Request) {}

func DeleteFileEP(w http.ResponseWriter, r *http.Request) {}

func GetFileList(w http.ResponseWriter, r *http.Request) {}

// Image Middleware
func ServeImage(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("@inside the actual func")

	data, err := downloadFile("https://images.pexels.com/photos/16664503/pexels-photo-16664503/free-photo-of-woman-in-hat-posing-on-shallow-water.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2")
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	fmt.Println("after")

	w.Write(data)

}

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
