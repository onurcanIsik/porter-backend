package handler

import (
	"crypto/subtle"
	"encoding/json"
	"net/http"
	"porter/config"
	"porter/internal/service"
	"porter/models"
	apprand "porter/pkg/crypto"
	"time"
)

var req struct {
	IDToken string `json:"id_token"`
}

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	googleConfig := config.GoogleConfig()

	state, err := apprand.GenerateRandomSafeString(32)
	if err != nil {
		http.Error(w, "Failed to generate state", http.StatusInternalServerError)
		return
	}

	url := googleConfig.AuthCodeURL(state)

	http.SetCookie(w, &http.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Expires:  time.Now().Add(10 * time.Minute),
		HttpOnly: true,
		Path:     "/",
	})

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func (h *UserHandler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	getStateFromUrl := r.URL.Query().Get("state")

	getStateFromCookie, err := r.Cookie("oauthstate")
	if err != nil {
		http.Error(w, "State cookie not found", http.StatusBadRequest)
		return
	}

	if subtle.ConstantTimeCompare([]byte(getStateFromUrl), []byte(getStateFromCookie.Value)) != 1 {
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	googleCode := r.URL.Query().Get("code")
	if googleCode == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	googleConfig := config.GoogleConfig()
	token, err := googleConfig.Exchange(r.Context(), googleCode)
	if err != nil {
		http.Error(w, "Failed to exchange code for token", http.StatusInternalServerError)
		return
	}

	client := googleConfig.Client(r.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var googleUser struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	err = json.NewDecoder(resp.Body).Decode(&googleUser)
	if err != nil {
		http.Error(w, "Failed to decode user info", http.StatusInternalServerError)
		return
	}

	accessToken, refreshToken, err := h.userService.LoginOrRegister(&models.UserModel{
		Provider:       "google",
		ProviderId:     googleUser.ID,
		UserMail:       googleUser.Email,
		UserName:       googleUser.Name,
		UserProfileUrl: "",
		UserJobTitle:   "",
		UserDeviceId:   "",
		UserCreatedAt:  time.Now(),
		UserUpdatedAt:  time.Now(),
	})
	if err != nil {
		http.Error(w, "Failed to login or register user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
