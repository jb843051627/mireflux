package api

import "net/http"

func (h *Handler) dashboard(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = writer.Write([]byte(`<!doctype html>
<html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>Mireflux Field Operations</title><style>
body{margin:0;font-family:Arial,sans-serif;background:#eef3ee;color:#17201b}header{padding:18px 24px;background:#174a3c;color:#fff}.shell{max-width:1120px;margin:0 auto;padding:24px}.grid{display:grid;grid-template-columns:repeat(auto-fit,minmax(220px,1fr));gap:14px}.panel{background:#fff;border:1px solid #c8d5cb;padding:16px;border-radius:6px}h1{margin:0;font-size:24px}h2{font-size:15px;margin:0 0 10px}.metric{font-size:28px;font-weight:700}small{color:#51615a}button{background:#174a3c;color:#fff;border:0;padding:8px 12px;border-radius:4px}</style></head>
<body><header><h1>Mireflux</h1><small>Closed-chamber peatland flux operations</small></header><main class="shell"><div class="grid"><section class="panel"><h2>Campaigns</h2><div class="metric" id="campaigns">-</div></section><section class="panel"><h2>Service</h2><div class="metric" id="service">-</div></section><section class="panel"><h2>Recorded readings</h2><div class="metric" id="readings">-</div></section></div></main>
<script>Promise.all([fetch('/api/campaigns').then(r=>r.json()),fetch('/healthz').then(r=>r.json())]).then(([campaigns,health])=>{document.getElementById('campaigns').textContent=campaigns.length;document.getElementById('service').textContent='ready';document.getElementById('readings').textContent=(health.metrics||{})['readings.recorded']||0}).catch(()=>{document.getElementById('service').textContent='unavailable'})</script></body></html>`))
}
