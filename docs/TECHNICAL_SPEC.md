# Technical Specification: PulseAI Insights

## 1. System Architecture Overview
[... previously defined ...]

## 2. The Scout (Data Ingestion)
[... previously defined ...]

## 3. The Editor (AI Processing)
[... previously defined ...]

## 4. The Publisher (Deployment & Social)

### Deployment Script (GitHub Actions + Vercel)
- **Build Command:** `hugo --minify`
- **Deploy Command:** `vercel deploy --prebuilt --token=$VERCEL_TOKEN` (using Vercel CLI in GitHub Actions).
- **Environment Variables:** `VERCEL_TOKEN`, `VERCEL_ORG_ID`, `VERCEL_PROJECT_ID`.

### Social Media Integration
- **X (Twitter):** Use `tweepy` library for API v2. Endpoint: `POST /2/tweets`.
- **LinkedIn:** Use `requests` to call the LinkedIn Marketing Developer Platform API. Endpoint: `POST /v2/ugcPosts`.
- **Authentication:** OAuth 2.0 with Refresh Tokens stored in GitHub Secrets.

### Social Post Templates
- **Format:** [Hook] + [Summary] + [Link] + [Hashtags]
- **Example:** "🚀 New breakthrough in AI Radiology: [Title]. AI is now [Summary snippet]. Read the full breakdown: [URL] #HealthAI #Radiology #PulseAI"

- [x] Define the Vercel/GitHub Pages deployment script.
- [x] Specify the X (Twitter) and LinkedIn API integration details.
- [x] Design the social post templates (hooks, hashtags).

## 5. Data Schema & Storage

### Article Markdown Frontmatter (Hugo)
```yaml
title: "AI-Driven Diagnostic Tool for Early Cancer Detection"
date: 2026-03-06T07:00:00Z
draft: false
summary: "A new study shows..."
categories: ["Diagnostics", "Oncology"]
tags: ["Cancer", "AI", "Radiology"]
source_url: "https://pubmed..."
affiliate_link: "https://amazon..."
```

### Tracking & Deduplication
- **File:** `src/data/processed_hashes.txt`
- **Mechanism:** Before processing, the Scout generates an MD5 hash of the article URL. If the hash exists in this file, the article is skipped. After successful publication, the hash is appended to the file.

- [x] Define the structure of the `articles.json` or Markdown frontmatter.
- [x] Specify how 'already published' articles are tracked to avoid duplicates.
