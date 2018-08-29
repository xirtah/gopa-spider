// Generated by ego.
// DO NOT EDIT

package common

import (
	"fmt"
	"github.com/xirtah/gopa/core/http"
	"html"
	"io"
	"net/http"
)

var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
func Body(w io.Writer) error {
	_, _ = io.WriteString(w, "\n\n</head>\n<body class=\"tm-background\">\n")
	return nil
}
func Footer(w io.Writer) error {
	_, _ = io.WriteString(w, "\n\n<div class=\"tm-footer\">\n    <div class=\"uk-container uk-container-center uk-text-center\">\n        <br>\n        <hr class=\"uk-article-divider\" >\n        <ul class=\"uk-subnav uk-subnav-line uk-flex-center\">\n            <li><a href=\"http://github.com/xirtah/gopa\">GitHub</a></li>\n            <li><a href=\"http://github.com/xirtah/gopa/issues\">Issues</a></li>\n            <li><a href=\"https://github.com/xirtah/gopa/releases\">Releases</a></li>\n            <li><a href=\"http://github.com/xirtah/gopa/blob/master/CHANGES.md\">Changelog</a></li>\n        </ul>\n\n        <div class=\"uk-panel\">\n            <p>Licensed under <a target=\"_blank\" href=\"https://github.com/xirtah/gopa/blob/master/LICENSE\">Apache License, Version 2.0</a>.</p>\n            <a href=\"/admin/\"><img src=\"/static/assets/img/logo.svg\" height=\"30\" title=\"GOPA\" alt=\"GOPA\"></a>\n        </div>\n\n    </div>\n</div>\n</body>\n<script src=\"/static/assets/js/ie_detect.js\"></script>\n</html>\n")
	return nil
}
func Head(w io.Writer, title string, customHeaderBlock string) error {
	_, _ = io.WriteString(w, "\n<!DOCTYPE html>\n<!--[if lt IE 7 ]> <html class=\"no-js ie6\" lang=\"en-US\"> <![endif]-->\n<!--[if IE 7 ]>    <html class=\"no-js ie7\" lang=\"en-US\"> <![endif]-->\n<!--[if IE 8 ]>    <html class=\"no-js ie8\" lang=\"en-US\"> <![endif]-->\n<!--[if (gte IE 9)|!(IE)]><!--> <html lang=\"en-US\"> <!--<![endif]-->\n<html>\n<head>\n  <title>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(title)))
	_, _ = io.WriteString(w, " - GOPA</title>\n\n  <meta content=IE=7 http-equiv=X-UA-Compatible>\n  <meta content=text/html;charset=utf-8 http-equiv=content-type>\n\n  <meta name=\"robots\" content=\"all\">\n  <meta name=\"license\" content=\"keep-copyright-footprint,no-KPI-shit,respect-first\">\n  <meta name=\"creator\" content=\"medcl\">\n  <meta name=\"generator\" content=\"https://github.com/xirtah/gopa\">\n  <meta name=\"copyright\" content=\"Apache License, Version 2.0\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n\n  <link rel=\"stylesheet\" href=\"/static/assets/uikit-2.27.1/css/uikit.min.css\" />\n  <link rel=\"icon\" href=\"/static/assets/img/favicon.ico\" type=\"image/x-icon\" />\n  <link rel=\"shortcut icon\" href=\"/static/assets/img/favicon.ico\" type=\"image/x-icon\" />\n\n  <script src=\"/static/assets/js/jquery.min.js\"></script>\n  <script src=\"/static/assets/uikit-2.27.1/js/uikit.min.js\"></script>\n  <script src=\"/static/assets/uikit-2.27.1/js/core/offcanvas.min.js\"></script>\n\n  <script src=\"/static/assets/js/d3-4.0.min.js\"></script>\n  <script src=\"/static/assets/js/vue.min.js\"></script>\n\n  <meta charset=\"utf-8\">\n  ")
	if len(customHeaderBlock) > 0 {
		_, _ = io.WriteString(w, "\n  ")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(customHeaderBlock)))
		_, _ = io.WriteString(w, "\n  ")
	}
	_, _ = io.WriteString(w, "\n")
	return nil
}
func Nav(w http.ResponseWriter, r *http.Request, current string) error {
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n\n<nav class=\"tm-navbar uk-navbar uk-navbar-attached\">\n  <div class=\"uk-container uk-container-center\">\n    ")

	logoUrl := "/admin/"
	if len(navs) > 0 {
		logoUrl = navs[0].url
	}

	_, _ = io.WriteString(w, "\n    <a class=\"uk-navbar-brand uk-hidden-small\" href=\"")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(logoUrl)))
	_, _ = io.WriteString(w, "\"><img class=\"uk-margin uk-margin-remove\" src=\"/static/assets/img/logo.svg\" width=\"90\" height=\"30\" title=\"GOPA\" alt=\"GOPA\"></a>\n\n    <ul class=\"uk-navbar-nav uk-hidden-small\">\n\n      ")
	if len(navs) > 0 {
		for _, obj := range navs {

			_, _ = io.WriteString(w, "\n      <li ")
			_, _ = fmt.Fprint(w, NavCurrent(current, obj.name))
			_, _ = io.WriteString(w, " ><a href=\"")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(obj.url)))
			_, _ = io.WriteString(w, "\">")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(obj.displayName)))
			_, _ = io.WriteString(w, "</a></li>\n      ")

		}
	}

	_, _ = io.WriteString(w, "\n    </ul>\n\n    <div class=\"uk-navbar-flip\">\n\n      <a href=\"#search-modal\" data-uk-modal=\"\" class=\"uk-navbar-toggle uk-navbar-toggle-alt uk-visible-small\"></a>\n      <div id=\"search-modal\" class=\"uk-modal\">\n        <div class=\"uk-modal-dialog\">\n          <a class=\"uk-modal-close uk-close\"></a>\n          <form  action=\"/\">\n            <input name=\"q\" type=\"text\" placeholder=\"Please type to search\" class=\"uk-form-large uk-width-1-1\">\n          </form>\n        </div>\n      </div>\n\n      <ul class=\"uk-navbar-nav uk-hidden-small\">\n        <li>\n          <div class=\"uk-form-icon\">\n              <form  action=\"/\">\n                <input name=\"q\" type=\"text\" placeholder=\"Please type to search\" class=\"uk-form-large uk-width-1-1\">\n              </form>\n          </div>\n        </li>\n        <li class=\"uk-parent\" data-uk-dropdown=\"\" aria-haspopup=\"true\" aria-expanded=\"false\">\n          <a href=\"\"><i class=\"uk-icon-plus-square\"></i>&nbsp;Menu</a>\n\n          <div class=\"uk-dropdown uk-dropdown-navbar uk-dropdown-bottom\" aria-hidden=\"true\" tabindex=\"\" style=\"top: 40px; left: -119px;\">\n            <ul class=\"uk-nav uk-nav-navbar\">\n              <li class=\"uk-nav-header\"><i class=\"uk-icon-tasks\"></i>&nbsp;Tasks</li>\n              <li><a href=\"#create-task-modal\" data-uk-modal=\"\">Create a task</a></li>\n              ")
	if api.IsAuthEnable() {
		user, _ := api.GetLoginInfo(w, r)

		_, _ = io.WriteString(w, "\n                <li class=\"uk-nav-divider\"></li>\n                <li><a href=\"/auth/logout/\"><i class=\"uk-icon-power-off\"></i> Logout(")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(user)))
		_, _ = io.WriteString(w, ")</a></li>\n              ")
	}
	_, _ = io.WriteString(w, "\n            </ul>\n          </div>\n\n        </li>\n      </ul>\n\n    </div>\n\n    <a href=\"#tm-offcanvas\" class=\"uk-navbar-toggle uk-visible-small\" data-uk-offcanvas=\"\"></a>\n\n    <div class=\"uk-navbar-brand uk-navbar-center uk-visible-small\"><img src=\"/static/assets/img/logo.svg\" width=\"90\" height=\"30\" title=\"GOPA\" alt=\"GOPA\"></div>\n\n  </div>\n</nav>\n\n<div style=\"height: 15px;clear: both\"></div>\n\n")
	OffCanvas(w, current)
	_, _ = io.WriteString(w, "\n\n")
	return nil
}
func OffCanvas(w io.Writer, current string) error {
	_, _ = io.WriteString(w, "\n<div id=\"tm-offcanvas\" class=\"uk-offcanvas\">\n\n  <div class=\"uk-offcanvas-bar\">\n\n    <ul class=\"uk-nav uk-nav-offcanvas uk-nav-parent-icon\" data-uk-nav=\"{multiple:true}\">\n      <li class=\"uk-parent uk-active uk-open\" aria-expanded=\"true\"><a href=\"#\">Menu</a>\n        <div style=\"overflow: hidden; height: auto; position: relative;\"><ul class=\"uk-nav-sub\" role=\"menu\">\n          ")
	if len(navs) > 0 {
		for _, obj := range navs {

			_, _ = io.WriteString(w, "\n          <li ")
			_, _ = fmt.Fprint(w, NavCurrent(current, obj.name))
			_, _ = io.WriteString(w, " ><a href=\"")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(obj.url)))
			_, _ = io.WriteString(w, "\">")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(obj.displayName)))
			_, _ = io.WriteString(w, "</a></li>\n          ")

		}
	}

	_, _ = io.WriteString(w, "\n        </ul></div>\n      </li>\n      <li class=\"uk-nav-divider\"></li>\n      <li><a href=\"http://github.com/xirtah/gopa\">GitHub</a></li>\n      <li><a href=\"http://github.com/xirtah/gopa/issues\">Issues</a></li>\n      <li><a href=\"https://github.com/xirtah/gopa/releases\">Releases</a></li>\n      <li><a href=\"http://github.com/xirtah/gopa/blob/master/CHANGES.md\">Changelog</a></li>\n    </ul>\n\n  </div>\n\n</div>\n")
	return nil
}
