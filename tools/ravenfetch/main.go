package main

import (
	"fmt"
	"html"
	"image"
	_ "image/jpeg"
	"os"
	"path/filepath"
	"strings"
)

const (
	cropX0, cropY0, cropX1, cropY1 = 132, 90, 330, 290
	cols, rows                     = 72, 38
	ramp                           = " .:-=+*#%@"
	fill                           = "func(raven)0x2EE6A6{go#mothership<-airlock}$rosetta&transponder%logbook=>kafka|nifi|k8s[01]@SSO;RBAC!0xDE1A7A"
)

type cell struct {
	ch    byte
	flame bool
}

type theme struct {
	name, bg0, bg1, bar, border, text, dim, key, a0, a1, flame, glow, prompt, ok, warn string
}

var themes = []theme{
	{"dark", "#0b0f17", "#121a2b", "#0f1522", "#1f2a3d", "#c9d1d9", "#6e7b8f", "#58a6ff", "#58a6ff", "#2ee6a6", "#ff7a45", "#ffb36b", "#2ee6a6", "#2ee6a6", "#febc2e"},
	{"light", "#ffffff", "#f3f6fa", "#eef2f7", "#d0d7de", "#24292f", "#6e7781", "#0969da", "#0969da", "#1a9e74", "#e8590c", "#f59f00", "#1a9e74", "#1a7f37", "#9a6700"},
}

type info struct{ k, v string }

var lines = []info{
	{"Name", "Ekdanai Kummee"},
	{"Role", "Application Engineer, Delta Electronics Thailand"},
	{"Focus", "DevOps, data platforms, full-stack web"},
	{"Platforms", "central Kafka and NiFi, OpenSearch log analytics"},
	{"Stack", "Go, TypeScript, Kafka, Kubernetes, React"},
	{"Building", "airlock, rosetta, transponder, logbook"},
	{"Projects", "ARK, tack, IdpForge"},
	{"Upstream", "20 merged PRs in 11 open source projects"},
	{"Uptime", "5 years of game servers, now factory data"},
	{"Education", "B.Sc., First Class Honors, Gold Medal"},
	{"Contact", "ekdanai.kk@gmail.com"},
}

type logLine struct{ kind, text string }

var runLog = []logLine{
	{"cmd", "$ ./mothership --launch"},
	{"dim", "[boot] loading CONTRACT.md ......................... ok"},
	{"dim", "[boot] identity: oidc issuer reachable, jwks cached"},
	{"cmd", "$ go build ./... && go test -race ./..."},
	{"ok", "ok   github.com/raven-clown/airlock/internal/gate      0.41s"},
	{"ok", "ok   github.com/raven-clown/rosetta/internal/bridge    0.37s"},
	{"ok", "ok   github.com/raven-clown/transponder/internal/router 0.52s"},
	{"ok", "ok   github.com/raven-clown/logbook/internal/recall    0.29s"},
	{"cmd", "$ kubectl apply -k deploy/ && kubectl rollout status deploy"},
	{"dim", "deployment.apps/airlock configured"},
	{"dim", "deployment.apps/transponder configured"},
	{"ok", "rollout: 3/3 replicas ready  [airlock transponder logbook]"},
	{"cmd", "$ transponder --check"},
	{"dim", "route chat-default  -> anthropic, openai, ollama  (fallback)"},
	{"dim", "budget user:*      $20.00/month  used $3.41"},
	{"dim", "limits free        20000 tok/min  60 req/min"},
	{"warn", "target openai: 503, cooling down 30s, trying anthropic"},
	{"ok", "200 OK  in=812 out=240 reasoning=96 cost=$0.0041"},
	{"cmd", "$ kafka-consumer-groups --describe --group mes-pipeline"},
	{"dim", "mes.events   p0  offset 18842071  lag 0"},
	{"dim", "mes.events   p1  offset 18839930  lag 0"},
	{"dim", "mes.events   p2  offset 18840455  lag 2"},
	{"cmd", "$ curl -s localhost:8080/readyz"},
	{"ok", "{\"status\":\"ready\",\"store\":\"postgres\",\"replicas\":3}"},
	{"cmd", "$ logbook recall \"deploy notes\" --verified"},
	{"dim", "[0.94] staging uses pgx pool size 20 (verified by ci)"},
	{"dim", "[0.88] rotate provider keys on the 1st (verified)"},
	{"cmd", "$ git push origin main"},
	{"ok", "main -> main  ci: 9/9 checks passed"},
	{"ok", "all systems nominal. see you next loop."},
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	grid := sample(img)
	for _, t := range themes {
		out := filepath.Join(os.Args[2], "ravenfetch-"+t.name+".svg")
		if err := os.WriteFile(out, []byte(render(grid, t)), 0o644); err != nil {
			panic(err)
		}
		fmt.Println("wrote", out)
	}
}

func sample(img image.Image) [][]cell {
	cw := float64(cropX1-cropX0) / cols
	ch := float64(cropY1-cropY0) / rows
	grid := make([][]cell, rows)
	k := 0
	for r := 0; r < rows; r++ {
		grid[r] = make([]cell, cols)
		for c := 0; c < cols; c++ {
			x0, y0 := cropX0+int(float64(c)*cw), cropY0+int(float64(r)*ch)
			x1, y1 := cropX0+int(float64(c+1)*cw), cropY0+int(float64(r+1)*ch)
			var dark, red, n int
			for y := y0; y < y1; y++ {
				for x := x0; x < x1; x++ {
					R, G, B, _ := img.At(x, y).RGBA()
					rr, gg, bb := R>>8, G>>8, B>>8
					n++
					if rr > 150 && gg < 110 && bb < 110 {
						red++
					} else if (rr+gg+bb)/3 < 110 {
						dark++
					}
				}
			}
			if n == 0 {
				continue
			}
			if red*12 >= n {
				grid[r][c] = cell{ch: '*', flame: true}
				continue
			}
			cov := float64(dark) / float64(n)
			if cov < 0.12 {
				continue
			}
			i := int(cov*float64(len(ramp)-1) + 0.35)
			if i >= len(ramp) {
				i = len(ramp) - 1
			}
			if i >= 6 {
				grid[r][c] = cell{ch: fill[k%len(fill)]}
				k++
				continue
			}
			grid[r][c] = cell{ch: ramp[i]}
		}
	}
	return grid
}

const (
	w, h       = 900, 490
	fs         = 11.5
	lh         = 9.0
	ilh        = 13.0
	charW      = 5.04
	infoCharW  = 6.9
	artX, artY = 30.0, 100.0
	infoX      = 430.0
	infoY      = 140.0
	cycle      = 16.0

	tCmd0, tCmd1   = 0.01, 0.05
	tArt0, tArt1   = 0.06, 0.40
	tInfo0, tInfo1 = 0.41, 0.52
	tEnter0        = 0.60
	tEnter1        = 0.66
	tGone          = 0.68
	tRun0, tRun1   = 0.69, 0.95
	tEnd           = 0.97
)

type svg struct{ strings.Builder }

func (b *svg) p(format string, a ...any) { fmt.Fprintf(b, format, a...) }

// reveal animates a clip rect's width: hidden until t0, wiped open by t1, held until hide, then gone.
func (b *svg) reveal(id string, x, y, height, full, t0, t1, hide float64) {
	b.p(`<clipPath id="%s"><rect x="%.1f" y="%.1f" height="%.1f" width="0"><animate attributeName="width" values="0;0;%.1f;%.1f;0;0" keyTimes="0;%.4f;%.4f;%.4f;%.4f;1" dur="%.0fs" repeatCount="indefinite"/></rect></clipPath>`,
		id, x, y, height, full, full, t0, t1, hide, hide+0.001, cycle)
}

// window shows an element from t0 to t1 of every cycle.
func window(t0, t1 float64) string {
	return fmt.Sprintf(`<animate attributeName="opacity" values="0;1;0" keyTimes="0;%.4f;%.4f" calcMode="discrete" dur="%.0fs" repeatCount="indefinite"/>`, t0, t1, cycle)
}

func render(grid [][]cell, t theme) string {
	var b svg
	b.p(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d" role="img" aria-label="ravenfetch: Ekdanai Kummee, Application Engineer. A raven in a top hat holding a lantern, typed out in ASCII.">`, w, h, w, h)
	b.p(`<defs>`)
	b.p(`<linearGradient id="bg" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient>`, t.bg0, t.bg1)
	b.p(`<linearGradient id="art" x1="%.0f" y1="%.0f" x2="%.0f" y2="%.0f" gradientUnits="userSpaceOnUse">`, artX, artY-lh, artX+cols*charW*0.5, artY+rows*lh)
	b.p(`<stop offset="0" stop-color="%s"><animate attributeName="stop-color" values="%s;%s;%s" dur="6s" repeatCount="indefinite"/></stop>`, t.a0, t.a0, t.a1, t.a0)
	b.p(`<stop offset="1" stop-color="%s"><animate attributeName="stop-color" values="%s;%s;%s" dur="6s" repeatCount="indefinite"/></stop>`, t.a1, t.a1, t.a0, t.a1)
	b.p(`</linearGradient>`)
	b.p(`<linearGradient id="bar" x1="0" y1="0" x2="1" y2="0"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient>`, t.a0, t.a1)
	b.p(`<radialGradient id="flameglow"><stop offset="0" stop-color="%s" stop-opacity="0.55"/><stop offset="1" stop-color="%s" stop-opacity="0"/></radialGradient>`, t.glow, t.glow)
	b.p(`<clipPath id="body"><rect x="20" y="44" width="%d" height="%d"/></clipPath>`, w-40, h-58)

	b.reveal("cmd", 58, 60, 20, 80, tCmd0, tCmd1, tGone)
	for r := range grid {
		y := artY + float64(r)*lh
		t0 := tArt0 + (tArt1-tArt0)*float64(r)/rows
		t1 := tArt0 + (tArt1-tArt0)*float64(r+1)/rows
		b.reveal(fmt.Sprintf("r%d", r), artX, y-lh+1, lh+1, cols*charW+2, t0, t1, tGone)
	}
	b.reveal("head", infoX, infoY-12, 16, 24*infoCharW, tInfo0, tInfo0+0.01, tGone)
	for i, l := range lines {
		y := infoY + ilh*1.4 + float64(i+1)*(ilh+6)
		t0 := tInfo0 + 0.01 + (tInfo1-tInfo0-0.01)*float64(i)/float64(len(lines))
		t1 := tInfo0 + 0.01 + (tInfo1-tInfo0-0.01)*float64(i+1)/float64(len(lines))
		full := float64(len(l.k)+2+len(l.v))*infoCharW + 4
		b.reveal(fmt.Sprintf("i%d", i), infoX, y-12, 16, full, t0, t1, tGone)
	}
	enterCmd := "./mothership --launch"
	b.reveal("enter", 58, float64(h)-41, 18, float64(len(enterCmd))*infoCharW+4, tEnter0, tEnter1, tGone)
	b.p(`</defs>`)
	b.p(`<style>text{font-family:ui-monospace,SFMono-Regular,Menlo,Consolas,"Liberation Mono",monospace;font-size:%.1fpx}.art text{font-size:8.4px}</style>`, fs)

	b.p(`<rect x="1" y="1" width="%d" height="%d" rx="14" fill="url(#bg)" stroke="%s"/>`, w-2, h-2, t.border)
	b.p(`<path d="M1 15a14 14 0 0 1 14-14h%dv0a14 14 0 0 1 14 14v19H1z" fill="%s"/>`, w-30, t.bar)
	b.p(`<rect x="1" y="34" width="%d" height="1" fill="%s"/>`, w-2, t.border)
	for i, c := range []string{"#ff5f57", "#febc2e", "#28c840"} {
		b.p(`<circle cx="%d" cy="18" r="6" fill="%s"/>`, 24+i*20, c)
	}
	b.p(`<text x="%d" y="22" text-anchor="middle" fill="%s">raven@mothership: ~</text>`, w/2, t.dim)

	b.p(`<g>%s`, window(0, tGone))
	b.p(`<text x="34" y="74" fill="%s">❯</text>`, t.prompt)
	b.p(`<text x="58" y="74" fill="%s" clip-path="url(#cmd)">ravenfetch</text>`, t.text)
	b.p(`</g>`)

	b.p(`<g class="art" fill="url(#art)">`)
	var flames [][2]float64
	for r, row := range grid {
		y := artY + float64(r)*lh
		var s strings.Builder
		for ci, c := range row {
			switch {
			case c.flame:
				s.WriteByte(' ')
				flames = append(flames, [2]float64{artX + float64(ci)*charW, y})
			case c.ch == 0:
				s.WriteByte(' ')
			default:
				s.WriteByte(c.ch)
			}
		}
		if line := strings.TrimRight(s.String(), " "); line != "" {
			b.p(`<text x="%.1f" y="%.1f" xml:space="preserve" clip-path="url(#r%d)">%s</text>`, artX, y, r, html.EscapeString(line))
		}
	}
	b.p(`</g>`)

	if len(flames) > 0 {
		var sx, sy float64
		for _, f := range flames {
			sx += f[0]
			sy += f[1]
		}
		cx, cy := sx/float64(len(flames))+charW/2, sy/float64(len(flames))-4
		frow := (flames[0][1] - artY) / lh
		fOn := tArt0 + (tArt1-tArt0)*(frow+1)/rows
		b.p(`<g opacity="0">%s`, window(fOn, tGone))
		b.p(`<circle cx="%.1f" cy="%.1f" r="26" fill="url(#flameglow)"><animate attributeName="r" values="24;30;22;28;24" dur="1.6s" repeatCount="indefinite"/></circle>`, cx, cy)
		b.p(`<g class="art" fill="%s">`, t.flame)
		for _, f := range flames {
			b.p(`<text x="%.1f" y="%.1f">*<animate attributeName="opacity" values="1;0.55;1;0.8;1" dur="0.9s" repeatCount="indefinite"/></text>`, f[0], f[1])
		}
		b.p(`</g></g>`)
	}

	b.p(`<text x="%.0f" y="%.1f" clip-path="url(#head)"><tspan fill="%s" font-weight="700">raven-clown</tspan><tspan fill="%s">@</tspan><tspan fill="%s" font-weight="700">mothership</tspan></text>`, infoX, infoY, t.a0, t.dim, t.a1)
	b.p(`<rect x="%.0f" y="%.1f" width="%.0f" height="2" rx="1" fill="url(#bar)" clip-path="url(#head)"/>`, infoX, infoY+ilh-5, 22*infoCharW)
	b.p(`<rect x="%.0f" y="%.1f" width="%.0f" height="2" rx="1" fill="url(#bar)" opacity="0">%s</rect>`, infoX, infoY+ilh-5, 22*infoCharW, window(tInfo0+0.01, tGone))
	for i, l := range lines {
		y := infoY + ilh*1.4 + float64(i+1)*(ilh+6)
		b.p(`<text x="%.0f" y="%.1f" clip-path="url(#i%d)"><tspan fill="%s" font-weight="700">%s</tspan><tspan fill="%s">: </tspan><tspan fill="%s">%s</tspan></text>`,
			infoX, y, i, t.key, html.EscapeString(l.k), t.dim, t.text, html.EscapeString(l.v))
	}
	sy := infoY + ilh*1.4 + float64(len(lines)+1)*(ilh+6) + 16
	for i, c := range []string{"#ff5f57", "#febc2e", "#28c840", t.a0, t.a1, "#a855f7", t.flame, t.text} {
		b.p(`<rect x="%.0f" y="%.1f" width="22" height="12" rx="2" fill="%s" opacity="0">%s</rect>`, infoX+float64(i)*26, sy-10, c, window(tInfo1+float64(i)*0.004, tGone))
	}

	py := float64(h) - 28
	b.p(`<g opacity="0">%s`, window(tInfo1, tGone))
	b.p(`<text x="34" y="%.1f" fill="%s">❯</text>`, py, t.prompt)
	b.p(`<text x="58" y="%.1f" fill="%s" clip-path="url(#enter)">%s</text>`, py, t.text, html.EscapeString(enterCmd))
	b.p(`<rect x="58" y="%.1f" width="8" height="14" fill="%s"><animate attributeName="x" values="58;58;%.1f;%.1f" keyTimes="0;%.4f;%.4f;1" dur="%.0fs" repeatCount="indefinite"/><animate attributeName="opacity" values="1;0;1" keyTimes="0;0.5;1" calcMode="discrete" dur="0.9s" repeatCount="indefinite"/></rect>`,
		py-11, t.text, 60+float64(len(enterCmd))*infoCharW, 60+float64(len(enterCmd))*infoCharW, tEnter0, tEnter1, cycle)
	b.p(`</g>`)

	const runLH = 16.0
	startY := float64(h) - 20
	travel := startY - 60 + float64(len(runLog))*runLH
	b.p(`<g clip-path="url(#body)"><g opacity="0">%s`, window(tRun0, tEnd))
	b.p(`<g><animateTransform attributeName="transform" type="translate" values="0 0;0 0;0 %.0f;0 %.0f" keyTimes="0;%.4f;%.4f;1" dur="%.0fs" repeatCount="indefinite"/>`, -travel, -travel, tRun0, tRun1, cycle)
	for i, l := range runLog {
		color := t.dim
		weight := ""
		switch l.kind {
		case "cmd":
			color, weight = t.text, ` font-weight="700"`
		case "ok":
			color = t.ok
		case "warn":
			color = t.warn
		}
		b.p(`<text x="34" y="%.1f" fill="%s"%s xml:space="preserve">%s</text>`, startY+float64(i)*runLH, color, weight, html.EscapeString(l.text))
	}
	b.p(`</g></g></g>`)

	b.p(`<g opacity="0">%s<text x="34" y="74" fill="%s">❯</text><rect x="58" y="63" width="8" height="14" fill="%s"><animate attributeName="opacity" values="1;0;1" keyTimes="0;0.5;1" calcMode="discrete" dur="0.6s" repeatCount="indefinite"/></rect></g>`, window(tEnd, 1), t.prompt, t.text)
	b.p(`</svg>`)
	return b.String()
}
