# ⚡ Cappit — API Rate Limiting & Quota Management

**Cappit** is a developer-first SaaS platform that provides **secure API key management**, **custom rate limiting**, and **reverse proxying** — all without the complexity of managing your own gateway infrastructure.

Whether you're running a public API or a SaaS platform, Cappit helps you protect your APIs from overuse, abuse, and unexpected surges.

---

## 🚀 What Cappit Does

### 🔑 API Key Management
- Create, manage, revoke, and rotate API keys per app or consumer
- Associate keys with custom usage plans
- Group keys under tenants for multi-user SaaS models

### 🚦 Flexible Rate Limiting Engine
- Enforce:
    - Per-minute
    - Per-hour
    - Per-day
    - Per-month
- Set quotas globally, per key, or per route
- Rate limiting powered by Redis — blazing fast and accurate

### 🔁 Hosted Reverse Proxy
- Cappit proxies your API requests securely
- Blocks requests that exceed usage limits
- Forwards valid requests to your real backend
- No need to change your backend code

### 📊 Usage Monitoring
- Real-time usage dashboards per key or tenant
- Webhooks on quota breach events
- Daily/monthly exports or API access to logs

### 🧱 Multi-Tenant Architecture
- Each of your customers can manage their own keys and limits
- Isolated usage tracking per tenant
- Ideal for SaaS startups offering developer APIs

---

## 🧭 Why Choose Cappit?

There are many API gateways and limiters — but Cappit stands out by focusing on **simplicity, speed, and developer experience**:

### ✅ Focused on Rate Limiting
We don’t try to be everything. Cappit does one thing extremely well — rate limits + key auth.

### ✅ Dev UX First
Clean HTTP APIs. Swagger docs. Fast onboarding. Built by devs, for devs.

### ✅ Works with Any Stack
REST, GraphQL, internal APIs — if it runs on HTTP, Cappit can protect it.

---