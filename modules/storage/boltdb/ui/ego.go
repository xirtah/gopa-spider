// Generated by ego.
// DO NOT EDIT

package ui

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/xirtah/gopa-spider/modules/ui/common"
	"html"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"unsafe"
)

var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
func Error(w io.Writer, err error) error {
	_, _ = io.WriteString(w, "\n\n<!DOCTYPE html>\n<html lang=\"en\">\n  <head>\n    <meta charset=\"utf-8\">\n    <title>boltd</title>\n  </head>\n\n  <body class=\"error\">\n    <div class=\"container\">\n      <div class=\"header\">\n        <h3 class=\"text-muted\">Error</h3>\n      </div>\n\n      An error has occurred: ")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(err)))
	_, _ = io.WriteString(w, "\n    </div> <!-- /container -->\n  </body>\n</html>\n")
	return nil
}
func Index(w io.Writer) error {
	_, _ = io.WriteString(w, "\n\n<!DOCTYPE html>\n<html lang=\"en\">\n  <head>\n    <meta http-equiv=\"refresh\" content=\"0; url=page\">\n  </head>\n\n  <body>redirecting...</body>\n</html>\n")
	return nil
}
func nav(w io.Writer, tx *bolt.Tx) error {
	_, _ = io.WriteString(w, "\n\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n\n")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(filepath.Base(tx.DB().Path()))))
	return nil
}
func Page(w http.ResponseWriter, r *http.Request, tx *bolt.Tx, indexes []int, directID int, showUsage bool) error {
	_, _ = io.WriteString(w, "\n\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n\n")

	p, ids, err := find(tx, directID, indexes)
	if err != nil {
		return err
	}

	// Generate page stats.
	pageSize := tx.DB().Info().PageSize
	stats := p.stats(pageSize)

	// Generate histogram of all nested page usage.
	var histogram map[int]int
	if showUsage {
		histogram = usage(tx, p.id)
	}

	_, _ = io.WriteString(w, "\n\n")
	common.Head(w, filepath.Base(tx.DB().Path()), "")
	_, _ = io.WriteString(w, "\n\n  <link rel=\"stylesheet\" href=\"/static/assets/css/tasks.css\" />\n  <script src=\"/static/assets/js/jquery.timeago.js\"></script>\n  <script src=\"/static/assets/js/page/tasks.js\"></script>\n <style>\n    table {\n      border-collapse:collapse;\n      word-break:break-all; word-wrap:break-word;\n    }\n\n    table, th, td {\n      border: 1px solid black;\n    }\n\n    th, td {\n      min-width: 100px;\n      padding: 2px 5px;\n    }\n  </style>\n\n")
	common.Body(w)
	_, _ = io.WriteString(w, "\n")
	common.Nav(w, r, "BoltDB")
	_, _ = io.WriteString(w, "\n\n<div class=\"uk-container uk-container-center\">\n\n    <div class=\"uk-grid\" data-uk-grid-margin>\n        <div class=\"uk-width-large-1-1 uk-visible-large\">\n            <div class=\"uk-alert\" >Database: ")
	nav(w, tx)
	_, _ = io.WriteString(w, "</div>\n\n\n          <h2>\n                ")
	for i, id := range ids {
		_, _ = io.WriteString(w, "\n                  ")
		if i > 0 {
			_, _ = io.WriteString(w, "&raquo;")
		}
		_, _ = io.WriteString(w, "\n                  <a href=\"")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(pagelink(indexes[:i+1]))))
		_, _ = io.WriteString(w, "\">#")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(id)))
		_, _ = io.WriteString(w, "</a>\n                ")
	}
	_, _ = io.WriteString(w, "\n              </h2>\n\n              <h3>Page Information</h3>\n              <p>\n                <strong>ID:</strong> ")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(comma(int(p.id)))))
	_, _ = io.WriteString(w, "<br/>\n                <strong>Type:</strong> ")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(fmt.Sprintf("%s (%x)", p.typ(), p.flags))))
	_, _ = io.WriteString(w, "<br/>\n                <strong>Overflow:</strong> ")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(p.overflow)))
	_, _ = io.WriteString(w, "<br/><br/>\n\n                <strong>Alloc:</strong> ")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(comma(stats.alloc))))
	_, _ = io.WriteString(w, "<br/>\n                <strong>In Use:</strong> ")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(comma(stats.inuse))))
	_, _ = io.WriteString(w, "<br/>\n                <strong>Utilization:</strong> ")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(fmt.Sprintf("%.2f%%", stats.utilization*100))))
	_, _ = io.WriteString(w, "<br/>\n              </p>\n\n              ")
	if (p.flags & branchPageFlag) != 0 {
		_, _ = io.WriteString(w, "\n                <h3>Branch Elements (")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(p.count)))
		_, _ = io.WriteString(w, ")</h3>\n                <table>\n                  <thead>\n                    <tr>\n                      <th align=\"left\">Key</th>\n                      <th align=\"left\">Page ID</th>\n                      <th align=\"left\">Size (k)</th>\n                      <th align=\"center\">%%Util</th>\n                    </tr>\n                  </thead>\n                  <tbody>\n                    ")
		for i := uint16(0); i < p.count; i++ {
			_, _ = io.WriteString(w, "\n                      ")
			e := p.branchPageElement(i)
			_, _ = io.WriteString(w, "\n                      ")
			subpage := pageAt(tx, e.pgid)
			_, _ = io.WriteString(w, "\n                      ")
			substats := subpage.stats(pageSize)
			_, _ = io.WriteString(w, "\n                      <tr>\n                        <td>")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(trunc(tostr(e.key()), 150))))
			_, _ = io.WriteString(w, "</td>\n                        <td><a href=\"")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(subpagelink(indexes, int(i)))))
			_, _ = io.WriteString(w, "\">")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(e.pgid)))
			_, _ = io.WriteString(w, "</a></td>\n                        <td>")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(len(e.key()))))
			_, _ = io.WriteString(w, "</td>\n                        <td align=\"right\">")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(fmt.Sprintf("%.2f%%", substats.utilization*100))))
			_, _ = io.WriteString(w, "</td>\n                      </tr>\n                    ")
		}
		_, _ = io.WriteString(w, "\n                  </tbody>\n                </table>\n\n              ")
	} else if (p.flags & leafPageFlag) != 0 {
		_, _ = io.WriteString(w, "\n                <h3>Leaf Elements (")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(p.count)))
		_, _ = io.WriteString(w, ")</h3>\n                <table>\n                  <thead>\n                    <tr>\n                      <th align=\"left\">Key</th>\n                      <th align=\"left\">Value</th>\n                      <th align=\"left\">Size (k/v)</th>\n                      <th align=\"center\">%%Util</th>\n                    </tr>\n                  </thead>\n                  <tbody>\n                    ")
		for i := uint16(0); i < p.count; i++ {
			_, _ = io.WriteString(w, "\n                      ")
			e := p.leafPageElement(i)
			_, _ = io.WriteString(w, "\n                      ")
			if (e.flags & bucketLeafFlag) != 0 {
				_, _ = io.WriteString(w, "\n                        ")
				b := ((*bucket)(unsafe.Pointer(&e.value()[0])))
				_, _ = io.WriteString(w, "\n                        ")

				util := "-"
				if b.root != 0 {
					substats := pageAt(tx, b.root).stats(pageSize)
					util = fmt.Sprintf("%.2f%%", substats.utilization*100)
				}

				_, _ = io.WriteString(w, "\n                        <tr>\n                          <td><strong>")
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(trunc(tostr(e.key()), 150))))
				_, _ = io.WriteString(w, "</strong></td>\n                          <td>\n                            &lt;bucket(root=")
				if b.root != 0 {
					_, _ = io.WriteString(w, "<a href=\"")
					_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(subpagelink(indexes, int(i)))))
					_, _ = io.WriteString(w, "\">")
				}
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(b.root)))
				if b.root != 0 {
					_, _ = io.WriteString(w, "</a>")
				}
				_, _ = io.WriteString(w, "; seq=")
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(b.sequence)))
				_, _ = io.WriteString(w, ")&gt;\n                          </td>\n                          <td>")
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(len(e.key()))))
				_, _ = io.WriteString(w, " / ")
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(len(e.value()))))
				_, _ = io.WriteString(w, "</td>\n                          <td align=\"right\">")
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(util)))
				_, _ = io.WriteString(w, "</td>\n                        </tr>\n                      ")
			} else {
				_, _ = io.WriteString(w, "\n                        <tr>\n                          <td>")
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(trunc(tostr(e.key()), 150))))
				_, _ = io.WriteString(w, "</td>\n                          <td>")
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(trunc(tostr(e.value()), 5000))))
				_, _ = io.WriteString(w, "</td>\n                          <td>")
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(len(e.key()))))
				_, _ = io.WriteString(w, " / ")
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(len(e.value()))))
				_, _ = io.WriteString(w, "</td>\n                          <td>&nbsp;</td>\n                        </tr>\n                      ")
			}
			_, _ = io.WriteString(w, "\n                    ")
		}
		_, _ = io.WriteString(w, "\n                  </tbody>\n                </table>\n              ")
	}
	_, _ = io.WriteString(w, "\n\n              ")
	if showUsage {
		_, _ = io.WriteString(w, "\n                ")

		mins, maxs, values := bucketize(histogram)
		vmax, maxlen := 0, 20
		for _, v := range values {
			if v > vmax {
				vmax = v
			}
		}

		_, _ = io.WriteString(w, "\n\n                <h3>Page Usage Histogram</h3>\n                <table>\n                  <thead>\n                    <tr>\n                      <th align=\"left\">Usage (bytes)</th>\n                      <th align=\"left\">Count</th>\n                      <th>&nbsp;</th>\n                    </tr>\n                  </thead>\n                  <tbody>\n                    ")
		for i := 0; i < len(values); i++ {
			_, _ = io.WriteString(w, "\n                      <tr>\n                        <td>")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(mins[i])))
			_, _ = io.WriteString(w, " - ")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(maxs[i])))
			_, _ = io.WriteString(w, "</th>\n                        <td>")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(values[i])))
			_, _ = io.WriteString(w, "</th>\n                        <td>")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(strings.Repeat("█", int((float64(values[i])/float64(vmax))*float64(maxlen))))))
			_, _ = io.WriteString(w, "</td>\n                      </tr>\n                    ")
		}
		_, _ = io.WriteString(w, "\n                  </tbody>\n                </table>\n              ")
	} else {
		_, _ = io.WriteString(w, "\n                ")

		u, q := r.URL, r.URL.Query()
		q.Set("usage", "true")
		u.RawQuery = q.Encode()

		_, _ = io.WriteString(w, "\n\n                <p><a href=\"")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(u.String())))
		_, _ = io.WriteString(w, "\">Show Page Usage</a></p>\n              ")
	}
	_, _ = io.WriteString(w, "\n\n              <br/><br/>\n              <form action=\"boltdb\" method=\"GET\">\n                Go to page: <input type=\"text\" name=\"id\"/>\n                <button type=\"submit\">Go</button>\n              </form>\n\n            </div>\n    </div>\n\n</div>\n\n")
	common.Footer(w)
	_, _ = io.WriteString(w, "\n")
	return nil
}