# Operational Guide: PulseAI Insights (AI Manager Manual)

## 1. P&L Management Logic

### Daily ROI Calculation
- **Formula:** `ROI = (Daily Ad Revenue + Daily Affiliate Revenue) / (Daily API Cost + Daily Hosting Pro-rata)`.
- **Target:** ROI > 1.5 (50% profit margin).
- **Budget Cap:** $1.00/day for API calls (hard limit in script).

### Daily Report Format
```markdown
### 📊 Daily P&L Summary: [Date]
- **Revenue:** $[Amount] (AdSense: $[...], Amazon: $[...])
- **Costs:** $[Amount] (Gemini: $[...], NewsAPI: $[...])
- **Net Profit:** $[Amount]
- **ROI:** [X.X]
- **Status:** [🟢 Healthy / 🟡 Warning / 🔴 Action Required]
```

- [x] Define the formula for Daily ROI calculation.
- [x] Specify the 'Budget Cap' for API spending.
- [x] Design the daily P&L summary report format for the user.

## 2. Error Handling & Recovery

### Retry Logic
- **Scout/Editor:** 3 retries with exponential backoff (5s, 30s, 5m).
- **Publisher:** If social API fails, log error and retry once after 1 hour. Do not block the site build.

### Emergency Stop
- **Condition:** If daily cost exceeds $2.00 or 3 consecutive runs fail.
- **Action:** Disable GitHub Action cron and notify user via `notify_user` tool.

- [x] Define retry logic for failed API calls (Scout/Editor/Publisher).
- [x] Specify the 'Emergency Stop' conditions (e.g., cost spike, 403 errors).
- [x] Create a 'Manual Override' procedure for the user.

## 3. Self-Optimization Protocols

### A/B Testing & Source Priority
- **Headline Testing:** AI generates 2 headlines; the one with higher CTR (from previous data) is used.
- **Source Weighting:** Sources that generate more clicks (via UTM parameters) are prioritized in the next day's 'Scout' run.

### Model Switching
- **Logic:** 
    - If ROI > 2.0: Use Gemini 2.0 Pro for higher quality summaries.
    - If ROI < 1.2: Switch to Gemini 2.0 Flash-Lite to cut costs.

- [x] Define the 'A/B Testing' logic for headlines and social hooks.
- [x] Specify how the AI adjusts source priority based on traffic data.
- [x] Design the 'Model Switching' logic (Flash vs. Pro) based on ROI.

## 4. Maintenance & Security
- **Secret Rotation:** Prompt user to rotate API keys every 90 days.
- **Stealth Check:** Script scans generated HTML for strings like "Agent Zero", "Localhost", or user's real name before deployment.

- [x] Define the secret rotation schedule (API keys).
- [x] Specify the backup procedure for the article database and P&L logs.
- [x] Outline the 'Stealth Check' (ensuring no identity leaks in metadata).
