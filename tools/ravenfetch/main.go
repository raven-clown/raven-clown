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
	{"cmd", "$ go test -race ./..."},
	{"ok", "ok  airlock/internal/gate         0.41s"},
	{"ok", "ok  rosetta/internal/bridge       0.37s"},
	{"ok", "ok  transponder/internal/router   0.52s"},
	{"ok", "ok  logbook/internal/recall       0.29s"},
	{"cmd", "$ kubectl rollout status deploy"},
	{"dim", "deployment/airlock configured"},
	{"dim", "deployment/transponder configured"},
	{"ok", "3/3 replicas ready"},
	{"cmd", "$ transponder --check"},
	{"dim", "route chat -> anthropic, openai, ollama"},
	{"dim", "budget user:*  $20.00/mo  used $3.41"},
	{"dim", "limits free  20k tok/min  60 req/min"},
	{"warn", "openai 503: cooldown 30s, -> anthropic"},
	{"ok", "200 OK  in=812 out=240 think=96"},
	{"cmd", "$ kafka-consumer-groups --describe"},
	{"dim", "mes.events  p0  lag 0"},
	{"dim", "mes.events  p1  lag 0"},
	{"dim", "mes.events  p2  lag 2"},
	{"cmd", "$ curl -s :8080/readyz"},
	{"ok", "{\"status\":\"ready\",\"replicas\":3}"},
	{"cmd", "$ logbook recall \"deploy notes\""},
	{"dim", "[0.94] pgx pool size 20 (verified)"},
	{"dim", "[0.88] rotate keys on the 1st"},
	{"cmd", "$ git push origin main"},
	{"ok", "ci: 9/9 checks passed"},
	{"ok", "all systems nominal."},
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

	paneX0, paneY0, paneX1, paneY1 = 22.0, 86.0, 405.0, 446.0
	promptY                        = 470.0
	runLH                          = 15.0
	runFS                          = 10.5

	// The left pane loops on a cycle that starts once the first command is typed.
	loopBegin = 1.2
	cycle     = 16.0
	tArt0     = 0.02
	tArt1     = 0.36
	tEnter0   = 0.50
	tEnter1   = 0.56
	tGone     = 0.58
	tRun0     = 0.60
	tRun1     = 0.92
	tEnd      = 0.97
)

type svg struct{ strings.Builder }

func (b *svg) p(format string, a ...any) { fmt.Fprintf(b, format, a...) }

func (b *svg) loopReveal(id string, x, y, height, full, t0, t1, hide float64) {
	b.p(`<clipPath id="%s"><rect x="%.1f" y="%.1f" height="%.1f" width="0"><animate attributeName="width" values="0;0;%.1f;%.1f;0;0" keyTimes="0;%.4f;%.4f;%.4f;%.4f;1" begin="%.1fs" dur="%.0fs" repeatCount="indefinite"/></rect></clipPath>`,
		id, x, y, height, full, full, t0, t1, hide, hide+0.001, loopBegin, cycle)
}

func (b *svg) onceReveal(id string, x, y, height, full, begin, dur float64) {
	b.p(`<clipPath id="%s"><rect x="%.1f" y="%.1f" height="%.1f" width="0"><animate attributeName="width" from="0" to="%.1f" begin="%.2fs" dur="%.2fs" fill="freeze"/></rect></clipPath>`,
		id, x, y, height, full, begin, dur)
}

func loopWindow(t0, t1 float64) string {
	return fmt.Sprintf(`<animate attributeName="opacity" values="0;1;0" keyTimes="0;%.4f;%.4f" calcMode="discrete" begin="%.1fs" dur="%.0fs" repeatCount="indefinite"/>`, t0, t1, loopBegin, cycle)
}

func onceShow(begin float64) string {
	return fmt.Sprintf(`<animate attributeName="opacity" from="0" to="1" begin="%.2fs" dur="0.25s" fill="freeze"/>`, begin)
}

func render(grid [][]cell, t theme) string {
	var b svg
	b.p(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d" role="img" aria-label="ravenfetch: Ekdanai Kummee, Application Engineer. A raven in a top hat holding a lantern, typed out in code.">`, w, h, w, h)
	b.p(`<defs>`)
	b.p(`<linearGradient id="bg" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient>`, t.bg0, t.bg1)
	b.p(`<linearGradient id="art" x1="%.0f" y1="%.0f" x2="%.0f" y2="%.0f" gradientUnits="userSpaceOnUse">`, artX, artY-lh, artX+cols*charW*0.5, artY+rows*lh)
	b.p(`<stop offset="0" stop-color="%s"><animate attributeName="stop-color" values="%s;%s;%s" dur="6s" repeatCount="indefinite"/></stop>`, t.a0, t.a0, t.a1, t.a0)
	b.p(`<stop offset="1" stop-color="%s"><animate attributeName="stop-color" values="%s;%s;%s" dur="6s" repeatCount="indefinite"/></stop>`, t.a1, t.a1, t.a0, t.a1)
	b.p(`</linearGradient>`)
	b.p(`<linearGradient id="bar" x1="0" y1="0" x2="1" y2="0"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient>`, t.a0, t.a1)
	b.p(`<radialGradient id="flameglow"><stop offset="0" stop-color="%s" stop-opacity="0.55"/><stop offset="1" stop-color="%s" stop-opacity="0"/></radialGradient>`, t.glow, t.glow)
	b.p(`<clipPath id="pane"><rect x="%.0f" y="%.0f" width="%.0f" height="%.0f"/></clipPath>`, paneX0, paneY0, paneX1-paneX0, paneY1-paneY0)

	b.onceReveal("cmd", 58, 60, 20, 80, 0.2, 0.7)
	for r := range grid {
		y := artY + float64(r)*lh
		t0 := tArt0 + (tArt1-tArt0)*float64(r)/rows
		t1 := tArt0 + (tArt1-tArt0)*float64(r+1)/rows
		b.loopReveal(fmt.Sprintf("r%d", r), artX, y-lh+1, lh+1, cols*charW+2, t0, t1, tGone)
	}
	infoBegin := 1.0
	b.onceReveal("head", infoX, infoY-12, 16, 24*infoCharW, infoBegin, 0.25)
	for i, l := range lines {
		y := infoY + ilh*1.4 + float64(i+1)*(ilh+6)
		full := float64(len(l.k)+2+len(l.v))*infoCharW + 4
		b.onceReveal(fmt.Sprintf("i%d", i), infoX, y-12, 16, full, infoBegin+0.3+float64(i)*0.22, 0.2)
	}
	enterCmd := "./mothership --launch"
	b.loopReveal("enter", 58, promptY-13, 18, float64(len(enterCmd))*infoCharW+4, tEnter0, tEnter1, tGone)
	b.p(`</defs>`)
	b.p(`<style>text{font-family:ui-monospace,SFMono-Regular,Menlo,Consolas,"Liberation Mono",monospace;font-size:%.1fpx}.art text{font-size:8.4px}.run text{font-size:%.1fpx}</style>`, fs, runFS)

	b.p(`<rect x="1" y="1" width="%d" height="%d" rx="14" fill="url(#bg)" stroke="%s"/>`, w-2, h-2, t.border)
	b.p(`<path d="M1 15a14 14 0 0 1 14-14h%dv0a14 14 0 0 1 14 14v19H1z" fill="%s"/>`, w-30, t.bar)
	b.p(`<rect x="1" y="34" width="%d" height="1" fill="%s"/>`, w-2, t.border)
	for i, c := range []string{"#ff5f57", "#febc2e", "#28c840"} {
		b.p(`<circle cx="%d" cy="18" r="6" fill="%s"/>`, 24+i*20, c)
	}
	b.p(`<text x="%d" y="22" text-anchor="middle" fill="%s">raven@mothership: ~</text>`, w/2, t.dim)
	b.p(`<rect x="%.0f" y="52" width="1" height="%.0f" fill="%s" opacity="0">%s</rect>`, paneX1+12, float64(h)-70, t.border, onceShow(0.8))

	b.p(`<text x="34" y="74" fill="%s">❯</text>`, t.prompt)
	b.p(`<text x="58" y="74" fill="%s" clip-path="url(#cmd)">ravenfetch</text>`, t.text)

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
		b.p(`<g opacity="0">%s`, loopWindow(fOn, tGone))
		b.p(`<circle cx="%.1f" cy="%.1f" r="26" fill="url(#flameglow)"><animate attributeName="r" values="24;30;22;28;24" dur="1.6s" repeatCount="indefinite"/></circle>`, cx, cy)
		b.p(`<g class="art" fill="%s">`, t.flame)
		for _, f := range flames {
			b.p(`<text x="%.1f" y="%.1f">*<animate attributeName="opacity" values="1;0.55;1;0.8;1" dur="0.9s" repeatCount="indefinite"/></text>`, f[0], f[1])
		}
		b.p(`</g></g>`)
	}

	b.p(`<text x="%.0f" y="%.1f" clip-path="url(#head)"><tspan fill="%s" font-weight="700">raven-clown</tspan><tspan fill="%s">@</tspan><tspan fill="%s" font-weight="700">mothership</tspan></text>`, infoX, infoY, t.a0, t.dim, t.a1)
	b.p(`<rect x="%.0f" y="%.1f" width="0" height="2" rx="1" fill="url(#bar)"><animate attributeName="width" from="0" to="%.0f" begin="%.2fs" dur="0.4s" fill="freeze"/></rect>`, infoX, infoY+ilh-5, 22*infoCharW, infoBegin+0.15)
	for i, l := range lines {
		y := infoY + ilh*1.4 + float64(i+1)*(ilh+6)
		b.p(`<text x="%.0f" y="%.1f" clip-path="url(#i%d)"><tspan fill="%s" font-weight="700">%s</tspan><tspan fill="%s">: </tspan><tspan fill="%s">%s</tspan></text>`,
			infoX, y, i, t.key, html.EscapeString(l.k), t.dim, t.text, html.EscapeString(l.v))
	}
	sy := infoY + ilh*1.4 + float64(len(lines)+1)*(ilh+6) + 16
	swBegin := infoBegin + 0.3 + float64(len(lines))*0.22
	for i, c := range []string{"#ff5f57", "#febc2e", "#28c840", t.a0, t.a1, "#a855f7", t.flame, t.text} {
		b.p(`<rect x="%.0f" y="%.1f" width="22" height="12" rx="2" fill="%s" opacity="0">%s</rect>`, infoX+float64(i)*26, sy-10, c, onceShow(swBegin+float64(i)*0.06))
	}

	b.p(`<g opacity="0">%s`, loopWindow(tArt1, tGone))
	b.p(`<text x="34" y="%.1f" fill="%s">❯</text>`, promptY, t.prompt)
	b.p(`<text x="58" y="%.1f" fill="%s" clip-path="url(#enter)">%s</text>`, promptY, t.text, html.EscapeString(enterCmd))
	endX := 60 + float64(len(enterCmd))*infoCharW
	b.p(`<rect x="58" y="%.1f" width="8" height="14" fill="%s"><animate attributeName="x" values="58;58;%.1f;%.1f" keyTimes="0;%.4f;%.4f;1" begin="%.1fs" dur="%.0fs" repeatCount="indefinite"/><animate attributeName="opacity" values="1;0;1" keyTimes="0;0.5;1" calcMode="discrete" dur="0.9s" repeatCount="indefinite"/></rect>`,
		promptY-11, t.text, endX, endX, tEnter0, tEnter1, loopBegin, cycle)
	b.p(`</g>`)

	// Terminal output: each step shows one more line at the bottom and pushes the rest up.
	n := len(runLog)
	firstY := paneY1 - 6 + runLH
	var kt, vals []string
	kt, vals = append(kt, "0"), append(vals, "0 0")
	for k := 1; k <= n; k++ {
		kt = append(kt, fmt.Sprintf("%.4f", tRun0+(tRun1-tRun0)*float64(k-1)/float64(n)))
		vals = append(vals, fmt.Sprintf("0 %.0f", -float64(k)*runLH))
	}
	b.p(`<g class="run" clip-path="url(#pane)"><g opacity="0">%s`, loopWindow(tRun0, tEnd))
	b.p(`<g><animateTransform attributeName="transform" type="translate" calcMode="discrete" values="%s" keyTimes="%s" begin="%.1fs" dur="%.0fs" repeatCount="indefinite"/>`, strings.Join(vals, ";"), strings.Join(kt, ";"), loopBegin, cycle)
	for i, l := range runLog {
		color, weight := t.dim, ""
		switch l.kind {
		case "cmd":
			color, weight = t.text, ` font-weight="700"`
		case "ok":
			color = t.ok
		case "warn":
			color = t.warn
		}
		b.p(`<text x="34" y="%.1f" fill="%s"%s xml:space="preserve">%s</text>`, firstY+float64(i)*runLH, color, weight, html.EscapeString(l.text))
	}
	b.p(`</g></g></g>`)
	b.p(`<g opacity="0">%s<text x="34" y="%.1f" fill="%s">❯</text><rect x="58" y="%.1f" width="8" height="14" fill="%s"><animate attributeName="opacity" values="1;0;1" keyTimes="0;0.5;1" calcMode="discrete" dur="0.6s" repeatCount="indefinite"/></rect></g>`,
		loopWindow(tRun1, 1), promptY, t.prompt, promptY-11, t.text)
	b.p(`</svg>`)
	return b.String()
}
