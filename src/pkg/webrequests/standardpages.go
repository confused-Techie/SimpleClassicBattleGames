package webrequests

import (
  "html/template"
  "net/http"
  "fmt"
  models "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/models"
  games "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/games"
  users "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/users"
)

var tmpl = make(map[string]*template.Template)

func returnTemplate(t string) string {
  return "./templates/" + t
}

func returnSubTemplate(t string) string {
  return "./templates/components/" + t
}

func errorPage(w http.ResponseWriter, r *http.Request, err string) {

  data := models.PageTemplate{
    Title: "Something went wrong.",
    Data: err,
  }

  templateArray := []string{
    returnTemplate("errorPage.go.html"),
    returnSubTemplate("head.go.html"),
  }

  tmpl["errorPage.html"] = template.Must(template.ParseFiles(templateArray...))
  templateError := tmpl["errorPage.html"].Execute(w, data)
  StandardPageError(templateError)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {

  data := models.PageTemplate{
    Title: "SimpleClassicBattleGames - Home",
    Data: games.GetGameCollection(),
  }

  templateArray := []string{
    returnTemplate("homePage.go.html"),
    returnSubTemplate("head.go.html"),
  }

  tmpl["homePage.html"] = template.Must(template.ParseFiles(templateArray...))

  templateError := tmpl["homePage.html"].Execute(w, data)
  StandardPageError(templateError)
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    // Parse our multipart form, 10 << 20 specifies a maximum upload of 10 MB files.
    r.ParseForm()

    user := r.Form["username"][0] // we only care about the first found valid instance
    pass := r.Form["password"][0]

    // now with the user we can check if they are a valid user.
    success, signInMsg := users.AttemptSignIn(user, pass)

    if success {
      // since the sign in was successful we want to add the cookies and redirect
      userDetails := users.GetUserDetails(user)

      CreateSignInCookie(userDetails.UserName, w)

      http.Redirect(w, r, "/", http.StatusSeeOther)
    } else {
      // with a bad sign in we need to respond.
      http.Redirect(w, r, "/sign-in?sign-in-error="+signInMsg, http.StatusSeeOther)
    }

  } else {
    fmt.Println("Header during Get", r.Header.Get("SignInMsg"))
    // For GET Requests.
    data := models.PageTemplate{
      Title: "SimpleClassicBattleGames - Sign In",
      SpecialMsg: GetSignInMsg(r),
    }

    templateArray := []string{
      returnTemplate("signin.go.html"),
      returnSubTemplate("head.go.html"),
    }

    tmpl["signin.html"] = template.Must(template.ParseFiles(templateArray...))

    templateError := tmpl["signin.html"].Execute(w, data)
    StandardPageError(templateError)
  }
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    r.ParseForm()

    user := r.Form["username"][0]
    pass := r.Form["password"][0]

    success, rtrn := users.CreateUser(user, pass)

    if success {
      CreateSignInCookie(user, w)
      http.Redirect(w, r, "/", http.StatusSeeOther)
    } else {
      http.Redirect(w, r, "/create-user?create-user-error="+rtrn, http.StatusSeeOther)

    }
  } else {
    data := models.PageTemplate{
      Title: "SimpleClassicBattleGames - Create User",
      SpecialMsg: GetCreateUserMsg(r),
    }

    templateArray := []string{
      returnTemplate("createuser.go.html"),
      returnSubTemplate("head.go.html"),
    }

    tmpl["createuser.html"] = template.Must(template.ParseFiles(templateArray...))

    templateError := tmpl["createuser.html"].Execute(w, data)
    StandardPageError(templateError)
  }
}
